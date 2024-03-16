package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Book struct {
	Kode          string `json:"kode"`
	Judul         string `json:"judul"`
	Pengarang     string `json:"pengarang"`
	Penerbit      string `json:"penerbit"`
	JumlahHalaman int    `json:"jumlah_halaman"`
	TahunTerbit   int    `json:"tahun_terbit"`
}

var booksDirectory = "books"
var pdfDirectory = "pdf"

func inputBook() {
	var newBook Book

	fmt.Println("\nMasukkan detail buku:")
	fmt.Print("Kode: ")
	fmt.Scanln(&newBook.Kode)

	// Check apakah kode buku sudah ada
	if _, err := os.Stat(filepath.Join(booksDirectory, "book-"+newBook.Kode+".json")); !os.IsNotExist(err) {
		fmt.Println("Kode buku sudah digunakan.")
		return
	}

	fmt.Print("Judul: ")
	fmt.Scanln(&newBook.Judul)
	fmt.Print("Pengarang: ")
	fmt.Scanln(&newBook.Pengarang)
	fmt.Print("Penerbit: ")
	fmt.Scanln(&newBook.Penerbit)
	fmt.Print("Jumlah Halaman: ")
	fmt.Scanln(&newBook.JumlahHalaman)
	fmt.Print("Tahun Terbit: ")
	fmt.Scanln(&newBook.TahunTerbit)

	// Simpan ke file JSON
	bookData, err := json.MarshalIndent(newBook, "", "  ")
	if err != nil {
		fmt.Println("Error encoding book data:", err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(booksDirectory, "book-"+newBook.Kode+".json"), bookData, 0644)
	if err != nil {
		fmt.Println("Error saving book data:", err)
		return
	}

	fmt.Println("Buku berhasil disimpan.")
}

func listBooks() {
	files, err := ioutil.ReadDir(booksDirectory)
	if err != nil {
		fmt.Println("Error reading books directory:", err)
		return
	}

	fmt.Println("\nDaftar Buku:")
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "book-") && strings.HasSuffix(file.Name(), ".json") {
			bookData, err := ioutil.ReadFile(filepath.Join(booksDirectory, file.Name()))
			if err != nil {
				fmt.Printf("Error reading book %s: %v\n", file.Name(), err)
				continue
			}

			var book Book
			err = json.Unmarshal(bookData, &book)
			if err != nil {
				fmt.Printf("Error decoding book %s: %v\n", file.Name(), err)
				continue
			}

			fmt.Printf("Kode: %s, Judul: %s, Pengarang: %s, Penerbit: %s, Jumlah Halaman: %d, Tahun Terbit: %d\n", book.Kode, book.Judul, book.Pengarang, book.Penerbit, book.JumlahHalaman, book.TahunTerbit)
		}
	}
}

func editBook() {
	var kode string
	fmt.Print("Masukkan kode buku yang ingin diubah: ")
	fmt.Scanln(&kode)

	filePath := filepath.Join(booksDirectory, "book-"+kode+".json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Buku dengan kode tersebut tidak ditemukan.")
		return
	}

	var newBook Book

	fmt.Println("\nMasukkan detail buku baru:")
	fmt.Print("Judul: ")
	fmt.Scanln(&newBook.Judul)
	fmt.Print("Pengarang: ")
	fmt.Scanln(&newBook.Pengarang)
	fmt.Print("Penerbit: ")
	fmt.Scanln(&newBook.Penerbit)
	fmt.Print("Jumlah Halaman: ")
	fmt.Scanln(&newBook.JumlahHalaman)
	fmt.Print("Tahun Terbit: ")
	fmt.Scanln(&newBook.TahunTerbit)

	// Simpan ke file JSON
	bookData, err := json.MarshalIndent(newBook, "", "  ")
	if err != nil {
		fmt.Println("Error encoding book data:", err)
		return
	}

	err = ioutil.WriteFile(filePath, bookData, 0644)
	if err != nil {
		fmt.Println("Error saving book data:", err)
		return
	}

	fmt.Println("Buku berhasil diubah.")
}

func deleteBook() {
	var kode string
	fmt.Print("Masukkan kode buku yang ingin dihapus: ")
	fmt.Scanln(&kode)

	filePath := filepath.Join(booksDirectory, "book-"+kode+".json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Buku dengan kode tersebut tidak ditemukan.")
		return
	}

	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting book:", err)
		return
	}

	fmt.Println("Buku berhasil dihapus.")
}

func printMenu() {
	fmt.Println("\nPrint Menu:")
	fmt.Println("1. Print Satu Buku")
	fmt.Println("2. Print Semua Buku")
	fmt.Println("3. Kembali ke Menu Utama")

	var choice int
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		printOneBook()
	case 2:
		printAllBooks()
	case 3:
		return
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func printOneBook() {
	var kode string
	fmt.Print("Masukkan kode buku yang ingin di-print: ")
	fmt.Scanln(&kode)

	filePath := filepath.Join(booksDirectory, "book-"+kode+".json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Buku dengan kode tersebut tidak ditemukan.")
		return
	}

	bookData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading book data:", err)
		return
	}

	var book Book
	err = json.Unmarshal(bookData, &book)
	if err != nil {
		fmt.Println("Error decoding book data:", err)
		return
	}

	printBookToPDF(book)
}

func printAllBooks() {
	files, err := ioutil.ReadDir(booksDirectory)
	if err != nil {
		fmt.Println("Error reading books directory:", err)
		return
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "book-") && strings.HasSuffix(file.Name(), ".json") {
			bookData, err := ioutil.ReadFile(filepath.Join(booksDirectory, file.Name()))
			if err != nil {
				fmt.Printf("Error reading book %s: %v\n", file.Name(), err)
				continue
			}

			var book Book
			err = json.Unmarshal(bookData, &book)
			if err != nil {
				fmt.Printf("Error decoding book %s: %v\n", file.Name(), err)
				continue
			}

			printBookToPDF(book)
		}
	}

	fmt.Println("Print semua buku selesai.")
}

func printBookToPDF(book Book) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Kode: "+book.Kode)
	pdf.Ln(10)
	pdf.Cell(0, 10, "Judul: "+book.Judul)
	pdf.Ln(10)
	pdf.Cell(0, 10, "Pengarang: "+book.Pengarang)
	pdf.Ln(10)
	pdf.Cell(0, 10, "Penerbit: "+book.Penerbit)
	pdf.Ln(10)
	pdf.Cell(0, 10, "Jumlah Halaman: "+strconv.Itoa(book.JumlahHalaman))
	pdf.Ln(10)
	pdf.Cell(0, 10, "Tahun Terbit: "+strconv.Itoa(book.TahunTerbit))

	fileName := fmt.Sprintf("%s-%s.pdf", strings.ReplaceAll(book.Judul, " ", "_"), time.Now().Format("2006-01-02_15-04-05"))
	err := pdf.OutputFileAndClose(filepath.Join(pdfDirectory, fileName))
	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	fmt.Println("Buku", book.Judul, "telah di-print ke file", fileName)
}

func main() {
	// Membuat direktori books jika belum ada
	err := os.MkdirAll(booksDirectory, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating books directory:", err)
		return
	}

	// Membuat direktori pdf jika belum ada
	err = os.MkdirAll(pdfDirectory, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating pdf directory:", err)
		return
	}

	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Input Buku")
		fmt.Println("2. List Buku")
		fmt.Println("3. Edit Buku")
		fmt.Println("4. Hapus Buku")
		fmt.Println("5. Print Buku")
		fmt.Println("6. Keluar")

		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			inputBook()
		case 2:
			listBooks()
		case 3:
			editBook()
		case 4:
			deleteBook()
		case 5:
			printMenu()
		case 6:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
