package utils

import (
	"time"
)

type User struct {
	Id             int
	Identificacion string
	Nombre         string
	Apellidos      string
	Email          string
	Password       string
	CreatedAt      time.Time
}

type User_json struct {
	Identificacion string `json:"identificacion"`
	Nombre         string `json:"nombre"`
	Apellidos      string `json:"apellidos"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

type User_id_json struct {
	Id int `json:"user_id"`
}

type Params_argon2 struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

type PaginationPage_JSON struct {
	Page int `json:"Page"`
}

type UserList_JSON_Pagination struct {
	Page       int    `json:"Page"`
	NextPage   int    `json:"NextPage"`
	BeforePage int    `json:"BeforePage"`
	UserList   []User `json:"UserList"`
}
