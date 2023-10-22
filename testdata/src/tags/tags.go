package tags

type sqlx struct {
	sqlx string `db:"db"` // want `Don't use tags`
}

type SQLX struct {
	SQLX string `db:"db"` // want `Don't use tags`
}

type JSON struct {
	JSON string `json:"JSON"` // want `Don't use tags`
}

type json struct {
	json string `json:"json"` // want `Don't use tags`
}

type MultiTags struct {
	MultiTags int `db:"db" json:"json"` // want `Don't use tags`
}

type multiTags struct {
	multiTags int `db:"db" json:"multiTags"` // want `Don't use tags`
}

type WithoutTags struct {
	WithoutTags string
}

type withoutTags struct {
	withoutTags string
}
