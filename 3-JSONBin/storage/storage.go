package storage

import (
	"encoding/json"
	"fmt"
	"homework/JSONBin/bins"
)

type Storager interface {
	Read() ([]byte, error)
	Write(any) error
}

type Storage struct {
	storage Storager
}

func NewStorage(s Storager) *Storage {
	return &Storage{storage: s}
}

func (s *Storage) LoadBins() (*bins.BinList, error) {
	data, err := s.storage.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read from storage: %w", err)
	}

	if len(data) == 0 {
		return bins.NewBinList(), nil
	}

	var binList bins.BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal bins: %w", err)
	}

	return &binList, nil
}

func (s *Storage) SaveBins(binList *bins.BinList) error {
	return s.storage.Write(binList)
}

func (s *Storage) AddBin(bin *bins.Bin) error {
	binList, err := s.LoadBins()
	if err != nil {
		return err
	}

	if err := binList.Add(bin); err != nil {
		return err
	}

	return s.SaveBins(binList)
}
