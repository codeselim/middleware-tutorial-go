package usersapi

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	//other properties are omitted in this example. Non present properties in our structures will get ignored
}
