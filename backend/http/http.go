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
	Connections []Connection
}

type Connection struct {
	Host     string
	Http     *http.Client
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
		body, err = ExecHttp(conn, conn.Host, key)
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
				body, err = ExecHttp(conn, conn.Host, key)
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

func ExecHttp(conn Connection, url string, key string) ([]byte, error) {
	payload := strings.NewReader(fmt.Sprintf(`{"key":"%s"}`, key))
	res, err := conn.Http.Post(url, "application/json", payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
