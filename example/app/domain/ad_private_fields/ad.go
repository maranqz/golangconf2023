package ad_private_fields

import (
	"ddd/example/app/domain/ad"
)

// Ad со всеми приватными полями.
// Приватные поля нам уверенность, что нельзя поломать данные снаружи.
// Однако, что значит снаружи, если например Картинка или другой Aggregate в этом же Ограниченном контексте будем менять приватные поля Ad, хорошо ли это?
// Ответ нет. Поэтому пока это можно отслеживать на уровне командных соглашений.
// Линтер gopublicfield не умеет такое отлавливать, но только "пока".
//
// Чтобы получить данные Ad:
// 1. можно создавать getter, IDE и некоторые редакторы умеют это делать.
// 2. создавать DTO.
type Ad struct {
	id         ad.AdID
	userID     ad.UserID
	categoryID ad.CategoryID
	status     ad.Status
	quantity   quantity
	/* и другие поля*/
}

func (a *Ad) ID() ad.AdID {
	return a.id
}

func (a *Ad) UserID() ad.UserID {
	return a.userID
}

func (a *Ad) CategoryID() ad.CategoryID {
	return a.categoryID
}

func (a *Ad) Status() ad.Status {
	return a.status
}

func (a *Ad) DTO() AdDTO {
	return AdDTO{
		ID:         int(a.id),
		CategoryID: int(a.categoryID),
		UserID:     int(a.userID),
		Status:     int(a.status),
		Quantity:   a.quantity.DTO(),
	}
}

type quantity struct {
	available int
	reserved  int
}

func (q quantity) DTO() QuantityDTO {
	return QuantityDTO{
		Available: q.available,
		All:       q.available + q.reserved,
	}
}
