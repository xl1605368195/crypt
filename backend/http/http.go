package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/xl1605368195/crypt/backend"
)

type Client struct {
	Hosts []string
}

func New(machines []string) (*Client, error) {
	hosts := make([]string, len(machines))
	for i, host := range machines {
		hosts[i] = host
	}
	return &Client{hosts}, nil
}

func (c *Client) Get(key string) ([]byte, error) {
	var body []byte = nil
	var err error = nil
	for _, host := range c.Hosts {
		body, err = ExecHttp(host, key)
		if err != nil {
			continue
		}
		break
	}
	return body, err
}

// 不需要list
func (c *Client) List(key string) (backend.KVPairs, error) {
	return nil, nil
}

// 不需要set
func (c *Client) Set(key string, value []byte) error {
	return nil
}

func (c *Client) Watch(key string, stop chan bool) <-chan *backend.Response {
	respChan := make(chan *backend.Response, 0)
	go func() {
		for {
			var body []byte = nil
			var err error = nil
			for _, conn := range c.Hosts {
				body, err = ExecHttp(conn, key)
				if err != nil {
					continue
				}
				break
			}

			// 读取失败
			if err != nil {
				respChan <- &backend.Response{nil, err}
				time.Sleep(time.Second * 5)
				continue
			}
			respChan <- &backend.Response{body, nil}
		}
	}()
	return respChan
}

func ExecHttp(url, key string) ([]byte, error) {
	payload := strings.NewReader(fmt.Sprintf(`{"key":"%s"}`, key))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
