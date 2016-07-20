package api

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	LastLogin string `json:"lastLogin"`
}
