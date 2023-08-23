package models

import "time"

type Casts struct {
	Id_cast    string     `db:"id_cast" form:"id_cast" json:"id_cast,omitempty" valid:"-"`
	Name_cast  string     `db:"name_cast" form:"name_cast,omitempty" json:"name_cast,omitempty" valid:"required~name cast is required,type(string)~name cast is string"`
	Created_at *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}
