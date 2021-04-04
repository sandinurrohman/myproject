package main

import (
	"fmt"
)

func main() {

	//INPUT
	var (
		j1, j2, j3, j4, j5 int
	)

	fmt.Println("(MASUKKAN JUMLAH BELANJAAN ANDA)")
	fmt.Print("Qty Sabun Cuci Piring : ")
	fmt.Scanln(&j1)
	fmt.Print("Qty Sabun Mandi : ")
	fmt.Scanln(&j2)
	fmt.Print("Qty Jumlah Shampoo : ")
	fmt.Scanln(&j3)
	fmt.Print("Qty Jumlah Pasta Gigi : ")
	fmt.Scanln(&j4)
	fmt.Print("Qty Jumlah Tissue : ")
	fmt.Scanln(&j5)
	fmt.Println("====================")

	//STRUCT
	type hargabarang struct {
		sabun1, sabun2, shampoo, pastagigi, tissue int
	}
	var h1 = hargabarang{}
	h1.sabun1 = 50000
	h1.sabun2 = 25000
	h1.shampoo = 35000
	h1.pastagigi = 10000
	h1.tissue = 7500

	type totalbarang struct {
		total1, total2, total3, total4, total5 int
	}
	var t1 = totalbarang{}
	t1.total1 = j1
	t1.total2 = j2
	t1.total3 = j3
	t1.total4 = j4
	t1.total5 = j5

	//RUMUS
	var (
		hargafix1, hargafix2, hargafix3, hargafix4, hargafix5, totalfix int
	)

	hargafix1 = h1.sabun1 * j1
	hargafix2 = h1.sabun2 * j2
	hargafix3 = h1.shampoo * j3
	hargafix4 = h1.pastagigi * j4
	hargafix5 = h1.tissue * j5
	totalfix = hargafix1 + hargafix2 + hargafix3 + hargafix4 + hargafix5

	//OUTPUT
	fmt.Println("(LIST BELANJA ANDA)")
	fmt.Println("Sabun Cuci Piring : Rp.", h1.sabun1, "X", t1.total1, "=", hargafix1)
	fmt.Println("Sabun Mandi : Rp.", h1.sabun2, "X", t1.total2, "=", hargafix2)
	fmt.Println("Shampoo : Rp.", h1.shampoo, "X", t1.total3, "=", hargafix3)
	fmt.Println("Pasta Gigi : Rp.", h1.pastagigi, "X", t1.total4, "=", hargafix4)
	fmt.Println("Tissue : Rp.", h1.tissue, "X", t1.total5, "=", hargafix5)
	fmt.Println("====================")
	fmt.Println("TOTAL YANG HARUS ANDA BAYARKAN : Rp.", totalfix)
}
