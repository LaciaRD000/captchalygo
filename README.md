# captchalygo
captchaly go library

## Supported CAPTCHA types:
* ReCaptchaV2
* ReCaptchaV3
* Turnstile
* HCaptcha
* HCaptchaEnterprise
* Geetest

## Installation
You don't need this source code unless you want to modify the package. If you just want to use the package, just run:
```
go get github.com/LaciaRD000/captchalygo@latest
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/LaciaRD000/captchalygo"
)

func main() {
	s := captchalygo.New("YOUR_CAPTCHA_SECRET")
	v, err := s.SolveHCaptcha("YOUR_URL", "YOUR_SITEKEY")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
```