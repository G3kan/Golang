package main

import "fmt"

// Go dilinde struct kullanımı
// Struct, birden fazla veriyi bir arada tutmak için kullanılan bir veri yapısıdır.
// Struct, bir nesnenin özelliklerini tanımlamak için kullanılır.
// Aşağıda bir struct tanımı ve kullanımı örneği verilmiştir.
// Struct tanımı
// Struct, bir veri yapısını temsil eder ve birden fazla alan içerebilir.
// Struct, bir nesnenin özelliklerini tanımlamak için kullanılır.

// Bir struct tanımlıyoruz
// Adddess bir struct ve biz bunu type gibi kullanabiliriz
type Address struct {
	City    string
	Country string
}

var evadresi Address = Address{
	City:    "İstanbul",
	Country: "Türkiye",
}

type Person struct {
	Name    string
	Age     int
	Address Address // İç içe struct
}

// Yeni bir kişi oluşturmak için bir fonksiyon
func NewPerson(name string, age int, city, country string) Person {
	return Person{
		Name: name,
		Age:  age,
		Address: Address{
			City:    city,
			Country: country,
		},
	}
}

// Kişinin yaşını artıran bir fonksiyon
func (p *Person) IncrementAge() {
	p.Age++
}
func (p *Person) UpdateAddress(city, country string) {
	p.Address.City = city
	p.Address.Country = country
}

// Kişinin bilgilerini yazdıran bir fonksiyon
func (p Person) PrintInfo() {
	fmt.Println("İsim:", p.Name)
	fmt.Println("Yaş:", p.Age)
	fmt.Println("Şehir:", p.Address.City)
	fmt.Println("Ülke:", p.Address.Country)
}

func main() {
	// Bir struct örneği oluşturuyoruz
	person := Person{
		Name: "Ahmet",
		Age:  30,
		Address: Address{
			City:    "İstanbul",
			Country: "Türkiye",
		},
	}

	// Struct alanlarına erişim
	fmt.Println("İsim:", person.Name)
	fmt.Println("Yaş:", person.Age)
	fmt.Println("Şehir:", person.Address.City)
	fmt.Println("Ülke:", person.Address.Country)

	// Alanları güncelleme
	person.Age = 31
	person.Address.City = "Ankara"
	fmt.Println("Güncellenmiş Yaş:", person.Age)
	fmt.Println("Güncellenmiş Şehir:", person.Address.City)

	person.PrintInfo()

	// Yaşı artırma
	person.IncrementAge()
	fmt.Println("\nYaş artirildi:")
	person.PrintInfo()

	// Adresi güncelleme
	person.UpdateAddress("Ankara", "Türkiye")
	fmt.Println("\nAdres güncellendi:")
	person.PrintInfo()

}
