package routing

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hetiansu5/urlquery"
)

type Request struct {
	request *http.Request
}

func NewRequest(req *http.Request) Request {
	return Request{request: req}
}

func (r *Request) DecodeBody(str interface{}) error {
	return json.NewDecoder(r.request.Body).Decode(&str)
}

func (r *Request) URI() string {
	return r.request.RequestURI
}

func (r *Request) GetQueryParam(key string) string {
	return r.request.URL.Query().Get(key)
}

func (r *Request) DecodeQuery(str interface{}) error {
	return urlquery.Unmarshal([]byte(r.request.URL.RawQuery), &str)
}

func (r *Request) Context() context.Context {
	return r.request.Context()
}

func (r *Request) GetUserIP() string {
	return r.request.RemoteAddr
}
