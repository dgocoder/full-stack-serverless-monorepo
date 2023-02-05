package types

type CreateUser struct {
	Email string
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
