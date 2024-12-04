package utils

import (
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dgrijalva/jwt-go"
)

type CustomResponse struct {
	Status  int         `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

type CustomResponseWithPagination struct {
	Status   int         `json:"Status"`
	Data     interface{} `json:"Data"`
	Page     int         `json:"Page"`
	PageSize int         `json:"PageSize"`
	Total    int64       `json:"Total"`
}

type JwtCustomClaims struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HashedPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func GenerateJWT(id int, data interface{}) (string, error) {
	expires := time.Now().Add(time.Minute * 60).Unix()

	claims := &JwtCustomClaims{
		ID:   id,
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type H map[string]interface{}

func PrefixLogFilename() string {
	now := time.Now().Format("2006-01-02")
	return now + ".log"
}

/*
Function GetSelectFields mengembalikan daftar kolom yang terisi dalam struct models
Jika isAllField true, maka semua kolom akan diambil, jika false, maka hanya kolom tidak kosong akan diambil
@author Roby Parlan
*/
func GetSelectFields(model interface{}, isAllField bool) []string {
	var fields []string
	val := reflect.ValueOf(model)

	// Pastikan model adalah pointer dan dereferensiasi
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Pastikan model adalah struct
	if val.Kind() != reflect.Struct {
		return nil // atau Anda bisa return error jika bukan struct
	}

	// Loop untuk memeriksa nilai kolom yang terisi
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		// Periksa jika field terisi (nilai tidak kosong)
		if field.IsValid() {
			// Ambil nama kolom dari tag GORM
			fieldName := val.Type().Field(i).Tag.Get("gorm")

			if fieldName != "" {
				// Pisahkan bagian tag GORM seperti column:name;type:varchar
				parts := strings.Split(fieldName, ";")

				// Ambil nama kolom setelah 'column:'
				columnName := ""
				for _, part := range parts {
					if strings.HasPrefix(part, "column:") {
						columnName = strings.TrimPrefix(part, "column:")
						break
					}
				}

				// Jika nama kolom ditemukan, tambahkan ke dalam fields
				if columnName != "" {
					// Jika `isAllField` true, tambahkan semua kolom
					// Jika `isAllField` false, hanya tambahkan kolom yang terisi
					if isAllField || !field.IsZero() {
						fields = append(fields, columnName)
					}
				}
			}
		}
	}

	return fields
}

// GRPCErrorToHTTP mengonversi error gRPC ke response HTTP yang sesuai
func GRPCErrorToHTTP(err error) (int, string) {
	if err == nil {
		return http.StatusOK, MESSAGE_OK
	}

	// Mendapatkan kode status gRPC dari error
	code := status.Code(err)
	st, _ := status.FromError(err)
	message := st.Message()

	// Menangani berbagai kode gRPC error
	switch code {
	case codes.InvalidArgument:
		return http.StatusBadRequest, message
	case codes.Unauthenticated:
		return http.StatusUnauthorized, message
	case codes.PermissionDenied:
		return http.StatusForbidden, message
	case codes.NotFound:
		return http.StatusNotFound, message
	case codes.Internal:
		return http.StatusInternalServerError, message
	case codes.Unavailable:
		return http.StatusServiceUnavailable, message
	case codes.DeadlineExceeded:
		return http.StatusRequestTimeout, message
	default:
		return http.StatusInternalServerError, message
	}
}
