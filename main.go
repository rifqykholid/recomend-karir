package main

import "fmt"

const NMAX = 100

type indentitas struct {
	nama, minat, keahlian string
}
type tabstring [NMAX]indentitas

type pekerjaan struct {
	industri string
	minat    [3]string
	keahlian [3]string
}
type tabPekerjaan [NMAX]pekerjaan

type hasilIndustri struct {
	industri       string
	totalPekerjaan int
	totalCocok     int
	persen         float64
}
type tabHasil [NMAX]hasilIndustri

func menu() {
	fmt.Println("+---------------------------------------------+")
	fmt.Println("		  MENU")
	fmt.Println("+---------------------------------------------+")
	fmt.Println("1. Menambahkan Identitas")
	fmt.Println("2. Mengubah Identitas")
	fmt.Println("3. Menghapus Indentitas")
	fmt.Println("4. Tampilkan Identitas")
	fmt.Println("5. Tamplikan Pekerjaan & Presentase")
	fmt.Println("0. Exit")
	fmt.Print("Pilih Menu : ")
}

func inputan(A *tabstring, n *int) {
	i := *n
	fmt.Print("Masukan nama: ")
	fmt.Scan(&A[i].nama)
	minat(A, i)
	keahlian(A, i)
	fmt.Println("Data berhasil ditambahkan")
}

func minat(A *tabstring, n int) {
	var x int

	fmt.Println("Pilih minat anda:")
	fmt.Println("1. Membaca")
	fmt.Println("2. Desain")
	fmt.Println("3. Musik")
	fmt.Println("4. Olahraga")
	fmt.Println("5. Kuliner")
	fmt.Println("6. Game")
	fmt.Println("7. Traveling")
	fmt.Println("8. Fotografi/Videografi")
	fmt.Println("9. Sosial media")
	fmt.Print("Masukan angka sesuai minat: ")
	fmt.Scan(&x)

	switch x {
	case 1:
		A[n].minat = "Membaca"
	case 2:
		A[n].minat = "Desain"
	case 3:
		A[n].minat = "Musik"
	case 4:
		A[n].minat = "Olahraga"
	case 5:
		A[n].minat = "Kuliner"
	case 6:
		A[n].minat = "Game"
	case 7:
		A[n].minat = "Traveling"
	case 8:
		A[n].minat = "Fotografi/Videografi"
	case 9:
		A[n].minat = "Sosial media"
	}
}

func keahlian(A *tabstring, n int) {
	var y int
	fmt.Println("Pilih keahlian anda:")
	fmt.Println("1. Komunikasi")
	fmt.Println("2. Kerja Tim")
	fmt.Println("3. Critical Thinking")
	fmt.Println("4. Self-Learning")
	fmt.Println("5. Coding")
	fmt.Println("6. Multitasking")
	fmt.Println("7. Kreativitas")
	fmt.Println("8. Public Speaking")
	fmt.Println("9. Leadership")
	fmt.Print("masukan angka sesuai keahlian: ")
	fmt.Scan(&y)

	switch y {
	case 1:
		A[n].keahlian = "Komunikasi"
	case 2:
		A[n].keahlian = "Kerja Tim"
	case 3:
		A[n].keahlian = "Critical Thinking"
	case 4:
		A[n].keahlian = "Self-Learning"
	case 5:
		A[n].keahlian = "Coding"
	case 6:
		A[n].keahlian = "Multitasking"
	case 7:
		A[n].keahlian = "Kreativitas"
	case 8:
		A[n].keahlian = "Public Speaking"
	case 9:
		A[n].keahlian = "Leadership"
	}
}

func pengubahInputan(A *tabstring, n int) {
	var i, x int
	fmt.Printf("Masukkan index yang ingin diubah dari 0 - %d : ", n-1)
	fmt.Scan(&i)

	fmt.Println("Pilih yang ingin diganti: ")
	fmt.Println("1. nama")
	fmt.Println("2. minat")
	fmt.Println("3. keahlian")
	fmt.Print("pilih angka yang ingin diganti: ")
	fmt.Scan(&x)

	switch x {
	case 1:
		fmt.Print("Masukan nama: ")
		fmt.Scan(&A[i].nama)
	case 2:
		minat(A, i)
	case 3:
		keahlian(A, i)
	}
}

func penghapusInputan(A *tabstring, n *int) {
	var i, j int
	var name string
	var status bool = false
	fmt.Print("Data yang ingin dihapus berdasarkan nama: ")
	fmt.Scan(&name)
	i = 0
	for i < *n {
		if A[i].nama == name {
			for j = i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			status = true
			*n--
			fmt.Println("Data telah dihapus")
		}
		i++
	}
	if !status {
		fmt.Println("Data tidak ditemukan")
	}
}

func daftarPekerjaan() (tabPekerjaan, int) {
	var D tabPekerjaan

	D[0] = pekerjaan{industri: "Keuangan", minat: [3]string{"Membaca", "Game", "Musik"}, keahlian: [3]string{"Critical Thinking", "Self-Learning", "Coding"}}
	D[1] = pekerjaan{industri: "Komunikasi", minat: [3]string{"Sosial media", "Desain", "Traveling"}, keahlian: [3]string{"Public Speaking", "Komunikasi", "Leadership"}}
	D[2] = pekerjaan{industri: "Kreatif", minat: [3]string{"Desain", "Musik", "Fotografi/Videografi"}, keahlian: [3]string{"Kreativitas", "Multitasking", "Komunikasi"}}
	D[3] = pekerjaan{industri: "Kuliner", minat: [3]string{"Kuliner", "Traveling", "Fotografi/Videografi"}, keahlian: [3]string{"Multitasking", "Kreativitas", "Self-Learning"}}
	D[4] = pekerjaan{industri: "Manajemen", minat: [3]string{"Sosial media", "Membaca", "Desain"}, keahlian: [3]string{"Public Speaking", "Komunikasi", "Critical Thinking"}}
	D[5] = pekerjaan{industri: "Media", minat: [3]string{"Membaca", "Musik", "Desain"}, keahlian: [3]string{"Critical Thinking", "Self-Learning", "Kreativitas"}}
	D[6] = pekerjaan{industri: "Olahraga", minat: [3]string{"Olahraga", "Game", "Traveling"}, keahlian: [3]string{"Multitasking", "Leadership", "Kerja Tim"}}
	D[7] = pekerjaan{industri: "Pemasaran", minat: [3]string{"Sosial media", "Membaca", "Desain"}, keahlian: [3]string{"Public Speaking", "Komunikasi", "Critical Thinking"}}
	D[8] = pekerjaan{industri: "Pendidikan", minat: [3]string{"Sosial media", "Game", "Olahraga"}, keahlian: [3]string{"Coding", "Public Speaking", "Kreativitas"}}
	D[9] = pekerjaan{industri: "Teknologi", minat: [3]string{"Game", "Fotografi/Videografi", "Musik"}, keahlian: [3]string{"Critical Thinking", "Coding", "Self-Learning"}}

	return D, 10
}

func hitungPersentase(A tabstring, n int, D tabPekerjaan, H *tabHasil, nPekerjaan int) {
	var nHasil int
	var status bool

	for i := 0; i < nPekerjaan; i++ {
		idx := -1
		status = false
		for j := 0; j < nHasil && !status; j++ {
			if H[j].industri == D[i].industri {
				idx = j
				status = true
			}
		}

		if idx == -1 {
			idx = nHasil
			H[nHasil].industri = D[i].industri
			nHasil++
		}

		H[idx].totalPekerjaan++

		for u := 0; u < n; u++ {
			cocok := 0
			status = false

			for j := 0; j < 3 && !status; j++ {
				if D[i].minat[j] == A[u].minat {
					cocok++
					status = true
				}
			}

			status = false
			for j := 0; j < 3 && !status; j++ {
				if D[i].keahlian[j] == A[u].keahlian {
					cocok++
					status = true
				}
			}

			H[idx].totalCocok += cocok
		}
	}

	for i := 0; i < nHasil; i++ {
		maksPoin := H[i].totalPekerjaan * n * 2
		H[i].persen = 0.0
		if maksPoin > 0 {
			H[i].persen = (float64(H[i].totalCocok) / float64(maksPoin)) * 100
		}
	}
}

func selectionSort(H *tabHasil, n int) {
	var i, j, idxMin int
	var temp hasilIndustri

	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if H[j].persen < H[idxMin].persen {
				idxMin = j
			}
		}
		if idxMin != i {
			temp = H[i]
			H[i] = H[idxMin]
			H[idxMin] = temp
		}
	}
}

func insertionSort(H *tabHasil, n int) {
	var i, j int
	var temp hasilIndustri

	for i = 1; i < n; i++ {
		temp = H[i]
		j = i - 1
		for j >= 0 && H[j].persen < temp.persen {
			H[j+1] = H[j]
			j--
		}
		H[j+1] = temp
	}
}

func binarySearchPekerjaan(D tabPekerjaan, n int) int {
	var x string
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, D[i].industri)
	}
	fmt.Print("Masukkan nama pekerjaan yang diinginkan: ")
	fmt.Scan(&x)

	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if D[mid].industri == x {
			return mid
		} else if D[mid].industri < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func tampilkanDetailHasil(A tabstring, D tabPekerjaan, H *tabHasil, n int, nPekerjaan int) {
	var pilihan int
	idxDipilih := binarySearchPekerjaan(D, nPekerjaan)
	if idxDipilih == -1 {
		fmt.Println("Pekerjaan tidak ditemukan.")
		return
	}
	pekerjaanDipilih := D[idxDipilih].industri

	var jumlahHasil int
	for i := 0; i < nPekerjaan; i++ {
		if H[i].industri != "" {
			jumlahHasil++
		}
	}

	fmt.Println("Pilih metode pengurutan")
	fmt.Println("1. Selection Sort (urutan naik)")
	fmt.Println("2. Insertion Sort (urutan turun)")
	fmt.Print("Masukkan pilihan (1/2): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		selectionSort(H, jumlahHasil)
	case 2:
		insertionSort(H, jumlahHasil)
	default:
		fmt.Println("Pilihan tidak valid, data tidak diurutkan.")
	}

	fmt.Println("===== DETAIL HASIL REKOMENDASI =====")
	for i := 0; i < n; i++ {
		fmt.Println("Nama                :", A[i].nama)
		fmt.Println("Pekerjaan Diinginkan:", pekerjaanDipilih)
		fmt.Println("Daftar Pekerjaan yang Didapat dan Persentasenya:")
		for j := 0; j < jumlahHasil; j++ {
			fmt.Printf("- %s (%.2f%%)\n", H[j].industri, H[j].persen)
		}
		fmt.Println("+---------------------------------------------+")
	}
}

func tampilkanIdentitas(A tabstring, n int) {
	var i int
	fmt.Printf(" %-14s  %-17s  %-24s\n", "Nama", "Minat", "Keahlian")
	for i = 0; i < n; i++ {
		fmt.Printf(" %-14s  %-17s  %-24s\n", A[i].nama, A[i].minat, A[i].keahlian)
	}
}

func main() {
	var A tabstring
	var H tabHasil
	var n, x int
	for {
		menu()
		fmt.Scan(&x)

		switch x {
		case 1:
			inputan(&A, &n)
			n++
		case 2:
			pengubahInputan(&A, n)
			fmt.Println("Data berhasil di ubah")
		case 3:
			penghapusInputan(&A, &n)
		case 4:
			tampilkanIdentitas(A, n)
		case 5:
			D, nPekerjaan := daftarPekerjaan()
			hitungPersentase(A, n, D, &H, nPekerjaan)
			tampilkanDetailHasil(A, D, &H, n, nPekerjaan)
		default:
			return
		}
	}
}
