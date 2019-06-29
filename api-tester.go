package goapitester

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emicklei/go-restful"
)

// NewAPITester creates a new helper that makes your life easier
func NewAPITester(requestBody string) *APITester {
	req, err := http.NewRequest("GET", "http://localhost:8080/resource", strings.NewReader(requestBody))
	if err != nil {
		panic(err)
	}
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	rec := httptest.NewRecorder()
	httpReq := restful.NewRequest(req)
	httpRsp := restful.NewResponse(rec)
	httpRsp.PrettyPrint(false)
	return &APITester{
		req:     req,
		rec:     rec,
		httpReq: httpReq,
		httpRsp: httpRsp,
	}
}

// APITester helps you testing go-micro APIs
type APITester struct {
	req     *http.Request
	rsp     *http.Response
	rec     *httptest.ResponseRecorder
	httpReq *restful.Request
	httpRsp *restful.Response
}

// Authorize sets the given bearer authorization token for the current request
func (a *APITester) Authorize(token string) {
	a.req.Header.Add("Authorization", "Bearer "+token)
}

// Run the requset with the given handler
func (a *APITester) Run(handler func(req *restful.Request, rsp *restful.Response)) {
	handler(a.httpReq, a.httpRsp)
	a.rsp = a.rec.Result()
}

// CompareBody compares the exepcted body with the response body
func (a *APITester) CompareBody(t *testing.T, expected string) {
	b, err := ioutil.ReadAll(a.rsp.Body)
	if err != nil {
		t.Fatalf("can not read response: %v", err)
	}
	a.rsp.Body.Close()
	body := string(bytes.TrimSpace(b))
	if body != expected {
		t.Fatalf("expected body %s; got %s", expected, body)
	}
}

// CompareStatus compares the expected status with the response status
func (a *APITester) CompareStatus(t *testing.T, extected int) {
	if a.rsp.StatusCode != extected {
		t.Fatalf("expected status %v; got %v", extected, a.rsp.Status)
	}
}
