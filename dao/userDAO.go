package dao

import (
	"log"

	"github.com/pattyvader/users_service/models"
)

func GetAllUsers() ([]models.User, error) {
	InitDB()
	defer CloseDB()

	users := []models.User{}

	rows, err := db.Query(`SELECT id, name, email, admin FROM users order by id`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	if err == nil {
		for rows.Next() {
			var ID int
			var Name string
			var Email string
			var Admin bool

			err = rows.Scan(&ID, &Name, &Email, &Admin)
			if err == nil {
				currentUser := models.User{ID: ID, Name: Name, Email: Email, Admin: Admin}
				users = append(users, currentUser)
			} else {
				return users, err
			}
		}
	} else {
		return users, err
	}
	return users, err
}
