package models

import (
	"time"
)

type User struct {
	Id_user      string     `db:"id_user" form:"id_user" json:"id_user,omitempty" valid:"-"`
	Email_user   string     `db:"email_user" form:"email_user" json:"email_user" valid:"email~enter valid email address"`
	Password     string     `db:"password" form:"password" json:"password,omitempty" valid:"stringlength(6|20)~password must be atleast 6 characters long"`
	Phone_number string     `db:"phone_number" form:"phone_number" json:"phone_number,omitempty" valid:"-"`
	Image_user   string     `db:"image_user"  json:"image_user" valid:"-"`
	Role         string     `db:"role" form:"role" json:"role" valid:"-"`
	First_name   string     `db:"first_name" form:"first_name" json:"first_name" valid:"-"`
	Last_name    string     `db:"last_name" form:"last_name" json:"last_name" valid:"-"`
	Created_at   *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}
