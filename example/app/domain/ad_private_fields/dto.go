package ad_private_fields

type AdDTO struct {
	ID         int
	CategoryID int
	UserID     int
	Status     int
	Quantity   QuantityDTO
}

type QuantityDTO struct {
	Available int
	All       int
}

type ImageDTO struct {
	ID      int
	Content []byte
	Order   int
}
