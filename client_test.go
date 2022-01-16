package gotado

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	oauth2int "github.com/flipkick/gotado/internal/oauth2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

// mockHTTPClient mocks an HTTPClient by returning the stored response when Do is called.
type mockHTTPClient struct {
	Response *http.Response
	Error    error
}

// Do returns the Response stored in the mockHTTPClient.
func (c mockHTTPClient) Do(_ *http.Request) (*http.Response, error) {
	return c.Response, c.Error
}

func TestWithCredentials(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	config := oauth2.Config{}
	mockConfig := oauth2int.NewMockConfigInterface(ctrl)
	oauth2int.NewConfig = func(clientID, clientSecret, authURL, tokenURL string, scopes []string) oauth2int.ConfigInterface {
		return mockConfig
	}
	forbiddenError := &url.Error{}

	token := &oauth2.Token{
		AccessToken:  "access_token",
		TokenType:    "token_type",
		RefreshToken: "refresh_token",
		Expiry:       time.Now(),
	}

	client := NewClient("test", "test")
	httpCtx := context.WithValue(ctx, oauth2.HTTPClient, client.http)

	mockConfig.EXPECT().PasswordCredentialsToken(gomock.AssignableToTypeOf(httpCtx), "username", "password").Return(token, nil)
	mockConfig.EXPECT().Client(httpCtx, token).Return(config.Client(httpCtx, token))

	_, err := client.WithCredentials(ctx, "username", "password")
	assert.NoError(t, err)

	mockConfig.EXPECT().PasswordCredentialsToken(gomock.AssignableToTypeOf(httpCtx), "username", gomock.Not("password")).Return(nil, forbiddenError)

	_, err = client.WithCredentials(ctx, "username", "wrong")
	assert.Exactly(t, fmt.Errorf("invalid credentials: %w", forbiddenError), err)
}

func makeResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Status:     http.StatusText(code),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestGet(t *testing.T) {
	type foobar struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}

	tests := map[string]struct {
		url        string
		mockResp   *http.Response
		mockErr    error
		wantFoobar *foobar
		wantErr    error
	}{
		"Simple": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:    nil,
			wantFoobar: &foobar{Foo: "foo", Bar: "bar"},
			wantErr:    nil,
		},
		"InvalidURL": {
			url:        "invalid://url%%",
			mockResp:   nil,
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to create http request: parse \"invalid://url%%%%\": invalid URL escape \"%%%%\""),
		},
		"HTTPClientError": {
			url:        "http://example.org",
			mockResp:   nil,
			mockErr:    fmt.Errorf("http client error"),
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to talk to tado° API: http client error"),
		},
		"UnparseableJson": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusOK, `{notjson}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to decode tado° API response: invalid character 'n' looking for beginning of object key string"),
		},
		"EmptyErrorList": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: API returned empty error"),
		},
		"SingleError": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"}]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: 1: One"),
		},
		"MultiError": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"},{"code":"2","title":"Two"}]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: 1: One, 2: Two"),
		},
		"UnparseableError": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusInternalServerError, `{errorjson}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: unable to decode API error: invalid character 'e' looking for beginning of object key string"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := NewClient("test", "test")
			client.http = mockHTTPClient{Response: tc.mockResp, Error: tc.mockErr}

			result := &foobar{}
			err := client.get(tc.url, result)

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Exactly(t, tc.wantFoobar, result)
			}

		})
	}
}

func TestPost(t *testing.T) {
	tests := map[string]struct {
		url      string
		mockResp *http.Response
		mockErr  error
		wantErr  error
	}{
		"Simple": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusNoContent, ""),
			mockErr:  nil,
			wantErr:  nil,
		},
		"InvalidURL": {
			url:      "invalid://url%%",
			mockResp: nil,
			mockErr:  nil,
			wantErr:  fmt.Errorf("unable to create http request: parse \"invalid://url%%%%\": invalid URL escape \"%%%%\""),
		},
		"HTTPClientError": {
			url:      "http://example.org",
			mockResp: nil,
			mockErr:  fmt.Errorf("http client error"),
			wantErr:  fmt.Errorf("unable to talk to tado° API: http client error"),
		},
		"UnexepctedResponseCode": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("unexpected tado° API response status: OK"),
		},
		"EmptyErrorList": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: API returned empty error"),
		},
		"SingleError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"}]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: 1: One"),
		},
		"MultiError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"},{"code":"2","title":"Two"}]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: 1: One, 2: Two"),
		},
		"UnparseableError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{errorjson}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: unable to decode API error: invalid character 'e' looking for beginning of object key string"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := NewClient("test", "test")
			client.http = mockHTTPClient{Response: tc.mockResp, Error: tc.mockErr}

			err := client.post(tc.url)

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestPut(t *testing.T) {
	type foobar struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}

	tests := map[string]struct {
		url        string
		data       interface{}
		mockResp   *http.Response
		mockErr    error
		wantFoobar interface{}
		wantErr    error
	}{
		"SimpleValue": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:    nil,
			wantFoobar: foobar{Foo: "foo", Bar: "bar"},
			wantErr:    nil,
		},
		"SimpleValuePtr": {
			url:        "http://example.org",
			data:       &foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:    nil,
			wantFoobar: &foobar{Foo: "foo", Bar: "bar"},
			wantErr:    nil,
		},
		"InvalidURL": {
			url:        "invalid://url%%",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   nil,
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to create http request: parse \"invalid://url%%%%\": invalid URL escape \"%%%%\""),
		},
		"HTTPClientError": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   nil,
			mockErr:    fmt.Errorf("http client error"),
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to talk to tado° API: http client error"),
		},
		"UnparseableJson": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusOK, `{notjson}`),
			mockErr:    nil,
			wantFoobar: foobar{Foo: "foo", Bar: "bar"},
			wantErr:    nil,
		},
		"UnparseableJsonPtr": {
			url:        "http://example.org",
			data:       &foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusOK, `{notjson}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to decode tado° API response: invalid character 'n' looking for beginning of object key string"),
		},
		"EmptyErrorList": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: API returned empty error"),
		},
		"SingleError": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"}]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: 1: One"),
		},
		"MultiError": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"},{"code":"2","title":"Two"}]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: 1: One, 2: Two"),
		},
		"UnparseableError": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusInternalServerError, `{errorjson}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: unable to decode API error: invalid character 'e' looking for beginning of object key string"),
		},
		"NilData": {
			url:        "http://example.org",
			data:       nil,
			mockResp:   makeResponse(http.StatusOK, `{}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    nil,
		},
		"EmptyData": {
			url:        "http://example.org",
			data:       foobar{Foo: "", Bar: ""},
			mockResp:   makeResponse(http.StatusOK, `{}`),
			mockErr:    nil,
			wantFoobar: foobar{Foo: "", Bar: ""},
			wantErr:    nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := NewClient("test", "test")
			client.http = mockHTTPClient{Response: tc.mockResp, Error: tc.mockErr}
			data := tc.data

			err := client.put(tc.url, data)

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Exactly(t, tc.wantFoobar, data)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := map[string]struct {
		url      string
		mockResp *http.Response
		mockErr  error
		wantErr  error
	}{
		"Simple": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusNoContent, ""),
			mockErr:  nil,
			wantErr:  nil,
		},
		"UnexepctedResponseCode": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("unexpected tado° API response status: OK"),
		},
		"InvalidURL": {
			url:      "invalid://url%%",
			mockResp: nil,
			mockErr:  nil,
			wantErr:  fmt.Errorf("unable to create http request: parse \"invalid://url%%%%\": invalid URL escape \"%%%%\""),
		},
		"HTTPClientError": {
			url:      "http://example.org",
			mockResp: nil,
			mockErr:  fmt.Errorf("http client error"),
			wantErr:  fmt.Errorf("unable to talk to tado° API: http client error"),
		},
		"EmptyErrorList": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: API returned empty error"),
		},
		"SingleError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"}]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: 1: One"),
		},
		"MultiError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"},{"code":"2","title":"Two"}]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: 1: One, 2: Two"),
		},
		"UnparseableError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{errorjson}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: unable to decode API error: invalid character 'e' looking for beginning of object key string"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := NewClient("test", "test")
			client.http = mockHTTPClient{Response: tc.mockResp, Error: tc.mockErr}

			err := client.delete(tc.url)

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
