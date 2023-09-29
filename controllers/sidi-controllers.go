package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Sidi/models"
	"github.com/josepnapitupulu/API_Sidi/utils"
)

var NewSidi models.Sidi
// var tmpl * template.Template
// func init(){
// 	tmpl = template.Must(template.ParseFiles("jemaat.html"))
// }

// func Index(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/jemaat.html")
// 	temp.Execute(w, nil)
// }

// func Create(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/createJemaat.html")
// 	temp.Execute(w, nil)
// }

func GetSidi(w http.ResponseWriter, r *http.Request) {
	newSidis := models.GetAllSidis()
	res, _ := json.Marshal(newSidis)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSidiById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sidiId := vars["sidiId"]
	NIK, err := strconv.ParseInt(sidiId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	sidiDetails, _ := models.GetSidibyId(NIK)
	res, _ := json.Marshal(sidiDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSidi(w http.ResponseWriter, r *http.Request) {
	CreateSidi := &models.Sidi{}
	utils.ParseBody(r, CreateSidi)
	b := CreateSidi.CreateSidi()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSidi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sidiId := vars["sidiId"]
	ID, err := strconv.ParseInt(sidiId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	sidi := models.DeleteSidi(ID)
	res, _ := json.Marshal(sidi)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSidi(w http.ResponseWriter, r *http.Request) {
    var updateSidi = &models.Sidi{}
    utils.ParseBody(r, updateSidi)
    vars := mux.Vars(r)
    sidiId := vars["sidiId"]
    ID, err := strconv.ParseInt(sidiId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    sidiDetails, db := models.GetSidibyId(ID)
    if updateSidi.NamaJemaat != "" {
        sidiDetails.NamaJemaat = updateSidi.NamaJemaat
    }
    if updateSidi.TanggalSidi != "" {
        sidiDetails.TanggalSidi = updateSidi.TanggalSidi
    }
    if updateSidi.TanggalLahir != "" {
        sidiDetails.TanggalLahir = updateSidi.TanggalLahir
    }
    if updateSidi.NamaPendeta != "" {
        sidiDetails.NamaPendeta = updateSidi.NamaPendeta
    }
    if updateSidi.Alamat != "" {
        sidiDetails.Alamat = updateSidi.Alamat
    }
    db.Save(&sidiDetails)
    res, _ := json.Marshal(sidiDetails)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}