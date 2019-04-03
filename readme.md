# VKAuth
Does the authorization, as the user does from the VKontakte application.  

- 2FA supports  
- Captcha supports  

# Examples

Simple. Without 2FA and captcha:  
```golang
func main() {
	vkctx := vkauth.NewContext(vklogin, vkpass)
	if err := vkctx.Auth(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your new token: %s", vkctx.Token)
}
```

By default Android client used.  
Way to change it to iPhone:  
```golang
func main() {
	vkctx := vkauth.NewContext(vklogin, vkpass).
	    SetClient(vkauth.ClientIPhone)
	if err := vkctx.Auth(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your new token: %s", vkctx.Token)
}
```

Or auth as another standalone client:  
```golang
func main() {
	vkctx := vkauth.NewContext(vklogin, vkpass).
	    SetCustomClient("3697615", "AlVXZFMUqyrnABp8ncuU")
	if err := vkctx.Auth(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your new token: %s", vkctx.Token)
}
```

2FA and captcha handling:
```golang
func main() {
	vkctx := vkauth.NewContext(vklogin, vkpass).
	    ExpectTwoFa(true).
	    ExpectCaptcha(true)
	err := vkctx.Auth()
	if err != nil {
		log.Fatal(err)
	}

	if vkctx.NeedCaptcha {
		fmt.Printf("Open the captcha: %s\n", vkctx.CaptchaUri)
		fmt.Print("Captcha text: ")
		var key string
		fmt.Scanln(&key)
		if err := vkctx.SetCaptcha(key).Auth(); err != nil {
			log.Fatalf("%s", err)
		}
	}

	if vkctx.NeedCode {
		fmt.Print("Code from SMS: ")
		var code string
		fmt.Scanln(&code)
		if err := vkctx.SetCode(code).Auth(); err != nil {
			log.Fatalf("%s", err)
		}
	}

	fmt.Printf("Your new token: %s", vkctx.Token)
}
```

# Raw queries examples

LOGIN - User login, e.g. phone number  
PASS - User password  
CLIENT - Standalone client ID  
SECRET - Standalone client secret key  
  
Client/Key from official VK clients:
```json
{
	"IPhone": {
		"client": "3140623",
		"secret": "VeWdmVclDCtn6ihuP1nt"
	},
	"WPhone": {
		"client": "3697615",
		"secret": "AlVXZFMUqyrnABp8ncuU"
	},
	"Android": {
		"client": "2274003",
		"secret": "hHbZxrka2uZ6jB1inYsH"
	}
}
```

### Regular:

```bash
curl -XGET 'https://oauth.vk.com/token?grant_type=password&client_id=CLIENT&client_secret=SECRET&username=LOGIN&password=PASS&v=5.40&2fa_supported=1'
```
```json
{
	"access_token": "e41deacb5932ad9c0eefbc88baef3890307a3cda4de6fee37cbed1c952247e7e3c4a5a0f1f0cb943a8635",
	"expires_in":    0,
	"user_id":       180828479,
	"trusted_hash": "bb2939a24e09915c92c43751571985fcac4f527584"
}
```

### Captcha

```bash
curl -XGET 'https://oauth.vk.com/token?grant_type=password&client_id=CLIENT&client_secret=SECRET&username=LOGIN&password=PASS&v=5.40&2fa_supported=1'
```
```json
{
	"error":       "need_captcha",
	"captcha_sid": "825644198168",
	"captcha_img": "http:\/\/api.vk.com\/captcha.php?sid=825644198168&s=1"
}
```

You must go to the browser or get a picture with captcha, and then repeat the request with the parameters:    
captcha_sid - captcha id and captcha_key - text from the captcha.  

```json
curl -XGET 'https://oauth.vk.com/token?grant_type=password&client_id=CLIENT&client_secret=SECRET&username=LOGIN&password=PASS&v=5.40&2fa_supported=1&captcha_sid=825644198168&captcha_key=ABCD1234'
{
	"access_token": "e41deacb5932ad9c0eefbc88baef3890307a3cda4de6fee37cbed1c952247e7e3c4a5a0f1f0cb943a8635",
	"expires_in":    0,
	"user_id":       180828479,
	"trusted_hash": "bb2939a24e09915c92c43751571985fcac4f527584"
}
```

### 2FA:

```bash
curl -XGET 'https://oauth.vk.com/token?grant_type=password&client_id=CLIENT&client_secret=SECRET&username=LOGIN&password=PASS&v=5.40&2fa_supported=1'
```
```json
{
	"error":             "need_validation",
	"validation_type":   "2fa_sms",
	"validation_sid":    "2fa_280838479_3162623_f6820147faf1af8266",
	"phone_mask":        "+7 *** *** ** 76",
	"error_description": "sms sent, use code param",
	"redirect_uri":      "https:\/\/m.vk.com\/login?act=authcheck&api_hash=02a80f3eaf756bfd3e"
}
```

At this moment SMS arrives with a password to the user.  
After that, you need to repeat the request with the code parameter from SMS:  

```bash
curl -XGET 'https://oauth.vk.com/token?grant_type=password&client_id=CLIENT&client_secret=SECRET&username=LOGIN&password=PASS&v=5.40&2fa_supported=1&code=032744'
```
```json
{
	"access_token": "e41deacb5932ad9c0eefbc88baef3890307a3cda4de6fee37cbed1c952247e7e3c4a5a0f1f0cb943a8635",
	"expires_in":    0,
	"user_id":       180828479,
	"trusted_hash": "bb2939a24e09915c92c43751571985fcac4f527584"
}
```