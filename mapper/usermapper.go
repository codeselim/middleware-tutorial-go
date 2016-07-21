package mapper

import (
	"github.com/codeselim/middleware-tutorial-go/contract/api"
	"github.com/codeselim/middleware-tutorial-go/contract/usersapi"
)

type User interface {
	GetDomainUser(remoteUser usersapi.User, data string) api.User
}

type UserMapper struct{}

func NewUserMapper() UserMapper {
	return UserMapper{}
}

func (um UserMapper) GetDomainUser(remoteUser usersapi.User, data string) api.User {
	user := api.User{}
	user.Id = remoteUser.Id
	user.Email = remoteUser.Email
	user.Name = remoteUser.Name
	user.Username = remoteUser.Username
	user.Data = data
	return user
}
