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

	fmt.Println("buku berhasil ditanmbahkan")
	fmt.Println("---------------------------------------------")
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

func UbahBuku() {
	fmt.Println("ubah buku")
}

func main() {
	for {
		var pilihMenu int

		fmt.Println("\nSistem manajemen buku")
		fmt.Println("---------------------------------------------")
		fmt.Println("Silahkan pilih menu:")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Tampil Buku")
		fmt.Println("3. Ubah Buku")
		fmt.Println("4. Keluar")
		fmt.Println("---------------------------------------------")

		fmt.Print("Masukkan pilihan: ")
		_, err := fmt.Scanln(&pilihMenu)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch pilihMenu {
		case 1:
			TambahBuku()
		case 2:
			TampilBuku()
		case 3:
			UbahBuku()
		case 4:
			fmt.Println("program selesai")
			os.Exit(0)
		default:
			fmt.Println("pilihan tidak valid")
		}
	}
}
