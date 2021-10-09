package gotado

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	oauth2int "github.com/gonzolino/gotado/internal/oauth2"
	"golang.org/x/oauth2"
)

const (
	authURL  = "https://auth.tado.com/oauth/authorize"
	tokenURL = "https://auth.tado.com/oauth/token"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client to access the tado° API
type Client struct {
	// ClientID specifies the client ID to use for authentication
	ClientID string
	// ClientSecret specifies the client secret to use for authentication
	ClientSecret string

	http HTTPClient
}

// NewClient creates a new tado° client
func NewClient(clientID, clientSecret string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		http:         http.DefaultClient,
	}
}

// WithHTTPClient configures the http client to use for tado° API interactions
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.http = httpClient
	return c
}

// WithCredentials sets the given credentials and scopes for the tado° API
func (c *Client) WithCredentials(ctx context.Context, username, password string) (*Client, error) {
	config := oauth2int.NewConfig(c.ClientID, c.ClientSecret, authURL, tokenURL, []string{"home.user"})

	httpContext := context.WithValue(ctx, oauth2.HTTPClient, c.http)
	token, err := config.PasswordCredentialsToken(httpContext, username, password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials: %w", err)
	}
	authClient := config.Client(httpContext, token)

	c.http = authClient

	return c, nil
}

// Do sends the given HTTP request to the tado° API.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to talk to tado° API: %w", err)
	}
	return resp, nil
}

// Request performs an HTTP request to the tado° API
func (c *Client) Request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("unable to create http request: %w", err)
	}
	return c.Do(req)
}

// get retrieves an object from the tado° API.
func (c *Client) get(url string, v interface{}) error {
	resp, err := c.Request(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := isError(resp); err != nil {
		return fmt.Errorf("tado° API error: %w", err)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("unable to decode tado° API response: %w", err)
	}

	return nil
}
