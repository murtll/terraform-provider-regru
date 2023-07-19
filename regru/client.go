package regru

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const defaultBaseURL = "https://api.reg.ru/api/regru2/"

type Client struct {
	username string
	password string

	baseURL    *url.URL
	HTTPClient *http.Client
}

func NewClient(username, password, apiEndpoint string) *Client {
	if apiEndpoint == "" {
		apiEndpoint = defaultBaseURL
	}

	baseURL, _ := url.Parse(apiEndpoint)

	return &Client{
		username:   username,
		password:   password,
		baseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
	}
}

func (c Client) doRequest(request any, fragments ...string) (*APIResponse, error) {
	endpoint := c.baseURL.JoinPath(fragments...)

	inputData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to create input data: %w", err)
	}

	query := endpoint.Query()
	query.Add("input_data", string(inputData))
	query.Add("input_format", "json")
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode/100 != 2 {
		return nil, parseError(req, resp)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp APIResponse
	err = json.Unmarshal(raw, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}

func parseError(req *http.Request, resp *http.Response) error {
	raw, _ := io.ReadAll(resp.Body)

	var errAPI APIResponse
	err := json.Unmarshal(raw, &errAPI)
	if err != nil {
		return err
	}

	return fmt.Errorf("status code: %d, %w", resp.StatusCode, errAPI)
}
