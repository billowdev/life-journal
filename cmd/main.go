package main

import (
	"fmt"

	"github.com/billowdev/life-journal/internal/core/domain"
	"github.com/billowdev/life-journal/pkg/argon2id"
	"github.com/billowdev/life-journal/pkg/uuidv7"
)

func main() {
	// Create a new user with a given ID, name, and email
	user := domain.User(domain.User{
		FirstName: "John Doe",
		LastName:  "John Doe",
		Email:     "johndoe@example.com",
		Username:  "john",
		Password:  argon2id.NewPassword("password"),
		BaseModel: domain.BaseModel{
			ID:        uuidv7.MustGenerateUUIDv7(),
			CreatedAt: domain.Now(),
			UpdatedAt: domain.Now(),
			DeletedAt: nil,
		},
	})

	// Print the user's information
	fmt.Printf("User ID: %d\n", user.ID)

}
