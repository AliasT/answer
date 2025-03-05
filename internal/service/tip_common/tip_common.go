package tipcommon

import (
	"context"
	"fmt"

	"github.com/apache/answer/internal/entity"
	"github.com/apache/answer/internal/schema"
	answercommon "github.com/apache/answer/internal/service/answer_common"
	questioncommon "github.com/apache/answer/internal/service/question_common"
)

type TipRepo interface {
	AddTip(ctx context.Context, tip *entity.Tip) (err error)
}

type TipCommonService struct {
	// Define the fields and methods for TipCommonService
	tipRepo      TipRepo
	questionRepo questioncommon.QuestionRepo
	answerRepo   answercommon.AnswerRepo
}

func NewTipCommonService(
	TipRepo TipRepo,
	questionRepo questioncommon.QuestionRepo,
	answerRepo answercommon.AnswerRepo,
) *TipCommonService {
	return &TipCommonService{
		tipRepo:      TipRepo,
		questionRepo: questionRepo,
		answerRepo:   answerRepo,
	}
}

func (ts *TipCommonService) AddTip(ctx context.Context, req *schema.TipReq) (*schema.TipRes, error) {
	var tip entity.Tip
	tip.ObjectID = req.ObjectID
	tip.TipType = req.TipType
	tip.ByUserID = req.UserID
	tip.Amount = req.Amount

	if tip.TipType == entity.TipObjectQuestion {
		question, exist, err := ts.questionRepo.GetQuestion(ctx, req.ObjectID)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, fmt.Errorf("question not exist")
		}
		tip.ToUserID = question.UserID
	} else if tip.TipType == entity.TipObjectAnswer {
		answer, exist, err := ts.answerRepo.GetAnswer(ctx, req.ObjectID)
		if err != nil {
			return nil, err
		}

		if !exist {
			return nil, fmt.Errorf("answer not exist")
		}

		tip.ToUserID = answer.UserID

	}

	err := ts.tipRepo.AddTip(ctx, &tip)

	if err != nil {
		return nil, err
	}

	return &schema.TipRes{}, err
}
