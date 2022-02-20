package http

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/d7561985/tel/v2"
	"github.com/d7561985/tel/v2/monitoring/metrics"
	"github.com/stretchr/testify/assert"
)

// want to checl if response will write to log via mw
const testString = "Hello World"

// want to check if it write to log via mw
const postContent = "XXX"

func TestTelemetry_HttpServerMiddlewareAll(t *testing.T) {
	c := tel.DefaultDebugConfig()
	c.LogLevel = "debug"
	srv, closer := tel.New(context.Background(), c)
	defer closer()

	ctx := srv.Ctx()
	buf := tel.SetLogOutput(tel.FromCtx(ctx))

	// key value helps check if our middleware not damage already existent context with own values
	type key struct{}

	handler := http.Handler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(testString))

		assert.NotNil(t, request.Context().Value(key{}))
		fmt.Println(request.Context().Value(key{}))
	}))

	m := metrics.NewHTTPMetric(metrics.DefaultHTTPPathRetriever())
	assert.NoError(t, m.SetUp())

	handler = NewServer(tel.FromCtx(ctx)).HTTPServerMiddlewareAll(m)(handler)

	// check context preservation already added fields
	handler = func(next http.Handler) http.Handler {
		fn := func(writer http.ResponseWriter, request *http.Request) {
			ctx := context.WithValue(request.Context(), key{}, "*****")
			next.ServeHTTP(writer, request.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}(handler)

	s := httptest.NewServer(handler)

	res, err := s.Client().Post(s.URL+"/", "", bytes.NewBufferString(postContent))
	assert.NoError(t, err)

	v, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(v), testString)
	assert.Contains(t, buf.String(), testString)
	assert.Contains(t, buf.String(), postContent)
	fmt.Println(buf.String())
}