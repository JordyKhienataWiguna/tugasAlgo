package main

import (
	"fmt"
)

type customer struct {
	name, address string
}

func main() {
	var customer, address, input string

	fmt.Println("=========== Selamat Datang di Toko Honda Pondok Indah ===========\n")
	// fmt.Printf("Masukkan Nama Anda: ")
	// fmt.Scan(&customer)
	// fmt.Printf("Silahkan Masukan Alamat Anda : ")
	// fmt.Scan(&address)

	fmt.Printf("Apa Yang Sedang anda cari ?\n1.Motor\n2.Sparepart\n3.Accessories\n")
	fmt.Printf("Masukkan Pilihan Anda : ")
	fmt.Scan(&input)

	fmt.Println("Bukti Order\n")
	fmt.Println("Order dengan atas nama", *&customer)
	fmt.Println("Pesanan akan di kirim ke ", *&address)

}
