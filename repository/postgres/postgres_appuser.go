package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"postingapp/model"
	"postingapp/repository"
	"postingapp/utils"
)

type IAppUserRepository struct {
	Ctx context.Context
}

func (rep IAppUserRepository) Store(user model.AppUser) error {
	q := `INSERT INTO app_user(name, password, avatar, create_at)
    VALUES($1, $2, $3, $4)`

	db := repository.GetConnection()
	defer db.Close()

	st, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer st.Close()
	r, err := st.Exec(utils.GetNUllString(user.Name),
		utils.GetNUllString(user.Password),
		utils.GetNUllString(user.Avatar),
		utils.GetNUllTime(user.CreateAt),
	)
	if err != nil {
		return err
	}
	rows, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("No rows affected")
	}
	return nil
}

func (rep IAppUserRepository) GetAll() ([]model.AppUser, error) {
	q := `SELECT id, name, password, avatar, create_at
         FROM app_user`
	db := repository.GetConnection()
	defer db.Close()

	na := sql.NullString{}
	pa := sql.NullString{}
	av := sql.NullString{}
	cr := pq.NullTime{}

	rows, err := db.Query(q)
	if err != nil {
		return []model.AppUser{}, err
	}
	defer rows.Close()
	al := []model.AppUser{}

	for rows.Next() {
		a := model.AppUser{}
		err := rows.Scan(&a.Id,
			&na,
			&pa,
			&av,
			&cr,
		)
		if err != nil {
			return []model.AppUser{}, err
		}
		a.Name = na.String
		a.Password = pa.String
		a.Avatar = av.String
		a.CreateAt = cr.Time
		al = append(al, a)
	}
	return al, nil
}

func (rep IAppUserRepository) GetById(id int) (model.AppUser, error) {
	q := `SELECT id, name, password, avatar, create_at
         FROM app_user WHERE id = $1`
	db := repository.GetConnection()
	defer db.Close()
	na := sql.NullString{}
	pa := sql.NullString{}
	av := sql.NullString{}
	cr := pq.NullTime{}
	m := model.AppUser{}

	err := db.QueryRowContext(rep.Ctx, q, id).Scan(
		&m.Id,
		&na,
		&pa,
		&av,
		&cr)
	if err != nil {
		return m, err
	}
	m.Name = na.String
	m.Password = pa.String
	m.Avatar = pa.String
	m.CreateAt = cr.Time

	return m, nil
}

func (rep IAppUserRepository) Delete(id int) error {
	q := `DELETE FROM app_user WHERE id = $1`
	db := repository.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rows, err := r.RowsAffected()
	if rows != 1 {
		return errors.New("Error: No rows affected")
	}
	return nil
}

func (rep IAppUserRepository) Update(appUser model.AppUser) error {
	q := `UPDATE app_user
          SET name = $1, password = $2, avatar = $3
          WHERE id = $4`

	db := repository.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(appUser.Name, appUser.Password, appUser.Avatar, appUser.Id)
	if err != nil {
		return err
	}

	rows, err := r.RowsAffected()
	if rows != 1 {
		return errors.New("Error: No rows affected")
	}
	return nil
}
