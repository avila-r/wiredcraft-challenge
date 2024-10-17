package users

import (
	"github.com/avila-r/wiredcraft-challenge/db"
	"github.com/avila-r/wiredcraft-challenge/domain/addresses"
	"github.com/avila-r/wiredcraft-challenge/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	addr_service *addresses.AddressService
	db           *sql.Conn
}

func NewService() *UserService {
	return &UserService{
		addr_service: addresses.NewService(),
		db:           db.Conn,
	}
}

func (s *UserService) Create(u User) (*User, error) {
	dob := pgtype.Date{
		Time:  u.DOB,
		Valid: true,
	}

	created, err := s.db.CreateUser(sql.CreateUserParams{
		Name:        u.Name,
		Dob:         dob,
		Description: u.Description,
	})

	if err != nil {
		return nil, err
	}

	address := u.Address.BindTo(created.ID)

	if _, err := s.addr_service.Create(address); err != nil {
		return nil, err
	}

	return s.GetByID(created.ID.Bytes)
}

func (s *UserService) List() ([]User, error) {
	var (
		users []User
		err   error
	)

	results, err := s.db.ListUsers()

	if err != nil {
		return users, err
	}

	for _, row := range results {
		u, _ := s.GetByID(row.ID.Bytes)

		users = append(users, *u)
	}

	return users, err
}

func (s *UserService) GetByID(id [16]byte) (*User, error) {
	target := s.parse(id)

	result, err := s.db.GetUser(target)

	if err != nil {
		return nil, err
	}

	address, err := s.addr_service.GetByUserID(id)

	uuid, _ := uuid.FromBytes(result.ID.Bytes[:])

	user := User{
		ID:          &uuid,
		Name:        result.Name,
		DOB:         result.Dob.Time,
		Address:     address,
		Description: result.Description,
		CreatedAt:   result.CreatedAt.Time,
	}

	return &user, err
}

func (s *UserService) DeleteByID(id uuid.UUID) error {
	uuid := s.parse(id)

	return s.db.DeleteUser(uuid)
}

func (s *UserService) parse(id [16]byte) pgtype.UUID {
	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
}
