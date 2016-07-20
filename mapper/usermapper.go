package mapper

import (
	"github.com/codeselim/middleware-tutorial-go/contract/api"
	"github.com/codeselim/middleware-tutorial-go/contract/usersapi"
	"time"
)

type UserMapper struct{}

func NewUserMapper() UserMapper {
	return UserMapper{}
}

func (um UserMapper) GetDomainUser(remoteUser usersapi.User) api.User {
	user := api.User{}
	user.Id = remoteUser.Id
	user.Email = remoteUser.Email
	user.Name = remoteUser.Name
	user.Username = remoteUser.Username
	user.LastLogin = time.Now().Format(time.RFC850)
	return user
}
