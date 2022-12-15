package auth

import (
	"github.com/alexperezortuno/go-auth/common"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/data_base"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/data_base/model"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/redis"
	"github.com/golang-jwt/jwt"
	"log"
)

var JWT_TOKEN_SECRET = "ZHTU3oHo6XButkt89ZkJRVUKcWPXDzbLU5UaGA3xYPpY6ASB873GXRJgXQp3pWTATNbNHtufS22xdLYrKf4NqCy5nNaKRryd"
var JWT_REFRESH_TOKEN = "YYELq6utge8Z9C8ynHawaHqcRpV6z33QT2mgfNdhL7porH4VPH3t3ppDSdprpzrMGNSKsmEK4aoFaarNmPByWFytEdLjBsLv"

type TokenResponse struct {
	Status       bool   `json:"status"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Nickname string `json:"nickname"`
	IdCard   string `json:"id_card"`
}

type CustomClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func ValidateUser(ar AuthRequest) (TokenResponse, string) {
	user := model.User{}
	_, err := user.ValidUser(data_base.Connection, ar.Email, ar.Password)

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	if err != nil {
		return TokenResponse{}, err.Error()
	}

	claimsToken := CustomClaim{
		ar.Email,
		jwt.StandardClaims{
			IssuedAt:  common.Now().Unix(),
			ExpiresAt: common.NowAdd(1, 0, 0).Unix(),
			Issuer:    "Makobe",
		},
	}

	claimsRefreshToken := CustomClaim{
		ar.Email,
		jwt.StandardClaims{
			Issuer: "Makobe",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)
	signedToken, err := t.SignedString([]byte(JWT_TOKEN_SECRET))
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	redis.SaveTokenInRedis(signedToken, ar.Email)

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
	refreshToken, err := rt.SignedString([]byte(JWT_REFRESH_TOKEN))
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	return TokenResponse{
		true,
		signedToken,
		refreshToken,
	}, ""
}

func CreateUser(ar UserRequest) (model.User, string) {
	user := model.User{
		Email:    ar.Email,
		Password: ar.Password,
		FullName: ar.FullName,
		Name:     ar.Name,
		LastName: ar.LastName,
		Nickname: ar.Nickname,
		IdCard:   ar.IdCard,
	}

	_, err := user.SaveUser(data_base.Connection)

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return model.User{}, err.Error()
	}

	return user, ""
}
