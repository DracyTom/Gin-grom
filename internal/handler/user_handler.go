package handler

import (
	"strconv"

	"github.com/dengxing/gin-grom/internal/middleware"
	"github.com/dengxing/gin-grom/internal/model"
	"github.com/dengxing/gin-grom/internal/service"
	"github.com/dengxing/gin-grom/pkg/auth"
	"github.com/dengxing/gin-grom/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

// Register 用户注册
func (h *UserHandler) Register(c *gin.Context) {
	var req model.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.INVALID_PARAMS, err.Error())
		return
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		if err.Error() == "username already exists" || err.Error() == "email already exists" {
			response.Error(c, response.INVALID_PARAMS, err.Error())
		} else {
			response.Error(c, response.ERROR, err.Error())
		}
		return
	}

	response.SuccessWithMessage(c, "User registered successfully", user.ToResponse())
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var req model.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.INVALID_PARAMS, err.Error())
		return
	}

	user, err := h.userService.Login(&req)
	if err != nil {
		response.Error(c, response.UNAUTHORIZED, err.Error())
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		response.Error(c, response.ERROR, "Failed to generate token")
		return
	}

	response.Success(c, gin.H{
		"token": token,
		"user":  user.ToResponse(),
	})
}

// GetProfile 获取当前用户信息
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.ErrorWithCode(c, response.UNAUTHORIZED)
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		response.Error(c, response.NOT_FOUND, err.Error())
		return
	}

	response.Success(c, user.ToResponse())
}

// UpdateProfile 更新当前用户信息
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.ErrorWithCode(c, response.UNAUTHORIZED)
		return
	}

	var req model.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.INVALID_PARAMS, err.Error())
		return
	}

	// 普通用户不能修改状态和角色
	req.Status = nil
	req.Role = ""

	user, err := h.userService.UpdateUser(userID, &req)
	if err != nil {
		response.Error(c, response.ERROR, err.Error())
		return
	}

	response.SuccessWithMessage(c, "Profile updated successfully", user.ToResponse())
}

// GetUsers 获取用户列表（管理员）
func (h *UserHandler) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	keyword := c.Query("keyword")
	role := c.Query("role")
	status := c.Query("status")

	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 10
	}

	// 构建搜索参数
	searchParams := map[string]interface{}{
		"keyword": keyword,
		"role":    role,
		"status":  status,
	}

	users, total, err := h.userService.GetUsersWithSearch(page, size, searchParams)
	if err != nil {
		response.Error(c, response.ERROR, err.Error())
		return
	}

	var userResponses []*model.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	response.SuccessWithPage(c, userResponses, total, page, size)
}

// GetUser 获取指定用户信息（管理员）
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, response.INVALID_PARAMS, "Invalid user ID")
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		response.Error(c, response.NOT_FOUND, err.Error())
		return
	}

	response.Success(c, user.ToResponse())
}

// UpdateUser 更新指定用户信息（管理员）
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, response.INVALID_PARAMS, "Invalid user ID")
		return
	}

	var req model.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.INVALID_PARAMS, err.Error())
		return
	}

	user, err := h.userService.UpdateUser(uint(id), &req)
	if err != nil {
		response.Error(c, response.ERROR, err.Error())
		return
	}

	response.SuccessWithMessage(c, "User updated successfully", user.ToResponse())
}

// DeleteUser 删除用户（管理员）
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, response.INVALID_PARAMS, "Invalid user ID")
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		response.Error(c, response.ERROR, err.Error())
		return
	}

	response.SuccessWithMessage(c, "User deleted successfully", nil)
}