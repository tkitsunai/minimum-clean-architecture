package response

import (
	"github.com/tkitsunai/minimum-clean-architecture/domain"
)

type JsonContext interface {
	JSON(code int, i interface{}) error
}

type UserResponsePort interface {
	Response(users domain.Users, jsonContext JsonContext) error
}
