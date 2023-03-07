package helper

import (
	"CloudStorage/core/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Second * time.Duration(second))},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("Invalid Token")
	}
	return uc, err
}

func MailSend(address, code string) error {
	e := email.NewEmail()
	e.From = define.EmailFrom
	e.To = []string{address}
	e.Subject = "Verification Code"
	e.HTML = []byte("Your Verification Code is below<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("1", define.EmailAccount, define.EmailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.BucketPath)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretId,
			SecretKey: define.SecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := GenerateUUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}
	return define.BucketPath + "/" + key, nil
}
