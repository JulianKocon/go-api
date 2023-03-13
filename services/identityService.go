package services

import (
	"example/go-api/models"
	"example/go-api/repositories"

	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type TokenRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type IdentityService interface {
	RegisterUser(models.User) error
	GenerateJWT(email string, username string) (tokenString string, err error)
	ValidateToken(signedToken string) (err error)
	GetToken(TokenRequest) (*string, error)
}

type identityService struct {
	identityRepostiory repositories.IdentityRepostiory
	mailService MailService
}

func NewIdentityService(repo repositories.IdentityRepostiory, mailService MailService) IdentityService {
	return &identityService{
		identityRepostiory: repo,
		mailService: mailService,
	}
}

func (service identityService) RegisterUser(user models.User)error{
	if(len(user.Email) < 5) {return errors.New("Email invalid")}
	if err := service.identityRepostiory.Register(user); err != nil{return err}
	if err := service.mailService.SendMail("<h1>Hello!</h1>","Szalom!", user.Email); err!= nil{return err}
	return nil
}

func (service identityService) GetToken(tokenRequest TokenRequest) (*string, error){
	user, err := service.identityRepostiory.GetUserByEmail(tokenRequest.Email)
	if err!= nil{return nil, err}
	credentialError := user.CheckPassword(tokenRequest.Password)
	if credentialError != nil {return nil, credentialError}
	tokenString, err:= service.GenerateJWT(user.Email, user.Username)
	if err != nil {return nil, err}
	return &tokenString, nil
}

func (identityService) GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims:= &JWTClaim{
		Email: email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func (identityService) ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}