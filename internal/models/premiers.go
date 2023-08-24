package models

import "time"

type Premiers struct {
	Id_premier     string     `db:"id_premier" form:"id_premier" json:"id_premier,omitempty" valid:"-"`
	Name_premier   string     `db:"name_premier" form:"name_premier,omitempty" json:"name_premier,omitempty" valid:"required~name premier is required,type(string)~name premier is string"`
	Image          string     `db:"image" json:"image,omitempty" valid:"-"`
	Count_row_seat string     `db:"count_row_seat" form:"count_row_seat" json:"count_row_seat,omitempty" valid:"-"`
	Count_col_seat string     `db:"count_col_seat" form:"count_col_seat" json:"count_col_seat,omitempty" valid:"-"`
	Created_at     *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at     *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}
