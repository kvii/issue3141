package data

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/encoding/form"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	vi "github.com/kvii/issue3141/api/helloworld/v1"
)

func TestHttpClient(t *testing.T) {
	// Execute `kratos run` first.

	ctx := context.Background()
	client, _ := http.NewClient(ctx, http.WithEndpoint("127.0.0.1:8000"))

	c := vi.NewMyServiceHTTPClient(client)
	_, err := c.Call(ctx, &vi.Request{
		Id: "id",
		Msg: []*vi.MsgInline{
			{},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	// error: code = 404 reason =  message =  metadata = map[] cause = proto: syntax error (line 1:1): unexpected token 404
}

func TestBind(t *testing.T) {
	// Codes in MyServiceHTTPClientImpl.Call:
	//
	// 	pattern := "/my-api/{id}"
	//	path := binding.EncodeURL(pattern, in, false)

	in := &vi.Request{
		Id: "id",
		Msg: []*vi.MsgInline{
			{},
		},
	}
	pattern := "/my-api/{id}"
	path := binding.EncodeURL(pattern, in, false)

	want := "/my-api/id"
	if path != want {
		t.Fatalf("path != want, path: %q want: %q", path, want)
	}
}

func TestEncode(t *testing.T) {
	// Codes in binding.EncodeURL:
	//	queryParams, _ := form.EncodeValues(msg) // <- Note here
	//	pathParams := make(map[string]struct{})
	//	path := reg.ReplaceAllStringFunc(pathTemplate, func(in string) string {
	//		key := in[1 : len(in)-1]
	//		pathParams[key] = struct{}{}
	//		return queryParams.Get(key) // <- Will get nothing is `form.EncodeValues` returns error
	//	})
	//	...
	// return path

	msg := &vi.Request{
		Id: "id",
		Msg: []*vi.MsgInline{
			{},
		},
	}
	_, err := form.EncodeValues(msg)
	if err != nil {
		t.Fatal(err)
	}
	// unsupported message type: "helloworld.v1.MsgInline"
}
