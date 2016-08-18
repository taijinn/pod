package main

import (
	"fmt"
	api "github.com/taijinn/pod/api"
	//"gopkg.in/mgo.v2"
	"encoding/json"
	//"log"
	"net/http"
	"bytes"
	time "time"
	"gopkg.in/mgo.v2/bson"
	//"os"
)
type DataGettingResInfo struct {
	PhoneNum string `json:"phone"`
	UserId string `json:"userId"`
}
type DataCheckinInfo struct {
	UserId string `json:"userId"`
	RestaurantId bson.ObjectId `json:"restaurantId"`
	Time time.Time `json:"time"`
	TableNum int `json:"tableNum"`
}
func getRestaurantInfoTest() {
	u := DataGettingResInfo{PhoneNum: "12345678", UserId: "taijin"}
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(u)
    //test for api.GetRestaurantInfo
    res, _ := http.Post("http://localhost:8080/getRestaurantInfo", "application/json; charset=utf-8", b)
    var body api.RestaurantInfo
    json.NewDecoder(res.Body).Decode(&body)
    fmt.Println("test side")
    fmt.Println(body)
    fmt.Println("test side")
}

func checkinTest(){
	u := DataCheckinInfo{UserId:"taijin", RestaurantId:bson.NewObjectId(), Time:time.Now(), TableNum:12}
	fmt.Println("test side")
    fmt.Println(u)
    fmt.Println("test side")
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	//test for api.CheckinInfo
	res, _ := http.Post("http://localhost:8080/checkin", "application/json; charset=utf-8", b)
	var body api.CheckinInfo
	json.NewDecoder(res.Body).Decode(&body)
	fmt.Println("test side")
    fmt.Println(body)
    fmt.Println("test side")
}

func main() {

    getRestaurantInfoTest()

}

