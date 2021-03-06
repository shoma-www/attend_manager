package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/core"
)

// Middleware routeのmiddleware
type Middleware struct {
	logger core.Logger
}

// NewMiddleware コンストラクタ
func NewMiddleware(l core.Logger) *Middleware {
	return &Middleware{
		logger: l,
	}
}

// AddUUIDWithContext Request のContextにUUIDを追加
func (m *Middleware) AddUUIDWithContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := xid.New().String()
		ctx := context.WithValue(context.Background(), core.UUIDContextKey, uuid)
		rctx := r.WithContext(ctx)
		next.ServeHTTP(w, rctx)
	})
}

// Logger Handlerの前後でログ出力を行う
func (m *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := m.logger.SetUUID(xid.New().String())

		ctx := core.SetLogger(r.Context(), l)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			l.Error("Failed read from r.body. err=%s", err.Error())
			next.ServeHTTP(w, r)
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		requestMap := map[string]string{
			"Host":     r.Host,
			"Method":   r.Method,
			"URI":      r.RequestURI,
			"Protocol": r.Proto,
			"UA":       r.UserAgent(),
			"Addr":     r.RemoteAddr,
			"Ref":      r.Referer(),
		}
		bs, err := json.Marshal(requestMap)
		if err != nil {
			l.Error("Failed read from request. err=%s", err.Error())
			next.ServeHTTP(w, r)
			return
		}

		l.Info("[Request] %s %s %s, json: %s", r.Method, r.RequestURI, r.Proto, string(bs))
		l.Debug("[Request] Body: %s", string(body))

		var sb strings.Builder
		wres := NewRWWrapper(w, &sb)

		next.ServeHTTP(wres, r.WithContext(ctx))

		l.Info("[Response] %s", sb.String())
	})
}

// http.ResponseWriterを実装した構造体
// Responseの書き込みを指定のio.Writerインターフェースを実装した
// 構造体に複製する
type rwWrapper struct {
	rw http.ResponseWriter
	w  io.Writer
}

// NewRWWrapper コンストラクタ
func NewRWWrapper(rw http.ResponseWriter, w io.Writer) http.ResponseWriter {
	mw := io.MultiWriter(rw, w)
	return &rwWrapper{
		rw: rw,
		w:  mw,
	}
}

func (rww *rwWrapper) Header() http.Header {
	return rww.rw.Header()
}

func (rww *rwWrapper) Write(p []byte) (int, error) {
	return rww.w.Write(p)
}

func (rww *rwWrapper) WriteHeader(statusCode int) {
	rww.rw.WriteHeader(statusCode)
}

// CORSMiddleware CORSヘッダーのミドルウェア
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, req)
	})
}
