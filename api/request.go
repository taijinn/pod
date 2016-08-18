package api

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type RestaurantInfo struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string `bson:"name"`
	Phone     string `bson:"phone"`
}

type Dish struct {
	Name string `bson:"name"`
	Pic string `bson:"pic"`
	Description string `bson:"description"`
	Availability bool `bson:"availability"`
}
type Menu struct {
	RestaurantId bson.ObjectId `bson:"restaurantId"`
	Breakfirst []Dish `bson:"breakfirst"`
	Lunch []Dish `bson:"lunch"`
	Dinner []Dish `bson:"dinner"`
	AllDay []Dish `bson:"allDay"`
}
type CheckinInfo struct {
	ID        bson.ObjectId `bson:"_id"`
	UserId string `bson:"userId"`
	RestaurantId bson.ObjectId `bson:"restaurantId"`
	Time time.Time `bson:"time"`
	TableNum int `bson:"tableNum"`
}
func GetRestaurantInfo(phone string, userId string, session *mgo.Session) *RestaurantInfo {
	// TODO for security
	c := session.DB("pod").C("restaurantInfo")
	result := RestaurantInfo{}
	err := c.Find(bson.M{"phone": phone}).One(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func Checkin(restaurantId bson.ObjectId, userId string, time time.Time, tableNum int, session *mgo.Session) *CheckinInfo {
	//TODO : for exception
	c := session.DB("pod").C("checkin")
	_id := bson.NewObjectId()
	err := c.Insert(&CheckinInfo{ID: _id, UserId: userId, RestaurantId: restaurantId, Time: time, TableNum: tableNum})
	if err != nil {
		panic(err)
	}
	result := CheckinInfo{}
	err = c.Find(bson.M{"_id":_id}).One(&result)
	if err != nil {
		panic(err)
	}
	return &result
}
/*

func GetMenu(tradeId *bson.ObjectId, userId string, restaurantInfo *RestaurantInfo) *Menu {
	c := r.ResSession.DB("pod").C("menu")
	
	index := mgo.Index{
		Key: bson.ObjectId {"restaurantId"}
		Unique: true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	
	result := Menu{}
	err := c.Find(bson.M{"restaurantId": restaurantInfo.ID}).One(&result)
	if err != nil {
		panic(err)
	}
	return &result
}
/*
func (r *Request) makeOrder(mealId string, orderid []int, order_status []bool){

}

func (r *Request) makeCheck(mealId string, userId string) bool {

}

func (r *Request) recommendMenu(mealId string, userId string, intpu []int, text string){

}

func (r *Request) checkout(restaurantId string, userId string){

}
*/