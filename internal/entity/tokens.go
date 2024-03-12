package entity

type Tokens struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}

type UserClaims struct {
	Email string
	Role  int64
}
