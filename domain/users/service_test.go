package users_test

import (
	"testing"
	"time"

	"github.com/avila-r/wiredcraft-challenge/domain/addresses"
	"github.com/avila-r/wiredcraft-challenge/domain/users"

	"github.com/google/uuid"
)

var (
	id = uuid.New()

	address = addresses.Address{
		ID:           uuid.New(),
		UserID:       id,
		AddressLine1: "123 Main St",
		AddressLine2: "Apt 4B",
		City:         "Springfield",
		State:        "IL",
		PostalCode:   "62701",
		Country:      "USA",
		CreatedAt:    time.Now(),
	}

	dob = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

	user = users.User{
		ID:          &id,
		Name:        "test-user",
		DOB:         dob,
		Address:     address,
		Description: "test-description",
		CreatedAt:   time.Now(),
	}
)

func Test_Create(t *testing.T) {
	service := users.NewService()

	result, err := service.Create(user)

	t.Cleanup(func() {
		service.DeleteByID(*result.ID)
	})

	if err != nil {
		t.Error(err.Error())
	}
}
