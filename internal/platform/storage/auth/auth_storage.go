package auth

import (
	"github.com/alexperezortuno/go-auth/common"
	"github.com/alexperezortuno/go-auth/common/environment"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/data_base"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/data_base/model"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/redis"
	"github.com/golang-jwt/jwt"
	"log"
	"strings"
)

var JWT_TOKEN_SECRET = "ZHTU3oHo6XButkt89ZkJRVUKcWPXDzbLU5UaGA3xYPpY6ASB873GXRJgXQp3pWTATNbNHtufS22xdLYrKf4NqCy5nNaKRryd"
var JWT_REFRESH_TOKEN = "YYELq6utge8Z9C8ynHawaHqcRpV6z33QT2mgfNdhL7porH4VPH3t3ppDSdprpzrMGNSKsmEK4aoFaarNmPByWFytEdLjBsLv"
var params = environment.Server()

type TokenResponse struct {
	Status       bool   `json:"status"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ValidationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
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

type TokenRequest struct {
	Token string `json:"token"`
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
			ExpiresAt: common.NowAdd(0, params.TokenLifeTime, 0).Unix(),
			Issuer:    "Makobe",
		},
	}

	claimsRefreshToken := CustomClaim{
		ar.Email,
		jwt.StandardClaims{
			IssuedAt:  common.Now().Unix(),
			ExpiresAt: common.NowAdd(params.RefreshTokenLifeTime, 0, 0).Unix(),
			Issuer:    "Makobe",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)
	signedToken, err := t.SignedString([]byte(JWT_TOKEN_SECRET))
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	redis.SaveToken(signedToken, ar.Email)

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
	refreshToken, err := rt.SignedString([]byte(JWT_REFRESH_TOKEN))
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	redis.SaveRefreshToken(refreshToken, ar.Email)

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

func ValidateToken(tr string) (ValidationResponse, string) {
	tr = strings.Replace(tr, "Bearer ", "", -1)
	log.Printf("[INFO] Validating token: %s", tr)
	validateInStorage := redis.GetToken(tr)

	if validateInStorage == "" {
		return ValidationResponse{Message: "invalid token", Status: false}, "invalid token"
	}

	token, err := jwt.ParseWithClaims(tr, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_TOKEN_SECRET), nil
	})

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		redis.DeleteToken(tr)
		return ValidationResponse{Message: "invalid token", Status: false}, err.Error()
	}

	if _, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return ValidationResponse{Message: "valid", Status: true}, ""
	} else {
		return ValidationResponse{Message: "invalid token", Status: false}, "invalid token"
	}
}

func RefreshToken(tr string) (TokenResponse, string) {
	log.Printf("[INFO] Refreshing token: %s", tr)

	token := redis.GetRefreshToken(tr)

	if token == "" {
		return TokenResponse{}, "not a valid token"
	}

	tokenRes, err := jwt.ParseWithClaims(tr, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_REFRESH_TOKEN), nil
	})

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	var email = tokenRes.Claims.(*CustomClaim).Email
	log.Println(email)

	claimsToken := CustomClaim{
		email,
		jwt.StandardClaims{
			IssuedAt:  common.Now().Unix(),
			ExpiresAt: common.NowAdd(params.RefreshTokenLifeTime, 0, 0).Unix(),
			Issuer:    "Makobe",
		},
	}

	claimsRefreshToken := CustomClaim{
		email,
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

	redis.SaveToken(signedToken, email)
	redis.DeleteRefreshToken(tr)

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
	refreshToken, err := rt.SignedString([]byte(JWT_REFRESH_TOKEN))
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return TokenResponse{}, err.Error()
	}

	redis.SaveRefreshToken(refreshToken, email)

	return TokenResponse{
		true,
		signedToken,
		refreshToken,
	}, ""
}
