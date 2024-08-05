package order

import (
	"database/sql"

	"github.com/ahenla/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) error {
	res, err := s.db.Exec(, args ...any)
}
