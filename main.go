package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Jumlah mahasiswa
	var jumlahMahasiswa int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scanln(&jumlahMahasiswa)

	// Slice untuk menyimpan nama dan rata-rata setiap mahasiswa
	namaMahasiswa := make([]string, jumlahMahasiswa)
	rataRataMahasiswa := make([]float64, jumlahMahasiswa)
	nilaiHurufMahasiswa := make([]string, jumlahMahasiswa)

	scanner := bufio.NewScanner(os.Stdin) // Untuk membaca input baris penuh

	// Input data
	for i := 0; i < jumlahMahasiswa; i++ {
		fmt.Printf("Masukkan nama mahasiswa ke-%d: ", i+1)
		scanner.Scan()
		namaMahasiswa[i] = scanner.Text()

		// Meminta input nilai
		var nilai [5]float64
		fmt.Print("Masukkan nilai untuk 5 mata kuliah: ")
		scanner.Scan() // Membaca satu baris penuh untuk nilai
		nilaiInput := strings.Fields(scanner.Text())

		total := 0.0
		for j := 0; j < 5; j++ {
			nilai[j], _ = strconv.ParseFloat(nilaiInput[j], 64)
			total += nilai[j]
		}

		// Hitung rata-rata
		rataRata := total / 5
		rataRataMahasiswa[i] = rataRata

		if rataRata >= 80 {
			nilaiHurufMahasiswa[i] = "A"
		} else if rataRata >= 70 {
			nilaiHurufMahasiswa[i] = "B"
		} else if rataRata >= 60 {
			nilaiHurufMahasiswa[i] = "C"
		} else if rataRata >= 50 {
			nilaiHurufMahasiswa[i] = "D"
		} else {
			nilaiHurufMahasiswa[i] = "E"
		}

		fmt.Println()
	}

	// Menampilkan hasil
	for i := 0; i < jumlahMahasiswa; i++ {
		fmt.Printf("Nama: %s, Rata-rata: %.1f, Nilai Huruf: %s\n", namaMahasiswa[i], rataRataMahasiswa[i], nilaiHurufMahasiswa[i])
	}

	// Mencari dan menampilkan mahasiswa dengan nilai rata-rata tertinggi
	tertinggiIndex := 0
	for i := 1; i < jumlahMahasiswa; i++ {
		if rataRataMahasiswa[i] > rataRataMahasiswa[tertinggiIndex] {
			tertinggiIndex = i
		}
	}
	fmt.Printf("Mahasiswa dengan nilai rata-rata tertinggi: %s dengan rata-rata %.1f\n", namaMahasiswa[tertinggiIndex], rataRataMahasiswa[tertinggiIndex])
}
