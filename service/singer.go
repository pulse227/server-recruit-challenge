package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type SingerService interface {
	GetSingerListService(ctx context.Context) ([]*model.Singer, error)
	GetSingerService(ctx context.Context, singerID model.SingerID) (*model.Singer, error)
	PostSingerService(ctx context.Context, singer *model.Singer) error
	DeleteSingerService(ctx context.Context, singerID model.SingerID) error
}

type singerService struct {
	singerRepository repository.SingerRepository
}

var _ SingerService = (*singerService)(nil)

func NewSingerService(singerRepository repository.SingerRepository) *singerService {
	return &singerService{singerRepository: singerRepository}
}

func (s *singerService) GetSingerListService(ctx context.Context) ([]*model.Singer, error) {
	singers, err := s.singerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return singers, nil
}

func (s *singerService) GetSingerService(ctx context.Context, singerID model.SingerID) (*model.Singer, error) {
	singer, err := s.singerRepository.Get(ctx, singerID)
	if err != nil {
		return nil, err
	}
	return singer, nil
}

func (s *singerService) PostSingerService(ctx context.Context, singer *model.Singer) error {
	if err := singer.Validate(); err != nil {
		return err
	}

	if err := s.singerRepository.Add(ctx, singer); err != nil {
		return err
	}
	return nil
}

func (s *singerService) DeleteSingerService(ctx context.Context, singerID model.SingerID) error {
	if err := s.singerRepository.Delete(ctx, singerID); err != nil {
		return err
	}
	return nil
}
