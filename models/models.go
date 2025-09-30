package models
import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
// go doesn't give json in default that why using json
//  bson used to talk to mongo db
// json for talking over http apis
type User struct{
	// using mongodb so using Primitive.objectID and bson ..... 
	// if sql ..remove bson and also replace primitive.... to uuid.UUID or uint `json:"id" gorm:"primaryKey;autoIncrement"`
    //  or for sql use gorm...
	// ----------------eg---------------
	// ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    // FirstName *string   `json:"first_name" gorm:"column:first_name"`
    // LastName  *string   `json:"last_name" gorm:"column:last_name"`
    // Email     *string   `json:"email" gorm:"unique;not null"`
    // Password  *string   `json:"password"`
    // CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    // UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
   ID   primitive.ObjectID 	`json:"_id" bson:"-_id"`
   First_Name *string		`json:"first_name"`
   Last_Name *string		`json:"last_name"`
   Password *string	 		`json:"password"`
   Email *string 			`json:"email"`
   Phone *string            `json:"phone"`
   Token *string			`json:"token"`
   Refresh_Token *string    `json:"refresh_token"`
   Created_At  time.Time    `json:"created_at"`
   Updated_At time.Time     `json:"updated_at"`
   User_ID string           `json:"user_id"`
//     using bson here as refring to other structs
   UserCart []ProductUser   `json:"user_cart" bson:"user_cart"`
   Address_Details []Address`json:"adress_details" bson:"adress_details"`
   Order_Details []Order    `json:"order_details" bson:"order_details"`
}
type Product struct{
  Product_ID primitive.ObjectID  `bson:"_id"`
  Product_Name string            `json:"product_name"`
  Price *uint64					 `json:"price"`
  Rating *uint					 `json:"rating"`
  Image  *string 				 `json:"image"`

}
type ProductUser struct{
    Product_ID primitive.ObjectID	`bson:"_id"`
    Product_Name *string			`json:"product_name" bson:"product_name"`
	Price *uint64                   `json:"price" bson:"price"`
	Rating *uint					`json:"rating" bson:"rating"`
	Image *string					`json:"image" bson:"image"`
}
type Address struct{
     Address_ID primitive.ObjectID  `bson:"_id"`
	 House *string					`json:"house_name" bson:"house_name"`
	 Street *string					`json:"street_name" bson:"house_name"`
	 City *string					`json:"city_name" bson:"city_name"`
	 Pincode *string				`json:"pincode" bson:"pincode"`
}
type Order struct{
	Order_ID primitive.ObjectID
	Order_Cart []ProductUser
	Order_At  time.Time
	Price     *uint64
	Discount  *int
	Payment_Method Payment
//  here havent created slice of payment mtd . as we are not givving multiple payment mtd
//  the use of * is .. as value can be nil at initally and further can set the value
}
type Payment struct{
	Digital bool
	COD bool

}
