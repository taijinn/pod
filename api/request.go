package api

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
	"golang.org/x/crypto/bcrypt"
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
type UserCredentials struct {
	Username	string  `json:"username"`
	Password	string	`json:"password"`
}

type UserInDatabase struct {
	Username string
	Password []byte
}

func GetEncriptedPass(userName string, session *mgo.Session) (bool, []byte) {
	c := session.DB("pod").C("userPass")
	var f UserInDatabase
	err := c.Find(bson.M{"username": userName}).One(&f)
	if err != nil {
		if err == mgo.ErrNotFound {
			return false, nil
		} else {
			log.Println(err)
		}
	} 
	return true, f.Password
}

func GetRestaurantInfo(phone string, userId string, session *mgo.Session) *RestaurantInfo {
	// TODO for security
	c := session.DB("pod").C("restaurantInfo")
	result := RestaurantInfo{}
	err := c.Find(bson.M{"phone": phone}).One(&result)
	if err != nil {
		log.Fatal(err)
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
func clear(b []byte) {
    for i := 0; i < len(b); i++ {
        b[i] = 0;
    }
}

func Crypt(password []byte) ([]byte, error) {
    defer clear(password)
    return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}
func GetSigninAPI(user *UserCredentials, session *mgo.Session) string {
	c := session.DB("pod").C("userPass")
	var f UserInDatabase
	err := c.Find(bson.M{"username": user.Username}).One(&f)
	if err != nil { 
		if err == mgo.ErrNotFound { // if no user registered
			//insert or update toBeActive collection
			c1 := session.DB("pod").C("toBeActive")
			err1 := c1.Find(bson.M{"username": user.Username}).One(&f)
			if err1 == nil {
				//update toBeActive
				key := bson.M{"username": user.Username}
				cryptedpass, _ := Crypt([]byte(user.Password))
				newpass := bson.M{"$set": bson.M{"password": cryptedpass}}
				err2 := c1.Update(key, newpass)
				if err2 != nil {
					panic(err2)
				}
				//send an email to username
				log.Println("update")
				return "sent"
			}else{
				if err1 == mgo.ErrNotFound {
					//insert into toBeActive collection
					cryptedpass, _ := Crypt([]byte(user.Password))
					err2 := c1.Insert(&UserInDatabase{Username:user.Username, Password:cryptedpass})
					if err2 != nil {
						panic(err2)
					}
					//send an email to username
					log.Println("insert")
					return "sent"
				}else{
					log.Fatal(err1)
				}
			}
		} else {
			log.Fatal(err)
		}
	} else{
		// if user is already in the user name
		//var p1, _ = Crypt([]byte(user.Password))
		//fmt.Println("user")
		//fmt.Println(p1)
		err := bcrypt.CompareHashAndPassword(f.Password, []byte(user.Password))
		if err == nil {//f.Password) {
			return "login"
		}else{
			return "forget"
		}
	}
	return "error"
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