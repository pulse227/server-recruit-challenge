package mysqldb

import (
	"database/sql"
)

func NewAlbumRepository(db *sql.DB) *albumRepository {
	return &albumRepository{
		db: db,
	}
}

type albumRepository struct {
	db *sql.DB
}
