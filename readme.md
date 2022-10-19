## 1 微信网页登录
### 1.1 配置
```toml
[wechat]
appID = "xxxx"
appSecret = "xxxx"
redirectUri = "https://xxx.com/code/wechat"
scope = "snsapi_login"
```

### 1.2 调用代码
```go
// 初始化
obj := ewechat.Load('key').Build()
oauthInvoker := obj.GetOauth()

// 跳转到微信授权登录
sEnc := base64.RawURLEncoding.EncodeToString(state)
c.Redirect(http.StatusFound, oauthInvoker.AuthCodeURL(econf.GetString("wechat.redirectUri"), econf.GetString("wechat.scope"), sEnc))

// 调用xxx/code 得到授权access
result, err := invoker.WechatOauth2.ExchangeTokenURL(code)
if err != nil {
c.JSONE(1, "exchange token err: "+err.Error(), err)
return
}

// 根据access，调用微信API获取用户信息
userInfo, err := invoker.WechatOauth2.GetUserInfo(result.AccessToken, result.OpenID, "")
if err != nil {
c.JSONE(1, "get user info err: "+err.Error(), err)
return
}
```
