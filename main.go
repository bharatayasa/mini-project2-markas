package main

import (
	"fmt"
	"os"
)

type Buku struct {
	Kode          string
	Judul         string
	Pengarang     string
	Penerbit      string
	JumlahHalaman int
	TahunTerbit   int
}

var ListBuku []Buku

func TambahBuku() {
	var KodeBuku, JudulBuku, PengarangBuku, PenerbitBuku string
	var JumlahHalamanBuku, TahunTerbitBuku int

	fmt.Print("masukkan kode buku: ")
	_, err := fmt.Scanln(&KodeBuku)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Print("masukkan judul buku: ")
	_, err = fmt.Scanln(&JudulBuku)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Print("masukkan pengarang buku: ")
	_, err = fmt.Scanln(&PengarangBuku)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Print("masukkan penerbit buku: ")
	_, err = fmt.Scanln(&PenerbitBuku)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Print("masukkan jumlah halaman buku: ")
	_, err = fmt.Scanln(&JumlahHalamanBuku)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Print("masukkan tahun terbit buku: ")
	_, err = fmt.Scanln(&TahunTerbitBuku)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	ListBuku = append(ListBuku, Buku{
		Kode:          KodeBuku,
		Judul:         JudulBuku,
		Pengarang:     PengarangBuku,
		Penerbit:      PenerbitBuku,
		JumlahHalaman: JumlahHalamanBuku,
		TahunTerbit:   TahunTerbitBuku,
	})

	fmt.Println("---------------------------------------------")
	fmt.Println("buku berhasil ditanmbahkan")
	fmt.Println("---------------------------------------------")
	fmt.Println("")
}

func TampilBuku() {
	if len(ListBuku) == 0 {
		fmt.Println("belum ada buku yang ditambah")
		fmt.Println("---------------------------------------------")
		return
	}

	fmt.Println("berikut merupakan daftar buku:")
	fmt.Println("------------------------------")
	for _, buku := range ListBuku {
		fmt.Printf("Kode: %s\nJudul: %s\nPengarang: %s\nPenerbit: %s\nJumlah Halaman: %d\nTahun Terbit: %d\n\n",
			buku.Kode,
			buku.Judul,
			buku.Pengarang,
			buku.Penerbit,
			buku.JumlahHalaman,
			buku.TahunTerbit,
		)
	}
}

func HapusBuku() {
	if len(ListBuku) == 0 {
		fmt.Println("Tidak ada buku yang tersedia untuk dihapus.")
		return
	}

	var kodeBuku string
	fmt.Print("Masukkan kode buku yang akan dihapus: ")
	_, err := fmt.Scan(&kodeBuku)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	found := false
	for i, buku := range ListBuku {
		if buku.Kode == kodeBuku {
			ListBuku = append(ListBuku[:i], ListBuku[i+1:]...)
			fmt.Println("Buku dengan kode", kodeBuku, "telah dihapus.")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Buku dengan kode", kodeBuku, "tidak ditemukan.")
	}
}

func EditBuku() {
	if len(ListBuku) == 0 {
		fmt.Println("Tidak ada buku yang tersedia untuk diedit.")
		return
	}

	var kodeBuku string
	fmt.Print("Masukkan kode buku yang akan diedit: ")
	_, err := fmt.Scan(&kodeBuku)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	found := false
	for i, buku := range ListBuku {
		if buku.Kode == kodeBuku {
			fmt.Println("Buku ditemukan, masukkan informasi baru:")

			fmt.Print("masukkan judul buku baru: ")
			_, err := fmt.Scan(&ListBuku[i].Judul)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			fmt.Print("masukkan pengarang baru: ")
			_, err = fmt.Scan(&ListBuku[i].Pengarang)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			fmt.Print("masukkan penerbit baru: ")
			_, err = fmt.Scan(&ListBuku[i].Penerbit)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			fmt.Print("masukkan jumlah halaman baru: ")
			_, err = fmt.Scan(&ListBuku[i].JumlahHalaman)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			fmt.Print("masukkan tahun terbit baru: ")
			_, err = fmt.Scan(&ListBuku[i].TahunTerbit)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			fmt.Println("Buku dengan kode", kodeBuku, "telah diperbarui.")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Buku dengan kode", kodeBuku, "tidak ditemukan.")
	}
}

func main() {
	for {
		var MenuOpsi int

		fmt.Print("---------------------------------------------")
		fmt.Println("\nSistem manajemen buku")
		fmt.Println("---------------------------------------------")
		fmt.Println("Silahkan pilih menu opsi:")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Tampil Buku")
		fmt.Println("3. Hapus Buku")
		fmt.Println("4. Ubah Buku")
		fmt.Println("5. Keluar")
		fmt.Println("---------------------------------------------")

		fmt.Print("Masukkan pilihan: ")
		_, err := fmt.Scanln(&MenuOpsi)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch MenuOpsi {
		case 1:
			TambahBuku()
		case 2:
			TampilBuku()
		case 3:
			HapusBuku()
		case 4:
			EditBuku()
		case 5:
			fmt.Println("program selesai")
			os.Exit(0)
		default:
			fmt.Println("pilihan tidak valid")
		}
		fmt.Println("tes")
	}
}
