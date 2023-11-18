package main

import (
	"fmt"
	"bufio"
	"os"
)

// Question struct untuk menyimpan pertanyaan dan jawaban
type Question struct {
	Pertanyaan string
	Pilihan    [4]string
	Jawaban    int
}

// Fungsi untuk menampilkan pertanyaan dan opsi jawaban
func tampilkanPertanyaan(q Question) {
	fmt.Println(q.Pertanyaan)
	for i, opsi := range q.Pilihan {
		fmt.Printf("%d. %s\n", i, opsi)
	}
}

// Fungsi untuk mengecek jawaban dan memberikan umpan balik
func cekJawaban(jawabanUser int, JawabanBenar int) bool {
	return jawabanUser == JawabanBenar
}

func main() {
	// Untuk menginput Nama
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan Nama:")

	scanner.Scan()
	nama := scanner.Text()

	// Inisialisasi beberapa pertanyaan
	pertanyaan1 := Question{
		Pertanyaan: "Yang merupakan member JKT 48?",
		Pilihan:    [4]string{"Jokowi", "Freya", "Nadin Amizah", "Igun Mama Takut"},
		Jawaban:    1,
	}

	pertanyaan2 := Question{
		Pertanyaan: "Nama Kakak Upin dan Ipin?",
		Pilihan:    [4]string{"Kak Rose", "Kak jeane", "Monkey D. Luffy", "Igun Mama Takut"},
		Jawaban:    0,
	}

	// Array pertanyaan
	pertanyaan := [2]Question{pertanyaan1, pertanyaan2}

	jawabanBenar := 0
	jawabanSalah := 0

	// Looping melalui setiap pertanyaan
	for _, q := range pertanyaan {
		// Menampilkan pertanyaan
		tampilkanPertanyaan(q)

		// Menerima jawaban dari pengguna
		var jawaban int
		fmt.Print("Masukkan nomor jawaban Anda (0,1,2,3): ")
		fmt.Scan(&jawaban)

		// Mengecek jawaban
		if cekJawaban(jawaban, q.Jawaban) {
			fmt.Print("Jawaban anda Benar!\n")
			fmt.Println("")
			jawabanBenar++

		} else {
			fmt.Print("Jawaban Anda salah. Jawaban yang benar adalah:", jawabanBenar, "\n")
			fmt.Println("")
			jawabanSalah++
		}

	}
	
	// Memberikan score, apa bila menjawab dengan benar
	score := jawabanBenar * 1;
	
	// Menampilkan Statistic Kusi
	fmt.Println("Statistic Kuis")
	fmt.Printf("Nama Peserta        = %s \n", nama)
	fmt.Printf("Total Jawaban Benar = %d \n", jawabanBenar)
	fmt.Printf("Total Jawaban Salah = %d \n", jawabanSalah)
	fmt.Printf("Total Score         = %d \n", score)
}