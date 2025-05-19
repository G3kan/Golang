package main

import "fmt"

//Arays
//Dizi, sabit boyutlu bir veri yapısıdır.
//Diziler, aynı türdeki birden fazla veriyi saklamak için kullanılır.
//Diziler, bellekte ardışık olarak saklanır ve her bir elemanına indeks numarası ile erişilir.
//Diziler, sabit boyutlu olduğu için boyutları tanımlandıktan sonra değiştirilemez.
//Diziler, elemanlar eklemek veya çıkarmak için kullanılmaz.
//Dizilerde elemanlar, sıfırdan başlayarak indekslenir.
//Diziler, elamanlar değiştirilebilir.

var dizi [5]int                     // dizinin boyutu belirtilir 5 tir elemanları 0'dır
var dizi2 = [5]int{1, 2, 3, 4, 5}   //dizinin boyutu belirtilir 5tir
var dizi3 = [...]int{1, 2, 3, 4, 5} // dizinin uzunluğu otomatik olarak belirlenir

func main() {

	fmt.Println(dizi)
	fmt.Println(dizi2)
	fmt.Println(dizi3)

	//range anahtar kelimesi ile dizinin elemanlarına erişim
	//i indeksi, v ise değeri temsil eder

	for i, v := range dizi2 {
		fmt.Printf("ind: %d, value: %d\n", i, v)
	}

	fmt.Println("Dizi 2'nin uzunluğu:", len(dizi2))

	//Dizinin elemanlarına erişim
	fmt.Println("Dizi 2'nin 0. elemanı:", dizi2[0])             //belirli elemana erişme
	fmt.Println("Dizi 2'nin 1. elemanı:", dizi2[:3])            //İlk 3 elemana erişme
	fmt.Println("Dizi 2'nin 2. elemanı:", dizi2[1:3])           //2. ve 3. elemana erişme
	fmt.Println("Dizi 2'nin 3. elemanı:", dizi2[3:])            //4. elemana erişme
	fmt.Println("Dizi 2'nin 4. elemanı:", dizi2[:])             //Tüm elemana erişme
	fmt.Println("Dizi 2'nin 5. elemanı:", dizi2[len(dizi2)-2:]) //Son 2 elemana erişme

	//Python'daki gibi arr[1:5:2]  dilimleme yoktur. Yani step yoktur.
	//For döngüleri ile yapılır. for i := 1; i < 5; i += 2

	//Dizinin elemanlarını değiştirme
	dizi2[0] = 10
	fmt.Println("Dizi 2'nin 0. elemanı:", dizi2[0])

}
