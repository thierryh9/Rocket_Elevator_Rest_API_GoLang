package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Elevator struct {
	Id             int    `json:"id"`
	Status         string `json:"status"`
	SerialNumber   string `json:"serialNumber"`
	InspectionDate string `json:"inspectiondate"`
	InstallDate    string `json:"installDate"`
	Certificat     string `json:"certificat"`
	Information    string `json:"information"`
	Note           string `json:"note"`
	Type           string `json:"type"`
	Column_id      int    `json:"column_id"`
	Category_id    int    `json:"category_id"`
}

type Column struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type Battery struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type Building struct {
	Id        int    `json:"id"`
	FullName  string `json:"fullName"`
	CellPhone string `json:"cellPhone"`
	Email     string `json:"email"`
	TechEmail string `json:"techEmail"`
	TechName  string `json:"techName"`
	TechPhone string `json:"techPhone"`
	Customer  int    `json:"customer_id"`
	Address   int    `json:"address_id"`
}

type Lead struct {
	Id             int    `json:"id"`
	FullName       string `json:"fullName"`
	EntrepriseName string `json:"entrepriseName"`
	CellPhone      string `json:"cellPhone"`
	ProjectName    string `json:"projectName"`
	Description    string `json:"description"`
	Type           string `json:"type"`
}

type Building_detail {
	Building_id: 	int    `json:"building_id"`
	InfoKey: 		string `json:"infoKey"`
	InfoValue: 		string `json:"infoValue"`

func main() {
	fmt.Println("allo")
	handleRequests()
}

func getElevatorList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	results, err := db.Query("SELECT e.id, s.name as 'status', e.serialNumber, e.inspectionDate, e.installDate, e.certificat, e.information, e.note, t.name AS 'type', e.column_id, e.category_id FROM elevators e JOIN statuses s ON s.id=e.status_id JOIN types t ON t.id=e.type_id WHERE s.name !='active';")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	elevators := []Elevator{}
	for results.Next() {
		var e Elevator
		// for each row, scan the result into our tag composite object
		err = results.Scan(&e.Id, &e.Status, &e.SerialNumber, &e.InspectionDate, &e.InstallDate, &e.Certificat, &e.Information, &e.Note, &e.Type, &e.Column_id, &e.Category_id)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		elevators = append(elevators, e)
	}
	json.NewEncoder(w).Encode(elevators)
}

func getElevator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	vars := mux.Vars(r)
	key := vars["id"]
	results, err := db.Query("SELECT e.id, s.name as 'status', e.serialNumber, e.inspectionDate, e.installDate, e.certificat, e.information, e.note, t.name AS 'type', e.column_id, e.category_id FROM elevators e JOIN statuses s ON s.id=e.status_id JOIN types t ON t.id=e.type_id WHERE e.id = " + key + ";")

	if results.Next() {
		var e Elevator
		// for each row, scan the result into our tag composite object
		err = results.Scan(&e.Id, &e.Status, &e.SerialNumber, &e.InspectionDate, &e.InstallDate, &e.Certificat, &e.Information, &e.Note, &e.Type, &e.Column_id, &e.Category_id)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		json.NewEncoder(w).Encode(e)
	}
}

func updateElevator(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	vars := mux.Vars(r)
	id := vars["id"]
	status := vars["status"]
	results, err := db.Query("UPDATE elevators c SET c.status_id = (SELECT s.id FROM statuses s WHERE s.name='" + status + "') WHERE c.id = " + id + ";")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	getElevator(w, r)
	defer results.Close()
}

func getColumn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	vars := mux.Vars(r)
	key := vars["id"]
	results, err := db.Query("SELECT c.id, s.name FROM columns c JOIN statuses s ON c.status_id=s.id WHERE c.id = " + key + ";")

	if results.Next() {
		var c Column
		// for each row, scan the result into our tag composite object
		err = results.Scan(&c.Id, &c.Status)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		json.NewEncoder(w).Encode(c)
	}
}

func updateColumn(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	vars := mux.Vars(r)
	id := vars["id"]
	status := vars["status"]
	results, err := db.Query("UPDATE columns c SET c.status_id = (SELECT s.id FROM statuses s WHERE s.name='" + status + "') WHERE c.id = " + id + ";")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	getColumn(w, r)
	defer results.Close()
}

func getBattery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	vars := mux.Vars(r)
	key := vars["id"]
	results, err := db.Query("SELECT b.id, s.name FROM batteries b JOIN statuses s ON b.status_id=s.id WHERE b.id =" + key + ";")

	if results.Next() {
		var b Battery
		// for each row, scan the result into our tag composite object
		err = results.Scan(&b.Id, &b.Status)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		json.NewEncoder(w).Encode(b)
	}
}

func updateBattery(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	vars := mux.Vars(r)
	id := vars["id"]
	status := vars["status"]
	results, err := db.Query("UPDATE batteries c SET c.status_id = (SELECT s.id FROM statuses s WHERE s.name='" + status + "') WHERE c.id = " + id + ";")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	getBattery(w, r)
	defer results.Close()
}

func getBuildingList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	results, err := db.Query("SELECT DISTINCT b.id, b.fullName, b.cellPhone, b.email,b.techEmail, b.techName, b.techPhone, b.customer_id, b.address_id FROM buildings b JOIN batteries b2 ON b2.building_id=b.id JOIN `columns` c ON b2.id=c.battery_id JOIN elevators e ON e.column_id=c.id WHERE b2.status_id=(SELECT s2.id FROM statuses s2 WHERE s2.name='intervention') OR c.status_id=(SELECT s2.id FROM statuses s2 WHERE s2.name='intervention') OR e.status_id=(SELECT s2.id FROM statuses s2 WHERE s2.name='intervention');")
	buildings := []Building{}
	for results.Next() {
		var e Building
		// for each row, scan the result into our tag composite object
		err = results.Scan(&e.Id, &e.FullName, &e.CellPhone, &e.Email, &e.TechEmail, &e.TechName, &e.TechPhone, &e.Customer, &e.Address)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		buildings = append(buildings, e)
	}
	json.NewEncoder(w).Encode(buildings)
}

func getLeadList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	results, err := db.Query("SELECT DISTINCT l.id, l.fullName, l.entrepriseName, l.cellPhone, l.projectName, l.description, t.name as 'type' FROM leads l JOIN customers c ON c.email!=l.email JOIN types t ON t.id=l.type_id WHERE DATEDIFF(l.created_at,CURDATE()) <= 30;")
	ls := []Lead{}
	for results.Next() {
		var l Lead
		// for each row, scan the result into our tag composite object
		err = results.Scan(&l.Id, &l.FullName, &l.EntrepriseName, &l.CellPhone, &l.ProjectName, &l.Description, &l.Type)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		ls = append(ls, l)
	}
	json.NewEncoder(w).Encode(ls)
}

func updateTechPhone(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	vars := mux.Vars(r)
	id := vars["id"]
	phone := vars["phone"]
	results, err := db.Query("UPDATE buildings SET techPhone='" + phone + "' WHERE id = " + id + ";")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	getBuildingList(w, r)
	defer results.Close()
}

func getBuildingDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "codeboxx:Codeboxx1!@tcp(codeboxx.cq6zrczewpu2.us-east-1.rds.amazonaws.com:3306)/ThierryHarvey")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	vars := mux.Vars(r)
	key := vars["id"]
	bd := []Building_detail{}
	results, err := db.Query("SELECT building_details.id, building_details.infoKey, building_details.infoValue FROM building_details JOIN buildings b ON building_details.id=b.id WHERE b.id =" + key + ";")
	for results.Next() {
		var b Building_detail
		// for each row, scan the result into our tag composite object
		err = results.Scan(&b.Building_id, &b.InfoKey, &b.InfoValue)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		bd=append(bd,b)
	}
	json.NewEncoder(w).Encode(bd)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/api/elevator/{id}", getElevator).Methods("GET")
	myRouter.HandleFunc("/api/elevator/", getElevatorList).Methods("GET")
	myRouter.HandleFunc("/api/column/{id}", getColumn).Methods("GET")
	myRouter.HandleFunc("/api/battery/{id}", getBattery).Methods("GET")
	myRouter.HandleFunc("/api/lead/", getLeadList).Methods("GET")
	myRouter.HandleFunc("/api/building/", getBuildingList).Methods("GET")
	myRouter.HandleFunc("/api/elevator/{id}/{status}", updateElevator).Methods("PUT")
	myRouter.HandleFunc("/api/column/{id}/{status}", updateColumn).Methods("PUT")
	myRouter.HandleFunc("/api/battery/{id}/{status}", updateBattery).Methods("PUT")
	myRouter.HandleFunc("/api/building/{id}/{phone}", updateTechPhone).Methods("PUT")
	myRouter.HandleFunc("/api/buildingdetails/{id}", getBuildingDetails).Methods("GET")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
} 


