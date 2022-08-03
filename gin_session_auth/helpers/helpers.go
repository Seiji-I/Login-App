package helpers

import (
	"log"
	"strings"
	"database/sql"
	_ "github.com/lib/pq"
)

func CheckUserPass(input_username, input_password string) bool {
	db, err := sql.Open("postgres", "user=root password=root host=postgres-db port=5432 dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
		
	var (
		user_id string
		user_name string
		user_password string
	)
	
	err = db.QueryRow("SELECT id, name, password FROM users WHERE name = $1", input_username).Scan(&user_id, &user_name, &user_password)
	if err != nil {
		if err == sql.ErrNoRows {
			user_id = ""
			user_name = ""
			user_password = ""
		} else {
			log.Fatal(err)
		}
	}
	is_match := strings.EqualFold(input_username, user_name) && strings.EqualFold(input_password, user_password)
	if is_match {
		return true
	} else {
		return false
	}
}

func EmptyUsePass(input_username, input_password string) bool {
	return strings.Trim(input_username, " ") == "" || strings.Trim(input_password, " ") == ""
}