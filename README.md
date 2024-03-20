# Go-Ticket

# BAYUAJI PRADIPTA ARINANDA

# FLowChart

1. User login ke sistem dengan yang sudah di register
2. User dapat melihat film apa saja yang sedang tayang dan memilih jadwal yang di inginkan 
3. User memilih kursi yang tersedia di dalam teater untuk jadwal tayang yang dpilih
4. jika User IYA melanjutkan, mereka akan pergi ke tahap pembayaran.
   Jika TIDAK, sistem akan memberi waktu tunggu sebelum kursi yang dipilih mejadi tidak tersedia
5. User memeriksa semua detail pemesanan mereka, seperti film, jadwal tayang, dan tempat duduk, sebelum melakukan pembayaran.
6. jika User TIDAK melanjutkan ke pembayaran dalam waktu tertentu, sistem akan menganggap pemesanan itu tidak dilanjutkan, dan tempat duduk akan dibuat tersedia lagi untuk pengguna lain.
7. Setelah pembayaran berhasil, e-tiket akan tersedia. User dapat melihat detail  atau menunjukkan e-tiket ini saat masuk ke bioskop.

# ERD
1. User: Berisi informasi tentang pengguna yang terdaftar, termasuk apakah mereka adalah admin atau bukan.
2. Movies: Menyimpan informasi tentang film yang ditayangkan.
3. ShowTime: Mengatur jadwal tayang untuk masing-masing film.
4. Seat: Mencatat detail kursi untuk masing-masing jadwal tayang, termasuk statusnya (tersedia atau dipesan).
5. Checkout (Booking): Mengelola pemesanan yang dilakukan oleh pengguna, menyertakan referensi ke pengguna, jadwal tayang, dan kursi yang dipilih.


# Problem and Solution

Untuk meminimalisir race condition dan sistem tetap mendapatkan performa yang ideal disini saya menggunakan 3 penedekatan yaitu:

1. Go-routines : Memastikan koneksi terhadap ke database dilakukan prosesnya secara lebih cepat
2. Transaction : Menggunakan prinsip ACID, yang meliputi Atomicity (Kesatuan), Consistency (Konsistensi), Isolation (Pengasingan), dan Durability (Ketahanan), sangat membantu dalam memastikan bahwa ketika banyak pengguna atau sistem mencoba melakukan perubahan data di saat yang sama, perubahan tersebut tidak akan saling mengganggu. Ini berarti sistem dapat mencegah data dari kekacauan atau kerusakan walaupun banyak tugas yang terjadi bersamaan.
3.Scheduled task : Dengan melakukan kombinasi Goroutines dan Ticker yang melakukan operasi pada interval waktu yang teratur yang membuat sistem lebih efisien dan responsif untuk menjalankan sistem secara teratur, tanpa perlu menunggu input pengguna atau event lain.
   - Dalam kasus sistem ini memungkinkan aplikasi untuk secara otomatis memperbarui status kursi yang telah melebihi batas waktu pemesanan yang ditentukan, membebaskan sumber daya yang mungkin terkunci oleh transaksi yang tidak selesai dan memastikan data tetap konsisten dan up-to-date. 