package main

import "fmt"

//                       MULTİPLE RETURN

/*
func main() {
	//FUNCTİONS

	//fonksiyonların faydası yeniden kullanılabilir olması
	//kod okunurluğunu arttırır
	//_, _, _, a := islem(12, 34)
	//fmt.Println(a)

}

/*
func islem(a int, b int) (int, int, int, int) {
	topla := a + b
	cikar := a - b
	carp := a * b
	bolme := a / b
	return topla, cikar, carp, bolme // multiple return
}*/

/*           //DİĞER TANIMLAMA YÖNTEMİ
func islem(a, b int) (topla, cikar, carp, bolme int) {
	topla = a + b
	cikar = a - b
	carp = a * b
	bolme = a / b
	return //yanına birşey yazmadık dönecek değerleri fonksiyonda tanımladık
}*/

//               FONKSİYONA DİZİ GÖNDERMEK
/*
func main() {
	a := []int{12, 36, 92}
	b := []int{12, 56, 43}

	fmt.Println(diziToplm(a,b))


}
func diziToplm(a,b []int) int {
	toplam := 0

	for _, v := range a {
		toplam += v
	}
	return toplam

}
*/

//                        FONSİYONA MAP GÖNDERME

func main() {
	a := map[string]int{
		"ali":  15,
		"rano": 63,
		"taha": 15}
	fmt.Println(a)
	fmt.Println(ortalama(a))
}

/*func ortalama(b map[string]int) int {

toplam := 0
for _, v := range b {
	toplam += v
}

return toplam / len(b)
*/

// diğer yöntem

func ortalama(b map[string]int) (ortalama float32) {
	for _, v := range b {
		ortalama += (float32(v) / float32(len(b)))
	}
	return
}
