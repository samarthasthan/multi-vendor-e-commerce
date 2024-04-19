// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package authentication

import (
	"database/sql"
	"time"
)

type Reset struct {
	Resetid    string
	User       sql.NullString
	Resettoken string
	Expiresat  time.Time
	Isused     sql.NullBool
	Createdat  time.Time
	Updatedat  time.Time
	Deleteat   sql.NullTime
}

type Role struct {
	Roleid    string
	Rolename  string
	Roledesc  sql.NullString
	Createdat time.Time
	Updatedat time.Time
	Deleteat  sql.NullTime
}

type User struct {
	Userid     string
	Firstname  string
	Lastname   string
	Email      string
	Isverified sql.NullBool
	Phoneno    sql.NullString
	Password   string
	Role       sql.NullString
	Blocked    sql.NullBool
	Createdat  time.Time
	Updatedat  time.Time
	Deleteat   sql.NullTime
}

type Verification struct {
	Verificationid string
	User           sql.NullString
	Verifytoken    string
	Expiresat      time.Time
	Isused         sql.NullBool
	Createdat      time.Time
	Updatedat      time.Time
	Deleteat       sql.NullTime
}