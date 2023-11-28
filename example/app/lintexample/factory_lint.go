package lintexample

import "ddd/example/app/domain/ad"

func InitiateStructWithoutFactory() {
	_ = ad.Ad{}
	_ = &ad.Quantity{}
}
