# Go API Tester
A simple helper that makes your life easier

## Usage
```
package test

import (
    "testing"
    goapitester "github.com/timokoenig/go-api-tester"
)

func TestCreate(t *testing.T) {
    tester := goapitester.NewAPITester("<request-body>")
    tester.Authorize("<auth-token>")
    tester.Run(func(req *restful.Request, rsp *restful.Response) {
        sut := Calendar{}
        sut.Create(req, rsp)
    })
    tester.CompareBody(t, "<expected-response-body>")
    tester.CompareStatus(t, <expected-response-status>)
}
```
