package main

import (
	"errors"
	"fmt"
)

func main()  {
	// error => tipe data bahasa go 
	// me-return nilai error di posisi terakhir dari nilai2 return sebuah function
	// zero value dr tipe data error yaitu nil

	// var number int
	// var err error

	// number, err = strconv.Atoi("123")

	// if err == nil {
	// 	fmt.Println(number) // 123
	// }else {
	// 	fmt.Println(err.Error()) // execute
	// }

	// // error custom 
	// var password string

	// fmt.Scanln(&password) // menangkap input

	// if valid, err := validPassword(password); err != nil {
	// 	fmt.Println(err.Error())
	// }else {
	// 	fmt.Println(valid)
	// }

	// panic => stack trace error menghentikan flow goroutine 
	// if valid, err := validPassword(password); err != nil {
	// 	panic(err.Error())
	// }else {
	// 	fmt.Println(valid)
	// }
	// recover => menangkap error yang terjadi pada goroutine dengan catatan memanggil defer

	defer catchErr()
	var password string
	fmt.Scanln(&password) 
	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	}else {
		fmt.Println(valid)
	}
}

// recover
func catchErr()  {
	if r := recover(); r != nil {
		fmt.Println("Error ", r)
	}else {
		fmt.Println("Application running")
	}
}

// custom error
func validPassword(password string) (string, error)  {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "valid password", nil
}