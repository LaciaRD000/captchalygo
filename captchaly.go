package captchalygo

import (
	"net/http"
	"time"
)

func New(token string) *Session {
	return &Session{
		Client: &http.Client{
			Timeout: 100 * time.Second,
		},
		Token: token,
	}
}
