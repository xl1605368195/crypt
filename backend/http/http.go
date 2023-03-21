package http

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/xl1605368195/crypt/backend"
)

type Client struct {
	Connections []Connection
}

type Connection struct {
	Host     string
	Http     *http.Client
	proxyURL *url.URL
}

func New(machines []string) (*Client, error) {
	connctions := make([]Connection, len(machines))
	for i, host := range machines {
		conn := &Connection{
			Host: host,
		}
		connctions[i] = *conn
	}
	return &Client{connctions}, nil
}

func (c *Client) Get(key string) ([]byte, error) {
	var body []byte = nil
	var err error = nil
	for _, conn := range c.Connections {
		resp, err := conn.Http.Post(conn.Host, "text/plain", strings.NewReader(key))
		if err != nil {
			continue
		}
		body, err = io.ReadAll(resp.Body)
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
			for _, conn := range c.Connections {
				resp, err := conn.Http.Post(conn.Host, "text/plain", strings.NewReader(key))
				if err != nil {
					continue
				}
				body, err = io.ReadAll(resp.Body)
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
