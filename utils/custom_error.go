package utils

import (
	"errors"
	"fmt"
	"net/http"
)

// CustomError struct untuk menyimpan informasi error
type CustomError struct {
	StatusCode int         `json:"StatusCode"` // Kode status HTTP
	Message    string      `json:"Message"`    // Pesan error
	Err        interface{} `json:"-"`          // Error asli (bisa tipe apa saja)
}

// NewCustomError membuat instance CustomError dari error
func NewCustomError(err error) *CustomError {
	// Default nilai jika error tidak memiliki detail
	statusCode := http.StatusInternalServerError
	message := "An unexpected error occurred"

	// Cek apakah error adalah tipe CustomError
	var customErr *CustomError
	if errors.As(err, &customErr) {
		// Gunakan detail dari CustomError
		statusCode = customErr.StatusCode
		message = customErr.Message
	} else if err != nil {
		// Gunakan pesan error asli jika bukan CustomError
		message = err.Error()
	}

	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

// Implementasikan metode Error untuk memenuhi interface error
func (ce *CustomError) Error() string {
	switch e := ce.Err.(type) {
	case nil:
		return ce.Message
	case error:
		return fmt.Sprintf("%d - %s: %v", ce.StatusCode, ce.Message, e)
	case string:
		return fmt.Sprintf("%d - %s: %s", ce.StatusCode, ce.Message, e)
	default:
		return fmt.Sprintf("%d - %s: %+v", ce.StatusCode, ce.Message, ce.Err)
	}
}

// Unwrap untuk mengambil error asli
func (ce *CustomError) Unwrap() error {
	if err, ok := ce.Err.(error); ok {
		return err
	}
	return nil
}
