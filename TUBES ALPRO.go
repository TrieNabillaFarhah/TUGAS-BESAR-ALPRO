package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Calon struct {
	Nama  string
	Partai string
	Suara int
}

type Pemilih struct {
	Nama         string
	NIK          string
	SudahMemilih bool
}

type Pemilu struct {
	Calon      []Calon
	Pemilih    []Pemilih
	BatasWaktu time.Time
	Threshold  int
}

func (p *Pemilu) TambahCalon(nama, partai string) {
	p.Calon = append(p.Calon, Calon{Nama: nama, Partai: partai, Suara: 0})
	fmt.Println("Calon berhasil ditambahkan.")
}

func (p *Pemilu) EditCalon(nomorUrut int, nama, partai string) {
	index := nomorUrut - 1 // Mengubah nomor urut menjadi indeks
	if index >= 0 && index < len(p.Calon) {
		p.Calon[index].Nama = nama
		p.Calon[index].Partai = partai
		fmt.Println("Calon berhasil diubah.")
	} else {
		fmt.Println("Nomor urut calon tidak valid.")
	}
}

func (p *Pemilu) HapusCalon(nomorUrut int) {
	index := nomorUrut - 1 // Mengubah nomor urut menjadi indeks
	if index >= 0 && index < len(p.Calon) {
		p.Calon = append(p.Calon[:index], p.Calon[index+1:]...)
		fmt.Println("Calon berhasil dihapus.")
	} else {
		fmt.Println("Nomor urut calon tidak valid.")
	}
}

func (p *Pemilu) TambahPemilih(nama, nik string) {
	p.Pemilih = append(p.Pemilih, Pemilih{Nama: nama, NIK: nik, SudahMemilih: false})
	fmt.Println("Pemilih berhasil ditambahkan.")
}

func (p *Pemilu) Pilih(nik string, nomorUrut int) {
	if time.Now().After(p.BatasWaktu) {
		fmt.Println("Waktu pemilihan telah habis.")
		return
	}
	index := nomorUrut - 1 // Mengubah nomor urut menjadi indeks
	for i, pemilih := range p.Pemilih {
		if pemilih.NIK == nik {
			if index >= 0 && index < len(p.Calon) {
				if !p.Pemilih[i].SudahMemilih {
					p.Calon[index].Suara++
					p.Pemilih[i].SudahMemilih = true
					fmt.Println("Pemilih berhasil memberikan suara.")
					return
				} else {
					fmt.Println("Pemilih sudah memberikan suara sebelumnya.")
					return
				}
			} else {
				fmt.Println("Nomor urut calon tidak valid.")
				return
			}
		}
	}
	fmt.Println("NIK tidak ditemukan atau pemilih tidak terdaftar.")
}

func (p *Pemilu) TampilkanHasil() {
	fmt.Println("Hasil Pemilu:")
	fmt.Printf("%-5s %-20s %-20s %-10s %-10s\n", "No", "Nama", "Partai", "Suara", "Status")
	fmt.Println(strings.Repeat("-", 70))
	for i, calon := range p.Calon {
		status := "Tidak Lolos"
		if calon.Suara >= p.Threshold {
			status = "Lolos"
		}
		fmt.Printf("%-5d %-20s %-20s %-10d %-10s\n", i+1, calon.Nama, calon.Partai, calon.Suara, status)
	}
}

func (p *Pemilu) CariCalon(keyword string) {
	fmt.Println("Hasil Pencarian Calon:")
	for _, calon := range p.Calon {
		if strings.Contains(strings.ToLower(calon.Nama), strings.ToLower(keyword)) || strings.Contains(strings.ToLower(calon.Partai), strings.ToLower(keyword)) {
			fmt.Printf("Nama: %s, Partai: %s, Suara: %d\n", calon.Nama, calon.Partai, calon.Suara)
		}
	}
}

func (p *Pemilu) TampilkanCalon() {
	fmt.Println("Daftar Calon:")
	for i, calon := range p.Calon {
		fmt.Printf("Nomor Urut: %d, Nama: %s, Partai: %s\n", i+1, calon.Nama, calon.Partai) // Menampilkan nomor urut
	}
}

func main() {
	// Inisialisasi Pemilu
	pemilu := Pemilu{
		BatasWaktu: time.Now().Add(24 * time.Hour), // Durasi waktu pemilihan adalah 24 jam dari sekarang
		Threshold:  10,                            // Ambang batas suara
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n-- Menu --")
		fmt.Println("1. Login sebagai Petugas KPU")
		fmt.Println("2. Login sebagai Pemilih")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			// Login Petugas KPU
			fmt.Print("Masukkan username petugas: ")
			scanner.Scan()
			username := scanner.Text()
			fmt.Print("Masukkan password petugas: ")
			scanner.Scan()
			password := scanner.Text()

			// Cek username dan password (simple check)
			if username == "admin" && password == "admin123" {
				fmt.Println("Login sebagai Petugas KPU berhasil.")
				for {
					// Menu untuk petugas KPU
					fmt.Println("\n-- Menu Petugas KPU --")
					fmt.Println("1. Tambah Calon")
					fmt.Println("2. Edit Calon")
					fmt.Println("3. Hapus Calon")
					fmt.Println("4. Tambah Pemilih")
					fmt.Println("5. Tampilkan Hasil")
					fmt.Println("6. Cari Calon")
					fmt.Println("7. Tampilkan Calon dengan Nomor Urut")
					fmt.Println("8. Logout")
					fmt.Print("Pilih menu: ")

					scanner.Scan()
					input := scanner.Text()
					switch input {
					case "1":
						fmt.Print("Masukkan nama calon: ")
						scanner.Scan()
						nama := scanner.Text()
						fmt.Print("Masukkan partai calon: ")
						scanner.Scan()
						partai := scanner.Text()
						pemilu.TambahCalon(nama, partai)
					case "2":
						fmt.Print("Masukkan nomor urut calon yang akan diedit: ")
						scanner.Scan()
						var nomorUrut int
						fmt.Sscan(scanner.Text(), &nomorUrut)
						fmt.Print("Masukkan nama baru: ")
						scanner.Scan()
						nama := scanner.Text()
						fmt.Print("Masukkan partai baru: ")
						scanner.Scan()
						partai := scanner.Text()
						pemilu.EditCalon(nomorUrut, nama, partai)
					case "3":
						fmt.Print("Masukkan nomor urut calon yang akan dihapus: ")
						scanner.Scan()
						var nomorUrut int
						fmt.Sscan(scanner.Text(), &nomorUrut)
						pemilu.HapusCalon(nomorUrut)
					case "4":
						fmt.Print("Masukkan nama pemilih: ")
						scanner.Scan()
						nama := scanner.Text()
						fmt.Print("Masukkan NIK pemilih: ")
						scanner.Scan()
						nik := scanner.Text()
						pemilu.TambahPemilih(nama, nik)
					case "5":
						pemilu.TampilkanHasil()
					case "6":
						fmt.Print("Masukkan kata kunci pencarian: ")
						scanner.Scan()
						keyword := scanner.Text()
						pemilu.CariCalon(keyword)
					case "7":
						pemilu.TampilkanCalon()
					case "8":
						fmt.Println("Logout berhasil. Kembali ke menu utama.")
						break
					default:
						fmt.Println("Pilihan tidak valid.")
					}
					if input == "8" {
						break
					}
				}
			} else {
				fmt.Println("Login gagal. Username atau password salah.")
			}

		case "2":
			// Login Pemilih
			fmt.Print("Masukkan NIK pemilih: ")
			scanner.Scan()
			nik := scanner.Text()

			// Cek apakah pemilih terdaftar
			var pemilihExist bool
			for _, pemilih := range pemilu.Pemilih {
				if pemilih.NIK == nik {
					pemilihExist = true
					break
				}
			}
			if pemilihExist {
				fmt.Println("Login sebagai Pemilih berhasil.")
				for {
					// Menu untuk pemilih
					fmt.Println("\n-- Menu Pemilih --")
					fmt.Println("1. Pilih Calon")
					fmt.Println("2. Tampilkan Calon")
					fmt.Println("3. Logout")
					fmt.Print("Pilih menu: ")

					scanner.Scan()
					input := scanner.Text()
					switch input {
					case "1":
						fmt.Print("Masukkan nomor urut calon: ")
						scanner.Scan()
						var nomorUrut int
						fmt.Sscan(scanner.Text(), &nomorUrut)
						pemilu.Pilih(nik, nomorUrut)
					case "2":
						pemilu.TampilkanCalon()
					case "3":
						fmt.Println("Logout berhasil. Kembali ke menu utama.")
						break
					default:
						fmt.Println("Pilihan tidak valid.")
					}
					if input == "3" {
						break
					}
				}
			} else {
				fmt.Println("NIK tidak ditemukan atau pemilih tidak terdaftar.")
			}

		case "9":
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
