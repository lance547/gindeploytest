package dao

import (
	"errors"
	"os"
	"strings"
)

//Selectusername函数用于查找是否有同名函数
//Testpassword函数用于检验密码是否正确

var UUU = "./dao/Username.txt"
var PPP = "./dao/Password.txt"

// test the username and password whether include the blank
func Testblank(inf string) error {
	for i := 0; i < len(inf); i++ {
		if inf[i:i+1] == " " {
			return errors.New("Incorrct text format")
		}
	}
	return nil
}

// add information into the datebase:
func Addinformation(inf string, datebase string) error {
	inf += " " //use the blank space to slice the information string
	file, err := os.OpenFile(datebase, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	byteinf := []byte(inf)
	_, err = file.Write(byteinf)
	if err != nil {
		return err
	}
	return nil
}

func Selectusername(username string, datebase string) (bool, error, int) {
	byteinf := Readdatebase(UUU)
	strinf := string(byteinf)
	sliceinf := strings.Split(strinf, " ")
	for i, v := range sliceinf {
		if v == username {
			return true, nil, i
		}
	}
	return false, nil, -1
}
func Testpassword(order int, datebase string, insertp string) (bool, error) {
	byteinf := Readdatebase(PPP)
	if byteinf == nil {
		return false, errors.New("read datebase failed")
	}
	strinf := string(byteinf)
	sliceinf := strings.Split(strinf, " ")
	if insertp == sliceinf[order] {
		return true, nil
	}
	return false, nil
}
func Readdatebase(datebase string) []byte {
	file, err := os.OpenFile(datebase, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil
	}
	defer file.Close()
	byteinf := make([]byte, 100)
	_, err = file.Read(byteinf)
	if err != nil {
		return nil
	}
	return byteinf
}
