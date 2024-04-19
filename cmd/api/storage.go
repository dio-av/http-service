package api

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

// TODO: move storage to the repository folder

type Storage struct {
	db *gorm.DB
}

func (Storage) GetAccountByNumber(int) (Account, error) {
	return Account{}, nil
}

func (Storage) GetAccounts() ([]Account, error) {
	return nil, nil
}

func (Storage) GetAccountByID(id int) (Account, error) {
	return Account{}, nil
}

func getID(_ *http.Request) (int, error) {
	return 0, nil
}

func (Storage) DeleteAccount(id int) error {
	return nil
}

func (Storage) CreateAccount(account *Account) error {
	return nil
}

type Reader interface {
	Get(id int) (*Account, error)
}

// Writer book writer
type Writer interface {
	Create(e *Account) (int, error)
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	Get(id int) (*Account, error)
	Create(ID int, FirstName string, LastName string, Number int64, EncryptedPassword string, Balance int64, CreatedAt time.Time) (int, error)
}
