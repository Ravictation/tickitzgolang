package repositories

import (
	"errors"
	"math"
	"strconv"

	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repo_Genres struct {
	*sqlx.DB
}

func NewGenres(db *sqlx.DB) *Repo_Genres {
	return &Repo_Genres{db}
}

func (r *Repo_Genres) Get_Data(data *models.Genres, page string, limit string) (*config.Result, error) {
	datas := []models.Genres{}
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

	count_data := r.Get_Count_Data()

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

	r.Select(&datas, `SELECT * FROM public.genres LIMIT $1 OFFSET $2`, limit_int, offset)
	if len(datas) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: datas, Meta: metas}, nil
}

func (r *Repo_Genres) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.genres WHERE id_genre=$1", id)
	return count_data
}

func (r *Repo_Genres) Get_Count_Data() int {
	var id int
	r.Get(&id, "SELECT count(*) FROM public.genres")
	return id
}

func (r *Repo_Genres) Insert_Data(data *models.Genres) (string, error) {
	query := `INSERT INTO public.genres(
			name_genre
		)VALUES(
			:name_genre
		);`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "add genre data successful", nil
}
func (r *Repo_Genres) Update_Data(data *models.Genres) (string, error) {
	query := `UPDATE public.genres SET
			name_genre=:name_genre,
			updated_at=now()
			WHERE id_genre=:id_genre;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update genre data successful", nil
}
func (r *Repo_Genres) Delete_Data(data *models.Genres) (string, error) {
	query := `DELETE FROM public.genres WHERE id_genre=:id_genre;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete genre data successful", nil
}
