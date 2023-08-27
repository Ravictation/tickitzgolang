package models

type Times_Scheduless struct {
	Id_time_schedule string   `json:"id_time_schedule" db:"id_time_schedule" valid:"-"`
	Time_schedule    string   `json:"time_schedule" db:"time_schedule" valid:"-"`
	Regency          string   `json:"regency" db:"regency" valid:"-"`
	Price            string   `json:"price" db:"price" valid:"-"`
	Set_date         string   `json:"set_date" db:"set_date" valid:"-"`
	Id_movie         string   `json:"id_movie" db:"id_movie" valid:"-"`
	Title            string   `json:"title" db:"title" valid:"-"`
	Image_movie      string   `json:"image_movie" db:"image_movie" valid:"-"`
	Release_date     *string  `json:"release_date" db:"release_date" valid:"-"`
	Id_premier       string   `json:"id_premier" db:"id_premier" valid:"-"`
	Name_premier     string   `json:"name_premier" db:"name_premier" valid:"-"`
	Image_premier    string   `json:"image_premier" db:"image_premier" valid:"-"`
	Count_row_seat   int      `json:"count_row_seat" db:"count_row_seat" valid:"-"`
	Count_col_seat   int      `json:"count_col_seat" db:"count_col_seat" valid:"-"`
	Genres           []Genres `json:"genres" valid:"-"`
}
