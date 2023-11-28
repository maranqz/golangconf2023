package crud

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Ad struct {
	ID           int64             `migration:"id" json:"string"`
	Title        string            `migration:"title" json:"title"`
	Description  string            `migration:"desc" json:"description"`
	Price        *int64            `migration:"price" json:"price"`
	SellerID     int64             `migration:"seller_id" json:"seller_id"`
	Images       []string          `migration:"images" json:"images"`
	Attrs        map[string]string `migration:"attrs" json:"attrs"`
	DeliveryFrom int64             `migration:"delivery_from" json:"delivery_from"`
	DeliveryTo   int64             `migration:"delivery_to" json:"delivery_to"`
	ArchivedAt   *time.Time        `migration:"archived_at" json:"archived_at"`
}

type Attrs map[string]string

func (a *Attrs) Scan(src any) error {
	if src == nil {
		return nil
	}

	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("attrs: migration value is not []byte")
	}

	err := json.Unmarshal(bytes, &a)
	if err != nil {
		return err
	}

	return nil
}

func (a Attrs) Value() (driver.Value, error) {
	return json.Marshal(a)
}
