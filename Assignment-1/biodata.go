package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	name string
	address string
	job string
	reasonChooseGolang string
}

func main()  {
	var biodata, error = getBiodata(data())

	if error != nil {
		fmt.Println(error.Error())
	}
	
	fmt.Printf("%#v \n",biodata)
}

func data() []Biodata {
	var biodata = []Biodata {
		{name: "Dicky AL Fattah", address: "Lebak Para", job: "Mahasiswa", reasonChooseGolang: "trying new things"}, 
		{name: "Adi Ahmadi", address: "Jendral Sudirman, Jakarta Selatan", job: "Mahasiswa", reasonChooseGolang: "backend developer"},
		{name: "Ayat Maulana", address: "Ciracas, Jakarta Timur", job: "Karyawan", reasonChooseGolang: "switch carrier to developer backend"},
		{name: "Edwin", address: "Cipayung, Jakarta Timur", job: "Karyawan", reasonChooseGolang: "expert developer"},
		{name: "Eko", address: "Cibubur, Jakarta Timur", job: "Mahasiswa", reasonChooseGolang: "learn new something programming"},
	}
	return biodata
}

func getBiodata(biodata []Biodata) (Biodata, error) {
	var getArgument string = os.Args[1]
	var index int

	if argumentInt, err := strconv.Atoi(getArgument); err == nil {
		index = argumentInt
	}

	return biodata[index], nil
}