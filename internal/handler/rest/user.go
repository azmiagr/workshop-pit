package rest

import (
	"net/http"
	"workshop-pit/model"
	"workshop-pit/pkg/response"

	"github.com/gin-gonic/gin"
)

func (r *Rest) RegisterUser(c *gin.Context) {
	var req model.UserRegisterParam

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind json", err)
		return
	}

	resp, err := r.service.UserService.RegisterUser(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to register user", err)
		return
	}

	response.Success(c, http.StatusOK, "success register user", resp)
}

func (r *Rest) LoginUser(c *gin.Context) {
	var req model.UserLoginParam

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind json", err)
		return
	}

	resp, err := r.service.UserService.LoginUser(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to login user", err)
		return
	}

	response.Success(c, http.StatusOK, "success login user", resp)
}
