package weather

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/zumosik/weatherapi-sdk-go/e"
)

const (
	host = "https://api.weatherapi.com/v1"

	endpointForecast = "/forecast.json"

	defaultTimeout = 5 * time.Second
)

var (
	errTokenIsEmpty = errors.New("token is empty")
	errInvalidData  = errors.New("invalid data")
)

type Client struct {
	client *http.Client
	token  string
}

func New(token string) (*Client, error) {
	if token == "" {
		return nil, errTokenIsEmpty
	}

	return &Client{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		token: token,
	}, nil
}

func (c *Client) GetForecast(q string, days int) (f ForecastResponse, err error) {
	defer func() {
		err = e.Wrap("can't get foreacst", err)
	}()

	if q == "" || days >= 10 || days <= 1 {
		return ForecastResponse{}, errInvalidData
	}
	// ?key=%s&q=%s&days=%s&aqi=%s&alerts=%s
	query := url.Values{}
	query.Add("key", c.token)
	query.Add("q", q)
	query.Add("days", strconv.Itoa(days))
	query.Add("aqi", "no")
	query.Add("alerts", "no")

	body, err := c.doRequest((host + endpointForecast), query)
	if err != nil {
		return ForecastResponse{}, err
	}

	var res ForecastResponse

	if err := json.Unmarshal(body, &res); err != nil {
		return ForecastResponse{}, err
	}

	return res, nil
}

func (c *Client) doRequest(url string, query url.Values) ([]byte, error) {
	const errorMsg = "can't do request"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, e.Wrap(errorMsg, err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, e.Wrap(errorMsg, err)

	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, e.Wrap(errorMsg, err)
	}

	return body, nil
}
