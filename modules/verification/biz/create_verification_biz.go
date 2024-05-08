package bizVerification

import (
	"context"
	"fastFood/common"
	"fastFood/components/tokenProvider"
	modelUser "fastFood/modules/user/model"
	modelVerification "fastFood/modules/verification/model"
	"fmt"
	"net/smtp"
)

type FindUserStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
}

type CreateVerificationStorage interface {
	InsertVerification(ctx context.Context, data *modelVerification.VerificationCreate) error
}

type createVerificationBiz struct {
	storeUser         FindUserStorage
	storeVerification CreateVerificationStorage
	tokenProvider     tokenProvider.Provider
	expiry            int
}

func NewCreateVerificationBiz(
	storeUser FindUserStorage,
	storeVerification CreateVerificationStorage,
	tokenProvider tokenProvider.Provider,
	expiry int,
) *createVerificationBiz {
	return &createVerificationBiz{
		storeUser:         storeUser,
		storeVerification: storeVerification,
		tokenProvider:     tokenProvider,
		expiry:            expiry,
	}
}

func (biz *createVerificationBiz) CreateVerification(ctx context.Context, params *modelVerification.ParamsVerification) error {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": params.Email})

	if err != nil {
		return modelUser.ErrCannotFindEntity(err)
	}

	payload := &common.TokenPayLoad{
		UId: user.Id,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return common.ErrInternal(err)
	}

	otpCode := "0000"
	timeExpired := biz.expiry / 60 / 60
	body := fmt.Sprintf("You have requested a password reset. Please enter the OTP to continue. Your OTP code is: %s. The OTP code is valid for %d hours. Please do not share the OTP code with anyone.", otpCode, timeExpired)
	fmt.Println("SendEmail.body", body)
	if err := SendEmail(user.Email, body); err != nil {
		return err
	}

	data := modelVerification.VerificationCreate{
		UserId:  user.Id,
		OTPCode: otpCode,
		Type:    modelVerification.TYPE_EMAIL,
		Token:   accessToken.GetToken(),
	}

	fmt.Println("data", data)
	// if err := biz.storeVerification.InsertVerification(ctx, &data); err != nil {
	// 	return modelVerification.ErrCannotCreateEntity(err)
	// }

	return nil
}

func SendEmail(to, body string) error {
	smtpUsername := "developer001104@gmail.com"
	smtpPassword := "agilejiejwhwizvd"
	smtpServer := "smtp.example.com"
	emailSubject := "Password Reset"
	smtpPort := 587
	emailFrom := "developer001104@gmail.com"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, emailSubject, body))

	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, emailFrom, []string{to}, msg)

	if err != nil {
		return modelVerification.ErrCannotCreateEntity(err)
	}

	return nil
}
