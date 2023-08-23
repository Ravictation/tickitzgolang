package models

import "time"

type Genres struct {
	Id_genre   string     `db:"id_genre" form:"id_genre" json:"id_genre,omitempty" valid:"-"`
	Name_genre string     `db:"name_genre" form:"name_genre,omitempty" json:"name_genre,omitempty" valid:"required~name genre is required,type(string)~name genre is string"`
	Created_at *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}
