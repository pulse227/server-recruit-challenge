package mysqldb

import (
	"context"
	"database/sql"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

func NewSingerRepository(db *sql.DB) *singerRepository {
	return &singerRepository{
		db: db,
	}
}

type singerRepository struct {
	db *sql.DB
}

var _ repository.SingerRepository = (*singerRepository)(nil)

func (r *singerRepository) GetAll(ctx context.Context) ([]*model.Singer, error) {
	singers := []*model.Singer{}
	query := "SELECT id, name FROM singers ORDER BY id ASC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		singer := &model.Singer{}
		if err := rows.Scan(&singer.ID, &singer.Name); err != nil {
			return nil, err
		}
		if singer.ID != 0 {
			singers = append(singers, singer)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return singers, nil
}

func (r *singerRepository) Get(ctx context.Context, id model.SingerID) (*model.Singer, error) {
	singer := &model.Singer{}
	query := "SELECT id, name FROM singers WHERE id = ?"
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&singer.ID, &singer.Name); err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if singer.ID == 0 {
		return nil, model.ErrNotFound
	}
	return singer, nil
}

func (r *singerRepository) Add(ctx context.Context, singer *model.Singer) error {
	query := "INSERT INTO singers (id, name) VALUES (?, ?)"
	if _, err := r.db.ExecContext(ctx, query, singer.ID, singer.Name); err != nil {
		return err
	}
	return nil
}

func (r *singerRepository) Delete(ctx context.Context, id model.SingerID) error {
	query := "DELETE FROM singers WHERE id = ?"
	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return err
	}
	return nil
}
