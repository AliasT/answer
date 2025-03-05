package schema

type TipReq struct {
	// question id
	ObjectID string `json:"object_id"`
	// amount
	Amount int `json:"amount"`
	// tip type
	TipType int `json:"tip_type"`
	// user id
	UserID string `json:"-"`

	ToUserID string `json:"-"`
}

type TipRes struct {
}
