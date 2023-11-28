package crud

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

var storage *sqlx.DB

func Archive(ctx context.Context, adID int64) error {
	var ad *Ad

	err := storage.GetContext(ctx, ad, "SELECT * FROM ad WHERE id = ?", adID)
	if err != nil {
		return err
	}

	now := time.Now()
	ad.ArchivedAt = &now

	_, err = storage.NamedExecContext(
		ctx,
		"UPDATE ad SET archived_at = :archived_at WHERE id = :id;",
		ad,
	)
	if err != nil {
		return err
	}

	return nil
}
