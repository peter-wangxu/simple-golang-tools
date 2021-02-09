package http

import (
	"bytes"
	"net/http"
)

type WrappedResponseWriter struct {
	http.ResponseWriter
	Buf  *bytes.Buffer
	Code *int
}

// TODO(peter) what if there are data/state already in the ResponseWriter?

func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	wrw := &WrappedResponseWriter{
		ResponseWriter: w,
		Buf:            &bytes.Buffer{},
		Code:           new(int),
	}
	*wrw.Code = 200
	////// Try to read the existing data
	//if c, ok := w.(*WrappedResponseWriter); ok {
	//	wrw.Buf.Write(c.Get())
	//	wrw.Code = c.Code
	//}
	return wrw
}

func (wrw *WrappedResponseWriter) Write(p []byte) (int, error) {
	wrw.ResponseWriter.Write(p)
	return wrw.Buf.Write(p)
}

func (wrw *WrappedResponseWriter) WriteHeader(code int) {
	*wrw.Code = code
	wrw.ResponseWriter.WriteHeader(code)
}

// Get returns all the written bytes, this make it
// possible to chain WrappedResponseWriter otherwise
// we lose the bytes written already
func (wrw *WrappedResponseWriter) Get() []byte {
	return wrw.Buf.Bytes()
}
