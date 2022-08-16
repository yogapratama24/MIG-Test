package repositories

import (
	"database/sql"
	"log"
	"mitramas_test/models"
)

type UserRepository interface {
	ReadUser() (*[]models.UserResponse, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) ReadUser() (*[]models.UserResponse, error) {
	db := r.db
	var result = []models.UserResponse{}

	sqlStatement := `SELECT id, user_name, email, created_at, updated_at FROM users`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Printf("Error get data users with err: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user = models.UserResponse{}

		err := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("Error get data user with err: %s", err)
			return nil, err
		}

		result = append(result, user)
	}

	return &result, nil
}
