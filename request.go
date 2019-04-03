package vkauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	vkAuthURI = "https://oauth.vk.com/token"
)

type authResp struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	UserID           int    `json:"user_id"`
	TrustedHash      string `json:"trusted_hash"`
	CaptchaSID       string `json:"captcha_sid"`
	CaptchaImg       string `json:"captcha_img"`
	ValidationType   string `json:"validation_type"`
	ValidationSID    string `json:"validation_sid"`
	PhoneMask        string `json:"phone_mask"`
	RedirectURI      string `json:"redirect_uri"`
}

//oauthRequest - Makes auth request to auth.vk.com and returns parsed response
func oauthRequest(values url.Values) (*authResp, error) {
	auth := &authResp{}
	requri := fmt.Sprintf("%s?%s", vkAuthURI, values.Encode())
	resp, err := http.Get(requri)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, auth); err != nil {
		return nil, err
	}
	return auth, nil
}
