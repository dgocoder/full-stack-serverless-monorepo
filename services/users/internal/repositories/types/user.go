package types

type CreateUser struct {
	Email string
}

type User struct {
	BaseEntity
	ID    string `json:"id"`
	Email string `json:"email"`
}

type BaseEntity struct {
	EntityType string `json:"entity_type"`
}

func (User) GetEntityType() string {
	return "USER"
}
