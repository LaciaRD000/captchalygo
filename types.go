package captchalygo

import (
	"net/http"
)

type Session struct {
	Token  string
	Client *http.Client
}

type GeeTestTokenResponse struct {
	Time     float32 `json:"time"`
	Duration float32 `json:"duration"`
	Deducted *string `json:"deducted"`
	Token    string  `json:"token"`
}

type HCaptchaResponse struct {
	Time     float32 `json:"time"`
	Duration float32 `json:"duration"`
	Deducted *string `json:"deducted"`
	Token    string  `json:"token"`
}

type RecaptchaResponse struct {
	Time     float32 `json:"time"`
	Duration float32 `json:"duration"`
	Deducted string  `json:"deducted"`
	Token    string  `json:"token"`
}

type TurnstileTokenResponse struct {
	Time     float32 `json:"time"`
	Duration float32 `json:"duration"`
	Deducted *string `json:"deducted"`
	Token    string  `json:"token"`
}

type SubcriptionResponse struct {
	ID             string  `json:"id"`
	Start          float32 `json:"start"`
	End            float32 `json:"end"`
	MaxConcurrency float32 `json:"max_concurrency"`
	Autorenew      bool    `json:"autorenew"`
	Active         bool    `json:"active"`
}

type FailedCaptchaResponse struct {
	Detail string `json:"detail"`
}

type UserResponse struct {
	ID      string  `json:"id"`
	Email   string  `json:"email"`
	Balance float32 `json:"balance"`
	JoinAt  float32 `json:"join_at"`
}

type HTTPValidationErrorResponse struct {
	Detail []ValidationError `json:"detail"`
}

type ValidationError struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}
