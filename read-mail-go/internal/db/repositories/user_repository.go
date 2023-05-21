package repositories

import (
	"auth-server/read-mail/internal/db"
	"auth-server/read-mail/internal/models"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() (*UserRepository, error) {
	// Establish a database connection
	dbConn, err := db.NewDBConnection()
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		db: dbConn,
	}, nil
}

func (ur *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO token (name, token)
		VALUES ($1, $2)
		ON CONFLICT (name) DO UPDATE
		SET token = $2
	`
	_, err := ur.db.Exec(query, user.Name, user.Token)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetByName(name string) (*models.User, error) {
	// Implement the logic to retrieve a user from the "token" table by ID
	row := ur.db.QueryRow("SELECT * FROM token WHERE name = $1", name)
	user := &models.User{}
	err := row.Scan(&user.Name, &user.Token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return user, nil
}

// ... Other repository methods such as Update and Delete
