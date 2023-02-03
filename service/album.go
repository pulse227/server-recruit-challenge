package service

import (
	"context"

	"github.com/Lupusdog/server-recruit-challenge-2024/model"
	"github.com/Lupusdog/server-recruit-challenge-2024/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.AlbumWithSinger, error)
	GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.AlbumWithSinger, error)
	PostAlbumService(ctx context.Context, album *model.Album) error
	DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error 
}

type albumService struct {
	albumRepository repository.AlbumRepository
}

var _ AlbumService = (*albumService)(nil)

func NewAlbumService(albumRepository repository.AlbumRepository) *albumService {
	return &albumService{albumRepository: albumRepository}
}

func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.AlbumWithSinger, error) {
	albums, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.AlbumWithSinger, error) {
	album, err := s.albumRepository.Get(ctx, albumID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (s *albumService) PostAlbumService(ctx context.Context, album *model.Album) error {
	if err := s.albumRepository.Add(ctx, album); err != nil {
		return err
	}
	return nil
}

func (s *albumService) DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error {
	if err := s.albumRepository.Delete(ctx, albumID); err != nil {
		return err
	}
	return nil
}

