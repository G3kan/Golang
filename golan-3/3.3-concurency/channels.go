package main

import (
	"fmt"
	"time"
)

// channel'lar, Go dilinde goroutine'ler arasında veri iletimi için kullanılan bir yapıdır.
// channel'lar, veri gönderme ve alma işlemlerini senkronize bir şekilde gerçekleştirir.
// channel'lar, veri türünü belirtmek için kullanılır ve bu türdeki verileri taşır.
// channel'lar, buffered (tamponlu) ve unbuffered (tamponsuz) olmak üzere iki türde olabilir.
// Unbuffered channel, veri gönderildiğinde alıcının da hazır olmasını bekler.
// Buffered channel ise belirli bir kapasiteye sahip olup, veri gönderildiğinde alıcının hazır olmasını beklemez.
// Buffered channel, belirli bir sayıda veriyi saklayabilir ve bu sayede veri gönderimi sırasında alıcının hazır olmasını beklemeden veri göndermeye devam edebilir.

// Unbuffered channel örneği
func unbufferedExample() {

	// chanel oluşturma
	ch := make(chan string) // Unbuffered channel

	go func() {

		// Gönderici goroutine
		// channel'a veri gönderme için <- operatörü kullanılır
		// Gönderici goroutine, alıcı goroutine'in hazır olmasını bekler
		ch <- "Merhaba, unbuffered channel!" // Gönderici burada bekler
	}()

	// Alıcı goroutine
	// channel'dan veri alma için <- operatörü kullanılır
	// Alıcı goroutine, gönderici goroutine'in veri göndermesini bekler
	/*
		Block (Bloklama):
		Bir goroutine, bir channel üzerinden veri göndermeye veya almaya çalıştığında ve karşı tarafta (alıcı veya gönderici) hazır bir goroutine yoksa, o goroutine bloklanır (yani beklemeye alınır). Örneğin, unbuffered channel'da gönderici, alıcı hazır olana kadar bekler; alıcı da gönderici veri gönderene kadar bekler.

		Deadlock (Kilitleme):
		Tüm goroutine'ler bloklandığında ve hiçbiri ilerleyemediğinde deadlock (kilitleme) oluşur. Örneğin, ana goroutine bir channel'dan veri almaya çalışırken, o channel'a veri gönderecek başka bir goroutine yoksa program deadlock ile sonlanır. Go runtime, deadlock tespit ettiğinde "fatal error: all goroutines are asleep - deadlock!" hatası verir.
	*/
	msg := <-ch // Alıcı burada bekler
	fmt.Println(msg)
}

/*
	Unbuffered channel'da, bir veri alınmadan diğeri gelirse ne olur?

	Unbuffered channel'lar, gönderici ve alıcı arasında tam bir senkronizasyon gerektirir. Yani, bir goroutine channel'a veri göndermeye çalıştığında, diğer bir goroutine o veriyi almaya hazır olana kadar gönderici bekler (bloklanır). Aynı şekilde, bir goroutine channel'dan veri almaya çalıştığında, gönderici veri gönderene kadar alıcı bekler.

	Eğer birden fazla veri gönderilmeye çalışılırsa ve arada veri alınmazsa, ilk gönderimden sonra gönderici goroutine bloklanır ve diğer gönderimler gerçekleşmez. Yani, unbuffered channel'da, bir veri alınmadan ikinci bir veri gönderilemez. Bu durum, programda deadlock (kilitlenme) oluşmasına neden olabilir.

	Özetle:
	- Unbuffered channel'da, her gönderim için bir alıcı olmalıdır.
	- Alıcı olmadan birden fazla veri gönderilmeye çalışılırsa, program bloklanır ve ilerlemez.
	- Bu nedenle, unbuffered channel'lar tam senkronizasyon gerektiren durumlarda kullanılır.
*/

// Buffered channel örneği
// Buffered channel'lar, belirli bir kapasiteye sahip channel'lardır.
// Gönderici, buffer dolu olmadığı sürece veriyi hemen gönderebilir ve bloklanmaz.
// Alıcı, buffer boş olmadığı sürece veriyi hemen alabilir ve bloklanmaz.
// Eğer buffer doluysa, gönderici bloklanır; buffer boşsa, alıcı bloklanır.
// Bu sayede, gönderici ve alıcı arasında tam senkronizasyon gerekmez ve daha esnek bir iletişim sağlanır.

func bufferedExample() {

	// Buffered channel oluşturma
	// Buffered channel, belirli bir kapasiteye sahip olup, veri gönderildiğinde alıcının hazır olmasını beklemez
	// Buffered channel, belirli bir sayıda veriyi saklayabilir ve bu sayede veri gönderimi sırasında alıcının hazır olmasını beklemeden veri göndermeye devam edebilir
	ch := make(chan string, 2) // Buffered channel, kapasite 2

	// Buffered channel'da blok ve deadlock nasıl olur?
	/*
		Blok (Blocking):
		- Buffered channel'ın kapasitesi dolduğunda, gönderici goroutine yeni veri göndermeye çalışırsa bloklanır (bekler).
		- Aynı şekilde, buffer boşsa ve alıcı veri almaya çalışırsa, alıcı bloklanır.

		Deadlock (Kilitleme):
		- Eğer tüm goroutine'ler buffered channel'a veri göndermeye veya almaya çalışırken bloklanır ve hiçbiri ilerleyemezse deadlock oluşur.
		- Örneğin, buffer kapasitesi kadar veri gönderilip, hiç veri alınmazsa ve başka bir goroutine yoksa, program deadlock ile sonlanır.
		- Go runtime, deadlock tespit ettiğinde "fatal error: all goroutines are asleep - deadlock!" hatası verir.
	*/

	ch <- "Birinci mesaj" // Hemen gönderilir
	ch <- "İkinci mesaj"  // Hemen gönderilir

	// ch <- "Üçüncü mesaj" // Bu satır açılırsa, buffer dolu olduğu için bekler

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func main() {
	fmt.Println("Unbuffered channel örneği:")
	unbufferedExample()

	time.Sleep(time.Second)

	fmt.Println("\nBuffered channel örneği:")
	bufferedExample()
}
