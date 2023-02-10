package testutils

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

// ClientTransport implements ClientTransport from github.com/go-openapi/runtime
// it abstracts the request/client operation handling.
type TestClientTransport struct {
	Server  *httptest.Server
	Handler http.Handler
}

func NewTestClientTransport(handler http.Handler) *TestClientTransport {
	testServer := httptest.NewServer(handler)
	return &TestClientTransport{Server: testServer, Handler: handler}
}

// Submit method handles a client operation.
// client operation is transformed into an http.Request and submitted against the test server
// response is parsed the same way as the generated client does.
func (t *TestClientTransport) Submit(op *runtime.ClientOperation) (v interface{}, err error) {
	recorder := httptest.NewRecorder()
	req, err := client.New(t.Server.URL, "/", nil).CreateHttpRequest(op)
	if err != nil {
		return
	}
	t.Handler.ServeHTTP(recorder, req)
	v, err = op.Reader.ReadResponse(response{recorder: recorder}, runtime.JSONConsumer())
	return
}

// response implements ClientResponse from github.com/go-openapi/runtime
type response struct {
	recorder *httptest.ResponseRecorder
}

func (r response) Code() int {
	return r.recorder.Code
}

func (r response) Message() string {
	return strconv.Itoa(r.recorder.Code) // TODO: change to status text
}

func (r response) GetHeader(name string) string {
	return r.recorder.HeaderMap.Get(name)
}

func (r response) GetHeaders(name string) []string {
	return r.recorder.HeaderMap.Values(name)
}

func (r response) Body() io.ReadCloser {
	return ioutil.NopCloser(r.recorder.Body)
}
