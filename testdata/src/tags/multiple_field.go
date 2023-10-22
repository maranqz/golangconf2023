package tags

type StructWithTags struct {
	sqlx        string `db:"db"` // want `Don't use tags`
	SQLX        string `db:"db"`
	JSON        string `json:"JSON"`
	json        string `json:"json"`
	MultiTags   int    `db:"db" json:"json"`
	multiTags   int    `db:"db" json:"multiTags"`
	WithoutTags string ``
	withoutTags string ``
}
