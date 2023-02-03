package repository

import (
	"context"

	"github.com/Lupusdog/server-recruit-challenge-2024/model"
)

type AlbumRepository interface {
	GetAll(ctx context.Context) ([]*model.AlbumWithSinger, error)
	Get(ctx context.Context, id model.AlbumID) (*model.AlbumWithSinger, error)
	Add(ctx context.Context, album *model.Album) error
	Delete(ctx context.Context, id model.AlbumID) error
	GetSingerInfo(ctx context.Context, singerID model.SingerID, singerRepo SingerRepository) (*model.Singer, error)
}