package auth

import (
	"github.com/aminkhn/portfolio/model"
	"github.com/google/uuid"
)

type BaseClaims struct {
	UUID     uuid.UUID
	ID       uint         `json:"id"`
	Name     string       `json:"name"`
	LastName string       `json:"last_name"`
	Role     string       `json:"role"`
	user     []model.User `json:"user"`
}
