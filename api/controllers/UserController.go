package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dimitryduarte/honestyapi/api/auth"
	"github.com/dimitryduarte/honestyapi/api/responses"
	"github.com/dimitryduarte/honestyapi/models"
	"github.com/dimitryduarte/honestyapi/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid <= 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetUserById(w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid <= 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	vars := mux.Vars(r)
	idUser, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}

	users, err := user.FindUserByID(server.DB, idUser)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid <= 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

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

	users, err := user.SaveUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, users)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if uid <= 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	userUpdate := models.User{}
	err = json.Unmarshal(body, &userUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Check if the user exist
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", userUpdate.IdUser).Take(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Usuário não encontrado!"))
		return
	}

	userUpdate.Prepare()

	userUpdate.IdUser = user.IdUser //this is important to tell the model the product id to update, the other update field are set above

	userUpdated, err := userUpdate.UpdateAUser(server.DB, userUpdate.IdUser)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, userUpdated)
}

func (server *Server) UpdateWallet(w http.ResponseWriter, r *http.Request) {

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if uid <= 0 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	userUpdate := models.User{}
	err = json.Unmarshal(body, &userUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Check if the user exist
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", userUpdate.IdUser).Take(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Usuário não encontrado"))
		return
	}

	userUpdate.Prepare()

	userUpdate.IdUser = user.IdUser //this is important to tell the model the product id to update, the other update field are set above

	userUpdated, err := userUpdate.UpdateWallet(server.DB, userUpdate.IdUser)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, userUpdated)
}

// func (server *Server) RechargeWallet(w http.ResponseWriter, r *http.Request) {

// 	//CHeck if the auth token is valid and  get the user id from it
// 	uid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}

// 	if uid <= 0 {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}

// 	// Read the data posted
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	// Start processing the request data
// 	userUpdate := models.Recharge{}
// 	err = json.Unmarshal(body, &userUpdate)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	// Check if the user exist
// 	user := models.User{}
// 	err = server.DB.Debug().Model(models.User{}).Where("id = ?", userUpdate.IdUser).Take(&user).Error
// 	if err != nil {
// 		responses.ERROR(w, http.StatusNotFound, errors.New("Usuário não encontrado"))
// 		return
// 	}

// 	userUpdate.Prepare()

// 	userUpdate.IdUser = user.IdUser //this is important to tell the model the product id to update, the other update field are set above

// 	userUpdated, err := userUpdate.UpdateWallet(server.DB, userUpdate.IdUser)

// 	if err != nil {
// 		formattedError := formaterror.FormatError(err.Error())
// 		responses.ERROR(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, userUpdated)
// }
