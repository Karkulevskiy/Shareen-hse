package ws

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// OTP dedcribes one time password
type OTP struct {
	Key     string    //secret key
	Created time.Time //time of creation
}

// RetentionMap dedcribes map of OTP
type RetentionMap map[string]OTP

// NewOTP creates new one time password
func (rm RetentionMap) NewOTP() OTP {
	otp := OTP{
		Key:     uuid.NewString(),
		Created: time.Now(),
	}

	rm[otp.Key] = otp

	return otp
}

// NewRetentionMap creates new retention map
func NewRetentionMap(ctx context.Context, retention time.Duration) *RetentionMap {
	rm := &RetentionMap{}

	go rm.Retention(ctx, retention)

	return rm
}

// VerifyOTP verifies one time password
func (rm RetentionMap) VerifyOTP(ctx context.Context, otp string) bool {
	if _, ok := rm[otp]; !ok {
		return false
	}

	delete(rm, otp)

	return true
}

// Retention verifies one time password
func (rm RetentionMap) Retention(ctx context.Context, retention time.Duration) {
	ticker := time.NewTicker(400 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			for _, otp := range rm {
				if otp.Created.Add(retention).Before(time.Now()) {
					delete(rm, otp.Key)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
