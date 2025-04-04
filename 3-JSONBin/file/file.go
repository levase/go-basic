package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	FilePermissions = 0644
	DirPermissions  = 0755
)

var (
	ErrNotJSONFile   = errors.New("file is not a JSON file")
	ErrIsDirectory   = errors.New("path is a directory")
	ErrInvalidJSON   = errors.New("file contains invalid JSON")
	ErrEmptyFilename = errors.New("filename cannot be empty")
)

type FileStorage struct {
	filename string
}

func NewFileStorage(filename string) (*FileStorage, error) {
	if filename == "" {
		return nil, ErrEmptyFilename
	}

	if !isJSONFile(filename) {
		return nil, fmt.Errorf("%w: '%s'", ErrNotJSONFile, filename)
	}

	return &FileStorage{
		filename: filename,
	}, nil
}

func (fs *FileStorage) Read() ([]byte, error) {
	info, err := os.Stat(fs.filename)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to stat file %s: %w", fs.filename, err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("%w: '%s'", ErrIsDirectory, fs.filename)
	}

	data, err := os.ReadFile(fs.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fs.filename, err)
	}

	if len(data) == 0 {
		return nil, nil
	}

	if !json.Valid(data) {
		return nil, fmt.Errorf("%w: '%s'", ErrInvalidJSON, fs.filename)
	}
	return data, nil
}

func (fs *FileStorage) Write(v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(fs.filename), DirPermissions); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	return os.WriteFile(fs.filename, data, FilePermissions)
}

func isJSONFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	base := filepath.Base(filename)
	return ext == ".json" && base != ".json"
}
