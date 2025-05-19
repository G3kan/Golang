package main

import (
	"fmt"
	"time"
)

// Concurrency, birden fazla işlemin aynı anda yürütülmesi anlamına gelir.
// Paralelizm ise, birden fazla işlemin aynı anda yürütülmesi anlamına gelir.

// Goroutine, Go dilinde eşzamanlı işlemleri gerçekleştirmek için kullanılan hafif iş parçacıklarıdır.
// Goroutine'ler, Go dilinin en güçlü özelliklerinden biridir ve çok sayıda eşzamanlı işlemi kolayca yönetmenizi sağlar.
// Goroutine'ler, "go" anahtar kelimesi ile başlatılır ve arka planda çalışır.
// Goroutine'ler, Go dilinin çalışma zamanı tarafından yönetilir ve hafif iş parçacıklarıdır.
// Bu nedenle, çok sayıda goroutine başlatmak, sistem kaynaklarını aşırı derecede tüketmez.
// Goroutine'ler, genellikle uzun süren işlemleri arka planda çalıştırmak için kullanılır.
// Örneğin, bir web sunucusu, her gelen isteği işlemek için bir goroutine başlatabilir.
// Bu sayede, sunucu aynı anda birden fazla isteği işleyebilir ve daha hızlı yanıt verebilir.
// Goroutine'ler, diğer goroutine'lerle iletişim kurmak için kanallar (channels) kullanabilir.
// Kanallar, goroutine'ler arasında veri iletmek için kullanılan bir yapıdır.
// Kanallar, goroutine'ler arasında veri iletmek için kullanılan bir yapıdır.

// Bir fonksiyon tanımlayalım
func yaz(mesaj string) {
	for i := 0; i < 3; i++ {
		fmt.Println(mesaj, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	// main fonksiyonu, Go dilinde programın başlangıç noktasıdır. Goroutine'ler, main fonksiyonu içinde başlatılabilir.
	// main ana Goroutine'dir ve programın başlangıç noktasıdır.
	// Ana goroutine, programın başlangıç noktasıdır ve diğer goroutine'lerin başlatılmasından sorumludur.

	// Goroutine başlatmak için "go" anahtar kelimesini kullanıyoruz
	go yaz("Goroutine")

	// Ana fonksiyonda da çalışıyor
	yaz("Main")

	// Program hemen bitmesin diye bekleyelim
	// time.Sleep kullanılmasaydı ana goroutine hemen bitip diğer goroutine'lerin çalışmasına izin vermezdi
	// Bu nedenle, programın bitmesini beklemek için bir süre bekliyoruz
	// bu kodda işlmeciye main goroutine verilir sonra yaz goroutine'i başlatılır. bekleme süresi sırasında ,işlemci concurrent şekilde yaz goroutine'ini çalıştırır.

	time.Sleep(time.Second * 2)
}

// Goroutine'ler işlemciye verilir ve işlemci tarafından yönetilir. çıktı kestirilemez
// birbirlerini beklemeden çalışabilirler. Bu nedenle, goroutine'lerin çıktısı her seferinde farklı olabilir.
// kontrol etmek için sync.WaitGroup kullanılabilir
// sync.WaitGroup, goroutine'lerin tamamlanmasını beklemek için kullanılan bir yapıdır.
// Çıktı:
// Goroutine 0
// Main 0
// Goroutine 1
// Main 1
// Goroutine 2
// Main 2
// Goroutine 0
// Goroutine 1
// Goroutine 2
// Main 0
// Main 1
// Main 2
