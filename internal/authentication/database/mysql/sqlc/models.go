// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type PermissionsActiontype string

const (
	PermissionsActiontypeRead   PermissionsActiontype = "read"
	PermissionsActiontypeWrite  PermissionsActiontype = "write"
	PermissionsActiontypeDelete PermissionsActiontype = "delete"
)

func (e *PermissionsActiontype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PermissionsActiontype(s)
	case string:
		*e = PermissionsActiontype(s)
	default:
		return fmt.Errorf("unsupported scan type for PermissionsActiontype: %T", src)
	}
	return nil
}

type NullPermissionsActiontype struct {
	PermissionsActiontype PermissionsActiontype
	Valid                 bool // Valid is true if PermissionsActiontype is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPermissionsActiontype) Scan(value interface{}) error {
	if value == nil {
		ns.PermissionsActiontype, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PermissionsActiontype.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPermissionsActiontype) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PermissionsActiontype), nil
}

type Permission struct {
	Permissionid   string
	Permissionname string
	Permissiondesc sql.NullString
	Actiontype     PermissionsActiontype
	Createdat      time.Time
	Updatedat      time.Time
	Deletedat      sql.NullTime
}

type Reset struct {
	Resetid    string
	User       sql.NullString
	Resettoken string
	Expiresat  time.Time
	Isused     sql.NullBool
	Createdat  time.Time
	Updatedat  time.Time
	Deletedat  sql.NullTime
}

type Role struct {
	Roleid    string
	Rolename  string
	Roledesc  sql.NullString
	Createdat time.Time
	Updatedat time.Time
	Deletedat sql.NullTime
}

type Rolepermission struct {
	Roleid       string
	Permissionid string
	Createdat    time.Time
	Updatedat    time.Time
	Deletedat    sql.NullTime
}

type User struct {
	Userid     string
	Firstname  string
	Lastname   string
	Email      string
	Isverified sql.NullBool
	Phoneno    string
	Password   string
	Roleid     string
	Blocked    sql.NullBool
	Createdat  time.Time
	Updatedat  time.Time
	Deletedat  sql.NullTime
}

type Verification struct {
	Verificationid string
	User           sql.NullString
	Verifytoken    string
	Expiresat      time.Time
	Isused         sql.NullBool
	Createdat      time.Time
	Updatedat      time.Time
	Deletedat      sql.NullTime
}
