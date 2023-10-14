package middleware

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type responseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := trace.SpanFromContext(ctx.Request.Context())
		if span.IsRecording() {
			span.SetAttributes(
				attribute.String("method", ctx.Request.Method),
				attribute.String("path", ctx.Request.URL.Path),
				attribute.String("request", ctx.Request.URL.RawQuery),
			)
		}

		ctx.Next()

		responsestatus := ctx.Writer.Status()

		if span.IsRecording() {
			span.SetAttributes(
				attribute.Int("response.status", responsestatus),
			)
			span.End()
		}
	}
}

func parseBody(d []byte) (string, error) {
	var bodyMap map[string]interface{}

	if err := json.Unmarshal(d, &bodyMap); err != nil {
		return "", err
	}

	// Replace the values of undesired fields with "*****".
	bodyMap = replaceSensitiveFields(bodyMap, "password", "access_token", "refresh_token")

	// Marshal the map back to a JSON string.
	newBody, err := json.Marshal(bodyMap)
	if err != nil {
		return "", err
	}

	return string(newBody), nil
}

// Replace the values of specified fields with "*****".
func replaceSensitiveFields(data map[string]interface{}, sensitiveFields ...string) map[string]interface{} {
	for _, field := range sensitiveFields {
		if _, exists := data[field]; exists {
			data[field] = "*****"
		}
	}
	return data
}
