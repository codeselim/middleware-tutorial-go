package connection

type UserConnection interface {
	GetUsers() (string, error)
	GetUserById(id string) (string, error)
}