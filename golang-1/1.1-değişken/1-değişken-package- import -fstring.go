package main

/*Go dilinde her .go dosyasının başında bir package adı bulunur.
main package’ı, çalıştırılabilir programların giriş noktasıdır.
Diğer package'lar (örneğin math, utils gibi) modüler yapılar oluşturur */

/* Bir Go programının sadece bir tane main package’ı olabilir.
main package’ı, programın çalıştırılabilir kısmıdır ve programın giriş noktası olan main() fonksiyonunu içerir.
*/

/* Diğer package'lar, programın farklı bölümlerini organize etmek için kullanılır.
Örneğin, math package'ı matematiksel işlemler, utils package'ı ise yardımcı işlevler sağlayabilir.
*/

/*import komutu, başka bir package’ı programınıza dahil etmek için kullanılır.
İçe aktarılan package, programda kullanılabilir ve o package’daki fonksiyonlar çağrılabilir.
*/

/*İçe aktarılan bir package’ın fonksiyonlarına erişim için, package adını kullanırsınız.
math.Add(), fmt.Println() gibi.
*/

// go dilinde package ve import anahtar kelimeleri ile birlikte kullanılan bazı önemli kütüphaneler şunlardır:

import (
	"encoding/json" // JSON formatındaki verilerle çalışmanızı sağlar. JSON serileştirme ve çözme işlemleri yapılır.
	"errors"        // Hata oluşturma ve hata kontrolü için kullanılır. Yeni hata mesajları oluşturabilirsiniz.
	"fmt"           // Formatlı giriş/çıkış işlemleri için kullanılır. Örneğin, ekrana yazdırma işlemleri yapılır.
	"math"          // Matematiksel hesaplamalar yapmanızı sağlar. Örneğin, karekök, trigonometrik hesaplamalar.
	"math/rand"     // Rastgele sayılar üretmek için kullanılır. Örneğin, 0 ile 100 arasında rastgele bir sayı oluşturmak.
	"net/http"      // HTTP istemci ve sunucu oluşturmak için kullanılır. Web üzerinden veri almak veya göndermek için.
	"os"            // Dosya sistemi ile ilgili işlemler ve işletim sistemiyle etkileşim sağlar. Örneğin, dosya açma, okuma, yazma işlemleri yapılır.
	"strings"       // String (metin) işlemleri yapmanıza olanak tanır. Örneğin, metinlerin büyük harfe dönüştürülmesi, bölünmesi.
	"time"          // Zaman ve tarih ile ilgili işlemler sağlar. Örneğin, geçerli zamanı almak, belirli bir süre beklemek.

	//  "io"   Girdi/çıktı işlemleri için temel fonksiyonlar sağlar. Örneğin, dosya okuma, yazma.
	"reflect" // Değişkenlerin türlerini ve değerlerini incelemek için kullanılır. Dinamik olarak tür bilgisi almak için.
)

// Global değişken
var globalVar = "Ben bir global değişkenim"

// Değişken tanımlama
var x int = 42 // Tür belirtilerek tanımlama
var y string = "You are a fool"
var z float32 = 3.14
var dogru bool = true

// Değişken tanımlama
var sayı = 42      // Tür otomatik olarak int olur
var hi = "Merhaba" // Tür otomatik olarak string olur

var a, b, c int = 1, 2, 3    // Aynı türde birden fazla değişken
var aa, bb = "Merhaba", 3.14 // Tür otomatik çıkarılır

// Varsayılan değerler, değişken tanımlanmadığında otomatik olarak atanır
// Varsayılan değerler: int -> 0, string -> "", bool -> false
var xx int    // Varsayılan değer 0
var yy string // Varsayılan değer ""
var zz bool   // Varsayılan değer false

// Değişken tanımlama
// Değeri değişmeyen sabitler const anahtar kelimesi ile tanımlanır:
const pi = 3.14
const greeting = "Merhaba"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, Universe!")
	fmt.Printf("Hello world secret is %d\n", x)                // fstring ile x değerini %d formatında yazdırıyoruz. %d integer değerini temsil eder.
	fmt.Printf("Hello world secret is %s\n", y)                // fstring ile y değerini %s formatında yazdırıyoruz. %s string değerini temsil eder.
	k := 42 + 3.14                                             // k değişkeni tanımlandı ve := sadece func içinde kullanılır.
	fmt.Printf("Hello world secret is %f\n", k)                // fstring ile k değerini %f formatında yazdırıyoruz. %f float değerini temsil eder.
	fmt.Printf("Hello world secret is %t\n", dogru)            // fstring ile dogru değerini %t formatında yazdırıyoruz. %t boolean değerini temsil eder.
	fmt.Printf("Hello world secret is %v\n", z)                // fstring ile z değerini %v formatında yazdırıyoruz. %v değişkenin türüne göre formatlanır.
	fmt.Printf("Hello world secret is %T\n and %s", x, y)      // fstring ile x ve y değerinin türünü sırayla yazdırıyoruz.
	fmt.Printf("Hello world secret is %T\n", yy)               // fstring ile yy değerinin türünü yazdırıyoruz.
	fmt.Printf("Hello world secret is %T\n", zz)               // fstring ile zz değerinin türünü yazdırıyoruz.
	fmt.Printf("Hello world secret is %T\n", aa)               // fstring ile aa değerinin türünü yazdırıyoruz.
	fmt.Printf("k değişkeninin tipi: %T\n", reflect.TypeOf(k)) // k değişkeninin türünü yazdırıyoruz.

	// Local değişken
	localVar := "Ben bir local değişkenim"

	fmt.Println(globalVar) // Global değişkene erişim
	fmt.Println(localVar)  // Local değişkene erişim

	// os örneği: Dosya açma
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Hata:", err)
	}
	defer file.Close()

	// math örneği: Kareköklü hesaplama
	result := math.Sqrt(16)
	fmt.Println("Karekök 16:", result)

	// strings örneği: String dönüşümü
	str := "hello"
	upperStr := strings.ToUpper(str)
	fmt.Println("Büyük harf:", upperStr)

	// time örneği: Geçerli zamanı al
	currentTime := time.Now()
	fmt.Println("Geçerli Zaman:", currentTime)

	// math/rand örneği: Rastgele sayı üretme
	randNum := rand.Intn(100)
	fmt.Println("Rastgele sayı:", randNum)

	// encoding/json örneği: JSON verisini serileştirme
	data := map[string]string{"name": "John", "age": "30"}
	jsonData, _ := json.Marshal(data)
	fmt.Println("JSON verisi:", string(jsonData))

	// net/http örneği: Web sayfasına istek gönderme
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("HTTP Hata:", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("HTTP Durum Kodu:", resp.StatusCode)
	}

	// errors örneği: Hata oluşturma
	errExample := errors.New("Bir hata oluştu")
	fmt.Println("Hata:", errExample)

	// io örneği: Dosya okuma (örnek)
	// (Burada basit bir dosya okuma işlemi yapılabilir, ancak daha fazla kod gerekir)
}
