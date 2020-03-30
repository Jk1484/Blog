package users

import (
	"blog/back/app/db"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func List() (users []User, err error) {
	query := `
		SELECT id, name
		FROM users
	`

	rows, err := db.Client().Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u User

		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return
		}

		users = append(users, u)
	}

	return
}

func Get(id int) (u User, err error) {
	query := `
		SELECT id, name
		FROM users
		WHERE id = $1
	`

	err = db.Client().QueryRow(query, id).Scan(&u.ID, &u.Name)
	if err != nil {
		return
	}

	return
}
