package storage

import (
	"LinkStrashBot/lib/e"
	"crypto/sha1"
	"fmt"
	"io"
	"time"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(UserName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}
type Page struct {
	URL      string
	UserName string
	Created  time.Time
}

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("cant do hash", err)
	}
	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("cant do hash", err)

	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
