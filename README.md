# 09.04.2022 DERS

## ÖZET

### Fonksiyonlar

- Parametre gönderimi  

- - a int
- - a, b int
- - a []int
- - a map[int]string
- - func fonksiyonIsmi() int{}
- - func fonksiyonIsmi() int, int{}
- - func fonksiyonIsmi() sonuc int{}

- - - Scope içerisinde sonucu tekrar tanımlamaya gerek yok
- - - return boş kalabilir
----
- ### variadic fonksiyonlar
- - Dizi gönderirken a... 
- - Parametre olarak alırken ...int  
- - Birden fazla parametre gönderilebilir
---
- ### Recursive fonksiyonlar
- - faktor(n int){  
    return (n-1)
}
---
- ### POİNTER / İŞARETLEYİCİ
- - &a => a değişkeninin referansını verir
- - *b => referans aldığı yerind eğerini verir
- - Fonksiyon parametre olarak alırsak func (a *int) şeklinde yazılır
---
- ANONİM FONKSİYONLAR
- - var a= func(){} şeklinde yazılır
- - func a() func() int{} şeklinde geri dönüşler de olabilir