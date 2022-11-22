/*
Kelompok 5 :
Devi Mikhael Empi
Jordy Khienata Wiguna
Kayla Berliana Aurellya
*/
package main

import (
	"fmt"
)

type cust struct {
	name, address, kartuKredit string

	// saldo float64
}

func main() {
	// var customer, address,
	var customer = cust{}
	var input, pilih, motor, Sparepart, Accessories, motorBebek, motorMatic, motorSport string
	var hargaMotor, hargaSparepart, hargaAccessories int

	fmt.Print("=========== Selamat Datang di Toko Honda Pondok Indah ===========\n\n")
	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scan(&customer.name)
	// fmt.Scan(&customer)
	fmt.Print("Silahkan Masukan Alamat Anda : ")
	fmt.Scan(&customer.address)
	// fmt.Scan(&address)
	fmt.Print("Silakan Masukkan Nomor Kartu Kredit Anda : ")
	fmt.Scan(&customer.kartuKredit)

	// fmt.Print("Silakan Masukkan Saldo Kartu Kredit Anda : ")
	// fmt.Scan(&customer.saldo)
menu:
	fmt.Print("\nApa Yang Sedang anda cari ?\n1.Motor\n2.Sparepart\n3.Accessories\n")
	fmt.Print("Masukkan Pilihan Anda : ")
	fmt.Scan(&input)

	switch input {
	case "1":
		fmt.Print("1.Motor Bebek.\n2.Motor Matic.\n3.Motor Sport.\n")
		fmt.Print("Silakan pilih motor anda : ")
		fmt.Scan(&motor)
		if motor == "1" {
			fmt.Print("1.Honda Supra X 125.\n2.Honda Revo.\n3.Honda GTR.\n")
			fmt.Print("Motor apa yang anda pilih ? ")
			fmt.Scan(&motorBebek)
			if motorBebek == "1" {
				hargaMotor = 19000000
				fmt.Print("Anda membeli Honda Supra X 125 dengan hargaMotor Rp ", hargaMotor, " Rupiah")
			} else if motorBebek == "2" {
				hargaMotor = 16000000
				fmt.Print("Anda membeli Honda Revo ", hargaMotor, " Rupiah ")
			} else if motorBebek == "3" {
				hargaMotor = 25000000
				fmt.Print("Anda membeli Honda GTR ", hargaMotor, " Rupiah")
			}

		} else if motor == "2" {
			fmt.Print("1.Honda Vario 160.\n2.Honda Genio.\n3.Honda BeAT.\n")
			fmt.Print("Motor apa yang anda pilih ? ")
			fmt.Scan(&motorMatic)
			if motorMatic == "1" {
				hargaMotor = 27000000
				fmt.Print("Anda membeli Honda Vario 160 ", hargaMotor, " Rupiah")
			} else if motorMatic == "2" {
				hargaMotor = 18880000
				fmt.Print("Anda membeli Honda Genio ", hargaMotor, " Rupiah")
			} else if motorMatic == "3" {
				hargaMotor = 17720000
				fmt.Print("Anda membeli Honda BeAT ", hargaMotor, " Rupiah")
			}
		} else if motor == "3" {
			fmt.Print("1.Honda Sonic 150R.\n2.Honda CB150R Streetfire.\n3.Honda CBR 150R.\n")
			fmt.Print("Motor apa yang anda pilih ? ")
			fmt.Scan(&motorSport)
			if motorSport == "1" {
				hargaMotor = 24132000
				fmt.Print("Anda membeli Honda Sonic 150R ", hargaMotor, " Rupiah")
			} else if motorSport == "2" {
				hargaMotor = 30111000
				fmt.Print("Anda membeli Honda CB150R Streetfire ", hargaMotor, " Rupiah")
			} else if motorSport == "3" {
				hargaMotor = 36741000
				fmt.Print("Anda membeli Honda CBR 150R ", hargaMotor, " Rupiah")
			}
		} else {
			fmt.Println("Anda salah memasukkan inputan!")
		}
		fmt.Print("\nApakah Anda ingin belanja lagi ? (Y/N)")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			goto menu
		} else {
			goto struk
		}

	case "2":
		fmt.Print("1.Aki.\n2.Gear Set.\n3.Shock Breaker.\n4.Piston Kit\n")
		fmt.Print("Silakan pilih motor anda : ")
		fmt.Scan(&Sparepart)
		if Sparepart == "1" {
			hargaSparepart = 335000
			fmt.Print("Anda membeli Aki dengan harga Rp.", hargaSparepart, " rupiah")
		} else if Sparepart == "2" {
			hargaSparepart = 150000
			fmt.Print("Anda membeli gear set dengan harga Rp.", hargaSparepart, " rupiah")
		} else if Sparepart == "3" {
			hargaSparepart = 1200000
			fmt.Print("Anda membeli Shock Breaker dengan harga Rp.", hargaSparepart, " rupiah")
		} else if Sparepart == "4" {
			hargaSparepart = 315000
			fmt.Print("Anda membeli Piston Kit dengan harga Rp.", hargaSparepart, " rupiah")
		} else {
			fmt.Print("Anda salah memasukkan inputan!")
		}
		fmt.Print("\nTotal pembayaran : Rp.", hargaSparepart, " Rupiah")
		fmt.Print("\nApakah Anda ingin belanja lagi ? (Y/N)")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			goto menu
		} else {
			goto struk
		}

	case "3":
		fmt.Print("1.Protector.\n2.Visor.\n3.Garnish Headlight.\n4.Smart Key Remote Cover\n")
		fmt.Print("Silakan pilih motor anda : ")
		fmt.Scan(&Accessories)
		if Accessories == "1" {
			hargaAccessories = 71000
			fmt.Print("Anda membeli Protector dengan harga Rp.", hargaAccessories, " rupiah")
		} else if Accessories == "2" {
			hargaAccessories = 51500
			fmt.Print("Anda membeli Visor dengan harga Rp.", hargaAccessories, " rupiah")
		} else if Accessories == "3" {
			hargaAccessories = 133000
			fmt.Print("Anda membeli Garnish Headlight dengan harga Rp.", hargaAccessories, " rupiah")
		} else if Accessories == "4" {
			hargaAccessories = 121000
			fmt.Print("Anda membeli Smart Key Remote Cover dengan harga Rp.", hargaAccessories, " rupiah")
		}
		fmt.Print("\nTotal pembayaran : Rp.", hargaAccessories, " Rupiah")
		fmt.Print("\nApakah Anda ingin belanja lagi ? (Y/N)")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			goto menu
		} else {
			goto struk
		}

	default:
		fmt.Print("Pesanan yang anda pilih tidak ada dalam kategori.")
		// totalBelanja = hargaAccessories + hargaMotor + hargaSparepart
		break
	}
	
struk:
	// Struck Pembayaran
	fmt.Print("\n\nStruk Pembayaran Honda Pondok Indah\n")
	// fmt.Println("Order dengan atas nama", *&customer)
	// fmt.Println("Pesanan akan di kirim ke ", *&address)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("Order dengan atas nama", customer.name)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("Pesanan akan di kirim ke ", customer.address)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("Total Motor: ", hargaMotor)
	fmt.Println("Total Sparepart: ", hargaSparepart)
	fmt.Println("Total Accessories: ", hargaAccessories)
	fmt.Println("Total belanja ", hargaAccessories + hargaMotor + hargaSparepart)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("Pembayaran akan ditarik dari Kartu Kredit Nomor ", customer.kartuKredit)
	fmt.Println("----------------------------------------------------------------------------")
	// fmt.Println("Saldo Anda tersisa ", customer.saldo)
	// fmt.Println("======================================")
	fmt.Print("\nApakah Anda ingin belanja lagi ? (Y/N)")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			goto menu
		} else {
			fmt.Println("Terimakasih Sudah Berbelanja di Toko Honda Pondok Indah")
		}
}
