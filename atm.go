package main

import (
	"fmt"
)

func main() {
	fmt.Println("SELAMAT DATANG DI BANK VOKASI")

	// Variabel
	var username string
	var password string

	// Masukkan input
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	// Operasi login
	if username == "Rifais" && password == "2406359140" {
		fmt.Println("Login Berhasil")
		MenuATM()
	} else {
		fmt.Println("Login Gagal")
	}
}

// Fungsi untuk menampilkan menu ATM
func MenuATM() {
	// Variabel ATM
	var pilihAngka int
	balance := 3500000.0
	transactionHistory := []string{}
	//Data akun
	username := "Rifais"
	tempat_lahir := "Jakarta"
	tanggal_lahir := "20-06-2006"
	status := "Belum Kawin"
	alamat := "Jl. Sungai Bambu no.2"
	pekerjaan := "Mahasiswa"
	tempat_bekerja := "Universitas Indonesia"

	for {
		fmt.Println("\n--- MENU ATM ---")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Tambah Saldo")
		fmt.Println("3. Tarik Saldo")
		fmt.Println("4. Lihat Akun")
		fmt.Println("5. Riwayat Transaksi")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih opsi (1-6): ")
		fmt.Scan(&pilihAngka)

		switch pilihAngka {
		case 1:
			fmt.Printf("Saldo Anda: %.2f\n", balance)
		case 2:
			var tambah_saldo float64
			fmt.Print("Masukkan jumlah Transaksi: ")
			fmt.Scan(&tambah_saldo)
			balance += tambah_saldo
			transactionHistory = append(transactionHistory, fmt.Sprintf("Tambah Saldo: %.2f", tambah_saldo))
			fmt.Printf("Jumlah yang disetor: %.2f\n", tambah_saldo)
			fmt.Printf("Saldo baru Anda: %.2f\n", balance)
		case 3:
			var tarik_saldo float64
			fmt.Print("Masukkan jumlah tarik tunai: ")
			fmt.Scan(&tarik_saldo)
			if tarik_saldo > balance {
				fmt.Println("Saldo tidak cukup.")
			} else {
				balance -= tarik_saldo
				transactionHistory = append(transactionHistory, fmt.Sprintf("Tarik Saldo: %.2f", tarik_saldo))
				fmt.Printf("Jumlah yang ditarik: %.2f\n", tarik_saldo)
				fmt.Printf("Saldo baru Anda: %.2f\n", balance)
			}
		case 4:
			fmt.Println("\n-----INFORMASI AKUN ANDA-----")
			fmt.Printf("username: %s\n", username)
			fmt.Printf("tempat lahir: %s\n", tempat_lahir)
			fmt.Printf("tanggal lahir: %s\n", tanggal_lahir)
			fmt.Printf("status: %s\n", status)
			fmt.Printf("alamat: %s\n", alamat)
			fmt.Printf("pekerjaan: %s\n", pekerjaan)
			fmt.Printf("tempat bekerja: %s\n", tempat_bekerja)

		case 5:
			fmt.Println("\n--- RIWAYAT TRANSAKSI ---")
			if len(transactionHistory) == 0 {
				fmt.Println("Tidak ada riwayat transaksi.")
			} else {
				for _, transaction := range transactionHistory {
					fmt.Println(transaction)
				}
			}
		case 6:
			fmt.Println("Terima kasih! Sampai jumpa.")
			return
		default:
			fmt.Println("Opsi tidak valid, silakan coba lagi.")
		}
	}
}
