package main

import "fmt"

type Calon struct {
	Nama   string
	Partai string
}

const maxCalon = 100

var calonList [maxCalon]Calon
var calonCount int

func tambahCalon(nama string, partai string) {
	if calonCount >= maxCalon {
		fmt.Println("Daftar calon sudah penuh.")
		return
	}
	calonList[calonCount] = Calon{Nama: nama, Partai: partai}
	calonCount++
	fmt.Println("Calon berhasil ditambahkan.")
}

func editCalon(index int, nama string, partai string) {
	index--
	if index < 0 || index >= calonCount {
		fmt.Println("Index calon tidak valid.")
		return
	}
	calonList[index].Nama = nama
	calonList[index].Partai = partai
	fmt.Println("Calon berhasil diubah.")
}

func hapusCalon(index int) {
	index--
	if index < 0 || index >= calonCount {
		fmt.Println("Index calon tidak valid.")
		return
	}
	for i := index; i < calonCount-1; i++ {
		calonList[i] = calonList[i+1]
	}
	calonCount--
	fmt.Println("Calon berhasil dihapus.")
}

func tampilkanCalon() {
	if calonCount == 0 {
		fmt.Println("Tidak ada calon yang terdaftar.")
		return
	}
	fmt.Println("\nDaftar Calon:")
	for i := 0; i < calonCount; i++ {
		fmt.Printf("%d. Nama: %s, Partai: %s\n", i+1, calonList[i].Nama, calonList[i].Partai) // Index + 1
	}
}
func main() {
	var pilihan int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Calon")
		fmt.Println("2. Edit Calon")
		fmt.Println("3. Hapus Calon")
		fmt.Println("4. Tampilkan Calon")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var nama, partai string
			fmt.Print("Masukkan nama calon: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan partai calon: ")
			fmt.Scan(&partai)
			tambahCalon(nama, partai)
		case 2:
			var index int
			var nama, partai string
			fmt.Print("Masukkan nomor calon yang ingin diubah: ")
			fmt.Scan(&index)
			fmt.Print("Masukkan nama baru calon: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan partai baru calon: ")
			fmt.Scan(&partai)
			editCalon(index, nama, partai)
		case 3:
			var index int
			fmt.Print("Masukkan nomor calon yang ingin dihapus: ")
			fmt.Scan(&index)
			hapusCalon(index)
		case 4:
			tampilkanCalon()
		case 5:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
