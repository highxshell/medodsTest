package jwt

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"io"
	models "medodsTest/services/jwt/domain/jwt"
	"os"
	"time"
)

type UseCaseImpl struct {
	refreshTokenCollection *mongo.Collection
}

func NewUseCaseImpl(coll *mongo.Collection) *UseCaseImpl {
	return &UseCaseImpl{
		refreshTokenCollection: coll,
	}
}

func (j *UseCaseImpl) CreateToken(guid string) (models.Token, error) {
	var err error

	claims := jwt.MapClaims{}
	claims["guid"] = guid
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	jwt := models.Token{}

	jwt.AccessToken, err = token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return jwt, err
	}

	return j.createRefreshToken(jwt)
}

func (*UseCaseImpl) ValidateToken(accessToken string) (string, error) {
	var guid string
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return guid, fmt.Errorf("failed to parse token: %v", token)
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		guid = payload["guid"].(string)

		return guid, nil
	}

	return guid, errors.New("invalid token")
}

func (j *UseCaseImpl) ValidateRefreshToken(model models.Token) (string, error) {
	var guid string
	sha1 := sha1.New()
	io.WriteString(sha1, os.Getenv("SECRET_KEY"))
	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return guid, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return guid, err
	}

	data, err := base64.URLEncoding.DecodeString(model.RefreshToken)
	if err != nil {
		return guid, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return guid, err
	}

	if string(plain) != model.AccessToken {
		return guid, errors.New("invalid token")
	}

	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(model.AccessToken, claims)

	if err != nil {
		return guid, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return guid, errors.New("invalid token")
	}
	guid = payload["guid"].(string)

	return guid, nil
}

func (j *UseCaseImpl) createRefreshToken(token models.Token) (models.Token, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, os.Getenv("SECRET_KEY"))

	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		fmt.Println(err.Error())

		return token, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return token, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return token, err
	}
	h, err := bcrypt.GenerateFromPassword([]byte(token.RefreshToken), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	_, err = j.refreshTokenCollection.InsertOne(context.TODO(), bson.M{"refreshTokenHash": h})
	if err != nil {
		return token, err
	}
	token.RefreshToken = base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(token.AccessToken), nil))

	return token, nil
}
