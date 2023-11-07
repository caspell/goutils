package jwtlib

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	JWT_EXP_TIME         = 0
	JWT_EXP_UNIT         = ""
	JWT_EXP_REFRESH_TIME = 0
	JWT_SECRET_KEY       = ""
)

type User struct {
	UserNo     int
	UserId     int
	UserName   string
	RoleId     int
	PhoneNo    int
	UpdateDate string
}

func GetTokenExpSeconds(exp int64, unit string) int64 {
	switch unit {
	case "MINUTE":
		return exp * 60
	case "HOUR":
		return exp * 60 * 60
	default:
		return 0
	}
}

// nologin(token 없음), unknown(token parsing 에러)
func GetUserIDFromToken(c *gin.Context) string {
	userID := "nologin"
	token := ExtractToken(c.Request)
	if token != "" {
		_, claims, err := GetClaimsFromTokenString(token)
		if err != nil {
			userID = "unknown"
		} else {
			userID = claims["userId"].(string)
		}
	}
	return userID
}

// nologin, unknown (-1)
func GetUserNoFromToken(c *gin.Context) int64 {
	token := ExtractToken(c.Request)

	_, claims, err := GetClaimsFromTokenString(token)
	if err != nil {
		return 0
	}
	userNo, _ := strconv.ParseInt(strconv.FormatFloat(claims["userNo"].(float64), 'f', -1, 64), 10, 64)
	return userNo
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func GetClaimsFromTokenString(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method. %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET_KEY), nil
	})
	return token, claims, err
}

func GetLoginInfo(c *gin.Context) *User {
	return nil
}

func NeedTokenUpdate(c *gin.Context) bool {
	tokenString := ExtractToken(c.Request)
	if tokenString == "" {
		return false
	}
	_, claims, _ := GetClaimsFromTokenString(tokenString)
	tokenModifyDate := claims["modifyDate"].(string)
	user := GetLoginInfo(c)
	return user.UpdateDate != tokenModifyDate
}

// CreateAccessToken
func CreateAccessToken(user *User) (string, error) {
	exp := GetTokenExpSeconds(JWT_EXP_TIME, JWT_EXP_UNIT)
	return makeTokenString(user, exp)
}

// CreateRefreshToken
func CreateRefreshToken(user *User) (string, error) {
	exp := GetTokenExpSeconds(JWT_EXP_REFRESH_TIME, JWT_EXP_UNIT)
	return makeTokenString(user, exp)
}

// make token string
func makeTokenString(user *User, exp int64) (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)

	claims["userNo"] = user.UserNo
	claims["userId"] = user.UserId
	claims["userName"] = user.UserName
	claims["roleId"] = user.RoleId
	claims["phoneNo"] = user.PhoneNo
	claims["modifyDate"] = user.UpdateDate
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Duration(exp) * time.Second).Unix()

	t, err := accessToken.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", fmt.Errorf("Failed to make token string. %s", err)
	}
	return t, nil
}
