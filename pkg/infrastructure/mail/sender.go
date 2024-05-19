package mail

import (
	"context"
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	"exchange_rate/pkg/domain/vo"
	"fmt"
	"net/smtp"

	"exchange_rate/pkg/packages/errors"

	"exchange_rate/pkg/utils"
)

type EmailSender struct {
	config  *Config
	auth    smtp.Auth
	address string
	subject string
}

func NewEmailService() (*EmailSender, *errors.Error) {
	emailAddress, err := utils.TryGetEnv[string]("EMAIL_ADDRESS")
	if err != nil {
		return nil, newErrorNoEnvVar("EMAIL_ADDRESS")
	}
	emailAppCode, err := utils.TryGetEnv[string]("EMAIL_APP_CODE")
	if err != nil {
		return nil, newErrorNoEnvVar("EMAIL_APP_CODE")
	}
	emailSubject, err := utils.TryGetEnv[string]("EMAIL_SUBJECT")
	if err != nil {
		return nil, newErrorNoEnvVar("EMAIL_SUBJECT")
	}
	smtpHost, err := utils.TryGetEnv[string]("SMTP_HOST")
	if err != nil {
		return nil, newErrorNoEnvVar("SMTP_HOST")
	}
	smtpPort, err := utils.TryGetEnv[string]("SMTP_PORT")
	if err != nil {
		return nil, newErrorNoEnvVar("SMTP_PORT")
	}
	mailConfig, err := newConfig(
		emailAddress,
		emailAppCode,
		emailSubject,
		smtpHost,
		smtpPort,
	)
	if err != nil {
		return nil, err
	}

	return &EmailSender{
		config:  mailConfig,
		auth:    smtp.PlainAuth("", mailConfig.address, mailConfig.appKey, mailConfig.smtpHost),
		address: fmt.Sprintf("%s:%s", mailConfig.smtpHost, mailConfig.smtpPort),
	}, nil
}

func (e *EmailSender) SendEmail(
	_ context.Context, data *rate_entity.Rate, receivers ...vo.Email,
) *errors.Error {
	rec := make([]string, 0, len(receivers))
	for _, receiver := range receivers {
		rec = append(rec, receiver.ToString())
	}
	err := e.sendEmail(e.crateTemplate(data), rec...)
	return err
}

func (e *EmailSender) sendEmail(message []byte, receiversEmail ...string) *errors.Error {
	if err := smtp.SendMail(
		e.address,
		e.auth,
		e.config.address,
		receiversEmail,
		message,
	); err != nil {
		return errSendMail(err)
	}

	return nil
}
