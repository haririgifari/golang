// Sebelum ada GO Modules
// 1. Saat kita membuat aplikasi biasanyya kita akan menggunakan library atau dependency dari project lain
// 2. sebelum ada Go Modules, management untuk dependency sangat sulit dilakukan di Go-Lang
// 3. Biasanya kita akan meng-copy semua source code library atau dependency lain ke project kitam hal ini membuat project kita menjadi bengkak karena penuh dengan library orang lain
// 4. atau biasanya library orang lain akan kita save di GOPATH dirctory, namun hal ini akan sulit jika ternyata beberapa aplikasi butuh library yang sama dengan versi yang berbeda

// Go-Modules
// 1. Go-Modules mulai dikenalkan di Go-Lang 1.11 dan 1.12
// 2. Dengan GO-Modules, kita bisa membuat project dengan mudah dan dependency management yang sangat mudah

// Membuat module
// 1. untuk membuat module baru, kita bisa menggunakan perintah : go mod init nama-module
// 2. Go-Lang akan secara otomatis membuat file go-mod yang berisikan informasi nama module dan juga versi Go-Lang yang kita gunakan

// Rilis Module
// 1. Go-Lang terintergrasi baik dengan Git
// 2. Untuk merilis module, kita hanya perlu membuat Tag di Git

// Major Upgrade
// 1. Major upgrade biasanya terjadi dikarenakan ada perubahan pada isi kode program kita, sehingga membuat tidak backward compatible
// 2. sebaiknya hal ini sebisa mungkin dihindari
// 3. Namun jika tidak bisa dihindari, strategy terbaik adalah merubah nama module
// tambah /v2 di module

// LANJUT DI FILE GO-SAY-HELLO