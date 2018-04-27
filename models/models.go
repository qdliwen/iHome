package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/* User Table
 * Table name = user
 */
 type User struct {
 	Id				 	int 			`json:"user_id"`
 	Name 			 	string 			`orm:"size(32);unique" json:"name"`
 	Password_hash    	string 			`orm:"size(128)" json:"password"`
 	Mobile 				string			`orm:"size(11);unique" json:"mobile"`
 	Real_Name			string			`orm:"size(32)" json:"real_name"`
 	Id_Card				string			`orm:"size(20)" json:"id_card"`
 	Avatar_url			string			`orm:"size(256)" json:"avatar_url"`
 	Houses				[]*House		`orm:"reverse(many)" json:"houses"`
 	Orders				[]*OrderHouse	`orm:"reverse(many)" json:"orders"`
 }

 /*
  * House Table
  */
type House struct {
	Id              int           `json:"house_id"`
	User            *User         `orm:"rel(fk)" json:"user_id"`
	Area            *Area         `orm:"rel(fk)" json:"area_id"`
	Title           string        `orm:"size(64)" json:"title"`
	Price           int           `orm:"default(0)" json:"price"`
	Address         string        `orm:"size(512)" orm:"default("")" json:"address"`
	Room_count      int           `orm:"default(1)" json:"room_count"`
	Acreage         int           `orm:"default(0)" json:"acreage"`
	Unit            string        `orm:"size(32)" orm:"default("")" json:"unit"`
	Capacity        int           `orm:"default(1)" json:"capacity"`
	Beds            string        `orm:"size(64)" orm:"default("")" json:"beds"`
	Deposit         int           `orm:"default(0)" json:"deposit"`
	Min_days        int           `orm:"default(1)" json:"min_days"`
	Max_days        int           `orm:"default(0)" json:"max_days"`
	Order_count     int           `orm:"default(0)" json:"order_count"`
	Index_image_url string        `orm:"size(256)" orm:"default("")" json:"index_image_url"`
	Facilities      []*Facility   `orm:"reverse(many)" json:"facilities"`
	Images          []*HouseImage `orm:"reverse(many)" json:"img_urls"`
	Orders          []*OrderHouse `orm:"reverse(many)" json:"orders"`
	Ctime           time.Time     `orm:"auto_now_add;type(datetime)" json:"ctime"`
}

/*
 *  table_name = area
 */
type Area struct {
	Id     int      `json:"aid"`
	Name   string   `orm:"size(32)" json:"aname"`
	Houses []*House `orm:"reverse(many)" json:"houses"`
}

/*
 *	table_name = "facility"
 */
type Facility struct {
	Id     int      `json:"fid"`
	Name   string   `orm:"size(32)"`
	Houses []*House `orm:"rel(m2m)"`
}

/*
 *	table_name = "house_image"
 */
type HouseImage struct {
	Id    int    `json:"house_image_id"`
	Url   string `orm:"size(256)" json:"url"`
	House *House `orm:"rel(fk)" json:"house_id"`
}

/*
 *	table_name = order
 */
type OrderHouse struct {
	Id          int       `json:"order_id"`
	User        *User     `orm:"rel(fk)" json:"user_id"`
	House       *House    `orm:"rel(fk)" json:"house_id"`
	Begin_date  time.Time `orm:"type(datetime)"`
	End_date    time.Time `orm:"type(datetime)"`
	Days        int
	House_price int
	Amount      int
	Status      string    `orm:"default(WAIT_ACCEPT)"`
	Comment     string    `orm:"size(512)"`
	Ctime       time.Time `orm:"auto_now_add;type(datetime)" json:"ctime"`
}

const (
	ORDER_STATUS_WAIT_ACCEPT  = "WAIT_ACCEPT"
	ORDER_STATUS_WAIT_PAYMENT = "WAIT_PAYMENT"
	ORDER_STATUS_PAID         = "PAID"
	ORDER_STATUS_WAIT_COMMENT = "COMMENT"
	ORDER_STATUS_COMPLETE     = "COMPLETE"
	ORDER_STATUS_CANCELED     = "CONCELED"
	ORDER_STATUS_REJECTED     = "REJECTED"
)


var HOME_PAGE_MAX_HOUSES int = 5


var HOUSE_LIST_PAGE_CAPACITY int = 2

/*
 *  init function
 */
 func init() {

 	// set default database
 	orm.RegisterDataBase("default", "mysql", "root:211828@tcp(127.0.0.1:3306)/lovehome?charset=utf8", 30)
 	// register model
	orm.RegisterModel(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse))
 	// create table
 	orm.RunSyncdb("default", false, true)
 }