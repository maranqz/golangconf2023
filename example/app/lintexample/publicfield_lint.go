package lintexample

import "ddd/example/app/domain/ad"

func ChangePublicField() {
	q, _ := ad.NewQuantity(1)
	q.Reserved = -5
}
