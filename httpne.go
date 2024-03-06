package httpne

import (
  "crypto/tls"
  "fmt"
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

type HttpRequest struct {
  Value *http.Request
}

type HttpResponse struct {
  Value *http.Response
}

func (client HttpClient) Do(req HttpRequest) HttpResponse {
  httpRes := utils.NilOfType(func(h HttpResponse) {})
  if HttpNoPanic {
    defer func() {
      if r := recover(); r != nil {
        fmt.Println("579 HttpClient.Do panic: ", r)
        status := fmt.Sprint("579 HttpClient.Do panic: ", r)
        body := io.NopCloser(strings.NewReader(status))
        httpRes = HttpResponse{Value: &http.Response{Status: status, StatusCode: 579, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: req.Value.Header, Body: body, ContentLength: -1, TransferEncoding: nil, Close: true, Uncompressed: true, Trailer: req.Value.Trailer, Request: req.Value, TLS: utils.NilOfType(func(t *tls.ConnectionState) {})}}
      }
    }()
  }
  res, err := client.Value.Do(req.Value)
  httpRes = HttpResponse{Value: res}
  if (res == nil) && (err == nil) && HttpNoNil {
    fmt.Println("559 HttpClient.Do nil")
    status := "559 HttpClient.Do nil"
    body := io.NopCloser(strings.NewReader(status))
    httpRes = HttpResponse{Value: &http.Response{Status: status, StatusCode: 559, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: req.Value.Header, Body: body, ContentLength: -1, TransferEncoding: nil, Close: true, Uncompressed: true, Trailer: req.Value.Trailer, Request: req.Value, TLS: utils.NilOfType(func(t *tls.ConnectionState) {})}}
  }
  if (err != nil) && HttpNoError {
    fmt.Println("569 HttpClient.Do error: ", err.Error())
    status := fmt.Sprint("569 HttpClient.Do error: ", err.Error())
    body := io.NopCloser(strings.NewReader(status))
    httpRes = HttpResponse{Value: &http.Response{Status: status, StatusCode: 569, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: req.Value.Header, Body: body, ContentLength: -1, TransferEncoding: nil, Close: true, Uncompressed: true, Trailer: req.Value.Trailer, Request: req.Value, TLS: utils.NilOfType(func(t *tls.ConnectionState) {})}}
  }
  return httpRes
}
