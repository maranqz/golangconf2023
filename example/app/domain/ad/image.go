package ad

import (
	"errors"
	"fmt"
	"time"

	lerrors "ddd/example/app/pkg/errors"
)

var (
	ErrImageID = fmt.Errorf("%w: imageID", Err)
)

type ImageID int

func NewImageID(in int) (ImageID, error) {
	if err := assertPositive(in); err != nil {
		return 0, lerrors.Nested(ErrImageID, err)
	}

	return ImageID(in), nil

}

// Image может быть как Value Object, так или Aggregate.
// Image делать сущностью скорее не целесообразно, хотя я показывал такой пример в презентации 😅.
// Если с Image работает только Ad, то это Value Object.
// Если c Image могут работать внешние системы или другой Ограниченный Контекст, то это Aggregate.
// Выносить Image в отдельный Ограниченный Контекст не нужно, если она не выделиться во что-то общее, которое можно использовать в разных сервисах.
type Image struct {
	AdID      AdID
	ID        ImageID
	Content   []byte
	Order     int
	CreatedAt time.Time
}

func NewImage(adID AdID, id ImageID, content []byte, order int) (*Image, error) {
	return &Image{
		AdID:      adID,
		ID:        id,
		Content:   content,
		Order:     order,
		CreatedAt: time.Now(),
	}, nil
}

func NewImageFromDB(adIDIn int, idIn int, content []byte, order int, createdAt time.Time) (*Image, error) {
	adID, err1 := NewAdID(adIDIn)
	id, err2 := NewImageID(idIn)
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}

	return &Image{
		AdID:      adID,
		ID:        id,
		Content:   content,
		Order:     order,
		CreatedAt: createdAt,
	}, nil
}

// Даже если поля публичные, методы можно делать приватными,
// чтобы предотвратить нарушения целостности Aggregate.
func (i *Image) reorder(id ImageID, order int) {
	if i.ID == id {
		i.Order = order
	} else if i.Order >= order {
		i.Order++
	}
}
