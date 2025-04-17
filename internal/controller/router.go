package controller

import (
	"context"
	"user-service/internal/entity"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id int) (*entity.User, error)
}

func NewRouter(engine *gin.Engine, usecase UserUsecase) {
	h := UserHandler{usecase}

	users := engine.Group("/users")
	{
		users.GET("/:id", h.GetUserByID)
	}
}
