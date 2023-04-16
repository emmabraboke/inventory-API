package emailService

import (
	"inventory/internals/entity/userEntity"
	"net/smtp"
)

type emailSrv struct {
	smtpHost     string
	smtpPort     string
	smtpPassword string
	smtpUser     string
}

type EmailService interface {
	SendEmail(req *userEntity.EmailReq) error
	smtpInstance() smtp.Auth
	SendBulkEmail(req *userEntity.BulkEmailReq) error
}

func NewEmailSrv(smtpHost, smtpPort, smtpPassword, smtpUser string) EmailService {
	return &emailSrv{smtpHost: smtpHost, smtpPort: smtpPort, smtpPassword: smtpPassword, smtpUser: smtpUser}
}

func (t *emailSrv) smtpInstance() smtp.Auth {
	auth := smtp.PlainAuth("", t.smtpUser, t.smtpPassword, t.smtpHost)
	return auth
}

func (t *emailSrv) SendEmail(req *userEntity.EmailReq) error {
	err := smtp.SendMail(t.smtpHost+":"+t.smtpPort, t.smtpInstance(), req.From, []string{req.To}, []byte(req.Body))

	return err
}

func (t *emailSrv) SendBulkEmail(req *userEntity.BulkEmailReq) error {
	err := smtp.SendMail(t.smtpHost+":"+t.smtpPort, t.smtpInstance(), req.From, req.To, []byte(req.Body))

	return err
}
