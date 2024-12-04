package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// LoggerInterface defines the logging methods
type LoggerInterface interface {
	LogRequest(c echo.Context, requestID string)
	LogResponse(c echo.Context, rw *ResponseWriter, requestID string, err error)
}

// Logger struct to encapsulate logrus.Logger
type Logger struct {
	Instance *logrus.Logger
}

// NewLogger initializes a new logger instance
func NewLogger(logFilePath string) (*Logger, error) {
	log := logrus.New()

	// Memastikan folder log ada
	logDir := logFilePath // Ganti dengan path folder yang sesuai jika perlu
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	// Check if running in test mode
	if strings.HasSuffix(os.Args[0], ".test") {
		log.SetOutput(os.Stdout)
		log.SetFormatter(&logrus.TextFormatter{})
		log.SetLevel(logrus.DebugLevel)
		return &Logger{Instance: log}, nil
	}

	// Log to file and stdout
	filename := logFilePath + PrefixLogFilename()
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	log.SetLevel(logrus.InfoLevel)

	return &Logger{Instance: log}, nil
}

// LogRequest logs the incoming HTTP request
func (l *Logger) LogRequest(c echo.Context, requestID string) {
	// Periksa apakah konten adalah multipart
	contentType := c.Request().Header.Get("Content-Type")
	isMultipart := strings.HasPrefix(contentType, "multipart/")

	var body interface{}

	// Jika bukan multipart, baca body dan batasi hingga 10MB
	if !isMultipart {
		bodyBytes, _ := io.ReadAll(io.LimitReader(c.Request().Body, 10*1024*1024)) // Limit to 10MB
		c.Request().Body = io.NopCloser(bytes.NewReader(bodyBytes))                // Restore body for further use

		if len(bodyBytes) > 0 {
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				body = map[string]interface{}{"raw": string(bodyBytes)}
			}
		} else {
			body = map[string]interface{}{} // Kosongkan body jika tidak ada isinya
		}
	} else {
		// Jika multipart, hanya mencatat bahwa ini adalah file upload
		body = map[string]interface{}{"multipart": "file upload detected, body not logged"}
	}

	// Log to stdout and file
	l.Instance.WithFields(logrus.Fields{
		"level":      "info",
		"action_msg": "REQUEST RECEIVED",
		"request_id": requestID,
		"method":     c.Request().Method,
		"url":        c.Request().URL.Path,
		"params":     c.ParamNames(),  // Mengambil nama parameter dari URL
		"query":      c.QueryString(), // Mengambil query string
		"body":       body,
		"headers":    c.Request().Header,
		"time":       time.Now().Format(time.RFC3339),
	}).Info()
}

// ResponseWriter wraps the original response writer to capture response body
type ResponseWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

// NewResponseWriter initializes a new ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		body:           new(bytes.Buffer),
	}
}

func (rw *ResponseWriter) Write(p []byte) (int, error) {
	rw.body.Write(p) // Capture response body
	return rw.ResponseWriter.Write(p)
}

// LogResponse logs the HTTP response
func (l *Logger) LogResponse(c echo.Context, rw *ResponseWriter, requestID string, err error) {
	responseBody := rw.body.String()
	statusCode := c.Response().Status

	// Potong responseBody jika lebih dari 3000 karakter
	if len(responseBody) > 3000 {
		responseBody = responseBody[:3000] + "...(truncated length of " + fmt.Sprint(len(responseBody)) + ")"
	}

	l.Instance.WithFields(logrus.Fields{
		"level":      "info",
		"action_msg": "RESPONSE SENT",
		"request_id": requestID,
		"method":     c.Request().Method,
		"url":        c.Request().URL.Path,
		"status":     statusCode,
		"body":       responseBody,
		"time":       time.Now().Format(time.RFC3339),
	}).Info()
}
