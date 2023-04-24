package main

import (
	"errors"
	"fmt"
)

type User struct{
	ID           int
    Name         string
    DNI          string
    Phone        string
    Address      string
}
var(
	UsearAlreadyExists = errors.New("User already exists")
	InvalidData = errors.New("Invalid Data")
)
var users = []User{}

func VerifyUserData(user *User) (bool,error){
	if user.ID == 0 || user.Name == "" || user.DNI == "" || user.Phone == "" || user.Address == "" {
        return false,InvalidData
    }
    return true,nil
}
func VerifyUserNotRepeat(u *User) []User{
	for _, user := range users {
		if user.DNI== u.DNI {
			panic(UsearAlreadyExists)
		}
	}
	_,err :=VerifyUserData(u)
	if err!=nil {
		panic(err)
	}
	
	users = append(users, *u)
	return users
}
func main()  {
	u := User{ID: 0, Name: "", DNI: "12345678", Phone: "555-5555", Address: "Av. Siempreviva 123"}
	
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Se detectaron varios errores en tiempo de ejecución")
			fmt.Println(r)
        } else {
            fmt.Println("Fin de la ejecución")
        }
    }()
	
	VerifyUserNotRepeat(&u)

}
 