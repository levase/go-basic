package bins

import (
	"errors"
	"time"
)

var (
	ErrBinIDEmpty   = errors.New("bin ID cannot be empty")
	ErrBinNameEmpty = errors.New("bin name cannot be empty")
	ErrBinIDExists  = errors.New("bin with this ID already exists")
)

type BinList struct {
	Bins []*Bin `json:"bins"`
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make([]*Bin, 0),
	}
}

func (bl *BinList) Add(bin *Bin) error {
	if err := bin.Validate(); err != nil {
		return err
	}

	if bl.ContainsID(bin.ID) {
		return ErrBinIDExists
	}

	bl.Bins = append(bl.Bins, bin)
	return nil
}

func (bl *BinList) ContainsID(id string) bool {
	for _, bin := range bl.Bins {
		if bin.ID == id {
			return true
		}
	}
	return false
}

type Bin struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	Private   bool      `json:"private"`
}

func NewBin(id, name string, private bool) (*Bin, error) {
	bin := &Bin{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		Private:   private,
	}
	return bin, bin.Validate()
}

func (b *Bin) Validate() error {
	if b.ID == "" {
		return ErrBinIDEmpty
	}
	if b.Name == "" {
		return ErrBinNameEmpty
	}
	return nil
}
