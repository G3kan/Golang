package main

import (
	"fmt"
	"sync"
)

// Granularity of concurrency: Aynı anda birden fazla goroutine çalıştığında,
// paylaşılan değişkenlere erişimde sorunlar yaşanabilir. Örneğin, iki goroutine
// aynı anda bir değişkeni güncellerse, beklenmeyen sonuçlar ortaya çıkabilir.

// Aşağıdaki örnekte, iki goroutine aynı anda 'counter' değişkenini artırmaya çalışıyor.
// Eğer senkronizasyon yapılmazsa, 'counter' beklenenden daha küçük bir değerde kalabilir.

func main() {
	var counter int
	var wg sync.WaitGroup

	wg.Add(2)

	// İlk goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			// counter++ işlemi atomik değildir, bu yüzden yarış durumu oluşabilir.
			counter++
		}
		wg.Done()
	}()

	// İkinci goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Counter (race condition):", counter)

	// Lock ve Unlock kullanımı:
	// Paylaşılan değişkenlere erişimi güvenli hale getirmek için mutex kullanılır.
	// Mutex ile aynı anda sadece bir goroutine'in değişkeni güncellemesine izin verilir.

	// Bu sayede, diğer goroutine'ler bu değişkene erişemez ve yarış durumu oluşmaz.
	// Mutex, "mutual exclusion" (karşılıklı dışlama) anlamına gelir.
	// Mutex, bir goroutine'in paylaşılan bir kaynağa erişimini kontrol eder.
	// Mutex kullanarak, bir goroutine'in paylaşılan bir kaynağa erişimini kontrol edebiliriz.

	//safeCounter değişkeni, paylaşılan bir değişkendir.
	// Bu değişken, iki goroutine tarafından güncellenmektedir.
	var safeCounter int

	// mu değişkeni, mutex'i temsil eder.
	var mu sync.Mutex
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()     // Kilidi al
			safeCounter++ // Güvenli artış
			mu.Unlock()   // Kilidi bırak
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {

			//Lock değişkeni işlemi kilitleyerek diğer goroutine'lerin erişimini engeller.
			// Bu sayede değişkenin değeri güvenli bir şekilde güncellenir.
			mu.Lock()
			safeCounter++
			// Unlock işlemi, kilidi bırakarak diğer goroutine'lerin erişimine izin verir.
			// Bu sayede diğer goroutine'ler de değişkeni güncelleyebilir.
			// Ancak bu işlem, kilidi bıraktıktan sonra yapılmalıdır.
			// Aksi takdirde, diğer goroutine'ler bu kilidi beklemek zorunda kalır.
			// Bu da performans kaybına neden olabilir.
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Counter (with mutex):", safeCounter)
}
