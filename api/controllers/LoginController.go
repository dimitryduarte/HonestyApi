package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dimitryduarte/honestyapi/api/auth"
	"github.com/dimitryduarte/honestyapi/api/responses"
	"github.com/dimitryduarte/honestyapi/models"
	"github.com/dimitryduarte/honestyapi/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (models.TokenDetails, error) {

	var err error

	user := models.User{}
	authToken := models.TokenDetails{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return authToken, err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return authToken, err
	}

	token, err := auth.CreateToken(uint32(user.IdUser))
	if err != nil {
		return authToken, err
	}

	authToken.AccessToken = token
	authToken.UserName = user.Name
	authToken.Company = user.Company
	authToken.Sector = user.Sector
	authToken.Wallet = float32(user.Wallet)

	return authToken, nil
}
