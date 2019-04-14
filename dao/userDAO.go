package dao

import (
	"github.com/pattyvader/users_service/models"
)

//GetAllUsers retrives all users
func GetAllUsers() ([]models.User, error) {
	InitDB()
	defer CloseDB()

	users := []models.User{}

	rows, err := db.Query(`SELECT id, name, email, admin FROM users order by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
			return nil, err
		}
	}
	return users, err
}

//GetUserByID retrieves only one user
func GetUserByID(userID int) (*models.User, error) {
	InitDB()
	defer CloseDB()

	user := models.User{}

	var ID int
	var Name string
	var Email string
	var Password string
	var Admin bool

	err := db.QueryRow(`SELECT id, name, email, password, admin FROM users where id = $1`,
		userID).Scan(&ID, &Name, &Email, &Password, &Admin)
	if err == nil {
		user = models.User{ID: ID, Name: Name, Email: Email, Password: Password, Admin: Admin}
	} else {
		return nil, err
	}
	return &user, err
}

// CreateUser creates a new user
func CreateUser(user models.User) (*models.User, error) {
	InitDB()
	defer CloseDB()

	var userID int

	err := db.QueryRow(`INSERT INTO users(name, email, password, admin) VALUES($1, $2, $3, $4) RETURNING id`,
		user.Name, user.Email, user.Password, user.Admin).Scan(&userID)
	if err != nil {
		return nil, err
	}

	user.ID = userID
	return &user, err
}

//UpdateUser update an user
func UpdateUser(user models.User, userID int) (models.User, int, error) {
	InitDB()
	defer CloseDB()

	result, err := db.Exec(`UPDATE users set name=$1, email=$2, password=$3, admin=$4 where id=$5 RETURNING id`,
		user.Name, user.Email, user.Password, user.Admin, userID)
	if err != nil {
		return user, 0, err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return user, 0, err
	}

	user.ID = userID
	return user, int(rowsUpdated), err
}

//RemoveUser remove an user
func RemoveUser(userID int) (int, error) {
	InitDB()
	defer CloseDB()

	result, err := db.Exec(`DELETE FROM users where id = $1`, userID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), err
}
