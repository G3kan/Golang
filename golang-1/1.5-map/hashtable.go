package main

import "fmt"

// map yapısı, anahtar-değer çiftlerini saklamak için kullanılır.
// Anahtarlar benzersiz olmalıdır ve genellikle string veya int gibi temel veri türleri kullanılır.
// Map'ler, hızlı erişim ve güncelleme işlemleri için idealdir.
// Map'ler, Go dilinde yerleşik bir veri yapısıdır ve dinamik boyutlandırma özelliğine sahiptir.
// Map'ler, hash tablosu olarak implement edilmiştir ve bu nedenle anahtarların hash değerleri kullanılarak erişilir.
// Map'ler, sıralı değildir ve anahtarların sırası her zaman değişebilir.
// Map'ler, nil olarak başlatılabilir ve daha sonra eleman eklenebilir.
// Map'ler, sıfır uzunluğunda başlatıldığında boş bir harita olarak kabul edilir.
// Map'ler, anahtarların sıralı olmadığı için, belirli bir sırayla dolaşmak mümkün değildir.

var kullanicilar = make(map[string]string) // Anahtar: Kullanıcı adı, Değer: Kullanıcı bilgileri nil olarak başlatıl
//nil: bir değişkenin başlangıç değeri yok demektir.

func main() {
	// Bir map oluşturma
	myMap := make(map[string]int)

	// nil değeri ile başlatma
	kullanicilar = make(map[string]string)

	// Değer ekleme
	myMap["Alice"] = 25
	myMap["Bob"] = 30

	// Değer okuma
	fmt.Println("Alice'in yaşı:", myMap["Alice"])

	// Değer güncelleme
	myMap["Alice"] = 26

	// Değer silme
	delete(myMap, "Bob")

	// Map'i dolaşma
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
