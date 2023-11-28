package tags

type StructWithTags struct {
	sqlx        string `migration:"migration"` // want `Don't use tags`
	SQLX        string `migration:"migration"`
	JSON        string `json:"JSON"`
	json        string `json:"json"`
	MultiTags   int    `migration:"migration" json:"json"`
	multiTags   int    `migration:"migration" json:"multiTags"`
	WithoutTags string ``
	withoutTags string ``
}
