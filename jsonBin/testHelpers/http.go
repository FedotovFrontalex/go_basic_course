package testHelpers

import (
	"bytes"
	"io"
	"net/http"
)

type FakeTransport struct {
	Resp *http.Response
	Err  error
	Req  *http.Request
}

func (ft *FakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ft.Req = req
	return ft.Resp, ft.Err
}

func CreateFakeTransport(respBytes []byte, err error, statusCode int) *FakeTransport {

	fakeResp := &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(bytes.NewBuffer(respBytes)),
		Header:     make(http.Header),
	}

	fakeTransport := &FakeTransport{
		Resp: fakeResp,
		Err:  err,
	}

	return fakeTransport
}
