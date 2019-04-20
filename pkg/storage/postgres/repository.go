package postgres

import (
	"database/sql"
	"fmt"

	"github.com/RainJoe/go-template/pkg/adding"
	"github.com/RainJoe/go-template/pkg/config"
	"github.com/RainJoe/go-template/pkg/listing"

	//pq
	_ "github.com/lib/pq"
)

// Storage stores  data in postgresql
type Storage struct {
	db *sql.DB
}

// NewStorage returns a new postgres storage
func NewStorage(cfg config.TOMLConfig) (*Storage, error) {
	var err error

	s := new(Storage)

	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
			cfg.DBConfig.Host,
			cfg.DBConfig.Port,
			cfg.DBConfig.UserName,
			cfg.DBConfig.DBName,
			cfg.DBConfig.PassWord))

	s.db = db

	return s, err
}

//AddUser saves the given user to the repository
func (s *Storage) AddUser(u adding.User) error {
	stmt, err := s.db.Prepare("INSERT INTO t_user (username, password) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.UserName, u.PassWord)
	if err != nil {
		return err
	}
	return nil
}

//GetAllUsers returns all users
func (s *Storage) GetAllUsers() ([]listing.User, error) {
	users := make([]listing.User, 0)
	rows, err := s.db.Query("SELECT id, username FROM t_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u listing.User
		rows.Scan(&u.ID, &u.UserName)
		users = append(users, u)
	}
	return users, nil
}
