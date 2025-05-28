package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PageResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
}

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	UNAUTHORIZED   = 401
	FORBIDDEN      = 403
	NOT_FOUND      = 404
)

var MsgFlags = map[int]string{
	SUCCESS:        "success",
	ERROR:          "error",
	INVALID_PARAMS: "invalid params",
	UNAUTHORIZED:   "unauthorized",
	FORBIDDEN:      "forbidden",
	NOT_FOUND:      "not found",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    SUCCESS,
		Message: GetMsg(SUCCESS),
		Data:    data,
	})
}

func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    SUCCESS,
		Message: message,
		Data:    data,
	})
}

func SuccessWithPage(c *gin.Context, data interface{}, total int64, page, size int) {
	c.JSON(http.StatusOK, PageResponse{
		Code:    SUCCESS,
		Message: GetMsg(SUCCESS),
		Data:    data,
		Total:   total,
		Page:    page,
		Size:    size,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(getHTTPStatusCode(code), Response{
		Code:    code,
		Message: message,
	})
}

func ErrorWithCode(c *gin.Context, code int) {
	c.JSON(getHTTPStatusCode(code), Response{
		Code:    code,
		Message: GetMsg(code),
	})
}

func getHTTPStatusCode(code int) int {
	switch code {
	case SUCCESS:
		return http.StatusOK
	case INVALID_PARAMS:
		return http.StatusBadRequest
	case UNAUTHORIZED:
		return http.StatusUnauthorized
	case FORBIDDEN:
		return http.StatusForbidden
	case NOT_FOUND:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}