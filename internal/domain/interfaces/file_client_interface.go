package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type FileClientUsecase interface {
	ListObjects(req *presenter.FileListReq) ([]presenter.FileListRes, int, error)
	DeleteObjects(ids []uint) error
}
type FileClientRepository interface {
	ListObjects(req *presenter.FileListReq) ([]presenter.FileListRes, float64, error)
	GetObjectsByID(id []uint) ([]models.File, error)
	GetURLsByID(id []uint) ([]string, error)
	DeleteObjects(id []uint) error
}
