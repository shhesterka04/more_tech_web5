package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"moretech-backend/more_tech_web5/maps"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var Offices []Office
var Atms []Atm

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
type Services struct {
	Wheelchair struct {
		ServiceCapability string `json:"serviceCapability"`
		ServiceActivity   string `json:"serviceActivity"`
	} `json:"wheelchair"`
	Blind struct {
		ServiceCapability string `json:"serviceCapability"`
		ServiceActivity   string `json:"serviceActivity"`
	} `json:"blind"`
	NfcForBankCards struct {
		ServiceCapability string `json:"serviceCapability"`
		ServiceActivity   string `json:"serviceActivity"`
	} `json:"nfcForBankCards"`
	QrRead struct {
		ServiceCapability string `json:"serviceCapability"`
		ServiceActivity   string `json:"serviceActivity"`
	} `json:"qrRead"`
	SupportsUsd struct {
		ServiceCapability string `json:"serviceCapability"`
		ServiceActivity   string `json:"serviceActivity"`
	} `json:"supportsUsd"`
	SupportsChargeRub struct {
		ServiceCapability string `json:"serviceCapability"`
		ServiceActivity   string `json:"serviceActivity"`
	} `json:"supportsChargeRub"`
	SupportsEur struct {
		ServiceCapability string `json:"serviceCapability"`
		ServiceActivity   string `json:"serviceActivity"`
	} `json:"supportsEur"`
	SupportsRub struct {
		ServiceCapability string `json:"serviceCapability"`
		ServiceActivity   string `json:"serviceActivity"`
	} `json:"supportsRub"`
}
type Atm struct {
	id        int
	Address   string   `json:"address"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	AllDay    bool     `json:"allDay"`
	Services  Services `json:"services"`
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
	data := make(map[string]string)

	data["isOffice"] = r.URL.Query().Get("isOffice")
	data["qr"] = r.URL.Query().Get("qr")
	data["nfc"] = r.URL.Query().Get("nfc")
	data["blind"] = r.URL.Query().Get("blind")
	data["wheelchair"] = r.URL.Query().Get("wheelchair")
	data["face"] = r.URL.Query().Get("face")
	data["allday"] = r.URL.Query().Get("allday")
	data["officetype"] = r.URL.Query().Get("officetype")
	flag := true

	for _, v := range data {
		if v != "" {
			flag = false
			break
		}
	}
	if flag {
		a, err := json.Marshal(Atms)
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.Marshal(Offices)

		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
		w.Write(a)
	}
	if data["isOffice"] == "1" {
		var Out []Office
		for i := 0; i < len(Offices); i++ {
			if data["officetype"] == "1" && (strings.Contains(Offices[i].OfficeType, "Да") || strings.Contains(Offices[i].OfficeType, "да")) ||
				Offices[i].SalePointFormat == "Универсальный" ||
				data["face"] == "0" && Offices[i].SalePointFormat == "Розничный" {
				Out = append(Out, Offices[i])
			}

		}
		b, err := json.Marshal(Out)

		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	} else {
		var Out []Atm
		for i := 0; i < len(Atms); i++ {
			if data["qr"] == "1" && Atms[i].Services.QrRead.ServiceCapability != "UNSUPPORTED" ||
				data["nfc"] == "1" && Atms[i].Services.NfcForBankCards.ServiceCapability != "UNSUPPORTED" ||
				data["blind"] == "1" && Atms[i].Services.Blind.ServiceCapability != "UNSUPPORTED" ||
				data["wheelchair"] == "1" && Atms[i].Services.Wheelchair.ServiceCapability != "UNSUPPORTED" ||
				data["allday"] == "1" && Atms[i].AllDay {
				Out = append(Out, Atms[i])
			}
		}
		b, err := json.Marshal(Out)

		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}
func GetRecomBranch(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // Закрыть тело запроса после чтения

	var data maps.MapRoute
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	result, err := maps.FetchRoute(data.Start, data.End, data.TransportType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	r := mux.NewRouter()
	file1, err := os.Open("offices.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()

	file2, err := os.Open("atms.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()
	err = json.NewDecoder(file1).Decode(&Offices)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(file2).Decode(&Atms)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(Offices); i++ {
		Offices[i].id = i
	}
	for i := 0; i < len(Atms); i++ {
		Atms[i].id = i
	}
	r.HandleFunc("/api/branch/{branchId}", GetOneBranch)
	r.HandleFunc("/api/branches", GetBranchesByFilter)
	r.HandleFunc("/api/branches/recommended", GetRecomBranch)
	{
		err := http.ListenAndServe(":80", handlers.CORS()(r))

		if err != nil {
			log.Fatal(err)
		}
	}

}
