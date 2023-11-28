package ad

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type AdService interface {
	CreateDraft(
		ctx context.Context,
		categoryID CategoryID,
		userID UserID,
		title Title,
		description Description,
	) (*Ad, error)
	// Если много параметров, может быть читабельнее их положить в DTO
	CreateDraftWithDTO(ctx context.Context, adDTO AdDTO) (*Ad, error)
}

type AdDTO struct{}

// AdRepo - PersistenceOriented репозиторий
type AdRepo interface {
	GetByID(ctx context.Context, ID AdID) (*Ad, error)
	// Create(ctx context.Context, ad *Ad) error // Необязательный метод
	Save(ctx context.Context, ad *Ad) error
	Delete(ctx context.Context, ad *Ad) error

	// NextID is not default part of repo interface
	NextID(ctx context.Context) (AdID, error)

	// SaveTr - Проверка работы depguard
	SaveTr(ctx context.Context, a *Ad, tx *sqlx.Tx) (err error)
}

type AdRepoPersistenceOrientedWithEntity interface {
	AdRepo

	SaveImage(ctx context.Context, image *Image, images ...*Image) error
	DeleteImage(ctx context.Context, image *Image, images ...*Image) error
}

type AdRepoCollectionOriented interface {
	GetByID(ctx context.Context, ID AdID) (*Ad, error)
	Add(ctx context.Context, ad *Ad) error
	Delete(ctx context.Context, ad *Ad) error
}

type UserRepo interface {
	GetByID(ctx context.Context, ID UserID) (User, error)
}
type User struct {
	id       UserID
	isBanned bool
}

func (u User) ID() UserID {
	return u.id
}

func (u User) IsBanned() bool {
	return u.isBanned
}

type CategoryRepo interface {
	GetByID(ctx context.Context, ID CategoryID) (Category, error)
}

type Category struct {
	id               CategoryID
	canManyAvailable bool
}

func (c Category) ID() CategoryID {
	return c.id
}

func (c Category) CanManyAvailable() bool {
	return c.canManyAvailable
}
