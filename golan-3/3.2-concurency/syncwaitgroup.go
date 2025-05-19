package main

import (
	"fmt"
	"sync"
	"time"
)

// Go’da main() fonksiyonu bittiği anda tüm goroutine'ler durur.
// Bunu önlemek ve ana programın, diğer goroutine’lerin bitmesini beklemesini sağlamak için sync.WaitGroup kullanılır.
// sync.WaitGroup, birkaç goroutine’in tamamlanmasını beklemek için kullanılan bir yapıdır.

// sync.WaitGroup, bir sayaç gibi çalışır. Bu sayaç, beklenen goroutine sayısını tutar.
// Goroutine’ler tamamlandıkça sayaç azalır. Sayaç sıfıra ulaştığında, ana program devam eder.
// wg.Add(n) ile beklenen goroutine sayısını artırırız.
// wg.Done() ile bir goroutine tamamlandığında sayaç 1 azalır.
// wg.Wait() ile ana program, sayaç sıfıra ulaşana kadar bekler.

func sayHello(wg *sync.WaitGroup) {
	// Goroutine içinde çalışacak fonksiyon
	// Bu fonksiyon, bir goroutine olarak çalışacak
	defer wg.Done() // İş bitince sayacı 1 azalt
	// defer anahtar kelimesi, fonksiyonun sonunda çalışacak olan bir işlemi belirtir
	// Goroutine tamamlandığında, wg.Done() çağrılır
	// Bu, ana programın beklemesini sağlar
	// Goroutine içinde yapılacak işlemler
	// Örnek olarak, 5 kez "Hello" yazdıracağız
	// ve her seferinde 500 ms bekleyeceğiz
	// Bu, goroutine'in çalıştığını ve ana programın beklediğini gösterir
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	var wg sync.WaitGroup // WaitGroup oluştur
	// WaitGroup'ü kullanmak için bir değişken oluşturuyoruz
	// WaitGroup, goroutine'lerin bitmesini beklemek için kullanılır
	// WaitGroup, bir sayaç gibi çalışır. Bu sayaç, beklenen goroutine sayısını tutar.
	// Goroutine’ler tamamlandıkça sayaç azalır. Sayaç sıfıra ulaştığında, ana program devam eder.

	wg.Add(1) // 1 tane goroutine bekleyeceğimizi söylüyoruz

	go sayHello(&wg) // Goroutine başlat ve wg'yi gönder

	wg.Wait() // sayHello bitene kadar burada bekler
	fmt.Println("Tüm işler bitti.")
}
