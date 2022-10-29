package models

/*
	Definiendo nuestro propios modelos de datos, con estructuras definimos lo que esperamos.
	Con json:"" estamos indicando un alias para los json del navegador
	y con bson:"" Es el alias de la DB.
*/

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type User struct {
	Name    string  `json:"name" bson:"user_name"`
	Age     int     `json:"age" bson:"user_age"`
	Address Address `json:"address" bson:"user_addres"`
}
