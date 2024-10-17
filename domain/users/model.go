package users

import (
	"time"

	"github.com/avila-r/wiredcraft-challenge/domain/addresses"
	"github.com/avila-r/wiredcraft-challenge/sql"

	"github.com/google/uuid"
)

type User struct {
	ID          *uuid.UUID        `json:"id,omitempty"` // user ID
	Name        string            `json:"name"`         // user name
	DOB         time.Time         `json:"dob"`          // date of birth
	Address     addresses.Address `json:"address"`      // user address
	Description string            `json:"description"`  // user description
	CreatedAt   time.Time         `json:"created_at"`   // user created date
}

func FromSQL(s sql.User) User {
	uuid, _ := uuid.FromBytes(s.ID.Bytes[:])

	return User{
		ID:          &uuid,
		Name:        s.Name,
		DOB:         s.Dob.Time,
		Description: s.Description,
		CreatedAt:   s.CreatedAt.Time,
	}
}

func (u User) OmitID() User {
	return User{
		ID:          nil,
		Name:        u.Name,
		DOB:         u.DOB,
		Address:     u.Address,
		Description: u.Description,
		CreatedAt:   u.CreatedAt,
	}
}
