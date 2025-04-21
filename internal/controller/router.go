package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/keshvan/user-service-sstu-forum/internal/entity"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id int) (*entity.User, error)
}

func NewRouter(engine *gin.Engine, usecase UserUsecase) {
	h := UserHandler{usecase}

	users := engine.Group("/users")
	{
		users.GET("/:id", h.GetByID)
	}
}
