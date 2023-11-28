package ad

import (
	"errors"
	"fmt"
	"time"

	lerrors "ddd/example/app/pkg/errors"
)

var (
	Err                     = errors.New("ad")
	ErrPublish              = fmt.Errorf("%w: publish", Err)
	ErrArchive              = fmt.Errorf("%w: achive", Err)
	ErrReserve              = fmt.Errorf("%w: reserve", Err)
	ErrReserveInvalidStatus = fmt.Errorf("%w: reserve: invalid Status", Err)
)

type Ad struct {
	ID          AdID
	UserID      UserID
	CategoryID  CategoryID
	Status      Status
	Title       Title
	Description Description
	Attrs       Attrs
	Quantity    Quantity
	Images      []*Image
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt *time.Time
	ArchivedAt  *time.Time
	/* и другие поля*/
}

func newAd(
	id AdID,
	userID UserID,
	categoryID CategoryID,
	title Title,
	description Description,
	now time.Time,
) (*Ad, error) {
	return &Ad{
		ID:          id,
		UserID:      userID,
		CategoryID:  categoryID,
		Status:      Draft,
		Title:       title,
		Description: description,
		Attrs:       emptyAttrs(),
		Quantity:    DefaultQuantity(),
		Images:      nil,
		CreatedAt:   now,
		UpdatedAt:   now,
		PublishedAt: nil,
		ArchivedAt:  nil,
	}, nil
}

func NewAdFromDB(
	idIn int,
	statusIn int,
	descriptionIn string,
	count Quantity,
	images []*Image,
	publishedAt *time.Time,
	archivedAt *time.Time,
) (*Ad, error) {
	adID, errs := lerrors.Check(NewAdID(idIn))(nil)
	status, errs := lerrors.Check(NewStatus(statusIn))(errs)
	desc, errs := lerrors.Check(NewDescription(descriptionIn))(errs)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return &Ad{
		ID:          adID,
		Status:      status,
		Description: desc,
		Quantity:    count,
		Images:      images,
		PublishedAt: publishedAt,
		ArchivedAt:  archivedAt,
	}, nil
}

func (a *Ad) Publish(now time.Time) error {
	if !a.Status.Can(Published) {
		return ErrPublish
	}

	a.Status = Published
	a.PublishedAt = &now

	return nil
}

func (a *Ad) Archive(now time.Time) error {
	if !a.Status.Can(Archived) {
		return ErrArchive
	}

	a.Status = Archived
	a.ArchivedAt = &now

	return nil
}

// Внутренние методы
func (a *Ad) Reorder(imageID ImageID, order int) error {
	for _, image := range a.Images {
		image.reorder(imageID, order)
	}

	return nil
}

func (a *Ad) Reserve(reserve int) error {
	if !a.Status.Is(Ban, Deleted) {
		return ErrReserveInvalidStatus
	}

	quantity, err := a.Quantity.Reserve(reserve)
	if err != nil {
		return lerrors.Nested(ErrReserve, err)
	}

	a.Quantity = quantity

	return nil
}

func (a *Ad) AddImage(imageID ImageID, content []byte) error {
	image, err := NewImage(a.ID, imageID, content, len(a.Images))
	if err != nil {
		return err
	}

	a.Images = append(a.Images, image)

	return nil
}

// IsEqual две сущности по ID
func (a *Ad) IsEqual(another *Ad) bool {
	return a.ID == another.ID
}

var (
	ErrUserID        = fmt.Errorf("%w: UserID", Err)
	ErrUserIDInvalid = fmt.Errorf("%w: invalid", ErrUserID)
)

type UserID int

func NewUserID[In ~int](in In) (UserID, error) {
	if in < 0 {
		return 0, ErrUserIDInvalid
	}

	return UserID(in), nil
}

func NewUserIDAssertBool[In ~int](in In) (UserID, error) {
	if assertPositiveBool(in) {
		return 0, ErrUserIDInvalid
	}

	return UserID(in), nil
}

func NewUserIDAssert[In ~int](in In) (UserID, error) {
	if err := assertPositive(in); err != nil {
		return 0, ErrUserIDInvalid
	}

	return UserID(in), nil
}

func NewUserIDAssertErrorf[In ~int](in In) (_ UserID, err error) {
	defer func() {
		if err != nil {
			err = lerrors.Nested(ErrUserID, err)
		}
	}()

	if err = assertPositive(in); err != nil {
		return 0, err
		// Если defer кажется громоздким или всего одна проверка, то можно оборачивать ее на месте.
		// Чтобы не пропустить ни одно оборачивание можно использовать линетр wrapcheck.
		// return 0, lerrors.Nested(ErrUserID, err)
	}

	return UserID(in), nil
}

var ErrCategoryIDInvalid = fmt.Errorf("%w: categoryID: invalid", Err)

type CategoryID int

var mpCategoryID = map[error]error{
	ErrNegative: ErrCategoryIDInvalid,
}

func NewCategoryID[In ~int](in In) (CategoryID, error) {
	if err := assertPositive(in); err != nil {
		return 0, mapError(err, mpCategoryID)
	}

	return CategoryID(in), nil
}

var ErrTitle = fmt.Errorf("%w: Title: invalid", Err)

type Title string

func NewTitle(in string) (Title, error) {
	if err := assertLength(in, 10, 512); err != nil {
		return "", fmt.Errorf("%w: %v", ErrTitle, err)
	}

	return Title(in), nil
}

var ErrDescription = fmt.Errorf("%w: Description: invalid", Err)

type Description string

func NewDescription(in string) (Description, error) {
	if err := assertLength(in, 10, 512); err != nil {
		return "", fmt.Errorf("%w: %v", ErrDescription, err)
	}

	return Description(in), nil
}

var (
	ErrAdID         = fmt.Errorf("%w: ID", Err)
	ErrAdIDNegative = fmt.Errorf("%w: negative", ErrAdID)
)

type AdID int

func NewAdID(id int) (AdID, error) {
	if id <= 0 {
		return 0, ErrAdIDNegative
	}

	return AdID(id), nil
}

var sequenceID = AdID(1)

func NextAdID() AdID {
	sequenceID++

	return sequenceID
}

var ErrStatus = fmt.Errorf("%w: Status", Err)
