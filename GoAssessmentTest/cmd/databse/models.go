// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import ()

type Post struct {
	PostID  int64
	Title   string
	Content string
}

type Session struct {
	Token  string
	Data   []byte
	Expiry float64
}

type User struct {
	UserID   int64
	Username string
	Password string
}
