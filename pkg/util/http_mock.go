package util

import (
	"fmt"
	"io"
	"net/http"
)

// DefaultRequestHeader contain some common header key and value.
var DefaultRequestHeader http.Header = http.Header{
	"Accept":[]string{"application/json"}, // Accept: application/json
	"Accept-Charset":[]string{"utf-8"}, // Accept-Charset: utf-8
}

type MockClient struct {
	client http.Client
}

// Get is the GET request, be used to get resource.
// Get method has no request-body, internal method is http.MethodGet.
func (m *MockClient) Get(url string, header http.Header) (*http.Response, error) {
	return m.DoRequest(url, http.MethodGet, header, nil)
}

// Head is the HEAD request
// Head method has no request-body, internal method is http.MethodHead
func (m *MockClient) Head(url string, header http.Header) (*http.Response, error) {
	return m.DoRequest(url, http.MethodHead, header, nil)
}

// Post is the POST request, be used to create resource.
// Post method has request-body, internal method is http.MethodPost.
func (m *MockClient) Post(url string, header http.Header, body io.Reader) (*http.Response, error) {
	return m.DoRequest(url, http.MethodPost, header, body)
}

// Put is the PUT request, be used to update all value of resource.
// Put method has request-body, internal method is http.MethodPut
func (m *MockClient) Put(url string, header http.Header, body io.Reader) (*http.Response, error) {
	return m.DoRequest(url, http.MethodPut, header, body)
}

// Patch is the PATCH request, be used to patch some value of resource. Like put, but not update all values.
// Patch method has request-body, internal method is http.MethodPatch
func (m *MockClient) Patch(url string, header http.Header, body io.Reader) (*http.Response, error) {
	return m.DoRequest(url, http.MethodPatch, header, body)
}

// Delete is DELETE request, be used to delete resource.
// Delete method has no request-body, internal method is http.MethodDelete
func (m *MockClient) Delete(url string, header http.Header) (*http.Response, error) {
	return m.DoRequest(url, http.MethodDelete, header, nil)
}

// todo implement the connect method of the http-request
//func(m*MockClient)Connect(url string, head http.Header, body io.Reader)(*http.Response, error){}

// todo implement the option method of the http-request
//func (m *MockClient) Option(url string, header http.Header, body io.Reader) (*http.Response, error) {
//	return m.DoRequest(url, http.MethodOptions, header, body)
//}

// todo implement the trace method of the http-request
//func (m *MockClient) Trace(url string, header http.Header, body io.Reader) (*http.Response, error) {
//	return m.DoRequest(url, http.MethodTrace, header, body)
//}

// DoRequest use to do a request with input the method, url, header and body.
// If target method is "", it will be set to GET(http.Client#Do implement it).
// DoRequest support do the special method like GET-Request with request body.
// If the header is empty, it will set default header to it.
func (m *MockClient) DoRequest(url, method string, header http.Header, body io.Reader) (*http.Response, error) {
	r, err := http.NewRequest(method, url, body)
	if err == nil {
		return nil, err
	}
	if r == nil{
		return nil, fmt.Errorf("request struct init fail, and can't patch any error")
	}
	if header == nil {
		header = DefaultRequestHeader
	}
	r.Header = header
	return m.client.Do(r)
}
