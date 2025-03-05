package entity

import "time"

const (
	// 阅读
	TipObjectRead = 1
	// 问题
	TipObjectQuestion = 2
	// 评论
	TipObjectAnswer = 3
	// 打赏
	TipObjectTip = 4
)

type Tip struct {
	ID        string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated_at TIMESTAMP"`
	//
	Amount  int `xorm:"not null default 0 INT(11) amount"`
	TipType int `xorm:"not null default 0 INT(11) tip_type"`
	// ObjectType TipObjectTip 为打赏人
	ByUserID string `xorm:"not null default 0 BIGINT(20)"`
	ToUserID string `xorm:"not null default 0 BIGINT(20)"`
	ObjectID string `xorm:"not null default 0 BIGINT(20)"`
}

func (Tip) TableName() string {
	return "tip"
}
