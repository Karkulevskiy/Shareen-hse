package ws

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type OTP struct {
	Key     string
	Created time.Time
}

type RetentionMap map[string]OTP

func (rm RetentionMap) NewOTP() OTP {
	otp := OTP{
		Key:     uuid.NewString(),
		Created: time.Now(),
	}

	rm[otp.Key] = otp

	return otp
}

func NewRetentionMap(ctx context.Context, retention time.Duration) *RetentionMap {
	return &RetentionMap{}
}
