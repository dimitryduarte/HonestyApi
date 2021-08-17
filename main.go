package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/dimitryduarte/honestyapi/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
var users models.Users
var products models.Product

var db_host = os.Getenv("DB_HOST")
var db_user = os.Getenv("DB_USER")
var db_password = os.Getenv("DB_PASSWORD")
var db_name = os.Getenv("DB_NAME")
var db_port = os.Getenv("DB_PORT")
var apisecret = os.Getenv("API_SECRET")
var tempToken = os.Getenv("TOKEN")

//var dsn = "test_user:123456@tcp(127.0.0.1:3306)/honestyapp"

//var dsn = "bc3ac486906125:9da0ccaf@tcp(us-cdbr-east-04.cleardb.com:3306)/heroku_94037f830475225"

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_host, db_port, db_name)

// função principal
func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")

	db.AutoMigrate(
		//&models.Users{},
		&models.Product{},
	)
	//Seed
	//db.&User{}).Save((&User{username: "admin", password: "123456", status: "A"}))

	router := mux.NewRouter()

	//Declaração dos Endpoints
	//GET
	router.HandleFunc("/product", GetProduct).Methods("GET")
	router.HandleFunc("/product/{id}", GetProductId).Methods("GET")

	//POST
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/product", CreateProduct).Methods("POST")

	//PUT
	router.HandleFunc("/product", UpdateProduct).Methods("PUT")

	//DELETE
	router.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, router))
}

//Endpoints

func Login(resp http.ResponseWriter, req *http.Request) {
	var token models.TokenDetails
	var login models.Logins

	reqBody, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(reqBody, &login)

	if login.Email == "honestybox@yahoo.com.br" && login.Password == "H0n3styB0X" {
		token.AccessToken = tempToken
		token.UserName = "Dimitry"
		token.Wallet = (rand.Float32() * 100)
		json.NewEncoder(resp).Encode(token)
	} else {
		json.NewEncoder(resp).Encode("Usuário ou senha inválido!")
	}
}

func GetProduct(resp http.ResponseWriter, req *http.Request) {
	if req.Header.Get("accessToken") != tempToken {
		json.NewEncoder(resp).Encode("errorMessage: A autenticação falhou, verifique o accessToken informado")
		return
	}

	var products []models.Product
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	result := db.Find(&products)

	if result.RowsAffected == 0 {
		json.NewEncoder(resp).Encode("{errorMessage: Não foram encontrados produtos para o Id informado}")
		return
	} else {
		json.NewEncoder(resp).Encode(products)
		return
	}
}

func GetProductId(resp http.ResponseWriter, req *http.Request) {

	if req.Header.Get("accessToken") != tempToken {
		json.NewEncoder(resp).Encode("errorMessage: A autenticação falhou, verifique o accessToken informado")
		return
	}

	params := mux.Vars(req)
	var product models.Product
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	result := db.Find(&product, params["id"])

	if result.RowsAffected == 0 {
		json.NewEncoder(resp).Encode("{errorMessage: Não foram encontrados produtos para o Id informado}")
		return
	} else {
		json.NewEncoder(resp).Encode(product)

	}
}

func CreateProduct(resp http.ResponseWriter, req *http.Request) {

	if req.Header.Get("accessToken") != tempToken {
		json.NewEncoder(resp).Encode("errorMessage: A autenticação falhou, verifique o accessToken informado")
		return
	}

	var newProduct models.Product

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &newProduct)

	db.Create(&newProduct)
	json.NewEncoder(resp).Encode(newProduct.IdProduct)
}

func UpdateProduct(rep http.ResponseWriter, req *http.Request) {

	if req.Header.Get("accessToken") != tempToken {
		json.NewEncoder(rep).Encode("errorMessage: A autenticação falhou, verifique o accessToken informado")
		return
	}

	var newproduct models.Product
	var produto models.Product
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &newproduct)

	db.Find(&produto, &newproduct.IdProduct)

	produto = newproduct

	db.Save(&produto)
}

func DeleteProduct(resp http.ResponseWriter, req *http.Request) {

	if req.Header.Get("accessToken") != tempToken {
		json.NewEncoder(resp).Encode("errorMessage: A autenticação falhou, verifique o accessToken informado")
		return
	}

	params := mux.Vars(req)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Delete(&models.Product{}, params["id"])
}
