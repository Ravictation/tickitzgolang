package models

import "time"

type Bookings struct {
	Id_booking       string             `json:"id_booking" db:"id_booking" valid:"-"`
	Id_time_schedule string             `json:"-" db:"id_time_schedule" valid:"-"`
	Id_user          string             `json:"id_user" db:"id_user" valid:"-"`
	Seats            string             `json:"seats" db:"seats" valid:"-"`
	Total            int                `json:"total" db:"total" valid:"-"`
	Schedule         []Times_Scheduless `json:"schedule" valid:"-"`
	Created_at       *time.Time         `db:"created_at" json:"created_at" valid:"-"`
	Updated_at       *time.Time         `db:"updated_at" json:"updated_at" valid:"-"`
}

type Bookingsset struct {
	Id_booking       string `json:"id_booking" db:"id_booking" valid:"-"`
	Id_time_schedule string `json:"id_time_schedule" db:"id_time_schedule" valid:"-"`
	Id_user          string `json:"id_user" db:"id_user" valid:"-"`
	Seats            string `json:"seats" db:"seats" valid:"-"`
	Total            int    `json:"total" db:"total" valid:"-"`
}
