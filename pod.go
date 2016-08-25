package main

import (
	api "github.com/taijinn/pod/api"
	mgo "gopkg.in/mgo.v2"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"time"
	"net/http"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/taijinn/pod/jwt-go"
	"fmt"
	"bytes"
	"io/ioutil"
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



type LoginResponse struct {
	Success bool `json:"success"`
	Token string `json:"token"`
	NoUser bool `json:"noUser"`
	WrongPass bool `json:"wrongPass"`
}



const (
	privKeyPath = "path/to/keys/app.rsa"
	pubKeyPath = "path/to/keys/app.rsa.pub"
)

var VerifyKey, SignKey []byte


func initKeys(){
	var err error

	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

func clear(b []byte) {
    for i := 0; i < len(b); i++ {
        b[i] = 0;
    }
}

func Crypt(password []byte) ([]byte, error) {
    defer clear(password)
    return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}


func loginHandler(w http.ResponseWriter, r *http.Request){
	var user api.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(w, "Error in request")
		return
	}
	session, err := mgo.Dial("localhost:27017")
    //session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }

    log.Printf("Connected to replica set %v!\n", session.LiveServers())
	isRegistered, password := api.GetEncriptedPass(user.Username, session)
	var respo LoginResponse
	if !isRegistered {
		respo = LoginResponse{false, "", true, false}
		log.Println("No such user")
	}else {
		//todo
		ctext, err := Crypt([]byte(user.Password))
		fmt.Println(user.Password)
		fmt.Println(ctext)
		if err != nil {
			log.Fatal(err)
		}
		if !bytes.Equal(password, ctext) {
			respo = LoginResponse{false, "", false, true}
			log.Println(user.Password + " is a wrong password")
		}else{
			signer := jwt.New(jwt.GetSigningMethod("RS256"))

			//set claims
			claims := make(jwt.MapClaims)
			claims["iss"] = "admin"
			claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
			claims["CustomUserInfo"] = struct {
				Name	string
				Role	string
			}{user.Username, "Member"}
			signer.Claims = claims

			tokenString, err := signer.SignedString(SignKey)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error while signing the token")
				log.Printf("Error signing token: %v\n", err)
			}
			respo = LoginResponse{true, tokenString, false, false}
			log.Println(tokenString)
		}
	}
	json.NewEncoder(w).Encode(respo)
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var user api.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(w, "Error in request")
		return
	}
	session, err := mgo.Dial("localhost:27017")
    //session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }

    log.Printf("Connected to replica set %v!\n", session.LiveServers())
    fmt.Println("init")
    fmt.Println(user.Password)
    var res = api.GetSigninAPI(&user, session)
    if res == "login" {
    	loginHandler(w, r)
    }else if res == "forget" {
    	 //forgetHandler(w, r)
    }
    json.NewEncoder(w).Encode(res)
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
	http.HandleFunc("/signin", signinHandler)
	http.HandleFunc("/login", loginHandler)
	//http.HandleFunc("/forgetPass", forgetPassHandler)
	//http.HandleFunc("/loginFB", loginFBHandler)
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
