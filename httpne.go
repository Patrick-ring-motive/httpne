package httpne

import (
	"crypto/tls"
	"fmt"
	. "github.com/Patrick-ring-motive/ione"
	utils "github.com/Patrick-ring-motive/utils"
	"io"
	"net/http"
	"strings"
)

var HttpNoNil = true
var HttpNoError = true
var HttpNoPanic = true

type HttpResponseWriter struct {
	Value *http.ResponseWriter
}

type HttpHeader struct {
	Value *http.Header
}

func (responseWriter HttpResponseWriter) Write(bytes []byte) int {
	length, err := (*responseWriter.Value).Write(bytes)
	if HttpNoPanic {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("HttpResponseWriter.Write panic: ", r)
				length = 0
			}
		}()
	}
	if (err != nil) && HttpNoError {
		fmt.Println("HttpResponseWriter.Write error: ", err.Error())
		length = 0
	}
	return length
}

func (responseWriter HttpResponseWriter) WriteHeader(statusCode int) {
	if HttpNoPanic {
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

func HttpNewRequest(method string, url string, body io.Reader) HttpRequest{
  httpReq := utils.NilOfType(func(h HttpRequest){})
  if HttpNoPanic {
    defer func() {
      if r := recover(); r != nil {
        req,err := http.NewRequest("GET","https://go.dev/blog/defer-panic-and-recover",nil)
        utils.AllowUnused(err)
        fmt.Println("HttpNewRequest recover: ", r)
        httpReq = HttpRequest{Value:req}
      }
    }()
  }
  req,erro := http.NewRequest(method,url,body)
  if (req == nil) && (erro == nil) && HttpNoNil {
    req,err := http.NewRequest("GET","https://go.dev/ref/spec#The_zero_value",nil)
    utils.AllowUnused(err)
    httpReq = HttpRequest{Value:req}
    fmt.Println("HttpNewRequest nil")
    return httpReq
  }
  if (erro != nil) && (req == nil) && HttpNoError {
    req,err := http.NewRequest("GET","https://go.dev/blog/error-handling-and-go",nil)
    utils.AllowUnused(err)
    httpReq = HttpRequest{Value:req}
    fmt.Println("HttpNewRequest error: ", erro.Error())
    return httpReq
  }
  httpReq = HttpRequest{Value: req}
  return httpReq
  
}

type HttpResponse struct {
	Value *http.Response
}

func NewHttpResponse(Status string, StatusCode int, Proto string, ProtoMajor int, ProtoMinor int, Header http.Header, Body io.ReadCloser, ContentLength int64, TransferEncoding []string, Close bool,
	Uncompressed bool, Trailer http.Header, Request *http.Request, TLS *tls.ConnectionState) HttpResponse {
	return HttpResponse{Value: &http.Response{Status: Status, StatusCode: StatusCode, Proto: Proto, ProtoMajor: ProtoMajor, ProtoMinor: ProtoMinor, Header: Header, Body: Body, ContentLength: ContentLength, TransferEncoding: TransferEncoding, Close: Close, Uncompressed: Uncompressed, Trailer: Trailer, Request: Request, TLS: TLS}}
}

func (httpRes HttpResponse) Body() IoReadCloser {
	return IoReadCloser{Value: &httpRes.Value.Body}
}

func (client HttpClient) Do(req HttpRequest) HttpResponse {
	httpRes := utils.NilOfType(func(h HttpResponse) {})
	if HttpNoPanic {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("579 HttpClient.Do panic: ", r)
				status := fmt.Sprint("579 HttpClient.Do panic: ", r)
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
			}
		}()
	}
	res, err := client.Value.Do(req.Value)
	httpRes = HttpResponse{Value: res}
	if (res == nil) && (err == nil) && HttpNoNil {
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
	if (err != nil) && ((res == nil) || (res.Body == nil)) && HttpNoError {
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
	return httpRes
}


func (client HttpClient) Get(url string) HttpResponse {
  httpRes := utils.NilOfType(func(h HttpResponse) {})
  if HttpNoPanic {
    defer func() {
      if r := recover(); r != nil {
        fmt.Println("579 HttpClient.Do panic: ", r)
        status := fmt.Sprint("579 HttpClient.Do panic: ", r)
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
      }
    }()
  }
  res, err := client.Value.Get(url)
  httpRes = HttpResponse{Value: res}
  if (res == nil) && (err == nil) && HttpNoNil {
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
  if (err != nil) && ((res == nil) || (res.Body == nil)) && HttpNoError {
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
  return httpRes
}
