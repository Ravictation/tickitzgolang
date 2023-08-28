package repositories

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repo_Bookings struct {
	*sqlx.DB
}

func NewBookings(db *sqlx.DB) *Repo_Bookings {
	return &Repo_Bookings{db}
}

func (r *Repo_Bookings) Get_Data(data *models.Bookings, page string, limit string, user string, time_schedule string, id_booking string) (*config.Result, error) {
	datas := []models.Bookings{}
	var metas config.Metas

	var offset int = 0
	var page_int, _ = strconv.Atoi(page)
	var limit_int, _ = strconv.Atoi(limit)
	if limit == "" {
		limit_int = 5
	}
	if page == "" {
		page_int = 1
	}
	if page_int > 0 {
		offset = (page_int - 1) * limit_int
	} else {
		offset = 0
	}

	count_data := r.Get_Count_Data(user, time_schedule, id_booking)

	if count_data <= 0 {
		metas.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			metas.Next = ""
		} else {
			metas.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		metas.Prev = ""
	} else {
		metas.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		metas.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		metas.Last_page = ""
	}

	if count_data != 0 {
		metas.Total_data = strconv.Itoa(count_data)
	} else {
		metas.Total_data = ""
	}

	if time_schedule == "" {
		time_schedule = ""
	} else {
		time_schedule = fmt.Sprintf(` AND id_time_schedule='%s'`, time_schedule)
	}

	if id_booking == "" {
		id_booking = ""
	} else {
		id_booking = fmt.Sprintf(` AND id_booking='%s'`, id_booking)
	}

	if user == "" {
		user = ""
	} else {
		user = fmt.Sprintf(` AND id_user='%s'`, user)
	}

	q := fmt.Sprintf(`SELECT * FROM public.bookings WHERE TRUE %s %s %s LIMIT %d OFFSET %d`, user, time_schedule, id_booking, limit_int, offset)

	r.Select(&datas, r.Rebind(q))
	if len(datas) == 0 {
		return nil, errors.New("data not found.")
	}
	for i := range datas {
		datas[i].Seats = strings.Replace(datas[i].Seats, " ", "", -1)

		var list_times_schedules []models.Times_Scheduless
		times_schedules_data := models.Times_Scheduless{}
		q := fmt.Sprintf(`select ts.id_time_schedule, ts.time_schedule, s.regency,s.price, s.set_date,m.id_movie,m.title,m.image as image_movie,m.release_date,p.name_premier, p.image as image_premier,p.count_row_seat,p.count_col_seat
	from times_schedules ts 
	left join schedules s on ts.id_schedule = s.id_schedule 
	left join movies m on s.id_movie = m.id_movie
	left join premiers p on s.id_premier = p.id_premier WHERE TRUE AND ts.id_time_schedule='%s'`, datas[i].Id_time_schedule)
		rows, err := r.Queryx(r.Rebind(q))
		if err != nil {
			log.Fatalln(err)
		}
		for rows.Next() {
			var list_movies_genres []models.Genres
			err := rows.StructScan(&times_schedules_data)
			if err != nil {
				log.Fatalln(err)
			}
			times_schedules_data.Release_date = &strings.Split(*times_schedules_data.Release_date, "T")[0]
			times_schedules_data.Set_date = strings.Split(times_schedules_data.Set_date, "T")[0]
			times_schedules_data.Time_schedule = strings.Split(strings.Split(times_schedules_data.Time_schedule, "T")[1], "Z")[0]
			rows1, _ := r.Queryx("select g.* from movies_genres mg left join genres g on mg.id_genre=g.id_genre  where mg.id_movie=$1", times_schedules_data.Id_movie)
			for rows1.Next() {
				var movies_genres models.Genres
				err1 := rows1.Scan(&movies_genres.Id_genre, &movies_genres.Name_genre, &movies_genres.Created_at, &movies_genres.Updated_at)
				if err1 != nil {
					log.Fatalln(err1)
				}
				// fmt.Println(movies_genres)
				list_movies_genres = append(list_movies_genres, movies_genres)
			}
			times_schedules_data.Genres = list_movies_genres

			// fmt.Println(times_schedules_data)
			list_times_schedules = append(list_times_schedules, times_schedules_data)
		}
		datas[i].Schedule = list_times_schedules
		rows.Close()

	}
	return &config.Result{Data: datas, Meta: metas}, nil
}

func (r *Repo_Bookings) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.bookings WHERE id_booking=$1", id)
	return count_data
}

func (r *Repo_Bookings) Get_Price(id string) int {
	var price int
	r.Get(&price, `SELECT s2.price FROM public.bookings s 
	left join times_schedules ts on s.id_time_schedule =ts.id_time_schedule 
	left join schedules s2 on ts.id_schedule =s2.id_schedule
	WHERE s.id_time_schedule=$1`, id)
	return price
}

func (r *Repo_Bookings) Get_Count_Data(user string, time_schedule string, id_booking string) int {
	var id int

	if user == "" {
		user = ""
	} else {
		user = fmt.Sprintf(` AND id_user='%s'`, user)
	}

	if time_schedule == "" {
		time_schedule = ""
	} else {
		time_schedule = fmt.Sprintf(` AND id_time_schedule='%s'`, time_schedule)
	}
	if id_booking == "" {
		id_booking = ""
	} else {
		id_booking = fmt.Sprintf(` AND id_booking='%s'`, id_booking)
	}

	q := fmt.Sprintf(`SELECT count(*) FROM public.bookings WHERE TRUE %s %s %s`, time_schedule, user, id_booking)
	r.Get(&id, r.Rebind(q))
	return id
}

func doPickid(wg *sync.WaitGroup, r Repo_Bookings) {
	defer wg.Done()
	values := ""

	tx := r.MustBegin()
	tx.Get(&values, "select gen_random_uuid()")
	fmt.Println(values)
}

func (r *Repo_Bookings) Insert_Data(data *models.Bookingsset) (string, string, error) {

	tx := r.MustBegin()
	var new_id string

	tx.Get(&new_id, "select gen_random_uuid()")

	data.Id_booking = new_id
	query := `INSERT INTO public.bookings(
		id_booking,
			id_user,
			id_time_schedule,
			seats,
			total
		)VALUES(
			:id_booking,
			:id_user,
			:id_time_schedule,
			:seats,
			:total
		);`
	_, err := tx.NamedExec(query, data)
	if err != nil {
		return "", "", err
	}
	tx.Commit()
	return new_id, "booking movie successful", nil
}
