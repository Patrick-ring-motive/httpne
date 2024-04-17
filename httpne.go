package httpne

import (
	"crypto/tls"
	"fmt"
	. "github.com/Patrick-ring-motive/ione"
	. "github.com/Patrick-ring-motive/traigo"
	utils "github.com/Patrick-ring-motive/utils"
	"io"
	"net/http"
	"strings"
)

type HttpOptions struct {
	HttpNoNil   bool
	HttpNoError bool
	HttpNoPanic bool
}

type HttpResponseWriter struct {
	Value *http.ResponseWriter
}

type HttpHeader struct {
	Value *http.Header
}

func (responseWriter HttpResponseWriter) Write(bytes []byte, options ...HttpOptions) int {
	httpNoError := true
	httpNoPanic := true
	if len(options) > 0 {
		httpNoError = options[0].HttpNoError
		httpNoPanic = options[0].HttpNoPanic
	}
	var length = 0
	var err = error(nil)
	Try(func() {
		length, err = (*responseWriter.Value).Write(bytes)
		if (err != nil) && httpNoError {
			fmt.Println("HttpResponseWriter.Write error: ", err.Error())
			length = 0
		}
	}, Catch(func(e interface{}) {
		if httpNoPanic {
			fmt.Println("HttpResponseWriter.Write panic: ", e)
			length = 0
		} else {
			panic(e)
		}
	}))
	return length
}

func (responseWriter HttpResponseWriter) WriteHeader(statusCode int, options ...HttpOptions) {
	httpNoPanic := true
	if len(options) > 0 {
		httpNoPanic = options[0].HttpNoPanic
	}
	if httpNoPanic {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("HttpResponseWriter.WriteHeader panic: ", r)
			}
		}()
	}
	(*responseWriter.Value).WriteHeader(statusCode)
}

func (responseWriter HttpResponseWriter) Header() HttpHeader {
	return HttpHeader{Value: utils.Ptr((*responseWriter.Value).Header())}
}

type HttpClient struct {
	Value *http.Client
}

var DefaultHttpClient = HttpClient{Value: &http.Client{}}

type HttpRequest struct {
	Value *http.Request
}

func HttpNewRequest(method string, url string, body io.Reader, options ...HttpOptions) HttpRequest {
	httpNoNil := true
	httpNoError := true
	httpNoPanic := true
	if len(options) > 0 {
		httpNoNil = options[0].HttpNoNil
		httpNoError = options[0].HttpNoError
		httpNoPanic = options[0].HttpNoPanic
	}
	httpReq := utils.NilOfType(func(h HttpRequest) {})
	Try(func() {
		req, erro := http.NewRequest(method, url, body)
		httpReq = HttpRequest{Value: req}
		if (req == nil) && (erro == nil) && httpNoNil {
			req, err := http.NewRequest("GET", "https://go.dev/ref/spec#The_zero_value", nil)
			utils.AllowUnused(err)
			httpReq = HttpRequest{Value: req}
			fmt.Println("HttpNewRequest nil")
			return
		}
		if (erro != nil) && (req == nil) && httpNoError {
			req, err := http.NewRequest("GET", "https://go.dev/blog/error-handling-and-go", nil)
			utils.AllowUnused(err)
			httpReq = HttpRequest{Value: req}
			fmt.Println("HttpNewRequest error: ", erro.Error())
			return
		}
	}, Catch(func(e interface{}) {
		if httpNoPanic {
			req, err := http.NewRequest("GET", "https://go.dev/blog/defer-panic-and-recover", nil)
			utils.AllowUnused(err)
			fmt.Println("HttpNewRequest recover: ", e)
			httpReq = HttpRequest{Value: req}
		} else {
			panic(e)
		}
	}))
	return httpReq

}

type HttpResponse struct {
	Value *http.Response
}

func NewHttpResponse(Status string, StatusCode int, Proto string, ProtoMajor int, ProtoMinor int, Header http.Header, Body io.ReadCloser, ContentLength int64, TransferEncoding []string, Close bool,
	Uncompressed bool, Trailer http.Header, Request *http.Request, TLS *tls.ConnectionState) HttpResponse {
	return HttpResponse{Value: &http.Response{Status: Status, StatusCode: StatusCode, Proto: Proto, ProtoMajor: ProtoMajor, ProtoMinor: ProtoMinor, Header: Header, Body: Body, ContentLength: ContentLength, TransferEncoding: TransferEncoding, Close: Close, Uncompressed: Uncompressed, Trailer: Trailer, Request: Request, TLS: TLS}}
}

func (httpRes HttpResponse) Body(ioReadCloser ...io.ReadCloser) IoReadCloser {
	if len(ioReadCloser) > 0 {
		httpRes.Value.Body = ioReadCloser[0]
	}
	return IoReadCloser{Value: httpRes.Value.Body}
}

func (client HttpClient) Do(req HttpRequest, options ...HttpOptions) HttpResponse {
	httpNoNil := true
	httpNoError := true
	httpNoPanic := true
	if len(options) > 0 {
		httpNoNil = options[0].HttpNoNil
		httpNoError = options[0].HttpNoError
		httpNoPanic = options[0].HttpNoPanic
	}
	httpRes := utils.NilOfType(func(h HttpResponse) {})
	Try(func() {
		res, err := client.Value.Do(req.Value)
		httpRes = HttpResponse{Value: res}
		if (res == nil) && (err == nil) && httpNoNil {
			fmt.Println("559 HttpClient.Do nil")
			status := "559 HttpClient.Do nil"
			body := io.NopCloser(strings.NewReader(status))
			httpRes = NewHttpResponse(
				status,
				559,
				"HTTP/1.0",
				1,
				0,
				req.Value.Header,
				body,
				-1,
				nil,
				true,
				true,
				req.Value.Trailer,
				req.Value,
				utils.NilOfType(func(t *tls.ConnectionState) {}),
			)
		}
		if (err != nil) && ((res == nil) || (res.Body == nil)) && httpNoError {
			fmt.Println("569 HttpClient.Do error: ", err.Error())
			status := fmt.Sprint("569 HttpClient.Do error: ", err.Error())
			body := io.NopCloser(strings.NewReader(status))
			httpRes = NewHttpResponse(
				status,
				569,
				"HTTP/1.0",
				1,
				0,
				req.Value.Header,
				body,
				-1,
				nil,
				true,
				true,
				req.Value.Trailer,
				req.Value,
				utils.NilOfType(func(t *tls.ConnectionState) {}),
			)
		}
	}, Catch(func(e interface{}) {
		if httpNoPanic {
			fmt.Println("579 HttpClient.Do panic: ", e)
			status := fmt.Sprint("579 HttpClient.Do panic: ", e)
			body := io.NopCloser(strings.NewReader(status))
			httpRes = NewHttpResponse(
				status,
				579,
				"HTTP/1.0",
				1,
				0,
				req.Value.Header,
				body,
				-1,
				nil,
				true,
				true,
				req.Value.Trailer,
				req.Value,
				utils.NilOfType(func(t *tls.ConnectionState) {}),
			)
		} else {
			panic(e)
		}
	}))
	return httpRes
}

func (client HttpClient) Get(url string, options ...HttpOptions) HttpResponse {
	httpNoNil := true
	httpNoError := true
	httpNoPanic := true
	if len(options) > 0 {
		httpNoNil = options[0].HttpNoNil
		httpNoError = options[0].HttpNoError
		httpNoPanic = options[0].HttpNoPanic
	}
	httpRes := utils.NilOfType(func(h HttpResponse) {})
	Try(func() {
		res, err := client.Value.Get(url)
		httpRes = HttpResponse{Value: res}
		if (res == nil) && (err == nil) && httpNoNil {
			fmt.Println("559 HttpClient.Do nil")
			status := "559 HttpClient.Do nil"
			body := io.NopCloser(strings.NewReader(status))
			httpRes = NewHttpResponse(
				status,
				559,
				"HTTP/1.0",
				1,
				0,
				make(http.Header),
				body,
				-1,
				nil,
				true,
				true,
				make(http.Header),
				utils.NilOfType(func(t *http.Request) {}),
				utils.NilOfType(func(t *tls.ConnectionState) {}),
			)
		}
		if (err != nil) && ((res == nil) || (res.Body == nil)) && httpNoError {
			fmt.Println("569 HttpClient.Do error: ", err.Error())
			status := fmt.Sprint("569 HttpClient.Do error: ", err.Error())
			body := io.NopCloser(strings.NewReader(status))
			httpRes = NewHttpResponse(
				status,
				569,
				"HTTP/1.0",
				1,
				0,
				make(http.Header),
				body,
				-1,
				nil,
				true,
				true,
				make(http.Header),
				utils.NilOfType(func(t *http.Request) {}),
				utils.NilOfType(func(t *tls.ConnectionState) {}),
			)
		}
	}, Catch(func(e interface{}) {
		if httpNoPanic {
			fmt.Println("579 HttpClient.Do panic: ", e)
			status := fmt.Sprint("579 HttpClient.Do panic: ", e)
			body := io.NopCloser(strings.NewReader(status))
			httpRes = NewHttpResponse(
				status,
				579,
				"HTTP/1.0",
				1,
				0,
				make(http.Header),
				body,
				-1,
				nil,
				true,
				true,
				make(http.Header),
				utils.NilOfType(func(t *http.Request) {}),
				utils.NilOfType(func(t *tls.ConnectionState) {}),
			)
		} else {
			panic(e)
		}
	}))
	return httpRes
}
