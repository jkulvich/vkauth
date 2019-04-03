package vkauth

import (
	"net/url"
)

const (
	ClientIPhone = iota
	ClientWPhone
	ClientAndroid
)

type authParams struct {
	GrantType      string
	ClientID       string
	ClientSecret   string
	Username       string
	Password       string
	Version        string
	TwoFaSupported bool
	CaptchaSID     string
	CaptchaKey     string
	Code           string
}

//ToURL - Marshal the structure to URl values
func (p *authParams) toURL() url.Values {
	u := url.Values{}
	u.Add("grant_type", p.GrantType)
	u.Add("client_id", p.ClientID)
	u.Add("client_secret", p.ClientSecret)
	u.Add("username", p.Username)
	u.Add("password", p.Password)
	u.Add("v", p.Version)
	if p.TwoFaSupported {
		u.Add("2fa_supported", "1")
	}
	if p.CaptchaSID != "" {
		u.Add("captcha_sid", p.CaptchaSID)
	}
	if p.CaptchaKey != "" {
		u.Add("captcha_key", p.CaptchaKey)
	}
	if p.Code != "" {
		u.Add("code", p.Code)
	}
	return u
}

//newBasicParas - Generate new auth structure with default values
func newBasicParams(user, pass string) *authParams {
	p := &authParams{
		GrantType:      "password",
		Username:       user,
		Password:       pass,
		Version:        "5.40",
	}
	return p
}

//bindClient - Binds standalone client to auth
func (p *authParams) bindClient(client int) {
	switch client {
	case ClientIPhone:
		{
			p.ClientID = "3140623"
			p.ClientSecret = "VeWdmVclDCtn6ihuP1nt"
		}
	case ClientWPhone:
		{
			p.ClientID = "3697615"
			p.ClientSecret = "AlVXZFMUqyrnABp8ncuU"
		}
	case ClientAndroid:
		{
			p.ClientID = "2274003"
			p.ClientSecret = "hHbZxrka2uZ6jB1inYsH"
		}
	default:
		p.bindClient(ClientAndroid)
	}
}

//bindCode - Binds SMS code to auth
func (p *authParams) bindCode(code string) {
	p.Code = code
}

//bindCaptcha - Binds captcha key to auth
func (p *authParams) bindCaptcha(sid, key string) {
	p.CaptchaSID = sid
	p.CaptchaKey = key
}
