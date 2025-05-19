package main

import (
	"fmt"
	"time"
)

//select ifadesi ile birden fazla kanaldan veri almak mümkündür.
// Bu örnekte iki farklı goroutine oluşturuluyor ve her biri farklı bir kanala veri gönderiyor.
// select ifadesi ile bu kanallardan gelen veriler dinleniyor. İlk hangi kanaldan veri gelirse, o case bloğu çalışır ve mesaj ekrana yazdırılır.
// Bu işlem her iki kanal için de tekrarlanır.

// main fonksiyonu, Go'nun select ifadesinin birden fazla kanal işlemini aynı anda nasıl yönettiğini gösterir.
// İki goroutine başlatılır, her biri belirli bir gecikmeden sonra farklı bir kanala mesaj gönderir.
// select ifadesi her iki kanalı da dinler ve her iterasyonda ilk gelen mesajı ekrana yazar.
//
// Select ifadesi, bir goroutine'in birden fazla iletişim işlemini beklemesini sağlar.
// Hiçbir case hazır değilse, select ifadesi bloklanır ve bir case hazır olduğunda o case çalışır.
//
// Select ifadesinde 'default' case'i kullanılırsa, hiçbir case hazır değilse bile default case hemen çalışır.
// Bu, kanal işlemlerini bloklamadan devam etmek veya zaman aşımı gibi durumları yönetmek için kullanışlıdır.
//
// 'abort' (veya 'done', 'quit' gibi) adında özel bir kanal kullanılarak goroutine'lere erken durdurma sinyali gönderilebilir.
// Select ifadesinde bu kanala ait bir case eklenirse, abort sinyali geldiğinde goroutine'ler düzgün şekilde sonlandırılabilir.
//
// 'default' kullanım örneği:
//
//	select {
//	case msg := <-ch:
//	    // mesajı işle
//	default:
//	    // mesaj yok, bloklanmadan devam et
//	}
//
// 'abort' kullanım örneği:
//
//	select {
//	case msg := <-ch:
//	    // mesajı işle
//	case <-abort:
//	    // abort sinyali geldi, goroutine'den çık
//	}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine ile ch1'e veri gönder
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "ch1'den mesaj"
	}()

	// Goroutine ile ch2'ye veri gönder
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "ch2'den mesaj"
	}()

	//case

	// select ile kanalları dinle
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Aldım:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Aldım:", msg2)
		default:
			fmt.Println("Hiçbir kanaldan mesaj yok, devam ediyorum...")
		}
	}
}
