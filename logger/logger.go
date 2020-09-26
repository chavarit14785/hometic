package logger

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type logkey string

const key logkey = "logger"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := zap.NewExample()
		l = l.With(zap.Namespace("hometic"), zap.String("I'm", "gopher"))
		l.Info("Middleware Start")
		log := context.WithValue(r.Context(), "logger", l)
		next.ServeHTTP(w, r.WithContext(log))
		l.Info("Middleware Stop")
	})
}

func Get(ctx context.Context) *zap.Logger {
	val := ctx.Value(key)
	if val == nil {
		return zap.NewExample()
	}
	if l, ok := val.(*zap.Logger); ok {
		return l
	}
	return zap.NewExample()
}
