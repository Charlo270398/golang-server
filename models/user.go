package models

import (
	"fmt"
	"strconv"
	"time"

	util "../utils"
)

var p = &util.Params_argon2{
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: 1,
	SaltLength:  16,
	KeyLength:   32,
}

func LoginUser(email string, password string) bool {
	user, err := GetUserByEmail(email)
	if err != nil {
		util.PrintErrorLog(err)
		return false
	}
	if user.Password == "" {
		util.PrintLog("El usuario no existe")
		return false
	}
	match, err := util.Argon2comparePasswordAndHash(password, user.Password)
	if err != nil {
		util.PrintErrorLog(err)
	}
	return match
}

func InsertUser(user util.User_json) (inserted bool, err error) {
	//ARGON2
	encodedHash, err := util.Argon2generateFromPassword(user.Password, p)
	if err != nil {
		fmt.Println(err)
		util.PrintErrorLog(err)
		return false, err
	}
	//INSERT
	createdAt := time.Now()
	_, err = db.Exec(`INSERT INTO usuarios (dni, nombre, apellidos, email, password, created_at) VALUES (?, ?, ?, ?, ?, ?)`, user.Identificacion,
		user.Nombre, user.Apellidos, user.Email, encodedHash, createdAt)
	if err == nil {
		return true, nil
	} else {
		fmt.Println(err)
		util.PrintErrorLog(err)
	}
	return false, nil
}

func DeleteUser(user_id int) bool {
	_, err := db.Exec(`DELETE FROM usuarios where id = ` + strconv.Itoa(user_id))
	if err == nil {
		return true
	} else {
		fmt.Println(err)
		util.PrintErrorLog(err)
	}
	return false
}

func GetUserById(id int) (user util.User, err error) {
	row, err := db.Query(`SELECT id, dni, nombre, apellidos, email, password, created_at FROM usuarios where id = ` + strconv.Itoa(id)) // check err
	if err == nil {
		defer row.Close()
		row.Next()
		row.Scan(&user.Id, &user.Identificacion, &user.Nombre, &user.Apellidos, &user.Email, &user.Password, &user.CreatedAt)
		return user, err
	} else {
		fmt.Println(err)
		util.PrintErrorLog(err)
		return user, err
	}
}

func GetUserByEmail(email string) (user util.User, err error) {
	row, err := db.Query(`SELECT id, dni, nombre, apellidos, email, password, created_at FROM usuarios where email = '` + email + `'`) // check err
	if err == nil {
		defer row.Close()
		row.Next()
		row.Scan(&user.Id, &user.Identificacion, &user.Nombre, &user.Apellidos, &user.Email, &user.Password, &user.CreatedAt)
		return user, err
	} else {
		fmt.Println(err)
		util.PrintErrorLog(err)
		return user, err
	}
}

func GetUsersPagination(page int) []util.User {
	firstRow := strconv.Itoa(page * 10)
	lastRow := strconv.Itoa((page * 10) + 10)
	rows, err := db.Query("SELECT id, dni, nombre, apellidos, email, created_at FROM usuarios LIMIT " + firstRow + "," + lastRow)
	if err == nil {
		defer rows.Close()
		var users []util.User
		for rows.Next() {
			var u util.User
			rows.Scan(&u.Id, &u.Identificacion, &u.Nombre, &u.Apellidos, &u.Email, &u.CreatedAt)
			users = append(users, u)
		}
		return users
	} else {
		fmt.Println(err)
		return nil
	}
}
