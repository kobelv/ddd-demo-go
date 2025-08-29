package httpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/avast/retry-go/v4"
)

// todo 需要设置各种超时参数
type MyHttpClient struct {
	Timeout int
	Retry   int
}

func NewMyHttpClient() *MyHttpClient {
	return &MyHttpClient{}
}

func (f *MyHttpClient) makeHTTPRequest(ctx context.Context,
	client http.Client,
	reqURL string,
	isPost bool,
	params []byte,
	headers http.Header) (string, error) {
	var result []byte
	if headers == nil {
		headers = http.Header{}
		headers.Set("Content-Type", "application/json; charset=utf-8")
	}

	method := "GET"
	if isPost {
		method = "POST"
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, bytes.NewReader([]byte(params)))
	if err != nil {
		return "", err
	}

	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err = io.ReadAll(resp.Body)
	return string(result), nil
}

func (f *MyHttpClient) SendRequest(ctx context.Context,
	client http.Client,
	reqURL string,
	isPost bool,
	params []byte,
	headers http.Header) (string, error) {
	var result string
	err := retry.Do(
		func() error {
			resp, err := f.makeHTTPRequest(ctx, client, reqURL, isPost, params, headers)
			if err != nil {
				return err
			}
			result = resp
			return nil
		},
		retry.Attempts(uint(f.Retry)),
	)
	if err != nil {
		return "", err
	}
	return result, nil
}
