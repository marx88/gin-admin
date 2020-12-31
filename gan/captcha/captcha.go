package captcha

import (
	"errors"
	"strings"

	"github.com/mojocn/base64Captcha"
)

var captcha *base64Captcha.Captcha

func init() {
	captcha = base64Captcha.NewCaptcha(
		(&base64Captcha.DriverMath{
			Height:     39,
			Width:      116,
			NoiseCount: 50,
		}).ConvertFonts(),
		base64Captcha.DefaultMemStore,
	)
}

// Get 获取验证码
func Get() (string, string, error) {
	uuid, img, err := captcha.Generate()
	if err != nil {
		return "", "", err
	}

	imgs := strings.Split(img, ";base64,")
	if len(imgs) != 2 {
		return "", "", errors.New("验证码生成失败")
	}

	return uuid, imgs[1], nil
}

// Verify 验证图形验证码
func Verify(uuid string, answer string) bool {
	return base64Captcha.DefaultMemStore.Verify(uuid, answer, true)
}
