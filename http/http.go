package http

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	ErrServerError = errors.New("http code not 200")
	readTimeout    = 30 * time.Second
	cliPoolSize    = 10

	tr http.RoundTripper = &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	version string
)

type Config struct {
}

func SetVersion(v string) {
	version = v
}

func SetRoundTripper(rt http.RoundTripper) {
	tr = rt
}

func getTransport() http.RoundTripper {
	return tr
}

func Init(timeout int64, poolSize int) error {
	if timeout > 0 {
		readTimeout = time.Duration(timeout) * time.Second
	}
	if poolSize > 0 {
		cliPoolSize = poolSize
	}
	tr = &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DisableCompression:  true,
		MaxIdleConnsPerHost: cliPoolSize,
	}
	return nil
}

func GetDataWithHeader(url string, params url.Values, headers map[string]string) ([]byte, *http.Response, error) {
	cli := &http.Client{
		Timeout:   readTimeout,
		Transport: getTransport(),
	}
	buf := bytes.NewBufferString(url)
	if params != nil {
		buf.WriteString("?")
		buf.WriteString(params.Encode())
	}
	req, err := http.NewRequest(http.MethodGet, buf.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("User-Agent", "Golang SDK v"+version)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}

func PostDataWithHeader(url string, params url.Values, headers map[string]string) ([]byte, *http.Response, error) {
	cli := &http.Client{
		Timeout:   readTimeout,
		Transport: getTransport(),
	}
	var ioParams io.Reader
	if params != nil {
		ioParams = bytes.NewReader([]byte(params.Encode()))
	}
	req, err := http.NewRequest(http.MethodPost, url, ioParams)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Golang SDK v"+version)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return buff, resp, nil
}
