package main

import (
	"bufio"         // bufio I/O için kullanılır
	"encoding/json" // json paketi ile JSON verilerini işleyebiliriz
	"fmt"
	"io"       // io paketi ile dosya okuma ve yazma işlemleri yapabiliriz
	"net"      // net paketini kullanarak TCP/IP ve UDP protokollerini kullanabiliriz
	"net/http" // http paketi ile HTTP sunucusu oluşturabiliriz
	"os"       // os paketi ile işletim sistemi ile etkileşimde bulunabiliriz
)

// RFC (Request for Comments):
// RFC, internet protokolleri, sistemler ve diğer teknik standartlar hakkında resmi belgeler serisidir.
// İnternet standartlarının geliştirilmesi ve belgelenmesi için kullanılır.
// Örneğin, HTTP, TCP/IP, ve DNS gibi protokoller için RFC belgeleri bulunmaktadır.
// Her RFC, belirli bir standardı veya öneriyi açıklar ve genellikle bir numara ile tanımlanır (örneğin, RFC 2616 - HTTP/1.1).

// Protokoller:
// Protokoller, bilgisayar ağları üzerinde veri iletimi için belirli kurallar ve standartlardır.
// Protokoller, veri iletimini düzenler, hata kontrolü yapar ve veri paketlerinin sırasını korur.
// Protokoller, istemci ve sunucu arasındaki iletişimi sağlar.

// IP (Internet Protocol):
// IP, internet üzerindeki cihazlar arasında veri paketlerinin iletilmesini sağlayan bir protokoldür.
// Her cihaz için benzersiz bir adres (IP adresi) kullanarak verilerin doğru hedefe ulaşmasını sağlar.
// IPv4 (32-bit) ve IPv6 (128-bit) olmak üzere iki temel sürümü vardır.

// TCP (Transmission Control Protocol):
// TCP, IP ile birlikte çalışan ve verilerin güvenilir bir şekilde iletilmesini sağlayan bir protokoldür.
// Bağlantı tabanlıdır (connection-oriented) ve veri iletiminde hata kontrolü yapar.
// Paketlerin sırasını korur ve eksik veya bozuk paketleri yeniden gönderir.
// Örnek kullanım alanları: HTTP/HTTPS, e-posta (SMTP, IMAP), dosya transferi (FTP).

// UDP (User Datagram Protocol):
// UDP, TCP'ye alternatif olarak kullanılan, daha hızlı ancak güvenilirlik sağlamayan bir protokoldür.
// Bağlantısızdır (connectionless) ve hata kontrolü yapmaz.
// Paketlerin sırasını garanti etmez, ancak daha az gecikme sağlar.
// Örnek kullanım alanları: Video/oyun akışı, VoIP, DNS sorguları.

// HTTP (Hypertext Transfer Protocol):
// HTTP, web tarayıcıları ve sunucuları arasında veri iletimi için kullanılan bir protokoldür.
// İstemci-sunucu mimarisine dayanır ve genellikle metin tabanlıdır.

//// HTTPS (HTTP Secure):
// HTTPS, HTTP'nin güvenli bir versiyonudur ve SSL/TLS ile şifreleme sağlar.
// HTTPS, veri iletimini güvenli hale getirir ve kimlik doğrulama sağlar.
// HTTPS, web siteleri için güvenli bir bağlantı sağlar ve kullanıcı bilgilerini korur.
// net paketi örneği: Basit bir TCP sunucusu

// net paketi örneği: TCP istemcisi
// net paketi, TCP/IP ve UDP protokollerini kullanarak ağ iletişimi sağlar.
func netDialExample() {
	conn, err := net.Dial("tcp", "localhost:8080") // TCP bağlantısı başlat
	if err != nil {
		fmt.Println("Bağlantı hatası:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintln(conn, "Merhaba, TCP sunucusu!") // Sunucuya mesaj gönder

	// Sunucudan gelen yanıtı oku
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("Okuma hatası:", err)
		return
	}
	fmt.Println("Sunucudan gelen yanıt:", string(response[:n]))
}

// net paketi örneği: Basit bir TCP sunucusu
// net paketi, TCP/IP ve UDP protokollerini kullanarak ağ iletişimi sağlar.
func netExample() {
	listener, err := net.Listen("tcp", ":8080") // TCP dinleyici oluşturma iki çıktı verir listener bağlantısı ve err hata
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP sunucusu çalışıyor...")
	for {
		conn, err := listener.Accept() // bağlantı kabul etme iki çıktı verir conn bağlantısı ve err hata
		if err != nil {
			fmt.Println("Bağlantı hatası:", err)
			continue
		}
		go func(c net.Conn) { // yeni bir goroutine başlatma
			defer c.Close() // bağlantıyı kapatma
			fmt.Fprintln(c, "Merhaba, TCP istemcisi!")
			c.Close()
		}(conn)
	}
}

// http paketi örneği: Basit bir HTTP sunucusu
func httpGetExample() {
	// URL'ye GET isteği gönder
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("HTTP GET isteği hatası:", err)
		return
	}
	defer resp.Body.Close() // Yanıt gövdesini kapatmayı unutmayın

	// HTTP durum kodunu kontrol et
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Hata: HTTP durum kodu %d\n", resp.StatusCode)
		return
	}

	// Yanıt gövdesini okuyup ekrana yazdır
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Yanıt okuma hatası:", err)
		return
	}
	fmt.Println("Sunucudan gelen yanıt:")
	fmt.Println(string(body))
}

// JSON (JavaScript Object Notation):
// JSON, veri değişimi için kullanılan hafif ve insan tarafından okunabilir bir veri formatıdır.
// Anahtar-değer çiftlerinden oluşur ve genellikle web uygulamaları arasında veri iletimi için kullanılır.
// JSON şu veri türlerini destekler: string, number, boolean, array, object ve null.
// JSON, dil bağımsızdır ve birçok programlama dili tarafından desteklenir.
// JSON, GO'da maplare benzer bir yapıya sahiptir.
// JSON, veri yapısını temsil etmek için anahtar-değer çiftlerini kullanır.
// JSON, veri iletiminde yaygın olarak kullanılır ve RESTful API'lerde sıkça tercih edilir.
// Örnek bir JSON verisi:
// {
//   "name": "Ali",
//   "age": 25,
//   "isStudent": true,
//   "courses": ["Math", "Physics", "Chemistry"],
//   "address": {
//     "city": "Istanbul",
//     "country": "Turkey"
//   }
// }

// encoding/json paketi örneği: JSON serileştirme ve seriden çıkarma
func jsonExample() {
	type Person struct {
		Name string `json:"name"` // JSON anahtar adı "name" ile eşleşir
		Age  int    `json:"age"`  // JSON anahtar adı "age" ile eşleşir
	}

	// JSON serileştirme (struct -> JSON)
	// JSON serileştirme, bir struct'ı JSON formatında bir byte dizisine dönüştürür.
	person := Person{Name: "Ali", Age: 25}
	jsonData, err := json.Marshal(person) // struct'ı JSON'a dönüştür
	if err != nil {
		fmt.Println("JSON serileştirme hatası:", err)
		return
	}
	fmt.Println("JSON:", string(jsonData)) // JSON verisini string olarak yazdır

	// JSON seriden çıkarma (JSON -> struct)
	// JSON seriden çıkarma, bir JSON byte dizisini struct'a dönüştürür.
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson) // JSON'u struct'a dönüştür
	if err != nil {
		fmt.Println("JSON seriden çıkarma hatası:", err)
		return
	}
	fmt.Println("Kişi:", newPerson) // struct olarak yazdır
}

// bufio paketi, tamponlu girdi ve çıktı işlemleri için kullanılır.
// Örneğin, dosyaları satır satır okumak veya yazmak için kullanılır.

func bufioExample() {
	file, err := os.Open("example.txt") // open iki çıktı verir file dosyası ve err hata
	if err != nil {                     // dosya açılamadıysa hata veriri err nil dönerse hata yoktur.
		fmt.Println("Hata:", err)
		return
	}
	defer file.Close() // dosya kapatılır
	// bufio.Scanner kullanarak dosyayı satır satır okuma

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Okuma hatası:", err)
	}
}

// os paketi temel fonksiyonları:
// 1. os.Create(filename string) (*os.File, error):
//    - Yeni bir dosya oluşturur. Eğer dosya zaten varsa, üzerine yazar.
//    - Dönen değerler:
//      - *os.File: Dosya nesnesi.
//      - error: Hata oluşursa döner.
//    - Örnek: file, err := os.Create("output.txt")

// 2. os.Open(filename string) (*os.File, error):
//    - Var olan bir dosyayı okuma modunda açar.
//    - Dönen değerler:
//      - *os.File: Dosya nesnesi.
//      - error: Hata oluşursa döner.
//    - Örnek: file, err := os.Open("example.txt")

// 3. os.Remove(filename string) error:
//    - Belirtilen dosyayı siler.
//    - Dönen değer:
//      - error: Hata oluşursa döner.
//    - Örnek: err := os.Remove("output.txt")

// 4. os.Rename(oldpath string, newpath string) error:
//    - Bir dosyanın adını değiştirir veya başka bir konuma taşır.
//    - Dönen değer:
//      - error: Hata oluşursa döner.
//    - Örnek: err := os.Rename("old.txt", "new.txt")

// 5. os.Mkdir(name string, perm os.FileMode) error:
//    - Yeni bir dizin (klasör) oluşturur.
//    - Dönen değer:
//      - error: Hata oluşursa döner.
//    - Örnek: err := os.Mkdir("new_folder", 0755)

// 6. os.ReadFile(filename string) ([]byte, error):
//    - Bir dosyanın içeriğini okur ve bir byte dizisi olarak döner.
//    - Dönen değerler:
//      - []byte: Dosya içeriği.
//      - error: Hata oluşursa döner.
//    - Örnek: data, err := os.ReadFile("example.txt")

// 7. os.WriteFile(filename string, data []byte, perm os.FileMode) error:
//    - Bir dosyaya veri yazar. Eğer dosya yoksa oluşturur.
//    - Dönen değer:
//      - error: Hata oluşursa döner.
//    - Örnek: err := os.WriteFile("output.txt", []byte("Merhaba, Dünya!"), 0644)
//    - 3 parametre alır:
//      - filename: Dosya adı.
//      - data: Yazılacak veri.
//      - perm: Dosya izinleri.
//      - perm: 0644, dosya izinlerini belirtir (okuma ve yazma izinleri).

// Dosya İzinleri (Permissions):
// - Dosya izinleri, Unix/Linux sistemlerinde dosya ve dizinlere erişim kontrolünü sağlar.
// - İzinler üç gruba ayrılır: Sahip (Owner), Grup (Group), Diğerleri (Others).
// - İzinler şu şekilde ifade edilir:
//   - İlk rakam: Özel bayraklar (genellikle `0` olarak bırakılır).
//   - İkinci rakam: Dosya sahibinin (Owner) izinleri.
//   - Üçüncü rakam: Grubun (Group) izinleri.
//   - Dördüncü rakam: Diğer kullanıcıların (Others) izinleri.
// - İzinler şu değerlerle ifade edilir:
//   - `4`: Okuma izni (read)
//   - `2`: Yazma izni (write)
//   - `1`: Çalıştırma izni (execute)
// - İzinler toplama yöntemiyle belirtilir:
//   - `7`: Okuma, yazma ve çalıştırma (4 + 2 + 1)
//   - `6`: Okuma ve yazma (4 + 2)
//   - `5`: Okuma ve çalıştırma (4 + 1)
//   - `4`: Sadece okuma
// - Örnekler:
//   - `0644`:
//       - İlk rakam `0`: Özel bayraklar (kullanılmıyor).
//       - İkinci rakam `6`: Sahip için okuma ve yazma izni (4 + 2).
//       - Üçüncü rakam `4`: Grup için sadece okuma izni.
//       - Dördüncü rakam `4`: Diğer kullanıcılar için sadece okuma izni.
//   - `0755`:
//       - İlk rakam `0`: Özel bayraklar (kullanılmıyor).
//       - İkinci rakam `7`: Sahip için tüm izinler (okuma, yazma, çalıştırma).
//       - Üçüncü rakam `5`: Grup için okuma ve çalıştırma izni.
//       - Dördüncü rakam `5`: Diğer kullanıcılar için okuma ve çalıştırma izni.
//   - `0700`:
//       - İlk rakam `0`: Özel bayraklar (kullanılmıyor).
//       - İkinci rakam `7`: Sahip için tüm izinler (okuma, yazma, çalıştırma).
//       - Üçüncü rakam `0`: Grup için hiçbir izin yok.
//       - Dördüncü rakam `0`: Diğer kullanıcılar için hiçbir izin yok.

// 8. os.Stat(name string) (os.FileInfo, error):
//    - Bir dosya veya dizin hakkında bilgi alır.
//    - Dönen değerler:
//      - os.FileInfo: Dosya bilgisi.
//      - error: Hata oluşursa döner.
//    - Örnek: info, err := os.Stat("example.txt")

// 9. os.Exit(code int):
//    - Programı belirtilen çıkış koduyla sonlandırır.
//    - Örnek: os.Exit(1)

// os paketi örneği: Bir dosyayı byte olarak okuma
func osReadFileExample() {
	// Dosyayı okuma
	data, err := os.ReadFile("example.txt") // Dosyanın tüm içeriğini byte dizisi olarak okur
	if err != nil {
		fmt.Println("Dosya okuma hatası:", err)
		return
	}

	// Byte verisini string'e dönüştürerek ekrana yazdır
	fmt.Println("Dosya içeriği:")
	fmt.Println(string(data))
}

// os paketi örneği: Bir dosyayı parça parça okuma
func osReadChunksExample() {
	// Dosyayı aç
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Dosya açma hatası:", err)
		return
	}
	defer file.Close() // İşlem tamamlandıktan sonra dosyayı kapat

	// 1024 byte'lık bir buffer oluştur
	buffer := make([]byte, 1024)

	fmt.Println("Dosya içeriği:")
	for {
		// Dosyadan buffer boyutunda veri oku
		n, err := file.Read(buffer)
		if err != nil {
			// Eğer hata EOF ise (dosya sonu), döngüden çık
			if err == io.EOF {
				break
			}
			fmt.Println("Okuma hatası:", err)
			return
		}

		// Okunan veriyi ekrana yazdır
		fmt.Print(string(buffer[:n]))
	}
}

// os paketi örneği: Bir dosya oluşturma ve yazma
func osExample() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Merhaba, Dünya!")
	if err != nil {
		fmt.Println("Yazma hatası:", err)
	}
}

// io paketi örneği: Bir dosyayı başka bir dosyaya kopyalama
func ioExample() {
	srcFile, err := os.Open("source.txt")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Kopyalama hatası:", err)
	}
}

func main() {
	// net paketi örneği: TCP istemcisi
	netDialExample()

	// net paketi örneği: TCP sunucusu
	// netExample()

	// http paketi örneği: HTTP GET isteği
	httpGetExample()

	// encoding/json paketi örneği: JSON serileştirme ve seriden çıkarma
	jsonExample()

	// bufio paketi örneği: Dosyayı satır satır okuma
	bufioExample()

	// os paketi örneği: Dosyayı byte olarak okuma
	osReadFileExample()

	// os paketi örneği: Dosyayı parça parça okuma
	osReadChunksExample()

	// os paketi örneği: Dosya oluşturma ve yazma
	osExample()

	// io paketi örneği: Dosyayı başka bir dosyaya kopyalama
	ioExample()
}
