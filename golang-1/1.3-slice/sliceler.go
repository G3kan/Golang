package main

import "fmt"

// Slice'lar, Go'da dinamik boyutlu dizilerdir.
// Array'lerin aksine, slice'ların uzunluğu değiştirilebilir.
// Slice'lar, aslında bir array’in belirli bir bölümüne işaret eden bir referans yapısıdır.

var arr = [...]string{"a", "b", "c", "d", "e"} // Array tanımı

var slice1 []string // Slice tanımı
var slice2 []string

/*
Aşağıdaki kod blogu çalışmaz çünkü execute gerekiyor slice için func lazım.
slice1=arr[1:3] // 1. ve 2. elemanları alır
slice2=arr[2:5]  // 2. elemandan 4. elemana kadar alır
*/

// Array'ler sabit uzunlukludur ve tanımlandıktan sonra boyutları değiştirilemez.
// Slice'lar ise array'lere kıyasla daha esnektir ve gerektiğinde büyüyebilir.

// Array'ler doğrudan elemanları bellekte tutar, slice'lar ise bir array'in belirli bir bölümüne referans verir.
// Bu nedenle slice’lar, aynı array üzerinde farklı bölümleri gösterebilir ve elemanları değiştirebilir.

// Slice’lar ile çalışırken kapasite (`cap`) ve uzunluk (`len`) kavramları ayrıdır.
// Uzunluk (`len`), slice içindeki mevcut eleman sayısını belirtirken, kapasite (`cap`), slice’in dayandığı array'in kullanılabilecek toplam kapasitesini ifade eder.

// Slice'lar hafıza açısından daha verimli olabilir, çünkü yeni bir kopya oluşturulmadan kullanılabilirler.
// Ancak, slice'in değiştirilmesi durumunda bağlı olduğu array de etkilenebilir.
// Eğer slice’in bağımsız olması istenirse, yeni bir slice oluşturup `copy()` fonksiyonu ile veri aktarılmalıdır.

/*Slice'ın Özellikleri:
1. Pointer: Slice'ın başlangıç elemanını gösterir.
2. Length: Slice'daki eleman sayısını belirtir.
3. Capacity: Slice'ın maksimum eleman sayısını gösterir; bu, başlangıç elemanının konumuna bağlıdır.*/

func main() {

	//len, bir slice'ın içindeki mevcut eleman sayısını döndürür.
	slice3 := []int{1, 2, 3, 4}
	fmt.Println(len(slice3)) // Çıktı: 4

	//cap, bir slice'ın kapasitesini döndürür.
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice4 := arr1[1:4]      // 2. elemandan 4. elemana kadar
	fmt.Println(cap(slice4)) // Çıktı: 4 (arr[1:]'in kapasitesi)

	// 1. Diziden Slice Oluşturma
	arr := [5]int{1, 2, 3, 4, 5}            // Dizi
	slice := arr[1:4]                       // 2. elemandan 4. elemana kadar (4 dahil değil)
	fmt.Println("Slice from array:", slice) // Çıktı: [2 3 4]

	// 2. make ile Slice Oluşturma
	slice2 := make([]int, 3, 5)                     // Uzunluk: 3, Kapasite: 5
	fmt.Println("Slice created with make:", slice2) // Çıktı: [0 0 0]

	// 2-1. make ile Slice Oluşturma (Örnek)
	slice20 := make([]int, 10)                       // Uzunluk: 10 uzunluk ve kapasite aynı olur Kapasite: 10
	fmt.Println("Slice created with make:", slice20) // Çıktı: [0 0 0 0 0 0 0 0 0 0]

	// 3. Literal ile Slice Tanımlama
	slice5 := []int{10, 20, 30}
	fmt.Println("Slice literal:", slice5) // Çıktı: [10 20 30]

	// 4. Eleman Ekleme (append)
	slice3 = append(slice3, 40, 50)
	fmt.Println("After append:", slice3) // Çıktı: [10 20 30 40 50]

	// 5. Eleman Silme
	slice3 = append(slice3[:2], slice3[3:]...) // 3. elemanı sil
	fmt.Println("After deletion:", slice3)     // Çıktı: [10 20 40 50]

	// 6. Slice Kopyalama (copy)
	src := []int{1, 2, 3}
	dest := make([]int, len(src))
	copy(dest, src)
	fmt.Println("Source slice:", src)  // Çıktı: [1 2 3]
	fmt.Println("Copied slice:", dest) // Çıktı: [1 2 3]

	//append ile slice'e eleman ekleme
	slice6 := []int{1, 2, 3}
	slice6 = append(slice6, 4, 5)        // 4 ve 5'i ekle
	fmt.Println("After append:", slice6) // Çıktı: [1 2 3 4 5]
	//append ile slice'e başka bir slice ekleme
	slice7 := []int{6, 7, 8}
	slice6 = append(slice6, slice7...)                    // slice7'yi ekle
	fmt.Println("After appending another slice:", slice6) // Çıktı: [1 2 3 4 5 6 7 8]

	// 7. Slice'ı Kopyalama
	slice8 := []int{1, 2, 3}
	slice9 := make([]int, len(slice8))
	copy(slice9, slice8)                   // slice8'i kopyala
	fmt.Println("Original slice:", slice8) // Çıktı: [1 2 3]
	fmt.Println("Copied slice:", slice9)   // Çıktı: [1 2 3]

	// 8. Slice'ı Sıralama
	slice10 := []int{5, 3, 4, 1, 2}
	fmt.Println("Before sorting:", slice10) // Çıktı: [5 3 4 1 2]
	// sort.Ints(slice10) // Sıralama işlemi için sort paketi kullanılabilir
	// fmt.Println("After sorting:", slice10) // Çıktı: [1 2 3 4 5]

	// 9. Slice'ı Ters Çevirme

	slice11 := []int{1, 2, 3, 4, 5}
	fmt.Println("Before reversing:", slice11) // Çıktı: [1 2 3 4 5]
	for i, j := 0, len(slice11)-1; i < j; i, j = i+1, j-1 {
		slice11[i], slice11[j] = slice11[j], slice11[i]
	}
	fmt.Println("After reversing:", slice11) // Çıktı: [5 4 3 2 1]

	// 10. Slice'ı Elemanları ile Doldurma
	slice12 := make([]int, 5)
	for i := range slice12 {
		slice12[i] = i * 2 // Her elemanı 2 ile çarp
	}
	fmt.Println("Filled slice:", slice12) // Çıktı: [0 2 4 6 8]
	// 11. Slice'ı Elemanları ile Doldurma (Örnek)
	slice13 := make([]int, 5)
	for i := range slice13 {
		slice13[i] = i + 1 // Her elemanı 1 artır
	}
	fmt.Println("Filled slice:", slice13) // Çıktı: [1 2 3 4 5]

	// 12. Slice dan eleman silme
	slice14 := []int{1, 2, 3, 4, 5}
	slice14 = append(slice14[:2], slice14[3:]...) // 3. elemanı sil
	fmt.Println("After deletion:", slice14)       // Çıktı: [1 2 4 5]

}
