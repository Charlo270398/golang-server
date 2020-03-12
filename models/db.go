package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //variable db común a todos

func ConnectDB(){
	var err error
    db, err = sql.Open("mysql", "golang:golang@(127.0.0.1:3306)/golang?parseTime=true")
    if err != nil {
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }
}
func query(query string) bool{

	// Executes the SQL query in our database. Check err to ensure there was no error.
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
		return false;
	}
	return true;
}

func CreateDB() {
	ConnectDB()
	//CREATE TABLES
	query(USUARIOS_TABLE)
	query(ROLES_TABLE)
	query(USERS_ROLES_TABLE)

	fmt.Println("Database OK")
}

var USUARIOS_TABLE string = `
CREATE TABLE IF NOT EXISTS usuarios (
	id INT AUTO_INCREMENT,
	dni VARCHAR(9) UNIQUE,
	nombre VARCHAR(25) NOT NULL,
	apellidos VARCHAR(50) NOT NULL,
	email VARCHAR(59) UNIQUE,
	password VARCHAR(100) NOT NULL,
	created_at DATETIME,
	PRIMARY KEY (id)
);`

var ROLES_TABLE string = `
CREATE TABLE IF NOT EXISTS roles (
	id INT AUTO_INCREMENT,
	nombre VARCHAR(20) UNIQUE,
	descripcion VARCHAR(50),
	PRIMARY KEY (id)
);`

var USERS_ROLES_TABLE string = `
CREATE TABLE IF NOT EXISTS usuarios_roles (
	usuario_id INT,
	rol_id INT,
	created_at DATETIME,
	PRIMARY KEY (usuario_id, rol_id),
	FOREIGN KEY(usuario_id) REFERENCES usuarios(id),
	FOREIGN KEY(rol_id) REFERENCES roles(id)
);`
