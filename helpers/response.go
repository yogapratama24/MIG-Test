package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	// Data    interface{} `json:"data,omitempty"`
	Data interface{} `json:"data"`
}

type HandlerValidationResponse struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHandlerResponse(message string, data interface{}) *HandlerResponse {
	return &HandlerResponse{
		Message: message,
		Data:    data,
	}
}

func NewHandlerValidationResponse(message, data interface{}) *HandlerValidationResponse {
	return &HandlerValidationResponse{
		Message: message,
		Data:    data,
	}
}

func (response *HandlerResponse) Success(c *gin.Context) {
	response.Status = true
	c.JSON(http.StatusOK, response)
}

func (response *HandlerResponse) SuccessCreate(c *gin.Context) {
	response.Status = true
	c.JSON(http.StatusCreated, response)
}

func (response *HandlerResponse) Failed(c *gin.Context) {
	response.Status = false
	c.JSON(http.StatusInternalServerError, response)
}

func (response *HandlerResponse) BadRequest(c *gin.Context) {
	response.Status = false
	c.JSON(http.StatusBadRequest, response)
}

func (response *HandlerResponse) Unauthorized(c *gin.Context) {
	response.Status = false
	c.JSON(http.StatusUnauthorized, response)
}

func (response *HandlerValidationResponse) Failed(c *gin.Context) {
	response.Status = false
	c.JSON(http.StatusInternalServerError, response)
}

func (response *HandlerValidationResponse) BadRequest(c *gin.Context) {
	response.Status = false
	c.JSON(http.StatusBadRequest, response)
}
