package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Offices []Office

type Office struct {
	id            int
	SalePointName string `json:"salePointName"`
	Address       string `json:"address"`
	Status        string `json:"status"`
	OpenHours     []struct {
		Days  string `json:"days"`
		Hours string `json:"hours"`
	} `json:"openHours"`
	RKO                 string `json:"rko"`
	OpenHoursIndividual []struct {
		Days  string `json:"days"`
		Hours string `json:"hours"`
	} `json:"openHoursIndividual"`
	OfficeType            string  `json:"officeType"`
	SalePointFormat       string  `json:"salePointFormat"`
	SUOAvailability       string  `json:"suoAvailability"`
	HasRamp               string  `json:"hasRamp"`
	Latitude              float64 `json:"latitude"`
	Longitude             float64 `json:"longitude"`
	MetroStation          *string `json:"metroStation"`
	Distance              int     `json:"distance"`
	KEP                   bool    `json:"kep"`
	MyBranch              bool    `json:"myBranch"`
	NumWindowsIndividuals int     `json:"numWindowsIndividuals"`
	NumWindowsEntities    int     `json:"numWindowsEntities"`
	NumWindowsPrime       int     `json:"numWindowsPrime"`
	NumIndividualClients  int     `json:"numIndividualClients"`
	NumEntitiesClients    int     `json:"numEntitiesClients"`
	NumWindowsPrivelege   int     `json:"numWindowsPrivelege"`
	NumPrivelegeClients   int     `json:"numPrivelegeClients"`
}

// TO DO
func GetNumById(branchId string) int {

	return 2
} // получение загруженности отделения по его id

type OneBranchInfo struct {
	Cur  int     // текущая загруженность
	Tble [][]int // загруженность на каждый час на 7 дней
}

func GetOneBranch(w http.ResponseWriter, r *http.Request) {
	branchId := mux.Vars(r)["branchId"]
	cur := GetNumById(branchId)
	tble := make([][]int, 7)

	for i := 0; i < 7; i++ {
		tble[i] = make([]int, 24)
	}

	res := &OneBranchInfo{}
	res.Cur = cur
	res.Tble = tble

	b, err := json.Marshal(res)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", b)
}

func GetBranchesByFilter(w http.ResponseWriter, r *http.Request) {
	isOffice := r.URL.Query().Get("isOffice")
	qr := r.URL.Query().Get("qr")
	nfc := r.URL.Query().Get("nfc")
	blind := r.URL.Query().Get("blind")
	wheelchair := r.URL.Query().Get("wheelchair")
	face := r.URL.Query().Get("face")
	allday := r.URL.Query().Get("allday")

}
func GetRecomBranch(w http.ResponseWriter, r *http.Request) {

}
func main() {
	r := mux.NewRouter()
	file, err := os.Open("updated_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Offices)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(Offices); i++ {
		Offices[i].id = i
	}

	r.HandleFunc("/api/branches/{branchId}", GetOneBranch)
	r.HandleFunc("/api/branches", GetBranchesByFilter)
	r.HandleFunc("/api/branches/recommended", GetRecomBranch)
	{
		err := http.ListenAndServe(":80", r)

		if err != nil {
			log.Fatal(err)
		}
	}

}
