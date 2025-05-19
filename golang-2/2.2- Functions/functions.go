package main

import "fmt"

// Fonksiyonlar, belirli bir işlevi yerine getiren kod bloklarıdır.
// go returnde değer belirtmek zorunda değiliz.
// eğer değer belirmezsek go isimlerdirilmiş değişkenleri döndürecektir.
// Fonksiyonlar, parametmetre, argüman ve dönüş olabilir.

// Fonksiyon tanımı şu şekildedir:
func functionName(parameterName parameterType) returnType {
	// Fonksiyonun gövdesi
	return value
}

// İki tam sayı alan ve bunların toplamını döndüren basit bir fonksiyon
func add(a int, b int) int {
	return a + b
}

// Argüman almayan ve bir string döndüren bir fonksiyon
func greet() string {
	return "Hello, World!"
}

// Bir string alan ve onu ekrana yazdıran bir fonksiyon
func printMessage(message string) {
	fmt.Println(message)
}

// Bir fonksiyon tanımı, ancak hiçbir şey yapmıyor
// Bu fonksiyonun dönüş tipi yok
// ve hiçbir argüman almıyor
func doNothing() {
	fmt.Println("Hi, I do nothing!")
}

func execute(f func(string)) {
	f("Merhaba, anonim fonksiyon!")
}

// Bir fonksiyon döndüren fonksiyon
func multiplier(factor int) func(int) int {
	return func(value int) int {
		return factor * value
	}
}

// var ile bir fonksiyon tanımlama
var subtract func(a int, b int) int

// Fonksiyonu bir anonim fonksiyon ile başlatma
subtract = func(a int, b int) int {
	return a - b
}

// Argüman olarak bir fonksiyon alan ve onu çağıran bir fonksiyon
func callFunction(f func(int, int) int, x int, y int) int {
	return f(x, y)
}

// Bir fonksiyon tanımlayan fonksiyon
func createAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Closure örneği
// Closure, bir fonksiyonun dışındaki değişkenlere erişim sağlayan bir fonksiyondur.
// Closure, fonksiyonun tanımlandığı yerin kapsamını (scope) hatırlar.
func closureExample() {
	// Dış değişken
	counter := 0

	// Closure tanımı
	increment := func() int {
		counter++
		return counter
	}

	// Closure'ı çağırma
	fmt.Println("Counter:", increment()) // 1
	fmt.Println("Counter:", increment()) // 2
	fmt.Println("Counter:", increment()) // 3
}

// Değişken sayıda argüman alan bir fonksiyon
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// defer kullanımı
// defer, bir fonksiyonun sonunda çalıştırılacak ifadeleri tanımlamak için kullanılır.
// Genellikle kaynakları serbest bırakmak veya temizleme işlemleri yapmak için kullanılır.
func deferExample() {
	fmt.Println("Defer örneği başlıyor.")

	// defer ile bir ifade tanımlama
	defer fmt.Println("Bu ifade en son çalıştırılacak.")

	fmt.Println("Defer örneği devam ediyor.")

	// Birden fazla defer kullanımı
	defer fmt.Println("Bu ifade ikinci sırada çalıştırılacak.")
	defer fmt.Println("Bu ifade üçüncü sırada çalıştırılacak.")

	fmt.Println("Defer örneği bitiyor.")
}



// Function as a parameter and return example
func functionAsParameterAndReturn() {
	// Bir fonksiyonu parametre olarak alan ve döndüren fonksiyon
	execute := func(f func(int) int, value int) func(int) int {
		fmt.Println("Fonksiyon çağrıldı, değer:", f(value))
		return f
	}

	// Bir fonksiyon tanımlama
	double := func(x int) int {
		return x * 2
	}

	// execute fonksiyonunu çağırma
	newFunc := execute(double, 5)

	// Döndürülen fonksiyonu çağırma
	fmt.Println("Yeni fonksiyon sonucu:", newFunc(10)) // 20
}

func main() {
	// add fonksiyonunu çağırma
	sum := add(5, 3)
	fmt.Println("Toplam:", sum)

	// greet fonksiyonunu çağırma
	greeting := greet()
	fmt.Println(greeting)

	// printMessage fonksiyonunu çağırma
	printMessage("Bu özel bir mesajdır.")

	//anonim fonksiyonlar : bir isme sahip olmayan ve
	// genellikle bir değişkene atanarak kullanılan fonksiyonlardır.
	add := func(a int, b int) int {
		return a + b
		fmt.Println("Toplam:", add(3, 5))
	}

	// multiplier fonksiyonunu çağırarak bir fonksiyon döndür
	timesTwo := multiplier(2)
	timesThree := multiplier(3)

	// Döndürülen fonksiyonları kullan
	fmt.Println("2'nin 5 katı:", timesTwo(5))  // 10
	fmt.Println("3'ün 4 katı:", timesThree(4)) // 12

	// var ile bir fonksiyon tanımlama
	var multiply func(a int, b int) int

	// Fonksiyonu bir anonim fonksiyon ile başlatma
	multiply = func(a int, b int) int {
		return a * b
	}

	// Fonksiyonu çağırma
	result := multiply(4, 5)
	fmt.Println("Çarpım:", result)

}

/*
-İşlev İsimlendirme
İşlevlerin anlaşılır ve açıklayıcı isimlerle adlandırılması, kodun daha kolay anlaşılmasını sağlar.
Parametrelerin de iyi bir şekilde adlandırılması, işlevin ne yaptığını anlamayı kolaylaştırır.

-Fonksiyonel Kohezyon
Her işlevin yalnızca bir işlem gerçekleştirmesi gerektiği vurgulanır; bu, kodun daha anlaşılır olmasını sağlar.
Farklı işlemleri aynı işlevde birleştirmek, kodun karmaşıklaşmasına neden olabilir.

-Parametre Sayısını Azaltma
İşlevlerin aldığı parametre sayısının sınırlandırılması, hata ayıklamayı kolaylaştırır.
İlgili argümanların yapılandırmalar içinde gruplanması, parametre sayısını azaltmanın bir yoludur.

-Fonksiyon Karmaşıklığı
Fonksiyonların karmaşıklığı, genellikle fonksiyon uzunluğu ile ölçülür; kısa fonksiyonlar genellikle daha basit kabul edilir.
Fonksiyonları basit tutmak için, karmaşık kod parçalarını daha küçük, yönetilebilir fonksiyonlara ayırmak önemlidir.

-Kontrol Akışı Karmaşıklığı
Kontrol akışı, bir fonksiyonun başlangıcından sonuna kadar olan yolları ifade eder; if ifadeleri ve döngüler gibi yapılar, kontrol akışını karmaşıklaştırabilir.
Fonksiyonlardaki kontrol akışı yollarını azaltmak, kodun hata ayıklanmasını kolaylaştırır.

-Fonksiyon Hiyerarşisi
Fonksiyonlar arasında bir hiyerarşi oluşturarak, karmaşıklığı yönetmek mümkündür; bir fonksiyon diğer fonksiyonları çağırarak daha basit hale getirilebilir.
Fonksiyonları mantıklı bir şekilde gruplamak, her birinin karmaşıklığını azaltarak genel kod kalitesini artırır.
*/
