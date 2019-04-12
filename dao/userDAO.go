package dao

import (
	"log"

	"github.com/pattyvader/users_service/models"
)

//GetAllUsers retorna todos os usuários que estao cadastrados no banco de dados
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

//GetUserByID retorna somente 1 usuário cadastrado no banco
func GetUserByID(userID string) (models.User, error) {
	InitDB()
	defer CloseDB()

	user := models.User{}

	var ID int
	var Name string
	var Email string
	var Password string
	var Admin bool

	err := db.QueryRow(`SELECT id, name, email, password, admin FROM users where id = $1`, userID).Scan(&ID, &Name, &Email, &Password, &Admin)
	if err == nil {
		user = models.User{ID: ID, Name: Name, Email: Email, Password: Password, Admin: Admin}
	}
	return user, err

}

// CreateUser cria um novo usuário no banco de dados
func CreateUser(user models.User) (models.User, error) {
	InitDB()
	defer CloseDB()

	var userID int

	err := db.QueryRow(`INSERT INTO users(name, email, password, admin) VALUES($1, $2, $3, $4) RETURNING id`,
		user.Name, user.Email, user.Password, user.Admin).Scan(&userID)
	if err != nil {
		return user, err
	}

	user.ID = userID
	return user, err
}

//UpdateUser atualiza um usuário
func UpdateUser(user models.User, userID string) (int, error) {
	InitDB()
	defer CloseDB()

	response, err := db.Exec(`UPDATE users set name=$1, email=$2, password=$3, admin=$4 where id=$5 RETURNING id`,
		user.Name, user.Email, user.Password, user.Admin, userID)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	rowsUpdated, err := response.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

//RemoveUser remove um usuário
func RemoveUser(userID string) (int, error) {
	InitDB()
	defer CloseDB()

	response, err := db.Exec(`DELETE FROM users where id = $1`, userID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := response.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
