package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RipulHandoo/jwt/internal/database"
	_ "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegisteredUser(w http.ResponseWriter,r *http.Request){
	type parameters struct{
		Email string
		Passwrod string
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil{
		ResponseWithError(w,http.StatusInternalServerError,err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(params.Passwrod),10)

	if err != nil{
		hash = []byte(params.Passwrod)
	}

	apiConfig := DBInstance()
	user, errToDB := apiConfig.CreateUser(r.Context(),database.CreateUserParams{
		Email: params.Email,
		Password: string(hash),
	})

	if errToDB != nil{
		ResponseWithError(w,http.StatusInternalServerError,err)
		return
	}

	token, expiryTime, jwtTokenError := GetJwt(Credentials{
		Email: params.Email,
	})

	if jwtTokenError != nil {
		ResponseWithError(w, http.StatusUnauthorized, jwtTokenError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "auth_token",
		Value:   token,
		Expires: expiryTime,
		Path:    "/",
	})



	ResponseWithJson(w,200,user)

}

func login(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Email string
		Passwrod string
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil{
		ResponseWithError(w,http.StatusInternalServerError,err)
		return
	}

	apiConfig := DBInstance()
	user, err := apiConfig.GetUserByEmail(r.Context(),params.Email)

	if err != nil{
		ResponseWithError(w,http.StatusInternalServerError,err)
		return
	}

	hashPassword:= bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(params.Passwrod))

	if hashPassword != nil {
		err = fmt.Errorf("incorrect password")
		ResponseWithError(w, http.StatusUnauthorized, err)
		return
	}
	

	jwtToken, expiryTime, tokenErr := GetJwt(Credentials{
		Email: user.Email,
	})

	if tokenErr != nil {
		ResponseWithError(w, http.StatusForbidden, tokenErr)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "auth_token",
		Value:   jwtToken,
		Expires: expiryTime,
		Path:    "/",
	})


	ResponseWithJson(w,200,user)

}



func HandleUserLogout(w http.ResponseWriter, req *http.Request, user database.Jwtuser) {
	//clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "auth_token",
		Value: "",
		Path:  "/",
	})
	ResponseWithJson(w, http.StatusAccepted, user)
}