package storage

import (
	"encoding/json"
	"fmt"
	"homework/JSONBin/bins"
	"os"
	"path/filepath"
)

const (
	FilePermissions = 0644
	DirPermissions  = 0755
)

type Storage struct {
	path string
}

func NewStorage(path string) (*Storage, error) {
	if err := os.MkdirAll(filepath.Dir(path), DirPermissions); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}
	return &Storage{path: path}, nil
}

func (storage *Storage) SaveJSON(binList *bins.BinList) error {
	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal bin list: %w", err)
	}

	if err = os.WriteFile(storage.path, data, FilePermissions); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", storage.path, err)
	}
	return nil
}

func (s *Storage) ReadJSON() (bins.BinList, error) {
	data, err := os.ReadFile(s.path)
	if os.IsNotExist(err) {
		return bins.BinList{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", s.path, err)
	}

	var binList bins.BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}
	return binList, nil
}
