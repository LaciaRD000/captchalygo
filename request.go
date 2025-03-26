package captchalygo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (s *Session) SolveRecaptchaV2(url, sitekey string) (v RecaptchaResponse, err error) {
	err = s.do(endpointRecaptchaV2(url, sitekey), &v)
	return v, err
}

func (s *Session) SolveRecaptchaV3(url, sitekey, action string, fast bool) (v RecaptchaResponse, err error) {
	err = s.do(endpointRecaptchaV3(url, sitekey, action, fast), &v)
	return v, err
}

func (s *Session) SolveTurnstile(url, sitekey string) (v TurnstileTokenResponse, err error) {
	err = s.do(endpointTurnstile(url, sitekey), &v)
	return v, err
}

func (s *Session) SolveHCaptcha(url, sitekey string, proxy ...string) (v HCaptchaResponse, err error) {
	reqURL := endpointHCaptcha(url, sitekey)
	if len(proxy) > 0 {
		reqURL += "&proxy=" + proxy[0]
	}
	err = s.do(reqURL, &v)
	return v, err
}

func (s *Session) SolveHCaptchaEnterprise(url, sitekey string, proxy ...string) (v HCaptchaResponse, err error) {
	reqURL := endpointHCaptchaEnterPrise(url, sitekey)
	if len(proxy) > 0 {
		reqURL += "&proxy=" + proxy[0]
	}
	err = s.do(reqURL, &v)
	return v, err
}

func (s *Session) SolveGeetest(url, captchaID string) (v GeeTestTokenResponse, err error) {
	err = s.do(endpointGeetest(url, captchaID), &v)
	return v, err
}

func (s *Session) GetAccount() (v UserResponse, err error) {
	err = s.do(endpointAccount, &v)
	return v, err
}

func (s *Session) GetSubscriptions() (v []SubcriptionResponse, err error) {
	err = s.do(endpointSubscriptions, &v)
	return v, err
}

func (s *Session) do(url string, v interface{}) error {
	resp, err := s.request(url)
	if err != nil {
		return err
	}
	return s.unmarshal(resp, &v)
}

func (s *Session) unmarshal(resp *http.Response, v interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		err = json.Unmarshal(body, v)
		if err != nil {
			return err
		}
	} else {
		return errors.New("StatusCode: " + resp.Status + " Body: " + string(body))
	}

	return nil
}

func (s *Session) request(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+s.Token)
	return s.Client.Do(req)
}
