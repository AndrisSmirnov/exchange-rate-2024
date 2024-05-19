package mail_usecase

import (
	"context"
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	"exchange_rate/pkg/domain/vo"
	"exchange_rate/pkg/packages/errors"
)

func (u *MailUC) SendUpdatedCurrencyEmail(
	ctx context.Context,
	rate *rate_entity.Rate,
	emails []vo.Email,
) *errors.Error {
	return u.infra.Mail.SendEmail(ctx, rate, emails...)
}
