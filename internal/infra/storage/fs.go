package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Storage interface {
	Store(ctx context.Context, value io.ReadCloser) error
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

func (f *FileSystemStorage) Store(ctx context.Context, value io.ReadCloser) error {
	file, err := os.Open(filepath.Join(f.path, "a.png"))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, value)

	return err
}

func (f *FileSystemStorage) FindByID(ctx context.Context, id uuid.UUID) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}
