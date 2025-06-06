package gapi

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	startTime := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(startTime)

	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	logger := log.Info()
	if err != nil {
		logger = log.Error().Err(err)
	}

	logger.Str("protocol", "grpc").
		Str("method", info.FullMethod).
		Int("status_code", int(statusCode)).
		Str("status_text", statusCode.String()).
		Dur("duration", duration).
		Msg("received a gRPC request")

	return result, err
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (rr *ResponseRecorder) WriteHeader(statusCode int) {
	rr.StatusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}

func (rr *ResponseRecorder) Write(b []byte) (int, error) {
	rr.Body = b
	return rr.ResponseWriter.Write(b)
}

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		rr := &ResponseRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		handler.ServeHTTP(rr, r)
		duration := time.Since(startTime)
		logger := log.Info()
		if rr.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("body", rr.Body)
		}
		logger.Str("protocol", "http").
			Str("method", r.Method).
			Str("path", r.RequestURI).
			Int("status_code", rr.StatusCode).
			Str("status_text", http.StatusText(rr.StatusCode)).
			Dur("duration", duration).
			Msg("received a HTTP request")
	})
}
