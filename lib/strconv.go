package lib

import (
	"fmt"
	"strconv"
)

func KonversiWithStrconv() {
	// konversi menggunakan strconv
	str := "1234"

	// function untuk mengubah tipe data dari string ke int
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v %T\n", num, num)

	// function untuk mengubah tipe data dari int ke string
	num = 1234
	str = strconv.Itoa(num)
	fmt.Printf("%v %T\n", str, str)

	// function untuk konversi string ke int basis tertentu
	{
		str := "ff"
		num, err := strconv.ParseInt(str, 16, 32)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%v %T\n", num, num)
		}

		str = "0b1010"
		num, err = strconv.ParseInt(str, 0, 8)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%v %T\n", num, num)
		}
	}

	// function untuk konversi string ke numeric desimal dengan lebar data bisa di tentukan
	{
		str := "24.12"
		num, err := strconv.ParseFloat(str, 32)

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%f %T\n", num, num)

		check := float32(num)
		fmt.Printf("%f %T\n", check, check)
	}

	// function untuk konversi string ke bool
	str = "F"
	if boolean, err := strconv.ParseBool(str); err == nil {
		fmt.Printf("%T %v\n", boolean, boolean)
	} else {
		fmt.Println(err.Error())
	}
}
