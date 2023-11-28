package tags

type sqlx struct {
	sqlx string `migration:"migration"` // want `Don't use tags`
}

type SQLX struct {
	SQLX string `migration:"migration"` // want `Don't use tags`
}

type JSON struct {
	JSON string `json:"JSON"` // want `Don't use tags`
}

type json struct {
	json string `json:"json"` // want `Don't use tags`
}

type MultiTags struct {
	MultiTags int `migration:"migration" json:"json"` // want `Don't use tags`
}

type multiTags struct {
	multiTags int `migration:"migration" json:"multiTags"` // want `Don't use tags`
}

type WithoutTags struct {
	WithoutTags string
}

type withoutTags struct {
	withoutTags string
}
