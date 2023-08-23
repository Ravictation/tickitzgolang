package repositories

import (
	"errors"
	"math"
	"strconv"

	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repo_Directors struct {
	*sqlx.DB
}

func NewDirectors(db *sqlx.DB) *Repo_Directors {
	return &Repo_Directors{db}
}

func (r *Repo_Directors) Get_Data(data *models.Directors, page string, limit string) (*config.Result, error) {
	sizes_data := []models.Directors{}
	var meta_size config.Metas

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
		meta_size.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_size.Next = ""
		} else {
			meta_size.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_size.Prev = ""
	} else {
		meta_size.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_size.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_size.Last_page = ""
	}

	if count_data != 0 {
		meta_size.Total_data = strconv.Itoa(count_data)
	} else {
		meta_size.Total_data = ""
	}

	r.Select(&sizes_data, `SELECT * FROM public.directors LIMIT $1 OFFSET $2`, limit_int, offset)
	if len(sizes_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: sizes_data, Meta: meta_size}, nil
}

func (r *Repo_Directors) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.directors WHERE id_director=$1", id)
	return count_data
}

func (r *Repo_Directors) Get_Count_Data() int {
	var id int
	r.Get(&id, "SELECT count(*) FROM public.directors")
	return id
}

func (r *Repo_Directors) Insert_Data(data *models.Directors) (string, error) {
	query := `INSERT INTO public.directors(
			name_director
		)VALUES(
			:name_director
		);`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "add director data successful", nil
}
func (r *Repo_Directors) Update_Data(data *models.Directors) (string, error) {
	query := `UPDATE public.directors SET
			name_director=:name_director,
			update_at=now()
			WHERE id_director=:id_director;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update director data successful", nil
}
func (r *Repo_Directors) Delete_Data(data *models.Directors) (string, error) {
	query := `DELETE FROM public.directors WHERE id_director=:id_director;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete director data successful", nil
}
