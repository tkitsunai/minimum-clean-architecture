package domain

type Users []User

func NewEmptyUsers() Users {
	return make([]User, 0)
}

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
