package database

import "math/big"

type User struct {
	id       big.Int `database:"id"`
	role     string  `database:"role"`
	password string  `database:"password"`
	email    string  `database:"email"`
}
