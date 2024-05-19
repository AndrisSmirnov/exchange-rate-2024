package mail_usecase

import (
	"context"
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	"exchange_rate/pkg/domain/vo"
)

func (u *MailUC) SendUpdatedCurrencyEmail(
	ctx context.Context,
	rate *rate_entity.Rate,
	emails []vo.Email,
) error {
	return u.infra.Mail.SendEmail(ctx, rate, emails...)
}
