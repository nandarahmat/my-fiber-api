-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               11.6.2-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             12.8.0.6908
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table db_golang.alamat
CREATE TABLE IF NOT EXISTS `alamat` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `judul` varchar(255) NOT NULL,
  `nama_penerima` varchar(255) NOT NULL,
  `no_telp` varchar(255) NOT NULL,
  `detail_alamat` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `alamat_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table db_golang.alamat: ~2 rows (approximately)
INSERT INTO `alamat` (`id`, `id_user`, `judul`, `nama_penerima`, `no_telp`, `detail_alamat`, `created_at`, `updated_at`) VALUES
	(1, 10, 'Rumah', 'Budi Santoso', '081234567890', 'Jl. Merdeka No. 10, Jakarta Pusat', '2025-03-31 07:58:52', '2025-03-31 07:58:52'),
	(3, 10, 'North Domenicaview', 'Scott Blick', '509-507-8196', 'Conner Haven, 865 Kreiger Street, Albabury', '2025-03-31 08:19:16', '2025-03-31 08:19:16');

-- Dumping structure for table db_golang.category
CREATE TABLE IF NOT EXISTS `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table db_golang.category: ~7 rows (approximately)
INSERT INTO `category` (`id`, `name`, `created_at`, `updated_at`) VALUES
	(1, 'Elektronik', '2025-03-30 20:15:04.000', '2025-03-30 20:15:04.000'),
	(2, 'Pakaian', '2025-03-30 20:15:04.000', '2025-03-30 20:15:04.000'),
	(3, 'Makanan', '2025-03-30 20:15:04.000', '2025-03-30 20:15:04.000'),
	(4, 'Buku', '2025-03-30 20:15:04.000', '2025-03-30 20:15:04.000'),
	(5, 'Perabotan', '2025-03-30 20:15:04.000', '2025-03-30 20:15:04.000'),
	(7, 'test', '2025-03-31 05:25:26.271', '2025-03-31 05:25:26.271'),
	(8, 'test', '2025-03-31 13:27:05.805', '2025-03-31 13:27:05.805'),
	(9, 'test', '2025-03-31 13:27:15.408', '2025-03-31 13:27:15.408');

-- Dumping structure for table db_golang.detail_trx
CREATE TABLE IF NOT EXISTS `detail_trx` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_trx` int(11) NOT NULL,
  `id_log_produk` int(11) NOT NULL,
  `id_toko` int(11) NOT NULL,
  `kuantitas` int(11) NOT NULL,
  `harga_total` decimal(10,2) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_trx` (`id_trx`),
  KEY `id_log_produk` (`id_log_produk`),
  KEY `id_toko` (`id_toko`),
  CONSTRAINT `detail_trx_ibfk_1` FOREIGN KEY (`id_trx`) REFERENCES `trx` (`id`),
  CONSTRAINT `detail_trx_ibfk_2` FOREIGN KEY (`id_log_produk`) REFERENCES `log_produk` (`id`),
  CONSTRAINT `detail_trx_ibfk_3` FOREIGN KEY (`id_toko`) REFERENCES `toko` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table db_golang.detail_trx: ~0 rows (approximately)
INSERT INTO `detail_trx` (`id`, `id_trx`, `id_log_produk`, `id_toko`, `kuantitas`, `harga_total`, `created_at`, `updated_at`) VALUES
	(1, 1, 3, 2, 3, 40500000.00, '2025-04-04 03:46:53', '2025-04-04 03:46:53');

-- Dumping structure for table db_golang.foto_produk
CREATE TABLE IF NOT EXISTS `foto_produk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_produk` int(11) NOT NULL,
  `url` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_produk` (`id_produk`),
  CONSTRAINT `foto_produk_ibfk_1` FOREIGN KEY (`id_produk`) REFERENCES `produk` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table db_golang.foto_produk: ~2 rows (approximately)
INSERT INTO `foto_produk` (`id`, `id_produk`, `url`, `created_at`, `updated_at`) VALUES
	(1, 1, 'https://png.pngtree.com/png-vector/20220726/ourmid/pngtree-pura-balinese-hindu-temple-silhouette-png-image_6052600.png', '2025-03-31 13:01:16', '2025-03-31 13:01:16'),
	(2, 1, 'https://png.pngtree.com/png-vector/20220726/ourmid/pngtree-pura-balinese-hindu-temple-silhouette-png-image_6052600.png', '2025-03-31 13:01:16', '2025-03-31 13:01:16');

-- Dumping structure for table db_golang.log_produk
CREATE TABLE IF NOT EXISTS `log_produk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_produk` int(11) NOT NULL,
  `nama_produk` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `harga_reseller` decimal(10,2) NOT NULL,
  `harga_konsumen` decimal(10,2) NOT NULL,
  `deskripsi` text DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `id_toko` int(11) NOT NULL,
  `id_category` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_produk` (`id_produk`),
  KEY `id_toko` (`id_toko`),
  KEY `id_category` (`id_category`),
  CONSTRAINT `log_produk_ibfk_1` FOREIGN KEY (`id_produk`) REFERENCES `produk` (`id`),
  CONSTRAINT `log_produk_ibfk_2` FOREIGN KEY (`id_toko`) REFERENCES `toko` (`id`),
  CONSTRAINT `log_produk_ibfk_3` FOREIGN KEY (`id_category`) REFERENCES `category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table db_golang.log_produk: ~0 rows (approximately)
INSERT INTO `log_produk` (`id`, `id_produk`, `nama_produk`, `slug`, `harga_reseller`, `harga_konsumen`, `deskripsi`, `created_at`, `updated_at`, `id_toko`, `id_category`) VALUES
	(1, 1, 'Laptop Gaming X1', 'laptop-gaming-x1', 12000000.00, 13500000.00, 'Laptop gaming dengan spesifikasi tinggi untuk kebutuhan profesional dan hiburan.', '2025-04-04 03:42:55', '2025-04-04 03:42:55', 2, 1),
	(2, 1, 'Laptop Gaming X1', 'laptop-gaming-x1', 12000000.00, 13500000.00, 'Laptop gaming dengan spesifikasi tinggi untuk kebutuhan profesional dan hiburan.', '2025-04-04 03:46:21', '2025-04-04 03:46:21', 2, 1),
	(3, 1, 'Laptop Gaming X1', 'laptop-gaming-x1', 12000000.00, 13500000.00, 'Laptop gaming dengan spesifikasi tinggi untuk kebutuhan profesional dan hiburan.', '2025-04-04 03:46:53', '2025-04-04 03:46:53', 2, 1);

-- Dumping structure for table db_golang.produk
CREATE TABLE IF NOT EXISTS `produk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama_produk` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `harga_reseller` varchar(255) NOT NULL,
  `harga_konsumen` varchar(255) NOT NULL,
  `stok` int(11) NOT NULL DEFAULT 0,
  `deskripsi` text DEFAULT NULL,
  `id_toko` int(11) NOT NULL,
  `id_category` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `id_toko` (`id_toko`),
  KEY `id_category` (`id_category`),
  CONSTRAINT `produk_ibfk_1` FOREIGN KEY (`id_toko`) REFERENCES `toko` (`id`) ON DELETE CASCADE,
  CONSTRAINT `produk_ibfk_2` FOREIGN KEY (`id_category`) REFERENCES `category` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table db_golang.produk: ~1 rows (approximately)
INSERT INTO `produk` (`id`, `nama_produk`, `slug`, `harga_reseller`, `harga_konsumen`, `stok`, `deskripsi`, `id_toko`, `id_category`, `created_at`, `updated_at`) VALUES
	(1, 'Laptop Gaming X1', 'laptop-gaming-x1', '12000000', '13500000', 10, 'Laptop gaming dengan spesifikasi tinggi untuk kebutuhan profesional dan hiburan.', 2, 1, '2025-03-31 12:49:15', '2025-03-31 12:49:15');

-- Dumping structure for table db_golang.toko
CREATE TABLE IF NOT EXISTS `toko` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `nama_toko` varchar(255) NOT NULL,
  `url_foto` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `toko_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table db_golang.toko: ~0 rows (approximately)
INSERT INTO `toko` (`id`, `id_user`, `nama_toko`, `url_foto`, `created_at`, `updated_at`) VALUES
	(2, 10, 'toko update', '/public/uploads/toko_10_1743411171.PNG', '2025-03-31 06:28:02', '2025-03-31 08:52:51');

-- Dumping structure for table db_golang.trx
CREATE TABLE IF NOT EXISTS `trx` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) NOT NULL,
  `alamat_pengiriman` int(11) NOT NULL,
  `harga_total` int(11) NOT NULL,
  `kode_invoice` varchar(255) NOT NULL,
  `method_bayar` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  KEY `alamat_pengiriman` (`alamat_pengiriman`),
  CONSTRAINT `trx_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `user` (`id`),
  CONSTRAINT `trx_ibfk_2` FOREIGN KEY (`alamat_pengiriman`) REFERENCES `alamat` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table db_golang.trx: ~0 rows (approximately)
INSERT INTO `trx` (`id`, `id_user`, `alamat_pengiriman`, `harga_total`, `kode_invoice`, `method_bayar`, `created_at`, `updated_at`) VALUES
	(1, 10, 1, 40500000, 'INV-20250404034653', 'Transfer Bank', '2025-04-04 03:46:53', '2025-04-04 03:46:53');

-- Dumping structure for table db_golang.user
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) NOT NULL,
  `kata_sandi` varchar(255) NOT NULL,
  `no_telp` varchar(255) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `jenis_kelamin` varchar(255) NOT NULL,
  `tentang` text DEFAULT NULL,
  `pekerjaan` varchar(255) DEFAULT NULL,
  `email` varchar(255) NOT NULL,
  `id_provinsi` varchar(255) NOT NULL,
  `id_kota` varchar(255) NOT NULL,
  `is_admin` tinyint(1) DEFAULT 0,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `no_telp` (`no_telp`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Dumping data for table db_golang.user: ~1 rows (approximately)
INSERT INTO `user` (`id`, `nama`, `kata_sandi`, `no_telp`, `tanggal_lahir`, `jenis_kelamin`, `tentang`, `pekerjaan`, `email`, `id_provinsi`, `id_kota`, `is_admin`, `created_at`, `updated_at`) VALUES
	(10, 'test fajar', '$2a$10$wJYDnsnEzhgLB6FCjtTZKO4gehUh9P/BMutnRfP/t/OyUKDpKLhoK', '08961231235', '2006-05-02', 'laki-laki', 'test', 'developer', 'tes5t@mail.com', '11', '1101', 0, '2025-03-31 06:28:02', '2025-03-31 07:48:06');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
