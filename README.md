# Sistem manajemen stock barang

Bahasa yang digunakan adalah golang. Database sqlite. Tabel - tabel yang ada di database diantaranya :
**products** : Tabel yang digunakan untuk menyimpan barang - barang yang saat ini ada di gudang beserta stok nya.
**product_in** : Tabel yang digunakan untuk melakukan pencatatan barang masuk
**product_out** : Tabel yang digunakan untuk melakukan pencatatan barang keluar
**product_history**: Tabel yang digunakan untuk melakukan pencatatan keseluruhan transaksi yang dilakukan oleh user. Di tabel ini terdapat tiga tipe data :
 -  In : adalah semua data barang masuk.
 - Out: adalah semua data barang keluar.
 - Opname: adalah penyesuaian data barang antara barang di sistem dan barang yang ada di gudang.

# Daftar API
|Endpoint| Tipe  | Keterangan|
|--|--|--|
| /get_barang_masuk | GET | Menampilkan keseluruhan data barang masuk |
| /submit_barang_masuk | POST | Input data barang masuk |
| /update_barang_masuk | POST | Update data barang masuk |
| /delete_barang_masuk | POST | Delete data barang masuk |
|--|--|
| /get_barang_keluar | GET | Menampilkan keseluruhan data barang keluar
| /submit_barang_keluar | POST | Input data barang keluar
| /update_barang_keluar | POST | Update data barang keluar
| /delete_barang_keluar |  POST| Delete data barang keluar
|--|--|
| /get_stok_barang | GET | Menampilkan keseluruhan data stok barang
| /update_stok_barang | POST | Update data stok barang
| /submit_stok_barang | POST | Input data stok barang
| ---|---|---
| /laporan_nilai_barang | GET | laporan nilai barang
| /laporan_penjualan | GET | Laporan penjualan
| /download_csv | GET | Download csv

Mengenai parameter yang digukan lebih lengkapnya terdapat di Salestock.postman_collection.json 

## Petunjuk untuk menjalankan aplikasi

 1. git clone ke environment go lokal. Pastikan anda berada di folder go/src/api-inventory/. Jalankan 
	> git clone https://github.com/akhdakhz03/ims.git
 2. Masuk ke folder api-inventory lalu buka terminal dan jalankan
	 > go run main.go
 3. Buka postman yang sudah ada di repo dan coba masing - masing api yang tersedia.
 4. Hasil dari file csv akan masuk ke folder csv