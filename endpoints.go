package captchalygo

import "strconv"

var APIVersion = "1"

var (
	endpointCaptchalyAPI = "https://v" + APIVersion + ".captchaly.com"

	endpointRecaptchaV2 = func(url, sitekey string) string {
		return endpointCaptchalyAPI + "/recaptchav2?url=" + url + "&sitekey=" + sitekey
	}
	endpointRecaptchaV3 = func(url, sitekey, action string, fast bool) string {
		return endpointCaptchalyAPI + "/recaptchav3?url=" + url + "&sitekey=" + sitekey + "&action=" + action + "&fast=" + strconv.FormatBool(fast)
	}
	endpointTurnstile = func(url, sitekey string) string {
		return endpointCaptchalyAPI + "/turnstile?url=" + url + "&sitekey=" + sitekey
	}
	endpointHCaptcha = func(url, sitekey string) string {
		return endpointCaptchalyAPI + "/hcaptcha?url=" + url + "&sitekey=" + sitekey
	}
	endpointHCaptchaEnterPrise = func(url, sitekey string) string {
		return endpointCaptchalyAPI + "/hcaptcha-enterprise?url=" + url + "&sitekey=" + sitekey
	}
	endpointGeetest = func(url, captchaID string) string {
		return endpointCaptchalyAPI + "/geetest?url=" + url + "&captchaId=" + captchaID
	}

	endpointAccount       = endpointCaptchalyAPI + "/account"
	endpointSubscriptions = endpointCaptchalyAPI + "/subscriptions"
)
