package addresses

import (
	"github.com/avila-r/wiredcraft-challenge/db"
	"github.com/avila-r/wiredcraft-challenge/sql"

	"github.com/jackc/pgx/v5/pgtype"
)

type AddressService struct {
	db *sql.Conn
}

func NewService() *AddressService {
	return &AddressService{
		db: db.Conn,
	}
}

func (s *AddressService) Create(a Address) (*Address, error) {
	params := ToSQL(a)

	result, err := s.db.CreateUserAddress(params)

	address := FromSQL(result)

	return &address, err
}

func (s *AddressService) GetByUserID(id [16]byte) (Address, error) {
	uuid := s.parse(id)

	result, err := s.db.GetUserAddress(uuid)

	address := FromSQL(result)

	return address, err
}

func (s *AddressService) parse(id [16]byte) pgtype.UUID {
	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
}
