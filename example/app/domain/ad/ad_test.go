package ad

import (
	"fmt"
	"time"

	lerrors "ddd/example/app/pkg/errors"
)

var tzUTC3, _ = time.LoadLocation("Turkey")

var (
	userID, _     = NewUserID(2)
	categoryID, _ = NewCategoryID(1)

	title = lerrors.Must(NewTitle("title_test"))
	desc  = lerrors.Must(NewDescription("description_test"))
	now   = time.Now()
)

func testAd() *Ad {
	return lerrors.Must(newAd(
		NextAdID(),
		userID,
		categoryID,
		title,
		desc,
		now,
	))
}

func ExampleAd_Publish() {
	ad := testAd()

	err := ad.Publish(time.Date(2023, 11, 28, 13, 30, 0, 0, tzUTC3))
	fmt.Println(err) // <nil>

	fmt.Println(ad.Status)      // Published
	fmt.Println(ad.PublishedAt) // 2023-11-28 13:30:00 +0300 +03

	err = ad.Publish(time.Now())
	fmt.Println(err) // ad: publish

	// Output: <nil>
	// 2
	// 2023-11-28 13:30:00 +0300 +03
	// ad: publish
}

func ExampleAd_Lifecicle() {
	ad := testAd()

	err := ad.Publish(time.Now())
	fmt.Println(err) // <nil>
	err = ad.Archive(time.Now())
	fmt.Println(err) // <nil>

	fmt.Println(ad.Status) // Archived

	// Output: <nil>
	// <nil>
	// 3
}

func ExampleAd_CompareByID() {
	sameID := NextAdID()
	ad := lerrors.Must(newAd(
		sameID,
		userID,
		categoryID,
		title,
		desc,
		now,
	))
	adSame := lerrors.Must(newAd(
		sameID,
		userID,
		categoryID,
		title,
		desc,
		now,
	))

	fmt.Println(ad == adSame)       // false, сущности сравниваем по ID
	fmt.Println(ad.ID == adSame.ID) // true
	fmt.Println(ad.IsEqual(adSame)) // true

	adAnother := lerrors.Must(newAd(
		sameID,
		userID,
		categoryID,
		title,
		desc,
		now,
	))
	fmt.Println(ad.IsEqual(adAnother)) // false

	// Output: false
	// true
	// true
	// false
}

func ExampleQuantity_CompareByValue() {
	count1, _ := NewQuantity(1)    // Quantity{1, 0}
	count2As1, _ := NewQuantity(1) // Quantity{1, 0}
	count2, _ := NewQuantity(2)    // Quantity{2, 0}

	fmt.Println(count1 == count2As1) // true – Quantity{1, 0} == Quantity{1, 0}
	fmt.Println(count1 == count2)    // false – Quantity{1, 0} == Quantity{2, 0}

	// Output: true
	// false
}
