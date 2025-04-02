package bins

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt string
	name      string
}

func NewBin(id, name string, private bool) (*Bin, error) {
	bin := &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now().Format(time.DateTime),
		name:      name,
	}

	if err := bin.validate(); err != nil {
		return nil, err
	}
	return bin, nil
}

func (b *Bin) validate() error {
	if err := b.validateID(); err != nil {
		return err
	}
	if err := b.validateName(); err != nil {
		return err
	}
	if err := b.validateCreatedAt(); err != nil {
		return err
	}
	return nil
}

func (b *Bin) validateID() error {
	if b.id == "" {
		return errors.New("ID cannot be empty")
	}
	return nil
}

func (b *Bin) validateName() error {
	if b.name == "" {
		return errors.New("name cannot be empty")
	}

	if len(b.name) > 100 {
		return errors.New("name is too long (max 100 characters)")
	}

	validNameRegex := regexp.MustCompile(`^[a-zA-Z0-9 _-]+$`)
	if !validNameRegex.MatchString(b.name) {
		return errors.New("name contains invalid characters")
	}
	return nil
}

func (b *Bin) validateCreatedAt() error {
	if b.createdAt == "" {
		return errors.New("createdAt cannot be empty")
	}

	_, err := time.Parse(time.DateTime, b.createdAt)
	if err != nil {
		return errors.New("createdAt has invalid format")
	}
	return nil
}

func (b *Bin) String() string {
	return fmt.Sprintf("{id:%s private:%v createdAt:%s name:%s}",
		b.id, b.private, b.createdAt, b.name)
}

type BinList []*Bin
