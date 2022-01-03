package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Storage interface {
	Store(ctx context.Context, id uuid.UUID, value io.ReadCloser) error
	FindByID(ctx context.Context, id uuid.UUID) (io.ReadCloser, error)
}

type FileSystemStorage struct {
	path string
}

func NewFileSystemStorage(path string) (*FileSystemStorage, error) {
	return &FileSystemStorage{
		path: path,
	}, nil
}

func (f *FileSystemStorage) Store(_ context.Context, id uuid.UUID, value io.ReadCloser) error {
	file, err := os.Create(filepath.Join(f.path, id.String()+".png"))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, value)

	return err
}

func (f *FileSystemStorage) FindByID(_ context.Context, id uuid.UUID) (io.ReadCloser, error) {
	return os.Open(filepath.Join(f.path, id.String()+".png"))
}
