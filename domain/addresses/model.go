package addresses

import (
	"time"

	"github.com/avila-r/wiredcraft-challenge/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Address struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	AddressLine1 string    `json:"address_line_1"`
	AddressLine2 string    `json:"address_line_2"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	PostalCode   string    `json:"postal_code"`
	Country      string    `json:"country"`
	CreatedAt    time.Time `json:"created_at"`
}

func (a Address) BindTo(id pgtype.UUID) Address {
	return Address{
		ID:           a.ID,
		UserID:       id.Bytes,
		AddressLine1: a.AddressLine1,
		AddressLine2: a.AddressLine2,
		City:         a.City,
		State:        a.State,
		PostalCode:   a.PostalCode,
		Country:      a.Country,
		CreatedAt:    a.CreatedAt,
	}
}

func FromSQL(s sql.UserAddress) Address {
	return Address{
		ID:           s.ID.Bytes,
		UserID:       s.UserID.Bytes,
		AddressLine1: s.AddressLine1,
		AddressLine2: s.AddressLine2.String,
		City:         s.City,
		State:        s.State,
		PostalCode:   s.PostalCode,
		Country:      s.Country,
		CreatedAt:    s.CreatedAt.Time,
	}
}

func ToSQL(u Address) sql.CreateUserAddressParams {
	return sql.CreateUserAddressParams{
		UserID:       UUID(u.UserID),
		AddressLine1: u.AddressLine1,
		AddressLine2: text(u.AddressLine2),
		City:         u.City,
		State:        u.State,
		PostalCode:   u.PostalCode,
		Country:      u.Country,
	}
}

func UUID(bytes uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: bytes,
		Valid: true,
	}
}

func text(str string) pgtype.Text {
	return pgtype.Text{
		String: str,
		Valid:  true,
	}
}
