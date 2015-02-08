# Programming in GO

##BAB 1. PENDAHULUAN

Go adalah bahasa pemrogaman yang dirancang oleh Google dan mudah untuk dipelajari oleh anda yang baru pertama kali belajar bahasa pemrogaman. Selain itu, Go adalah bahasa pemrogaman yang mudah untuk membangun perangkat lunak sederhana, handal dan efisien. Keunggulan dari Go ini adalah memiliki fitur yang canggih dan ketersediaan yang luas pada berbagai *platform*.

Tahapan yang digunakan dalam membuat aplikasi dengan Go cukup mudah, yaitu :

1.	Mengumpulkan persyaratan
2.	Mencari solusi
3.	Menuliskan *source code* yang mengimplementasikan solusi
4.	Mengompilasi *source code*
5.	Menjalankan program untuk memastikan apakah program tersebut bekerja

###1. File dan Folder

File adalah kumpulan data yang disimpan sebagai unit dengan nama tertentu. Semua jenis file disimpan dengan cara yang sama pada komputer, memiliki nama, ukuran dan jenis yang terkait. Biasanya jenis file ditandai dengan ekstensi file setelah nama pada bagian terakhir. Contoh : 

    contoh.txt
  
* contoh : nama file
* txt : ekstensi yang digunakan untuk mewakili data tekstual

Folder atau direktori digunakan untuk sekumpulan file secara bersama-sama. Dapat juga berisi folder lainnya. Contoh :

    D:\apps\golang\contoh.txt

* contoh.txt adalah nama file yang terkandung didalam folder ‘golang’, folder ‘golang tersebut terkandung didalam folder ‘apps’ yang disimpan di drive D

Pada dasarnya dalam pengaturan file dan folder yang terkait dengan pembuatan aplikasi menggunakan bahasa pemrograman Go berhubungan dengan proses preparasi yang terdiri dari beberapa tahapan, yaitu :

1.	Penginstallan Go
2.	Pengaturan properties pada komputer
3.	Penginstallan Git
4.	Penginstallan Mercurial (hg)
5.	Penginstallan Gorilla

###2. Terminal

Sebagian besar interaksi kita dengan komputer dilakukan melalui interaksi antarmuka sebagai pengguna grafis (GUI). Digunakan keyboard, mouse dan layar sentuh untuk berinteraksi dengan tombol visual atau jenis lainnya dari kontrol yang ditampilkan pada layar. Sebelum GUI berkembang, terlebih dahulu digunakan terminal yaitu sebuah tekstual antarmuka  sederhana untuk komputer yang digunakan dengan memberikan perintah dan menerima balasan untuk memanipulasi tombol pada layar komputer. 

Walaupun sebagian besar dunia komputasi telah meninggalkan terminal, kebenarannya adalah terminal masih fundamental yang digynakan oleh sebagian besar bahasa pemrogaman. Bahasa pemrogaman Go tidak berbeda, sehingga sebelum kita menulis program pada Go, terlebih dahulu kita perlu memiliki pemahaman dasar bagaimana terminal bekerja.  

###3. *Text Editor*

Sebagai programmer, alat utama yang digunakan untuk menulis program adalah Text Editor. Text Editor mirip dengan pengolahan kata seperti Microsoft Word. Text Editor yang biasa digunakan dalam bahasa Go adalah Sublime Text. 

###4. *Tools*

Go adalah bahasa pemrogaman yang dikompilasi, artinya *source code* diterjemahkan ke dalam bahasa yang dapat dimengerti oleh komputer. Terkait akan hal tersebut, sebelum program Go dibuat, kita membutuhkan Go compiler. Installer akan mengatur Go untuk anda secara otomatis. Informasi lebih lanjut dapat ditemukan di http://www.golang.org


##BAB 2. PROGRAM PERTAMA

Didalam berbagai bahasa pemrogaman, biasanya anda diminta untuk membuat program sederhana “Hello World”. Yaitu sebuah program yang output nya hanya menampilkan “Hello World” pada terminal anda.

Untuk menjadi awalan dalam mempelajari Go, mari kita menulis program menggunakan Go. Hal pertama yang dilakukan adalah membuat folder baru untuk menyimpan program yang akan ditulis. Untuk mempermudah, pada teriminal anda berilah perintah dengan memasukkan : 

    mkdir apps/src/golang

Artinya, anda memerintahkan untuk membuat direktori baru (mkdir) misal pada drive D, dengan folder ‘apps’, didalamnya terdapat pula folder ‘src’, didalam folder src dibuat juga folder ‘golang’.

Pada text editor anda, buatlah file baru didalam folder golang, kemudian simpan file tersebut dengan nama main.go, lalu buatlah program berikut :

```go
package main
    
import "fmt"
// this is a comment
    
func main() {
  fmt.Println("Hello World")
}
```

Kemudian jalankan program tersebut pada terminal anda, tapi arahkan terlebih dahulu terminal anda ke direktori file anda disimpan.
Untuk menjalankan program, berikan perintah berikut pada terminal anda :

    go run main.go

Output yang akan ditampilkan adalah :

    Hello World

Perintah go run artinya menjalankan program dengan memanggil file yang dipisahkan dengan spasi, terlebih dahulu di compile menjadi executable yang disimpan dalam direktori sementara program berjalan. Apabila program anda tidak berjalan, maka anda membuat kesalahan dan Go compiler akan memberi petunjuk dimana letak kesalahan anda. Compiler Go tidak memiliki toleransi untuk kesalahan sekecil apapun.

###Cara Membaca Program Go

Selanjutnya, kita akan melihat program ini secara lebih terperinci. Program Go dibaca dari atas ke bawah dan dari kiri ke kanan. Baris pertama menyatakan ini :

```go
package main
```

Package main dikenal sebagai “package declaration”. Setiap program Go harus dimulai dengan mendeklarasikan sebuah package. Dimana package adalah cara pengorganisasian oleh Go dan code yang digunakan kembali. Ada dua jenis program Go, yaitu :

1.	Executables yaitu jenis program yang dapat dijalankan langsung dari terminal
2.	Libraries yaitu koleksi code dimana package yang kita buat dapat bersama-sama digunakan pada program lain

Selanjutnya :

```go
import "fmt"
```

Kaca kunci import adalah bagaimana kita menyertakan code dari package lain yang digunakan pada program ini. Package fmt (singkatan untuk format) adalah menerapkan format untuk input dan output. Tanda kutip ganda pada fmt (“fmt”) dikenal sebagai string literal” yang merupakan jenis ekspresi.

Selanjutnya, kita akan melihat deklarasi fungsi berikut :

```go
func main() {
  fmt.Println("Hello World")
}
```

*Functions* adalah blok bangunan dari program Go, dimana memiliki input, output dan serangkaian langkah-langkah yang disebut pernyataan kemudian dijalankan dalam kerangka tertentu.

##BAB 3. *TYPES*
Tipe data dikategorikan didalam suatu set yang terkait dengan nilai-nilai, dimana setiap nilai menggambarkan operasi yang dapat dilakukan oleh fungsi tersebut. Tipe data pada program Go sama halnya dengan tipe data pada umumnya yaitu terdiri dari 3 jenis, yaitu :

1.	Number
2.	String
3.	Booleans 

###*1. Number* 
Tipe data number dibagi menjadi dua jenis yaitu *integers* dan *floating*.

**Integers**

*Integers* adalah bilangan bulat seperti halnya dalam matematika. Jenis angka ini tanpa komponen desimal. Tipe *integers* digunakan dalam program untuk mendeklarasikan suatu variable yang termasuk bilangan bulat. Dengan kelompok :

- int/int32	: digunakan untuk bilangan bulat yang jumlah digitnya 2 pangkat 32
- int64		: digunakan untuk bilangan bulat yang jumlah digitnya 2 pangkat 64

contoh program Go nya :
```go
package main

import “fmt”
    
func main (){
        var x int
        x := 9
        fmt.Println(x)
    }
```

**Floating**

*Floating* adalah tipe angka desimal, seperti halnya integer, jumlah digit untuk data float dideklarasikan dengan :

- float/float32	: digunakan untuk bilangan bulat yang jumlah digitnya  2 pangkat 32
- float64		: digunakan untuk bilangan bulat yang jumlah digitnya  2 pangkat 64
    
contoh program Go nya :

```go
package main
    
    import “fmt”
    
    func main (){
        var x float
        x := 9.358463
        fmt.Println(x)
    }
```

###*2. String*
Adalah tipe data yang bias dikatakan sebagai karakter tertentu. Setiap karakter mempunyai kode nya tersendiri. Seperti \n yang berarti *newline* dan \t berarti *tab*, dan lainnya.
Pada bab sebelumnya, dimana anda diminta untuk membuat program dengan output “Hello Word”. Tipe data tersebut termasuk ke dalam string. 

Contoh program Go nya :
```go
package main
    
import "fmt"
    
func main() {
    fmt.Println(len("Hello World"))
    }
```

###*3. Booleans*
*Booleans* disini dapat dikatakan sebagai operator logika yang dapat menghasilkan output ‘TRUE’ dan ‘FALSE’. Nilai tersebut dapat dihasilkan dengan pengoperasian menggunakan operator ‘AND’, ‘OR’, dan ‘NOT’.

**AND (&&)**
Operator && akan bernilai ‘TRUE’ apabila semua persyaratan yang diminta bernilai ‘TRUE’, namun jika salah satu persyaratannya bernilai ‘FALSE’ maka akan menghasilkan ‘FALSE’.

**OR (||)**
Pada operator OR akan menghasilkan “FALSE” jika semua persyaratan bernili ‘FALSE’, dapat dikatan operator OR ini merupakan kebalikan dari operator AND.

**NOT (!)**
Operator NOT ini akan mengasilkan keluaran yang merupakan kebalikan dari nilai yang diperintahkan atau dioperasikan.
Contoh :

jika diberikan perintah :

    a>b menghasilkan ‘TRUE’
    
    dalam operator NOT menjadi :
    !(a>b) menghasilkan ‘FALSE’

## BAB 4. VARIABLES

*Variable* adalah sebuah lokasi penyimpanan suatu nilai dengan spesifikasi tipe dan nama yang terkait. Mudahnya bisa digambarkan ketika kita akan memberikan nilai ke dalam suatu lokasi, maka yang harus dilakukan adalah menyiapkan lokasi terlebih dahulu. Analoginya, seperti kita ingin mengambil air putih, sebelum kita menuangkan air putih terlebih dahulu tentunya kita akan menyiapkan gelas, bukan? Barulah air dapat terisi. 

Sudah bias dibayangkan sampai sini?

Jika sudah.. Baik, kita lanjutkan..

Coba perhatikan terlebih dahulu bentuk deklarasi *variable* pada program Go berikut :
```go
package main

import "fmt"

func main() {
    var x string = "Hello World"
    fmt.Println(x)
}
```

Dari program tersebut, bisa kita lihat *variable* di dalam Go dibuat dari kata kunci var. Deklarasi ‘var’ ini dibuat untuk mempermudah dalam membaca, yaitu hanya mengambil tiga buah huruf  di awal kata dari *‘variable’*. Bisa dipahami bahwa 'var' mendeklarasikan nama *variable* ‘x’ dengan tipe data ‘string’ (Hello Word). Dari pemahaman tersebut, disimpulkan bahwa pembuatan *variable* dalam Go memiliki aturanya tersendiri, diantaranya :
1. Nama *variable* dapat dibuat menggunakan huruf, angka atau *underscore*
2. Nama *variable* tidak boleh sama dengan kata kunci yang telah ada didalam Go 
3. Nama *variable* diawali oleh karakter huruf
4. Penulisan huruf capital dan huruf keci dapat mempengaruhi karena berbeda makna

Adapun hal yang membedakan Go dengan bahasa pemrograman lainnya adalah, penamaan *variable* dapat dipermudah atau dipersingkat dengan mendeklarasikan :

    x := 9
    artinya : var x int = 9

Yaitu *create and assign* dengan penulisan nama *variable* ‘x’, diikuti tanda titik dua dan sama dengan, angka 9 adalah isi dari *variable* yang dideklarasikan yaitu berfungsi untuk mengembalikan nilai. Penelisan seperti itu mempersingkat pendeklarasiaan *variable*.

###Cara Penamaan *Variable*

Didalam pengembangan perangkat lunak, penamaan suatu *variable* merupakan hal yang penting. Pembuatan nama *variable* bertujuan untuk menggambarkan tipe *variable* apa yang kita tuju. Dalam perkembangannya, Go memiliki aturan untuk penamaan *variable* dengan 2 suku kata, dimana pada Go tidak di tulis menjadi 2 suku kata terpisah tetapi 2 suku kata yang digabung, namun memiliki aturan :
1. Huruf pertama dari kata pertama adalah huruf kecil dan seterusnya
2. Huruf pertama dari kata-kata selanjutnya adalah huruf besar dan seterusnya huruf kecil. Contoh :
    
    catsName := “Brown”
    fmt.Println (“My Cat’s name is”, catasName)

Penulisan huruf kapital pada kata kedua untuk memanipulasi pemberian spasi.


##BAB 5. *CONTROL STRUCTURES*

*Control structures* dibutuhkan dalam mengerjakan suatu proses yang berulang-ulang (looping).

###1. For

For digunakan untuk mengulang daftar pernyataan (blok) beberapa kali sehingga kita tidak perlu menuliskannya secara berulang, cukup tuliskan perintahnya sekali saja. Contoh program untuk menampilkan angka 1 sampai 10:

```go
func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
```

* i merupakan variabel yang digunakan untuk mengontrol bilangan atau angka yang ingin di tampilkan.
* Alur iterasi program di atas adalah: Untuk i=1 di cek apakah nilai i masih kecil atau sama dengan 10. Jika i bernilai kecil atau sama dengan 10 (pernyataan *TRUE*) maka akan dijalankan perintah selanjutnya yaitu `fmt.Println(i)` yang artinya nilai i akan ditampilkan selama masih memenuhi syarat kecil atau sama dengan 10. Setelah bilangan pertama ditampilakan maka nilai i akan bertambah melalui perintah `i++`. Nilai i selanjutnya akan menjalankan perintah yang sama dengan yang pertama. *Looping* akan berhenti saat i bernilai lebih dari 10.

###2. If

If merupakan *statement* logika yang biasanya diikuti oleh *statement* `else`. Perintah `if` akan dijalankan apabila pernyataan bernilai *TRUE*, namun jika pernyataan bernilai *FALSE* maka yang akan dijalankan adalah perintah pada bagian `else`. `if` dan `else` biasanya digunakan bersamaan dengan `for`. Contohnya adalah untuk menampilkan angka 1 sampai 10 serta jenisnya apakan bilangan ganjil atau genap.

```go
func main() {
    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "genap")
        } else {
            fmt.Println(i, "ganjil")
        }
    }
}
```
* Alur iterasi program ini adalah: saat i=1 maka akan di cek terlebih dahulu apakan nilai i lebih kecil atau sama dengan 10. Apabila i kecil atau sama dengan 10 maka aka dilakukan perintah selanjutnya yaitu `if i % 2 == 0`. Jika pernyataan ini bernilai *TRUE* maka akan dijalankan perintah `fmt.Println(i, "genap")` yang akan menampilkan nilai i tersebut dan disertai dengan *string* 'genap'. Namun apabila pernyataan bernilai *FALSE* (i tidak habis dibagi 2) maka akan dijalankan perintah `fmt.Println(i, "ganjil")`. Begitu seterusnya hingga program akan berhenti apabila nilai i lebih besar dari 10.

###3. Switch

Misalkan kita ingin membuat program yang dapat menampilkan nama dari suatu angka. Agar lebih memudahkan kita dapat menuliskan programnya dengan menggunakan format `switch case`. Contoh: menampilkan nama angka dari 1 sampai 5.

```go
switch i {
case 0: fmt.Println("Nol")
case 1: fmt.Println("Satu")
case 2: fmt.Println("Dua")
case 3: fmt.Println("Tiga")
case 4: fmt.Println("Empat")
case 5: fmt.Println("Lima")
default: fmt.Println("Angka tidak diketahui")
}
```

*Statement* switch dimulai dengan kata kunci `switch` dan diikuti oleh sebuah ekspresi (dalam hal ini `i`) dan beberapa `case`. Nilai `i` akan dibandingkan dengan `case` yang cocok. Jika terdapat `case` yang sesuai dengan nilai `i` maka perintah dalam `case` tersebut akan dieksekusi. Namun apabila tidak ada `case` yang sesuai dengan nilai `i` maka yang akan dieksekusi adalah `default`.

##BAB 6. ARRAY, *SLICES*, DAN *MAPS*

###*1. Array*

Array adalah kumpulan dari nilai-nilai yang bertipe sama dan berurutan yang disimpan di dalam memori dengan jumlah element tertentu. 

Contoh array dalam Go :

    var x [9]int 
    
    - x adalah nama variable
    - [9] adalah jumlah elemen array
    - int adalah tipe data untuk bilangin bulat

Dapat kita lihat untuk *statement* diatas memiliki fungsi untuk memberi perintah kepada computer untuk menyiapkan 9 memori atau wadah yang akan diisi dengan bilangan bulat dan diberi nama x.

Selanjutnya, kita coba untuk menjalankan program berikut ya :
```go
package main

import "fmt"

func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
} 
```

Jangan lupa, dibiasakan sebelum menulis program, file disimpan terlebih dahulu dengan format ekstensi go. Program ini kita simpan dengan nama array1.go

Setelah di jalankan pada terminal, anda akan melihat output :

    [0 0 0 0 100]

Dari program diatas, dapat dipahami bahwa program memberikan perintah kepada komputer untuk menyiapkan 5 buah memori atau elemen berupa bilangan bulat. Kemudian dideklarasikan bahwa nilai dari x[4] = 100, lalu print nilai x setiap elemen array. Perlu diketahui bahwa dalam array, index array dimulai dari 0 (0 1 2 3 4, dst..). Sehingga index array ke 4, yaitu urutan array ke 5 bernilai 100. Dan elemen array lainnya bernilai 0, karena tidak dideklarasikan.

Apakah sudah cukup dipahami dengan penggambaran seperti itu?

Jika sudah, mari kita coba membuat program yang fungsinya lebih kompleks :

Simpan program berikut dengan nama array2.go
```go
package main

import “fmt”

func main() {
    var x [5]float64
    x[0] = 67
    x[1] = 78
    x[2] = 84
    x[3] = 98
    x[4] = 79
    var total float64 = 0.0
    for i := 0; i < len(x); i++ {
        total += x[i]
    }
    fmt.Println(total / float64(len(x)))
}
```

Program tersebut akan menghasilkan output :

    81.2

Dapat dipahami bahwa dalam program tersebut bertujuan untuk menjumlahkan semua elemen array kemudian dihitung nilai rata-ratanya.
Untuk melakukan operasi matematika tersebut, dipermudah dengan fungsi looping for seperti yang telah anda pelajari pada bab 5 sebelumnya. Len (x) menyatakan besar elemen array yang dideklarasikan. 

###*2. Slices*
Slices dapat dikatakan sebagai salah satu segmen array. Seperti halnya array, slices juga memiliki index dan panjang tertentu. Namun, slices ini panjangnya dapat diubah. 
Untuk lebih jelaskan, mari kita buat program berikut :

Disimpan dengan nama slice.go
```go
package main

import “fmt”
 
func main() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    
    fmt.Println(slice1, slice2)
}
```
Setelah dijalankan, output yang diperoleh adalah :

    [1,2,3] [1,2]

Bagaiamana? Apakah dapat dipahami?

Penjelasannya sebagai berikut :

Pertama slice1 := []int{1,2,3} memerintahkan untuk membuat slice dengan jumlah elemen yang tidak ditentukan namun sudah dideklarasikan secara langsung dengan := bernilai 1,2 dan 3. Kemudian, dibuat juga slice2 dengan jumlah array sebanyak 2 elemen. Lalu, disalin nilai pada slice1 yang jumlahnya disesuaikan dengan array pada slice2 yaitu hanya 2 elemen. Terakhir, diperintahkan untuk mencetak slice1 dan slice 2. Sehingga nilai pada slice1 adalah [1,2,3] dan slice2 [1,2].
Apakah sekarang sudah cukup dipahami? Untuk lebih yakin, silahkan coba program sederhana tersebut.
Selamat mencoba ;-)


###*3. Maps*

*Map* adalah koleksi atau sekumpulan pasangan *key-value*. Juga dikenal sebagai array asosiatif atau kamus. *Maps* digunakan untuk melihat *value* berdasarkan *key* yang berhubungan. Contoh: *Maps* beberapa unsur kimia yang di-indeks-kan oleh simbol nya.

```go
package main
import "fmt"
func main() {
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"
    fmt.Println(elements["Li"])
}
```


##BAB 7. FUNGSI

Fungsi merupakan bagian independen yang memetakan nol atau lebih parameter input ke nol atau lebih parameter output. Fungsi juga dikenal sebagai prosedur atau subrutin.

###1. Bekerja dengan Lebih dari Satu Fungsi

Pada program yang sebelumnya kita bekerja dengan hanya satu fungsi. Pada bagian ini akan dijelaskan bagaimana membuat program dengan lebih dari satu fungsi.

```go
package main
import "fmt"
func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}
func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))
}
```

Program di atas akan memberikan output berupa nilai rata-rata dari angka yang ada di dalam array `xs`. Fungsi main merupakan fungsi utama yang dijalankan pertama kali. Di dalam fungsi main diinisialisasi sebuah array `xs` yang berisi lima bilangan. Lalu di dalam fungsi main juga terdapat perintah `fmt.Println(average(xs))` yang memanggil fungsi average dan memasukkan nilai `xs` ke dalamnya. Setelah perintah di dalam fungsi average dieksekusi maka akan dijalankan perintah untuk menampilkan nilai rata-rata dari kelima bilangan tersebut.

###2. Mengembalikan Lebih dari Satu Nilai

Dalam bahasa Go juga memungkinkan kita untuk  mengembalikan beberapa nilai dari suatu fungsi. Contoh:

```go
package main
import "fmt"
func f() (int, int) {
    return 5, 6
}
func main() {
    x, y := f()
}
```

Fungsi f mengembalikan dua nilai `(int, int)` yaitu 5 dan 6 yang selanjutnya dipanggil dalam fungsi main.

###3. Fungsi *Variadic*

Fungsi jenis ini merupakan fungsi yang hanya dapat dijalankan pada bahasa Go. Dengan fungsi ini kita dapat membuat suatu variabel yang dapat berisi nol atau lebih nilai dengan menggunakan `...` sebelum tipe variabel.

```go
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}
func main() {
    fmt.Println(add(1,2,3))
}
```

Dengan program diatas kita dapat memasukkan nilai ke dalam fungsi dalam jumlah yang tidak perlu ditentukan terlebih dahulu.

###*4. Closure*

*Closure* merupakan suatu fungsi yang berisi fungsi lain di dalamnya. Contoh:

```go
package main
import "fmt"
func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1,1))
}
```

Di dalam fungsi main terdapat variabel `add` dengan tipe `func(int,int)int` (sebuah fungsi yang berisi dua `int` dan mengembalikan satu `int`.

###5. Rekursi

Suatu fungsi dapat memanggil dirinya sendiri dalam fungsi rekursi. Contohnya adalah dalam program penghitung nilai faktorial berikut:

```go
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}
func main(){
    fmt.Println(factorial(2))
}
```

Misalnya kita ingin menghitung nilai faktorial dari angka 2 melalui perintah: `fmt.Println(factorial(2))`

Fungsi main memanggil fungsi factorial dan memberikan nilai x=2. Dalam fungsi factorial nilai x tidak sama dengan 0 maka perintah selanjutnya yang akan dieksekusi adalah `return x * factorial(x-1)` yang berarti fungsi factorial tersebut akan memanggil dirinya sendiri dengan nilai x=1. Saat x=1 fungsi factorial akan memanggil kembali dirinya dengan x=0. Saat x=0 maka pernyataan `if` bernilai *TRUE* sehingga akan mengembalikan nilai 1. Pada akhirnya fungsi factorial akan mengembalikan nilai `2*1*1=2`. Jadi dapat dibuktikan bahwa nilai faktorial dari 2 adalah sama dengan 2.



























