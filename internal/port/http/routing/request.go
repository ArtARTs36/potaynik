package routing

import (
	"encoding/json"
	"net/http"
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
