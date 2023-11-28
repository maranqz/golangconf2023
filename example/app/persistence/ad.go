package persistence

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/jmoiron/sqlx"

	"ddd/example/app/domain/ad"
)

type adRepo struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
	trm    trm.Manager
}

func NewAd(db *sqlx.DB, c *trmsqlx.CtxGetter, trm trm.Manager) *adRepo {
	return &adRepo{
		db:     db,
		getter: c,
		trm:    trm,
	}
}

func (r *adRepo) GetByID(ctx context.Context, id ad.AdID) (*ad.Ad, error) {
	query := "SELECT * FROM ad WHERE id = ?;"
	row := adRow{}

	err := r.getter.DefaultTrOrDB(ctx, r.db).GetContext(ctx, &row, r.db.Rebind(query), id)
	if err != nil {
		return nil, err
	}

	queryImages := "SELECT * FROM image WHERE ad_id = ?;"
	var images []imageRow

	err = r.getter.DefaultTrOrDB(ctx, r.db).GetContext(ctx, &images, r.db.Rebind(queryImages), id)
	if err != nil {
		return nil, err
	}

	return r.toModel(row, images)
}

type Ad struct {
	/**/
}

func (r *adRepo) Save(ctx context.Context, a *ad.Ad) error {
	adRow, images := r.toRow(a)

	return r.trm.Do(ctx, func(ctx context.Context) error {
		query := `INSERT INTO ad (id, user_id, category_id, count /**/)
VALUES (:id, :user_id, :category_id, :count/**/)
ON CONFLICT (id)
    DO UPDATE SET COUNT = EXCLUDED.count /**/;`

		_, err := sqlx.NamedQueryContext(
			ctx,
			r.getter.DefaultTrOrDB(ctx, r.db),
			r.db.Rebind(query),
			adRow,
		)
		if err != nil {
			return err
		}

		// Удаляем существующие картинки, ну или высчитываем какие были удалены и делаем точечное изменение.
		// Для точечных изменений удобно можно использовать Domain Events, которые явно показывают что изменилось и по ним уже формировать запрос.
		_, err = sqlx.NamedQueryContext(
			ctx,
			r.getter.DefaultTrOrDB(ctx, r.db),
			r.db.Rebind(`DELETE FROM image WHERE ad_id = ?;`),
			a.ID,
		)
		if err != nil {
			return err
		}

		queryImages := `INSERT INTO image (id, ad_id, order /**/)
VALUES (:id, :ad_id, :order/**/)
ON CONFLICT (id)
    DO UPDATE SET ORDER = EXCLUDED.order /**/;`
		_, err = sqlx.NamedQueryContext(
			ctx,
			r.getter.DefaultTrOrDB(ctx, r.db),
			r.db.Rebind(queryImages),
			images,
		)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *adRepo) SaveTr(ctx context.Context, a *ad.Ad, tx *sqlx.Tx) (err error) {
	if tx == nil {
		tx, err := r.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		// Мы можем не вызывать Rollback().
		// Мы переводим Ad в adRow, тут не должно быть ошибок, т.к. мы уверены в своем домене.
		// А если нет ошибок в домене, то у нас могут быть только ошибки БД.
		// Значит, транзакция к этому времени уже закроется
		defer func() {
			if err != nil {
				return
			}

			err = tx.Commit()
		}()
	}

	/**/

	return nil
}

type adRow struct {
	ID          int         `migration:"id"`
	Status      int         `migration:"status"`
	Description string      `migration:"description"`
	Quantity    quantityRow `migration:"Quantity"`
	PublishedAt *time.Time  `migration:"published_at"`
	ArchivedAt  *time.Time  `migration:"archived_at"`
}

type imageRow struct {
	AdID      int
	ID        int
	Content   []byte
	Order     int
	CreatedAt time.Time
}

type quantityRow struct {
	Available int `migration:"available" json:"available"`
	Reserved  int `migration:"reserved" json:"reserved"`
}

func (c *quantityRow) Scan(src any) error {
	if src == nil {
		return nil
	}

	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("migration: count: value is not []byte")
	}

	err := json.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}

	return nil
}

func (c quantityRow) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (r *adRepo) toRow(model *ad.Ad) (adRow, []imageRow) {
	adRow := adRow{
		ID:          int(model.ID),
		Status:      int(model.Status),
		Description: string(model.Description),
		Quantity: quantityRow{
			Available: model.Quantity.Available,
			Reserved:  model.Quantity.Reserved,
		},
		PublishedAt: model.PublishedAt,
		ArchivedAt:  model.ArchivedAt,
	}

	images := make([]imageRow, 0, len(model.Images))
	for _, image := range model.Images {
		images = append(images, imageRow{
			AdID:      int(image.AdID),
			ID:        int(image.ID),
			Content:   image.Content,
			Order:     image.Order,
			CreatedAt: image.CreatedAt,
		})
	}

	return adRow, images
}

func (r *adRepo) toModel(adRow adRow, imageRows []imageRow) (*ad.Ad, error) {
	// Здесь я делаю воссоздание объекта с валидацией.
	// Если вы верите своей базе и у вас нет доступа в прод, то можно явно мапить все данные пропускаю валидацию.
	// Пример такой функции ad.NewAdFromDB, ad.NewImageFromDB, ad.QuantityFromDB.
	quantity, err := ad.NewCountWithReserve(adRow.Quantity.Available, adRow.Quantity.Reserved)
	if err != nil {
		return nil, err
	}

	images := make([]*ad.Image, 0, len(imageRows))
	for _, image := range imageRows {
		image, err := ad.NewImageFromDB(
			image.AdID,
			image.ID,
			image.Content,
			image.Order,
			image.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return ad.NewAdFromDB(
		adRow.ID,
		adRow.Status,
		adRow.Description,
		quantity,
		images,
		adRow.PublishedAt,
		adRow.ArchivedAt,
	)
}
