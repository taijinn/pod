package main

import (
	api "github.com/taijinn/pod/api"
	mgo "gopkg.in/mgo.v2"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"time"
	"net/http"
	"log"
)

const (
    AC = 0
    PENDING = 1
    DENY = 2
    PAID = 3
    CANCELLED = 4
)

type DataGettingResInfo struct {
	PhoneNum string `json:"phone"`
	UserId string `json:"userId"`
}

type DataForCheckin struct {
	RestaurantId bson.ObjectId `json:"restaurantId"`
	UserId string `json:"userId"`
	Time time.Time `json:"time"`
	TableNum int `json:"TableNum"`
}

type DataForOrder struct {
	CheckinId bson.ObjectId `json:"checkinId"`
	MenuMap map[int]int `json:"menuMap"` //menu id to number of each
	Status int `json:"status"`
}

type DataForPayment struct {
	OrderId bson.ObjectId `json:"orderId"`
	
}


func getRestaurantInfoHandler(w http.ResponseWriter, r *http.Request){
	var f DataGettingResInfo
	log.Println("access")
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

    err := json.NewDecoder(r.Body).Decode(&f)
    if err != nil {
    	http.Error(w, err.Error(), 400)
    	return
    }
    log.Println("pod start")
    log.Println(f)
    log.Println("pod end")
    
    session, err := mgo.Dial("localhost:27017")
    //session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }

    log.Printf("Connected to replica set %v!\n", session.LiveServers())
    restaurantInfo := api.GetRestaurantInfo(f.PhoneNum, f.UserId, session)
    log.Println("pod restaurantInfo")
    log.Println(*restaurantInfo)
    log.Println("pod restaurantInfo")
    json.NewEncoder(w).Encode(*restaurantInfo)
    
}
func checkinHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("in the checkin HandleFunc")
	var f DataForCheckin
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	
    err := json.NewDecoder(r.Body).Decode(&f)

    if err != nil {
    	http.Error(w, err.Error(), 400)
    	return
    }
    
 
    session, err := mgo.Dial("localhost:27017")
    //session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }

    log.Printf("Connected to replica set %v!\n", session.LiveServers())
    checkinInfo := api.Checkin(f.RestaurantId, f.UserId, f.Time, f.TableNum, session)
    log.Println("pod checkinInfo")
    log.Println(*checkinInfo)
    log.Println("pod checkinInfo")
    json.NewEncoder(w).Encode(*checkinInfo)
}
/*
func orderHandler(w http.ResponseWriter, r *http.Request) {
	var f DataForOrder
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }
    orderInfo := api.GetMenu()
    json.NewEncoder(w).Encode(*orderInfo)
}
*/
func main() {
	log.Println("Hello, new poder!")
	//http.HandleFunc("/signin", signinHandler)
	//http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/getRestaurantInfo", getRestaurantInfoHandler)
	http.HandleFunc("/checkin", checkinHandler)
	/*
	http.HandleFunc("/order", orderHandler)
	http.HandleFunc("/payment", paymentHandler)
	http.HandleFunc("/myOrderHistory", orderHistoryHandler)
	http.HandleFunc("/myPoint", pointHistoryHandler)

	http.HandleFunc("/myRestaurant", myRestaurantHandler)
	http.HandleFunc("/myRevenue", myRevenueHandler)
	*/
	http.ListenAndServe(":8080", nil)
}
