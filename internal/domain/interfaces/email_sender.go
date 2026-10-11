package interfaces

import (
	"context"

	utilsError "github.com/vucongthanh92/go-base-utils/helper/http_error"
)

type EmailSenderI interface {
	SendVerificationEmail(ctx context.Context, toEmail, token string) *utilsError.ErrorBuilder
}
