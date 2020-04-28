package model

type User struct {
	Id int 			`json:"Id"`
	UserName string `json:"UserName"`
   	Birthday string `json:"Birthday"`
	Sex  bool		`json:"Sex"`
}