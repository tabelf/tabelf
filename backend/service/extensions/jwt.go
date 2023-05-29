package extensions

type JWTConf struct {
	JwtKey    string `json:"jwt_key"`
	JwtExpire int64  `json:"jwt_expire"`
	JwtIssuer string `json:"jwt_issuer"`
}
