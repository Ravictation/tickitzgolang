package repositories

import (
	"errors"

	"github.com/Ravictation/tickitzgolang/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	query := `INSERT INTO public.users ( 
				email_user, 
				password,
				role) 
				VALUES(
					:email_user,
					:password,
					:role
				);`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data user created", nil
}

func (r *RepoUser) UpdateUser(data *models.User) (string, error) {
	query := `UPDATE public.users SET
				password = COALESCE(NULLIF(:password, ''), password),
				first_name = COALESCE(NULLIF(:first_name, ''), first_name),
				last_name = COALESCE(NULLIF(:last_name, ''), last_name),
				email_user = COALESCE(NULLIF(:email_user, ''), email_user),
				phone_number = COALESCE(NULLIF(:phone_number, ''), phone_number),
				updated_at = now()
			WHERE id_user = :id_user`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data user has been updated", nil
}

func (r *RepoUser) UpdateImageUser(data *models.User) (string, error) {
	query := `UPDATE public.users SET
				image_user = COALESCE(:image_user, image_user),
				updated_at = now()
			WHERE id_user = :id_user`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data user has been updated", nil
}

func (r *RepoUser) GetUser(data *models.User) (*models.User, error) {
	query := `SELECT email_user, phone_number, image_user,role FROM public.user WHERE username=$1;`
	var userModel models.User
	err := r.Get(&userModel, query, data.Email_user)
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}

func (r *RepoUser) GetAllUser(data *models.User) ([]models.User, error) {

	var users []models.User
	query := "SELECT username, image_user, phone_number, email_user,role FROM public.user"
	err := r.Select(&users, query)

	if err != nil {
		return nil, err
	}

	return users, err
}

func (r *RepoUser) DeleteUser(data *models.User) (string, error) {
	query := `DELETE FROM public.users WHERE Email = $1;`

	_, err := r.Exec(query, data.Email_user)
	if err != nil {
		return "", err
	}

	return "1 user has been Deleted", nil
}

func (r *RepoUser) GetAuthData(user string) (*models.User, error) {
	var result models.User
	q := `SELECT id_user, email_user, "role", "password" FROM public."users" WHERE email_user = ?`

	if err := r.Get(&result, r.Rebind(q), user); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil
}
