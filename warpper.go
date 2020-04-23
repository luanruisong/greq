package greq

import (
	"net/http"
	"strings"
)

type (
	Resp struct {
		Ok      bool
		Status  int
		Err     error
		Header  map[string][]string
		RawBody []byte
		Body    string
	}

	Client struct {
		c    *http.Client
		pool chan struct{}
	}

	Request struct {
		client *Client
		url    string
		header http.Header
		reader ReaderHandler
	}

	ReaderHandler func(interface{}) *strings.Reader
)
