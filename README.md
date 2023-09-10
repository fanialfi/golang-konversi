# Golang Konversi

## Konversi menggunakan `strconv`

package `strconv` berisi banyak function yang sangat berguna untuk melakukan konversi ke string dan dari string merepresentasikan basis data type seperti _numeric_, dan _string_.

### function `strconv.Atoi(s string)`

function ini dilakukan untuk konversi data dari _string_ ke _int_, function `strconv.Atoi()` menghasilkan 2 nilai kembalian, yaitu hasil konversi dan error (jika konversi sukses maka error berisi _nil_).

```go
package main

import (
  "fmt"
  "strconv"
)

func main(){
  str := "1234"

	// function untuk mengubah tipe data dari string ke int
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v %T\n", num, num)
}
```

### function `strconv.Itoa(i int)`

function ini merupakan kebalikan dari `strconv.Atoi()` berguna untuk konversi dari _int_ ke _string_, nilai kembaliannya hanya 1 yaitu hasil konversi nya itu sendiri.s

```go
package main

import (
  "fmt"
  "strconv"
)

func main(){
  // function untuk mengubah tipe data dari int ke string
	num = 1234
	str = strconv.Itoa(num)
	fmt.Printf("%v %T\n", str, str)
}
```

### function `strconv.Parseint(s string, base int, bitSize int)`

function ini digunakan untuk konversi dari _string_ numeric dengan basis tertentu (bisa biner, hexadecimal, dll) ke tipe numeric non-desimal (int64).

_string_ pada parameter pertama dapat dimulai dengan menambahkan "+" atau "-", merepresentasikan bahwa itu bilangan positif atau negatif.

_base_ pada parameter kedua menandakan numeric basis apa yang digunakan pada parameter pertama, jika _base_ argument berisi 0, maka akan mengikuti basis numeric yang digunakan oleh parameter pertama jika parameter pertama dimulai dengan prefix (jika diawali prefix) : 2 untuk "0b", 8 untuk "0" atau "0o", 16 untuk "0x" dan selain itu menggunakan basis 10.

argument _bitSize_ menandakan bahwa hasil dari konversi tersebut harus fit. 
bitSize 0, 8, 16, 32, dan 64 adalah alias dari _int_, _int8_, _int16_, _int32_, dan _int64_. 
jika bitSize dibawah 0 dan di atas 64 maka akan mengembalikan error.

pada contoh berikut string numeric yang berbasis 16 (hexadecimal) di konversi ke _int64_ dimana lebar data-nya harus di antara _int32_ (lihat pada parameter ke-tiga)

```go
package main

import (
  "fmt"
  "strconv"
)
func main(){
  str := "ff"
	num, err := strconv.ParseInt(str, 16, 32)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v %T\n", num, num) // 255 int64
	}
}
```

contoh berikutnya string berbasis biner (basis 2) di konversi dimana lebar data-nya harus di antara _int8_ :

```go
package main

import (
  "fmt"
  "strconv"
)

func main(){
	str := "0b1010"
	num, err := strconv.ParseInt(str, 0, 8)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v %T\n", num, num) // 10 int64
	}
}
```

### function `srtconv.ParseFloat(s string, bitSize int)`

digunakan untuk konversi dari string ke float64 dengan lebar data (precision) ditentukan oleh bitSize (32 untuk float32 or 64 untuk float64), ketika bizSize = 32 result akan tetap bertipe _float64_, namun ketika di konversi ke _float32_ tidak akan mengubah value.

```go
package main

import (
	"fmt"
	"strconv"
)

func main(){
	str := "24.12"
	num, err := strconv.ParseFloat(str, 32)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%f %T\n", num, num)

	check := float32(num)
	fmt.Printf("%f %T\n", check, check)
}
```

### function `strconv.ParseBool(s string)`

digunakan untuk konversi dari string ke boolean, `ParseBool()` memperbolehkan value 1,t,T,TRUE,True,true,0,f,F,False,false sebagai value dari string, selain value tersebut akan mengembalikan error.

```go
package main

import (
	"fmt"
	"strconv"
)

func main(){
	str = "F"
	if boolean, err := strconv.ParseBool(str); err == nil {
		fmt.Printf("%T %v\n", boolean, boolean)
	} else {
		fmt.Println(err.Error())
	}
}
```

## Konversi data menggunakan teknik casting

keyword type data bisa digunakan untuk casting (konversi antar type data), cara penggunaannya dengan menuliskan tipe data tujuan sebagai function, lalu menyisipkan data yang akan adi konversi tersebut sebagai parameter.

### Casting byte <--> byte

string sebenarnya adalah slice/array `byte`. 
di Golang sebuah karakter biasa (non unicode) di representasikan sebagai element slice byte, tiap element dari slice (tiap karakter dari string) berisi data int berbasis desimal yang merepresentasikan kode ASCII.

```go
package main

import (
	"fmt"
)

func main(){
	text := "😊😀😃😄😁😅😆😂🙃😉😊😇😎😍 hello world"

	a := []byte(text)
	a = append(a, 252, 169)
	for index, elm := range a {
		fmt.Printf("index ke-%d\t%d\t%s\n", index, elm, string(elm))
	}
	fmt.Println(string(a))

	fmt.Println()

	// tipe data rune bisa mengatas karakter ASCII non standar (seperti karakter emoji)
	b := []rune(text)
	b = append(b, 252, 169)
	for index, elm := range b {
		fmt.Printf("index ke-%d\t%d\t%s\n", index, elm, string(elm))
	}
	fmt.Println(string(b))
}
```

### Type Assertions pada interface kosong (`interface{}` / `any`)

type assertions pada interface kosong adalah cara untuk mengambil atau mengkonversi nilai yang di simpan dalam interface kosong ke tipe data aslinya.

```go
package main

import "fmt"

func main(){
	data := map[string]any{
		"nama":   "Fani Alfirdaus",
		"grade":  2,
		"height": 160,
		"isMale": true,
		"hobies": []string{"eating", "sleeping"},
	}
	fmt.Printf("%T\t%v\n", data["nama"].(string), data["nama"].(string))
	fmt.Printf("%T\t%v\n", data["grade"].(int), data["grade"].(int))
	fmt.Printf("%T\t%v\n", data["height"].(int), data["height"].(int))
	fmt.Printf("%T\t%v\n", data["isMale"].(bool), data["isMale"].(bool))
	fmt.Printf("%T\t%v\n", data["hobies"].([]string), data["hobies"].([]string))
	fmt.Println()
}
```

statement `data["nama"].(string)` maksudnya adalah nilai dari `data["nama"].(string)` yang bertipe `interface{}` diambil nilai aslinya dimana nilai aslinya adalah string.

Jika melakukan assertions ke tipe data yang tidak sesuai, maka akan mentrigger panic error. 
Saat melakukan assertions pada interface kosong (`interface{}` / `any`) lalu tidak tahu tipe data aslinya, maka bisa menggunakan teknik di bawah ini.

tipe data asli pada variabel `interface{}` / `any` bisa diketahui dengan cara mencasting ke tipe `type`, namun casting ini hanya bisa dilakukan di `switch`.

```go
package main

import "fmt"

func main(){
	data := map[string]any{
		"nama":   "Fani Alfirdaus",
		"grade":  2,
		"height": 160,
		"isMale": true,
		"hobies": []string{"eating", "sleeping"},
	}

	for _, value := range data {
		switch value.(type) {
		case string:
			fmt.Printf("%T\t%s\n", value.(string), value.(string))
		case int:
			fmt.Printf("%T\t%d\n", value.(int), value.(int))
		case bool:
			fmt.Printf("%T\t%t\n", value.(bool), value.(bool))
		case []string:
			fmt.Printf("%T\t%v\n", value.([]string), value.([]string))
		default:
			fmt.Println("value", value)
		}
	}
}
```