package main

import "fmt"

// Struct untuk menyimpan informasi ide
type Idea struct {
	ID        int
	Judul     string
	Kategori  string
	Upvote    int
	Tanggal   int // gunakan format YYYYMMDD
}

// Array penyimpanan ide dan variabel pendukung 
var daftarIde [100]Idea
var jumlahIde int = 0
var idCounter int = 1

// Fungsi untuk menambahkan ide baru
func tambahIde() {
	// Cek kapasitas 
	if jumlahIde >= len(daftarIde) {
		fmt.Println("Kapasitas penuh.")
		return
	}
	// Input data dari user 
	var judul, kategori string
	var tanggal int
	fmt.Print("Judul: ")
	fmt.Scan(&judul)
	fmt.Print("Kategori: ")
	fmt.Scan(&kategori)
	fmt.Print("Tanggal (YYYYMMDD): ")
	fmt.Scan(&tanggal)

	// Simpan ide ke array dengan ID otomatis 
	daftarIde[jumlahIde] = Idea{idCounter, judul, kategori, 0, tanggal}
	jumlahIde++
	idCounter++
	fmt.Println("Ide berhasil ditambahkan.")
}

// Fungsi untuk mengubah data ide berdasarkan ID 
func ubahIde() {
	var id int
	fmt.Print("Masukkan ID ide yang ingin diubah: ")
	fmt.Scan(&id)
	// Cari ide berdasarkan ID
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].ID == id {
			// Update data ide
			fmt.Print("Judul baru: ")
			fmt.Scan(&daftarIde[i].Judul)
			fmt.Print("Kategori baru: ")
			fmt.Scan(&daftarIde[i].Kategori)
			fmt.Print("Tanggal baru (YYYYMMDD): ")
			fmt.Scan(&daftarIde[i].Tanggal)
			fmt.Println("Ide berhasil diubah.")
			return
		}
	}
	fmt.Println("Ide tidak ditemukan.")
}

// Fungsi untuk menghapus ide berdasarkan ID 
func hapusIde() {
	var id int
	fmt.Print("Masukkan ID ide yang ingin dihapus: ")
	fmt.Scan(&id)
	// Cari ide berdasarkan ID
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].ID == id {

			// Geser elemen array ke kiri untuk menghapus ide 
			for j := i; j < jumlahIde-1; j++ {
				daftarIde[j] = daftarIde[j+1]
			}
			jumlahIde--
			fmt.Println("Ide berhasil dihapus.")
			return
		}
	}
	fmt.Println("Ide tidak ditemukan.")
}

// Fungsi untuk memberikan upvote pada ide 
func upvoteIde() {
	var id int
	fmt.Print("Masukkan ID ide yang ingin di-upvote: ")
	fmt.Scan(&id)
	// Cari ide berdasarkan ID dan tambahkan Upvote
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].ID == id {
			daftarIde[i].Upvote++
			fmt.Println("Upvote berhasil.")
			return
		}
	}
	fmt.Println("Ide tidak ditemukan.")
}

// Pencarian ide berdasarkan kategori menggunakan Sequential Search 
func cariSequential() {
	var kategori string
	fmt.Print("Masukkan kategori yang ingin dicari: ")
	fmt.Scan(&kategori)
	ditemukan := false
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].Kategori == kategori {
			tampilkanIde(daftarIde[i])
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada ide dengan kategori tersebut.")
	}
}

// Pencarian ide berdasarkan judul menggunakan Binary Search 
func cariBinary() {
	urutJudul() // Harus diurutkan dulu berdasarkan Judul
	var keyword string
	fmt.Print("Masukkan keyword judul (cocok persis): ")
	fmt.Scan(&keyword)
	low, high := 0, jumlahIde-1
	for low <= high {
		mid := (low + high) / 2
		if daftarIde[mid].Judul == keyword {
			tampilkanIde(daftarIde[mid])
			return
		} else if daftarIde[mid].Judul < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Ide tidak ditemukan.")
}

// Mengurutkan ide berdasarkan jumlah upvote tertinggi (Selection Sort) 
func urutUpvote() {
	// Selection Sort untuk mengurutkan berdasarkan upvote (descending) 
	for i := 0; i < jumlahIde-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahIde; j++ {
			if daftarIde[j].Upvote > daftarIde[maxIdx].Upvote {
				maxIdx = j
			}
		}
		daftarIde[i], daftarIde[maxIdx] = daftarIde[maxIdx], daftarIde[i]
	}
	fmt.Println("Data telah diurut berdasarkan upvote.")
	for i := 0; i < jumlahIde; i++ {
		tampilkanIde(daftarIde[i])
	}
}

// Mengurutkan ide berdasarkan tanggal (paling lama ke baru) menggunakan Insertion Sort 
func urutTanggal() {
	// Insertion Sort untuk mengurutkan berdasarkan tanggal (ascending)
	for i := 1; i < jumlahIde; i++ {
		key := daftarIde[i]
		j := i - 1
		for j >= 0 && daftarIde[j].Tanggal > key.Tanggal {
			daftarIde[j+1] = daftarIde[j]
			j--
		}
		daftarIde[j+1] = key
	}
	fmt.Println("Data telah diurut berdasarkan tanggal.")
	for i := 0; i < jumlahIde; i++ {
		tampilkanIde(daftarIde[i])
	}
}

// Menampilkan ide yang tanggalnya berada setelah batas tertentu 
func populerPeriode() {
	var batas int
	fmt.Print("Masukkan tanggal batas bawah (YYYYMMDD): ")
	fmt.Scan(&batas)
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].Tanggal >= batas {
			tampilkanIde(daftarIde[i])
		}
	}
}

// Mengurutkan ide berdasarkan judul secara alfabetis (Insertion Sort) 
func urutJudul() {
	for i := 1; i < jumlahIde; i++ {
		key := daftarIde[i]
		j := i - 1
		for j >= 0 && daftarIde[j].Judul > key.Judul {
			daftarIde[j+1] = daftarIde[j]
			j--
		}
		daftarIde[j+1] = key
	}
}

// Menampilkan informasi lengkap sebuah ide 
func tampilkanIde(i Idea) {
	fmt.Printf("ID: %d, Judul: %s, Kategori: %s, Upvote: %d, Tanggal: %d\n", i.ID, i.Judul, i.Kategori, i.Upvote, i.Tanggal)
}

// Fungsi utama 
func main() {
	var pilihan int
	for {
		// Menu utama 
		fmt.Println("\nAplikasi Pengelolaan Ide dan Brainstorming")
		fmt.Println("1. Tambah Ide")
		fmt.Println("2. Ubah Ide")
		fmt.Println("3. Hapus Ide")
		fmt.Println("4. Upvote Ide")
		fmt.Println("5. Cari Ide berdasarkan Kategori (Sequential Search)")
		fmt.Println("6. Cari Ide berdasarkan Judul (Binary Search)")
		fmt.Println("7. Urutkan Ide (Upvote - Selection Sort)")
		fmt.Println("8. Urutkan Ide (Tanggal - Insertion Sort)")
		fmt.Println("9. Tampilkan Ide Terpopuler dalam Periode")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahIde()
		case 2:
			ubahIde()
		case 3:
			hapusIde()
		case 4:
			upvoteIde()
		case 5:
			cariSequential()
		case 6:
			cariBinary()
		case 7:
			urutUpvote()
		case 8:
			urutTanggal()
		case 9:
			populerPeriode()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
} 
