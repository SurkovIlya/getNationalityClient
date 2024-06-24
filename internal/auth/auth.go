package auth

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"getNationalClient/internal/model"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type SingUp interface {
	Registration() model.Auth
}
type SingIn interface {
	Authorization(dataAuth model.Auth) (string, error)
}

type Auth struct {
	Reg   SingUp
	Auth  SingIn
	Users map[string]model.UserData
}

var user model.UserData

func NewAuth() *Auth {
	users := make(map[string]model.UserData)

	file, err := os.Open("././data/authData.csv")
	if err != nil {
		log.Println("File authData.csv is not found:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Could not read:", err)
		}
		user.Name = record[0]
		user.Login = record[1]
		user.Hash = []byte(record[2])
		users[user.Login] = user
	}

	return &Auth{
		Users: users,
	}
}

func (au *Auth) Registration(data model.Reg) (model.Auth, error) {
	var authData model.Auth

	pas := []byte(data.Password)
	hash, _ := bcrypt.GenerateFromPassword(pas, 10)
	user.Name = data.Name
	user.Login = data.Login
	user.Hash = hash

	if _, ok := au.Users[data.Login]; !ok {
		au.Users[data.Login] = user

		file, err := os.OpenFile("././data/authData.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Println("File authData.csv is not found:", err)

			return model.Auth{}, fmt.Errorf("регистрация пользователя невозможна! Обратитесь к разработчику")
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		defer writer.Flush()

		userData := [][]string{{user.Name, user.Login, string(user.Hash)}}
		writer.WriteAll(userData)

		authData.Login = data.Login
		authData.Password = data.Password

		return authData, nil
	}

	return model.Auth{}, fmt.Errorf("пользователь уже существует")
}

func (au *Auth) Authorization(dataAuth model.Auth) (string, error) {
	var userToken string
	const ttl = 86500 * time.Second

	if val, ok := au.Users[dataAuth.Login]; !ok {
		return "", fmt.Errorf("пользовтель не зарегистрирован")
	} else {
		err := bcrypt.CompareHashAndPassword(val.Hash, []byte(dataAuth.Password))
		if err != nil {
			return "", fmt.Errorf("неверный пароль")
		}

		tokenBytes := []byte(dataAuth.Login + dataAuth.Password)
		userToken = base64.StdEncoding.EncodeToString(tokenBytes)

		file, err := os.OpenFile("././data/dbToken.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Println("File dbToken.csv is not found:", err)

			return "", fmt.Errorf("ошибка авторизации")
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		defer writer.Flush()

		userData := [][]string{{userToken, fmt.Sprint(ttl)}}
		writer.WriteAll(userData)
	}

	return userToken, nil
}

func (au Auth) CheckToken(token string) error {
	file, err := os.Open("././data/dbToken.csv")
	if err != nil {
		log.Println("File dbToken.csv is not found:", err)

		return fmt.Errorf("пользователь не авторизирован")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Could not read:", err)

			return fmt.Errorf("пользователь не авторизирован")
		}

		if record[0] == token {
			return nil
		}
	}
	return fmt.Errorf("пользователь не авторизирован")
}
