package ad

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type AdBadExample struct {
	ID         AdID       `migration:"ID" gorm:"ID"`
	CategoryID CategoryID `migration:"category_id" gorm:"category_id"`
	Count      Quantity   `migration:"Quantity" gorm:"Quantity"`
	/**/
}

func (c *Quantity) Scan(src any) error {
	if src == nil {
		return nil
	}

	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("%w: migration value is not []byte", ErrCount)
	}

	err := json.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}

	return nil
}

func (c Quantity) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Quantity) UnmarshalJSON(bytes []byte) error { /**/
	return nil
}

func (c *Quantity) MarshalJSON() ([]byte, error) { /**/
	return nil, nil
}
