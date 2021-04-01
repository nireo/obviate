package obviate

import (
	"context"
	"log"
	"net/http"
	"time"
)

type key int
const (
	ContextOrgPath key = iota
	ContextReqStart
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextReqStart, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LogRequest(r *http.Request, code int) {
	ctx := r.Context()
	v := ctx.Value(ContextOrgPath)
	path, ok := v.(string)
	if !ok {
		path = r.URL.Path
	}

	v = ctx.Value(ContextReqStart)
	if v == nil {
		return
	}

	if s, ok := v.(time.Time); ok {
		log.Println(time.Since(s), code, r.Method, path)
	}
}