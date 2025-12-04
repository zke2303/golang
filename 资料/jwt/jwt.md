## 一.JWT 介绍

`JWT(json web token)`是为了在网络应用环境间传递声明而执行的一种基于JSON的开放标准。
JWT的声明一般被用来在身份提供者和服务提供者间传递被认证的用户身份信息，以便于从资源服务器获取资源。比如用户登录。在传统的用户登录认证中，因为http是无状态的，所以都是采用session方式。用户登录成功，服务端会保存一个session，服务端会返回给客户端一个sessionId，客户端会把sessionId保存在cookie中，每次请求都会携带这个sessionId。
cookie+session这种模式通常是保存在内存中，而且服务从单服务到多服务会面临的session共享问题。虽然目前存在使用Redis进行Session共享的机制，但是随着用户量和访问量的增加，Redis中保存的数据会越来越多，开销就会越来越大，多服务间的耦合性也会越来越大，Redis中的数据也很难进行管理，例如当Redis集群服务器出现Down机的情况下，整个业务系统随之将变为不可用的状态。而JWT不是这样的，只需要服务端生成token，客户端保存这个token，每次请求携带这个token，服务端认证解析就可。

- `Json web token` (`JWT`), 是为了在网络应用环境间传递声明而执行的一种基于JSON的开放标准（`RFC 7519`)
- 该token被设计为紧凑且安全的，特别适用于分布式站点的单点登录（SSO）场景
- JWT的声明一般被用来在身份提供者和服务提供者间传递被认证的用户身份信息，以便于从资源服务器获取资源
- 也可以增加一些额外的其它业务逻辑所必须的声明信息，该token也可直接被用于认证，也可被加密

- JWT是无状态的



## 二、JWT 的构成

JWT就是一段字符串, 由三段信息构成, 三段信息文本使用`.`(点) 拼接就构成了JWT字符串 :

- `eyJhbGciOiJIUzI1sNiIsIn.eyJzdWIiOiIxMjRG9OnRydWV9.TJVArHDcEfxjoYZgeFONFh7HgQ`
- 第一部分我们称它为头部 : `header`
- 第二部分我们称其为载荷 : `payload` (类似于飞机上承载的物品)
- 第三部分是签证 : `signature`

![在这里插入图片描述](assets/d5761caaf029235a132a0c5364bb39c8.png)



第一部分我们称它为头部（header),第二部分我们称其为载荷（payload)，第三部分是签证（signature)

- **header**
  jwt的头部承载两部分信息：
  1.声明类型，这里是jwt
  2.声明加密的算法 通常直接使用 HMAC SHA256
  完整的头部就像下面这样的JSON：

```bash
{
"typ": "JWT",
"alg": "HS256"
}
```

然后将头部进行base64加密（该加密是可以对称解密的),构成了第一部分：

- **playload**

载荷就是存放有效信息的地方。这个名字像是特指飞机上承载的货品，这些有效信息包含三个部分：标准中注册的声明、公共的声明、私有的声明。
标准中注册的声明 (建议但不强制使用) ：
iss: jwt签发者
sub: jwt所面向的用户
aud: 接收jwt的一方
exp: jwt的过期时间，这个过期时间必须要大于签发时间
nbf: 定义在什么时间之前，该jwt都是不可用的.
iat: jwt的签发时间
jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
公共的声明 ：
公共的声明可以添加任何的信息，一般添加用户的相关信息或其他业务需要的必要信息.但不建议添加敏感信息，因为该部分在客户端可解密。
私有的声明 ：
私有声明是提供者和消费者所共同定义的声明，一般不建议存放敏感信息，因为base64是对称解密的，意味着该部分信息可以归类为明文信息。
定义一个payload：

```bash{
"uid":"e12a34b56c78d9e0f",
"name":"ramostear",
"role":"admin"
}
```

然后将其进行base64加密，得到Jwt的第二部分：

eyJvcmciOiLku4rml6XlpLTmnaEiLCJuYW1lIjoiRnJlZeeggeWGnCIsImV4cCI6MTUxNDM1NjEwMywiaWF0IjoxNTE0MzU2MDQzLCJhZ2UiOiIyOCJ9


- **signature**
  jwt的第三部分是一个签证信息，这个签证信息由三部分组成：header (base64后的)、
  payload (base64后的)、secret。
  这个部分需要base64加密后的header和base64加密后的payload使用。连接组成的字符串，然后通过header中声明的加密方式进行加盐secret组合加密，然后就构成了jwt的第三部分：
  密钥secret是保存在服务端的，服务端会根据这个密钥进行生成token和验证，所以需要保护好。





# 三、使用

在go中，使用最多的JWT包是``.通过`go get -u github.com/golang-jwt/jwt/v5`进行安装。

### 默认Claim

如果我们直接使用JWT中默认的字段，没有其他定制化的需求则可以直接使用这个包中的和方法快速生成和解析token。

```go
// 用于签名的字符串
var mySigningKey = []byte("liwenzhou.com")

// GenRegisteredClaims 使用默认声明创建jwt
func GenRegisteredClaims() (string, error) {
	// 创建 Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间
		Issuer:    "qimi",                                             // 签发人
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(mySigningKey)
}

// ParseRegisteredClaims 解析jwt
func ValidateRegisteredClaims(tokenString string) bool {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil { // 解析token失败
		return false
	}
	return token.Valid
}
```



### 自定义Claims

我们需要定制自己的需求来决定JWT中保存哪些数据，比如我们规定在JWT中要存储`username`信息，那么我们就定义一个`MyClaims`结构体如下：

这个定制的结构体，必须内嵌 `jwt.RegisteredClaims`

```go
// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}
```

然后我们定义JWT的过期时间，这里以24小时为例：

```go
const TokenExpireDuration = time.Hour * 24
```

接下来还需要定义一个用于签名的字符串：

```go
// CustomSecret 用于加盐的字符串
var CustomSecret = []byte("夏天夏天悄悄过去")
```

----

### 生成JWT

我们可以根据自己的业务需要封装一个生成 token 的函数。

```go
// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "my-project", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}
```



### 解析JWT

根据给定的 JWT 字符串，解析出数据。

```go
// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
        if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
```

>```go
>token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
>	// 直接使用标准的Claim则可以直接使用Parse方法
>	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
>	return CustomSecret, nil
>})
>```
>
>在`jwt.ParseWithClaims`方法中, tokenString是需要解析的对象，`&CustomClaims{}`是根据需求定制的结构体对象的地址，最后一个参数，是一个方法类型，其作用是返回一个==密钥==，用于生成签名。

`ParseWithClaims(tokenString string, **claims jwt.Claims**, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error)`



```go
// jwt.Token结构体
type Token struct {
    Header    map[string]interface{}   // 头部
    Claims    Claims                   // 负载（你放的那些信息）
    Method    SigningMethod            // 使用的签名算法
    Signature string                   // 签名（生成前为空）
}
```

`claims, ok := token.Claims.(*CustomClaims);`这条语句的作用是：从token对象中，取出Claims，然后通过类型断言，判断这个token中的对象是否是 CustomClaims 结构体类型。
