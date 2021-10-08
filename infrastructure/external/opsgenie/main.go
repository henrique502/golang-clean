package opsgenie

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	BaseURL       *url.URL
	client        *http.Client
	UserAgent     string
	Authorization string
}

func New(httpClient *http.Client) *Client {
	host := os.Getenv("OPSGENIE_ENDPOINT")
	token := os.Getenv("OPSGENIE_TOKEN")
	baseURL, _ := url.Parse(host)

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{
		BaseURL:       baseURL,
		client:        httpClient,
		UserAgent:     "Pagar.me",
		Authorization: "GenieKey " + token,
	}
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	log.Infof("HTTP %s %s", method, urlStr)
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", c.Authorization)

	return req, nil
}

func (c *Client) GetList(endpoint string, nextURL *string, query map[string]string, v interface{}) error {
	if nextURL != nil && len(*nextURL) > 0 {
		endpoint = *nextURL
	} else {
		u, err := url.Parse(endpoint)
		if err != nil {
			return errors.WithMessage(err, "GetList cannot parse endpoint")
		}

		q := u.Query()
		q.Add("limit", "100")
		q.Add("offset", "0")

		for k, v := range query {
			q.Set(k, v)
		}

		u.RawQuery = q.Encode()
		endpoint = u.String()
	}

	request, err := c.NewRequest("GET", endpoint, nil)
	if err != nil {
		return errors.WithMessage(err, "GetList cannot create request")
	}

	response, err := c.client.Do(request)
	if err != nil {
		return errors.WithMessage(err, "GetList cannot do request")
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(v)
	if err != nil {
		return errors.WithMessage(err, "GetList cannot decode the body")
	}

	return nil
}
