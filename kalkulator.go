package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//deklarasi variable
	var numberin1 float64
	var numberin2 float64
	var result float64
	var operation string

	reader := bufio.NewReader(os.Stdin)

	//Welcome Text
	fmt.Println("--------------------------------------------")
	fmt.Println("|        Mau Hitung Apa Hari Ini???        |")
	fmt.Println("********************************************")
	fmt.Println("|--------Masukin BILANGAN Aja Yaaa!--------|")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++")

	//Input Operasi
	fmt.Println("|  Mau Operasi Apa? ( *, +, /, - ) ->  ")
	fmt.Scanf("%s", &operation)
	fmt.Println("--------------------------------------------")

	reader.ReadString('\n')

	//Input Bilangan 1
	fmt.Println("|  Masukkan BILANGAN ->  ")
	if _, err := fmt.Scanf("%f", &numberin1); err != nil {

		fmt.Println("--------------------------------------------")
		fmt.Println("|   Kalkulator Tidak Mengerti Input User   |")
		fmt.Println("--------------------------------------------")

		return
	}
	fmt.Println("--------------------------------------------")

	reader.ReadString('\n')

	//Input Bilangan 2
	fmt.Println("|  Masukkan BILANGAN ->  ")
	if _, err := fmt.Scanf("%f", &numberin2); err != nil {

		fmt.Println("--------------------------------------------")
		fmt.Println("|   Kalkulator Tidak Mengerti Input User   |")
		fmt.Println("--------------------------------------------")
		return
	}
	fmt.Println("--------------------------------------------")

	//Logic Operasi
	if operation == "*" {
		result = numberin1 * numberin2
		fmt.Println("|  Hasilnya Adalah ->  ", result)

	} else if operation == "+" {
		result = numberin1 + numberin2
		fmt.Println("|  Hasilnya Adalah ->  ", result)

	} else if operation == "-" {
		result = numberin1 - numberin2
		fmt.Println("|  Hasilnya Adalah ->  ", result)

	} else if operation == "/" {
		if numberin1 == 0 {
			fmt.Println("--------------------------------------------")
			fmt.Println("|   Kalkulator Tidak Mengerti Input User   |")
			fmt.Println("--------------------------------------------")
		}
		result = numberin1 / numberin2
		fmt.Println("|  Hasilnya Adalah ->  ", result)

	} else {
		fmt.Println("--------------------------------------------")
		fmt.Println("|            Input Operasi Salah           |")
		fmt.Println("--------------------------------------------")
		return
	}
	fmt.Println("--------------------------------------------")
}

// fmt.Println("Masukkan Bilangan ->")
// if _, err := fmt.Scanf("%f", &numberin2); err != nil {

// 	fmt.Println("Masukkan HANYA Bilangan")

// 	return

// }

// fmt.Println("INPUT SALAH!, Tolong Masukkan Bilangan!")

// fmt.Printf("Masukkan Bilangan ->")

// if _, err := fmt.Scanf("%f", &numberin2); err != nil {
// 	fmt.Println("INPUT SALAH!, Tolong Masukkan Bilangan!")
// 	return
// }

// fmt.Printf("Masukkan Operasi! (+, /, *, -) ->")
// fmt.Scanf("%s", &operation)

// if operation == "+" {
// 	fmt.Printf("Hasil --> %f\n", numberin1+numberin2)
// }

// bikin integrer if terus var , if ini data taype not integrer = false
