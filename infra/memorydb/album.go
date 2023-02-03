package memorydb

import (
	"context"
	"errors"
	"sync"

	"github.com/Lupusdog/server-recruit-challenge-2024/model"
	"github.com/Lupusdog/server-recruit-challenge-2024/repository"
)

type albumRepository struct {
	sync.RWMutex
	albumMap map[model.AlbumID]*model.Album // キーが AlbumID、値が model.Album のマップ
	singerRepo repository.SingerRepository
}

var _ repository.SingerRepository = (*singerRepository)(nil) 
var _ repository.AlbumRepository = (*albumRepository)(nil)

func NewAlbumRepository(singerRepo *singerRepository) *albumRepository {
	var initMap = map[model.AlbumID]*model.Album{
		1: {ID: 1, Title: "Alice's 1st Album", SingerID: 1},
		2: {ID: 2, Title: "Alice's 2nd Album", SingerID: 1},
		3: {ID: 3, Title: "Bella's 1st Album", SingerID: 2},
	}

	return &albumRepository{
		albumMap: initMap,
		singerRepo: singerRepo,
	}
}

func (r *albumRepository) GetAll(ctx context.Context) ([]*model.AlbumWithSinger, error) {
	r.RLock()
	defer r.RUnlock()

	albums := make([]*model.AlbumWithSinger, 0, len(r.albumMap))
	for _, a := range r.albumMap {
		singerData, err := r.GetSingerInfo(ctx, a.SingerID, r.singerRepo)
		if err != nil {
			return nil, err
		}
		albumData := &model.AlbumWithSinger{
			ID: a.ID,
			Title: a.Title,
			Singer: *singerData,
		}
		albums = append(albums, albumData)
	}
	return albums, nil
}

func (r *albumRepository) Get(ctx context.Context, id model.AlbumID) (*model.AlbumWithSinger, error) {
	r.RLock()
	defer r.RUnlock()

	a, ok := r.albumMap[id]
	if !ok {
		return nil, errors.New("not found")
	}
	singerData, err := r.GetSingerInfo(ctx, a.SingerID, r.singerRepo)
	if err != nil {
		return nil, err
	}
	albumData := &model.AlbumWithSinger{
		ID: a.ID,
		Title: a.Title,
		Singer: *singerData,
	}

	return albumData, nil
}

func (r *albumRepository) Add(ctx context.Context, album *model.Album) error {
	r.Lock()
	r.albumMap[album.ID] = album
	r.Unlock()
	return nil
}

func (r *albumRepository) Delete(ctx context.Context, id model.AlbumID) error {
	r.Lock()
	delete(r.albumMap, id)
	r.Unlock()
	return nil
}

//album情報にsinger情報を追加するために実行される
func (r *albumRepository) GetSingerInfo(ctx context.Context, singerID model.SingerID, singerRepo repository.SingerRepository) (*model.Singer, error) {
	r.RLock()
	defer r.RUnlock()
	// singerRepo := NewSingerRepository()
	singer, err := singerRepo.Get(ctx, singerID)
	if err != nil {
		return nil, errors.New("cannnot get singer information")
	}
	return singer, nil

}

