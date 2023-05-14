package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Book struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Roles struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserRoles struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
