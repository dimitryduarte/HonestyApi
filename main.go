package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	IdUser   uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	username string `gorm:"column:username"`
	password string `gorm:"column:password"`
}

type Logins struct {
	username string
	password string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}

type Company struct {
	IdCompany   int64  `json:"idcompany" gorm:"column:id;primaryKey;autoIncrement"`
	CompanyName string `json:"companyname" gorm:"column:companyname"`
	Name        string `json:"name" gorm:"column:name"`
	Cnpj        string `json:"cnpj" gorm:"column:cnpj"`
	TaxesId     int64  `json:"taxesid" gorm:"column:TaxesId"`
	Taxes       Taxes  `json:"taxes" gorm:"foreignKey:TaxesId;References:id"`
}

type Taxes struct {
	TaxesId int64   `json:"taxesid" gorm:"column:id;primaryKey;autoIncrement"`
	Taxa1   float64 `json:"taxa1" gorm:"column:taxa1"`
	Taxa2   float64 `json:"taxa2" gorm:"column:taxa2"`
	Taxa3   float64 `json:"taxa3" gorm:"column:taxa3"`
}

var err error
var company Company
var taxes Taxes
var dsn = "test_user:123456@tcp(127.0.0.1:3306)/honestyapp"
var client *redis.Client

//Endpoints

func GetCompany(resp http.ResponseWriter, req *http.Request) {
	var empresa []Company
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	result := db.Find(&empresa)

	if result.RowsAffected == 0 {
		json.NewEncoder(resp).Encode("{errorMessage: Não foram encontradas empresas para o Id informado}")
		return
	} else {
		json.NewEncoder(resp).Encode(empresa)
		return
	}
}

func GetTaxes(resp http.ResponseWriter, req *http.Request) {
	var taxes []Taxes
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	result := db.Find(&taxes)

	if result.RowsAffected == 0 {
		json.NewEncoder(resp).Encode("{errorMessage: Não foram encontradas taxas para o Id informado}")
		return
	} else {
		json.NewEncoder(resp).Encode(taxes)
		return
	}
}

func GetCompanyId(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var empresa Company
	var taxas Taxes
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Find(&empresa, params["id"])
	db.Find(&taxas, empresa.TaxesId)
	empresa.Taxes = taxas
	json.NewEncoder(resp).Encode(empresa)
}

func GetTaxesId(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var taxas Taxes
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Find(&taxas, params["id"])
	json.NewEncoder(resp).Encode(taxas)
}

func CreateCompany(resp http.ResponseWriter, req *http.Request) {
	var newcompany Company

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &newcompany)

	db.Create(&newcompany)
	json.NewEncoder(resp).Encode(newcompany.IdCompany)
}

func CreateTaxes(resp http.ResponseWriter, req *http.Request) {
	var newtaxes Taxes

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &newtaxes)

	db.Create(&newtaxes)

}

func UpdateCompany(rep http.ResponseWriter, req *http.Request) {
	var newcompany Company
	var empresa Company
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &newcompany)

	db.Find(&empresa, &newcompany.IdCompany)

	empresa = newcompany

	db.Save(&empresa)
}

func UpdateTaxes(rep http.ResponseWriter, req *http.Request) {
	var newtax Taxes
	var taxes Taxes
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &newtax)

	db.Find(&taxes, &newtax.TaxesId)

	taxes = newtax

	db.Save(&taxes)
}

func DeleteCompany(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Delete(&Company{}, params["id"])

}

func DeleteTax(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Delete(&Taxes{}, params["id"])

}

// função principal
func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(
		&Taxes{},
		&Company{},
		&Users{},
	)
	//Seed
	//db.&User{}).Save((&User{username: "admin", password: "123456", status: "A"}))

	router := mux.NewRouter()

	//Declaração dos Endpoints
	//GET
	router.HandleFunc("/company", GetCompany).Methods("GET")
	router.HandleFunc("/taxes", GetTaxes).Methods("GET")
	router.HandleFunc("/company/{id}", GetCompanyId).Methods("GET")
	router.HandleFunc("/taxes/{id}", GetTaxesId).Methods("GET")

	//POST
	router.HandleFunc("/company", CreateCompany).Methods("POST")
	router.HandleFunc("/taxes", CreateTaxes).Methods("POST")

	//PUT
	router.HandleFunc("/company", UpdateCompany).Methods("PUT")
	router.HandleFunc("/taxes", UpdateTaxes).Methods("PUT")

	//DELETE
	router.HandleFunc("/company/{id}", DeleteCompany).Methods("DELETE")
	router.HandleFunc("/taxes/{id}", DeleteTax).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
