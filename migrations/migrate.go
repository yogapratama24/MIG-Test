package migrations

import (
	"log"
	"mitramas_test/db"
	"mitramas_test/helpers"

	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
	db := db.Connect()
	defer db.Close()
	// User Migration
	sqlStatement := `CREATE TABLE IF NOT EXISTS users (
			id int AUTO_INCREMENT NOT NULL,
			user_name varchar(255) NOT NULL,
			password varchar(255) NOT NULL,
			email varchar(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			PRIMARY KEY(id),
			UNIQUE KEY (email)
		);`

	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Printf("Error create table users with err: %s", err)
		helpers.NewHandlerResponse("Error create table users", nil).Failed(c)
	}

	// Check In Migration
	sqlStatement = `CREATE TABLE IF NOT EXISTS check_ins (
			id int AUTO_INCREMENT NOT NULL,
			date_check_in TIMESTAMP NOT NULL,
			user_id int NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			PRIMARY KEY(id),
			FOREIGN KEY (user_id)
				REFERENCES users(id)
				ON UPDATE CASCADE
				ON DELETE CASCADE
		);`

	_, err = db.Exec(sqlStatement)

	if err != nil {
		log.Printf("Error create table check-in with err: %s", err)
		helpers.NewHandlerResponse("Error create table check-in", nil).Failed(c)
	}

	// Activity Migration
	sqlStatement = `CREATE TABLE IF NOT EXISTS activities (
			id int AUTO_INCREMENT NOT NULL,
			check_in_id int NOT NULL,
			description TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY(id),
			FOREIGN KEY (check_in_id)
				REFERENCES check_ins(id)
				ON UPDATE CASCADE
				ON DELETE CASCADE
		);`

	_, err = db.Exec(sqlStatement)

	if err != nil {
		log.Printf("Error create table activity in with err: %s", err)
		helpers.NewHandlerResponse("Error create table activity", nil).Failed(c)
	}

	// Check Out Migration
	sqlStatement = `CREATE TABLE IF NOT EXISTS check_outs (
			id int AUTO_INCREMENT NOT NULL,
			check_in_id int NOT NULL,
			date_check_out TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			PRIMARY KEY(id),
			FOREIGN KEY (check_in_id)
				REFERENCES check_ins(id)
				ON UPDATE CASCADE
				ON DELETE CASCADE
		);`

	_, err = db.Exec(sqlStatement)

	if err != nil {
		log.Printf("Error create table check-out in with err: %s", err)
		helpers.NewHandlerResponse("Error create table check-out", nil).Failed(c)
	}

	helpers.NewHandlerResponse("Successfully migrate databases", nil).Success(c)
}
