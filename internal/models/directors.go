package models

import "time"

type Directors struct {
	Id_director   string     `db:"id_director" form:"id_director" json:"id_director,omitempty" valid:"-"`
	Name_director string     `db:"name_director" form:"name_director,omitempty" json:"name_director,omitempty" valid:"required~name director is required,type(string)~name director is string"`
	Created_at    *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}
