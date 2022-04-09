package main

import "fmt"

/*
func main() {
	// Variadic function

	ortalama(2, 8, 36, 6)

}
func ortalama(ort ...int) {
	var a float32
	for _, k := range ort {
		a += (float32(k) / float32(len(ort)))
	}

	fmt.Println(a)
}
*/
/*
//                  RECURSİVE FUNC
func main() {
	// RECURSİVE FONKSİYONLAR
	//Özyinelemli
	//kendi kendini çağırır

	//Faktöriyel gibi şeylerde kullanılabilir

	fmt.Println(faktor(5))

}
func faktor(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktor(n-1)
}
*/

//                     POİNTER

/*func main() {

	a := 5
	b := &a
	*b = 45
	fmt.Println(a, *b)
// b nin değerinde a nın referansı bulunmaktadır.
//* b ile a değiştirilebilir.
//b yıldız kullanılmadan ancak konumunu geri dönderir

}*/

//              FONKSİYON VE POİNTER
//  & ile referans adresi gönderildiğinde * ile cevaplandığında adreteki sayı değişir.
/*
func main() {
	k := 22
	yeap(k)
	fmt.Println(k)
	hoop(&k)//referans adresi gider
	fmt.Println(k)
}
func hoop(a *int) {//referans adresi gelir. içerisi değiştirilir.
	*a = 55
}
func yeap(a int) {
	a = 55
}
*/

//                       ANONİM FUNC
/*
func main() {

	var a = func(m1, m2 int) int {
		toplam := m1 + m2
		return toplam
	}
	toplam := a(12, 23)
	fmt.Println(toplam)
}
*/
func a() func() int {
	k := 10
	return func() int {
		k++
		return k
	}
}
func main() {
	m := a()
	fmt.Println(m)
}
