package vkauth

import "fmt"

type Context struct {
	Token           string
	NeedCode        bool
	NeedCaptcha     bool
	CaptchaUri      string
	auth            *authParams
	sid             string
	captchaExpected bool
}

//NewContext - Returns new auth context
func NewContext(user, pass string) *Context {
	auth := newBasicParams(user, pass)
	auth.bindClient(ClientAndroid)
	return &Context{
		auth: auth,
	}
}

//Auth - Make auth request
func (c *Context) Auth() error {
	resp, err := oauthRequest(c.auth.toURL())
	if err != nil {
		return err
	}

	if resp.AccessToken != "" {
		c.Token = resp.AccessToken
		return nil
	}
	if resp.Error == "need_validation" {
		c.NeedCode = true
		if !c.auth.TwoFaSupported {
			return fmt.Errorf("2fa required to continue")
		}
		return nil
	}
	if resp.Error == "need_captcha" {
		c.NeedCaptcha = true
		c.CaptchaUri = resp.CaptchaImg
		c.sid = resp.CaptchaSID
		if !c.captchaExpected {
			return fmt.Errorf("captcha required to continue")
		}
		return nil
	}

	return fmt.Errorf("unknown login scenario: %s", resp.Error)
}

//SetClient - Sets official VK client for auth
func (c *Context) SetClient(client int) *Context {
	c.auth.bindClient(client)
	return c
}

//SetCustomClient - Sets custom client for auth
func (c *Context) SetCustomClient(clientID, clientSecret string) *Context {
	c.auth.ClientID = clientID
	c.auth.ClientSecret = clientSecret
	return c
}

//SetCode - Sets SMS code for auth
func (c *Context) SetCode(code string) *Context {
	c.auth.bindCode(code)
	return c
}

//SetCaptcha - Sets captcha key for auth
func (c *Context) SetCaptcha(key string) *Context {
	c.auth.bindCaptcha(c.sid, key)
	return c
}

//ExpectTwoFa - Sets 2FA support
func (c *Context) ExpectTwoFa(support bool) *Context {
	c.auth.TwoFaSupported = support
	return c
}

//ExpectCaptcha - Sets captcha support
func (c *Context) ExpectCaptcha(expected bool) *Context {
	c.captchaExpected = true
	return c
}
