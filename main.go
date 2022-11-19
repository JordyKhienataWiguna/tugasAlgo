package main

import (
	"fmt"
)

type cust struct {
	name, address      string
	kartuKredit, saldo float64
}

func main() {
	// var customer, address,
	var customer = cust{}
	var input string

	fmt.Print("=========== Selamat Datang di Toko Honda Pondok Indah ===========\n")
	fmt.Printf("Masukkan Nama Anda: ")
	fmt.Scan(&customer.name)
	// fmt.Scan(&customer)
	fmt.Printf("Silahkan Masukan Alamat Anda : ")
	fmt.Scan(&customer.address)
	// fmt.Scan(&address)
	fmt.Printf("Silakan Masukkan Nomor Kartu Kredit Anda : ")
	fmt.Scan(&customer.kartuKredit)

	fmt.Printf("Silakan Masukkan Saldo Kartu Kredit Anda : ")
	fmt.Scan(&customer.saldo)

	fmt.Printf("Apa Yang Sedang anda cari ?\n1.Motor\n2.Sparepart\n3.Accessories\n")
	fmt.Printf("Masukkan Pilihan Anda : ")
	fmt.Scan(&input)

	fmt.Print("Struk Pembayaran\n")
	// fmt.Println("Order dengan atas nama", *&customer)
	// fmt.Println("Pesanan akan di kirim ke ", *&address)
	fmt.Println("Order dengan atas nama", customer.name)
	fmt.Println("Pesanan akan di kirim ke ", customer.address)
	fmt.Println("Pembayaran akan ditarik dari Kartu Kredit Nomor ", customer.kartuKredit)
	fmt.Println("Saldo Anda tersisa ", customer.saldo)

}
