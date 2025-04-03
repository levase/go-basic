package bins

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

const (
	MaxNameLength    = 100
	NameRegexPattern = `^[a-zA-Z0-9](?:[a-zA-Z0-9 _-]*[a-zA-Z0-9])?$`
)

type BinList []*Bin

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin(id, name string, private bool) (*Bin, error) {
	now := time.Now()
	bin := &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: now,
		Name:      name,
	}

	if err := bin.validate(); err != nil {
		return nil, fmt.Errorf("invalid bin: %w", err)
	}
	return bin, nil
}

func (b *Bin) validate() error {
	if b.ID == "" {
		return errors.New("ID cannot be empty")
	}

	if b.Name == "" {
		return errors.New("name cannot be empty")
	}

	if len(b.Name) > MaxNameLength {
		return fmt.Errorf("name is too long (max %d characters)", MaxNameLength)
	}

	if !regexp.MustCompile(NameRegexPattern).MatchString(b.Name) {
		return errors.New("name contains invalid characters")
	}

	if b.CreatedAt.After(time.Now()) {
		return errors.New("createdAt cannot be in the future")
	}
	return nil
}

func (b *Bin) String() string {
	return fmt.Sprintf("{id:%s private:%v createdAt:%s name:%s}",
		b.ID, b.Private, b.CreatedAt, b.Name)
}
