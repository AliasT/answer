package tip

import (
	"context"

	"github.com/apache/answer/internal/base/data"
	"github.com/apache/answer/internal/base/reason"
	"github.com/apache/answer/internal/entity"
	tipcommon "github.com/apache/answer/internal/service/tip_common"
	"github.com/apache/answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
)

type tipRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

// NewTipRepo new repository
func NewTipRepo(
	data *data.Data,
	uniqueIDRepo unique.UniqueIDRepo,
) tipcommon.TipRepo {
	return &tipRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

func (tr *tipRepo) AddTip(ctx context.Context, tip *entity.Tip) (err error) {
	ID, err := tr.uniqueIDRepo.GenUniqueIDStr(ctx, tip.TableName())
	if err != nil {
		return
	}
	tip.ID = ID
	_, err = tr.data.DB.Context(ctx).Insert(tip)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
