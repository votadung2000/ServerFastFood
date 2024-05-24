-- -------------------------------------------------------------
-- TablePlus 6.0.4(556)
--
-- https://tableplus.com/
--
-- Database: go_fast_food_db
-- Generation Time: 2024-05-24 14:59:42.9440
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `image_id` int DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `delivery_address` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `recipient_name` varchar(100) NOT NULL,
  `street_address` varchar(255) NOT NULL,
  `country` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `city` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `postal_code` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `favorites` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `user_id` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `images` (
  `id` int NOT NULL AUTO_INCREMENT,
  `url` varchar(255) DEFAULT NULL,
  `width` int DEFAULT NULL,
  `height` int DEFAULT NULL,
  `cloud_name` varchar(255) DEFAULT NULL,
  `extension` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `order_item` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `quantity` int NOT NULL,
  `price` decimal(18,2) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `tax_fees` decimal(18,2) NOT NULL DEFAULT '0.00',
  `delivery_fee` decimal(18,2) NOT NULL DEFAULT '0.00',
  `total` decimal(18,2) NOT NULL,
  `coupon_id` int DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `canceled_at` datetime DEFAULT NULL,
  `completed_at` datetime DEFAULT NULL,
  `delivery_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `image_id` int DEFAULT NULL,
  `taste` text,
  `price` float NOT NULL,
  `category_id` int NOT NULL,
  `discount` float DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `quantity` int DEFAULT '0',
  `sold` int DEFAULT '0',
  `featured` int NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `salt` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `role` int NOT NULL DEFAULT '1',
  `avatar_id` int DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `verifications` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `otp_code` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `type` int NOT NULL DEFAULT '1',
  `status` int NOT NULL DEFAULT '1',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `categories` (`id`, `name`, `image_id`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Burger', 1, 1, '2024-03-05 09:25:48', '2024-03-06 07:37:57'),
(2, 'Pizza', 2, 1, '2024-03-05 09:26:08', '2024-03-06 07:38:04'),
(3, 'Sandwich', 3, 1, '2024-03-05 09:26:12', '2024-03-07 01:46:06'),
(4, 'Fruits', 4, 1, '2024-03-05 09:26:16', '2024-03-07 01:46:15');

INSERT INTO `delivery_address` (`id`, `user_id`, `status`, `recipient_name`, `street_address`, `country`, `city`, `postal_code`, `description`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'USER 1', 'Hẻm 59 Hồng Lạc Phường 10 Quận Tân Bình', 'Việt Nam', 'Thành phố Hồ Chí Minh', '72117', '', '2024-05-24 07:12:15', '2024-05-24 07:12:15');

INSERT INTO `favorites` (`id`, `product_id`, `user_id`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, '2024-03-19 06:43:25', '2024-03-19 06:43:25'),
(2, 2, 1, 1, '2024-03-19 07:09:59', '2024-03-19 07:09:59');

INSERT INTO `images` (`id`, `url`, `width`, `height`, `cloud_name`, `extension`, `created_at`, `updated_at`) VALUES
(1, 'static/1709775915665309000_burger.png', 512, 512, 'local', '.png', '2024-03-07 08:45:16', '2024-03-07 08:45:16'),
(2, 'static/1709775925414981000_pizza.png', 512, 512, 'local', '.png', '2024-03-07 08:45:25', '2024-03-07 08:45:25'),
(3, 'static/1709775939568126000_sandwich.png', 512, 512, 'local', '.png', '2024-03-07 08:45:40', '2024-03-07 08:45:40'),
(4, 'static/1709775947878244000_strawberry.png', 512, 512, 'local', '.png', '2024-03-07 08:45:48', '2024-03-07 08:45:48'),
(5, 'static/1709785899745518000_beef_big_size.png', 754, 800, 'local', '.png', '2024-03-07 11:31:40', '2024-03-07 11:31:40'),
(6, 'static/1709786019806174000_beef_chicken.png', 800, 800, 'local', '.png', '2024-03-07 11:33:40', '2024-03-07 11:33:40'),
(7, 'static/1709786063831989000_seafood_beef.png', 800, 800, 'local', '.png', '2024-03-07 11:34:24', '2024-03-07 11:34:24'),
(8, 'static/1709786112814811000_two_tier_beef.png', 800, 800, 'local', '.png', '2024-03-07 11:35:13', '2024-03-07 11:35:13'),
(9, 'static/1709786152917494000_two_tier_beef_bib_size.png', 800, 800, 'local', '.png', '2024-03-07 11:35:53', '2024-03-07 11:35:53'),
(10, 'static/1710748576617960000_mushroom_pizza.png', 464, 273, 'local', '.png', '2024-03-18 14:56:17', '2024-03-18 14:56:17'),
(11, 'static/1710748666443785000_chicken_pizza.png', 798, 408, 'local', '.png', '2024-03-18 14:57:46', '2024-03-18 14:57:46'),
(12, 'static/1710748768330323000_product_strawberry.png', 1920, 1200, 'local', '.png', '2024-03-18 14:59:28', '2024-03-18 14:59:28'),
(13, 'static/1710748820693118000_product_grape.png', 1000, 667, 'local', '.png', '2024-03-18 15:00:21', '2024-03-18 15:00:21'),
(14, 'static/1710748842693436000_product_apple.png', 1437, 1052, 'local', '.png', '2024-03-18 15:00:43', '2024-03-18 15:00:43'),
(15, 'static/1710748865320730000_product_banana.jpeg', 738, 360, 'local', '.jpeg', '2024-03-18 15:01:05', '2024-03-18 15:01:05'),
(16, 'static/1714970447476256000_chup-anh-chan-dung4.jpg', 1024, 682, 'local', '.jpg', '2024-05-06 11:40:47', '2024-05-06 11:40:47'),
(17, 'static/1714970542399826000_chup-anh-chan-dung4.jpg', 1024, 682, 'local', '.jpg', '2024-05-06 11:42:22', '2024-05-06 11:42:22'),
(18, 'static/1714970759315172000_chup-anh-chan-dung4.jpg', 1024, 682, 'local', '.jpg', '2024-05-06 11:45:59', '2024-05-06 11:45:59');

INSERT INTO `order_item` (`id`, `order_id`, `product_id`, `product_name`, `status`, `quantity`, `price`, `created_at`, `updated_at`) VALUES
(11, 9, 1, 'Beef Burger', 1, 2, 8.00, '2024-05-23 07:10:15', '2024-05-23 07:10:15'),
(12, 9, 3, 'Seafood Burger', 1, 4, 10.00, '2024-05-23 07:10:15', '2024-05-23 07:10:15');

INSERT INTO `orders` (`id`, `user_id`, `status`, `tax_fees`, `delivery_fee`, `total`, `coupon_id`, `created_at`, `canceled_at`, `completed_at`, `delivery_at`) VALUES
(9, 1, 1, 0.00, 0.00, 56.00, 0, '2024-05-23 07:10:15', NULL, NULL, NULL);

INSERT INTO `products` (`id`, `name`, `image_id`, `taste`, `price`, `category_id`, `discount`, `status`, `description`, `quantity`, `sold`, `featured`, `created_at`, `updated_at`) VALUES
(1, 'Beef Burger', 5, 'Spicy', 8, 1, 0, 1, 'Description Beef Burger', 100, 2, 1, '2024-03-07 04:33:13', '2024-05-23 07:10:15'),
(2, 'Chicken Burger', 6, 'Spicy', 7, 1, 0, 1, 'Description Chicken Burger', 100, 0, 1, '2024-03-07 04:34:09', '2024-05-23 04:42:54'),
(3, 'Seafood Burger', 7, 'Spicy', 10, 1, 0, 1, 'Description Seafood Burger', 100, 4, 2, '2024-03-07 04:34:44', '2024-05-23 07:10:15'),
(4, 'Double Beef Burger', 8, 'Spicy', 15, 1, 0, 1, 'Description Double Beef Burger', 50, 0, 1, '2024-03-07 04:35:37', '2024-05-23 04:42:54'),
(5, 'Beef Cheese Burger', 9, 'Spicy', 15, 1, 0, 1, 'Description Beef Cheese Burger', 50, 0, 2, '2024-03-07 04:36:14', '2024-05-23 04:42:54'),
(6, 'Mushroom Pizza', 10, 'Normal', 20, 2, 10, 1, 'Description Mushroom Pizza', 50, 0, 2, '2024-03-18 07:57:18', '2024-05-23 04:42:54'),
(7, 'Chicken Pizza', 11, 'Normal', 20, 2, 5, 1, 'Description Chicken Pizza', 50, 0, 2, '2024-03-18 07:58:21', '2024-05-23 04:42:54'),
(8, 'Strawberry', 12, 'Normal', 12, 4, 5, 1, 'Description Strawberry', 100, 0, 2, '2024-03-18 08:00:09', '2024-05-23 04:42:54'),
(9, 'Grape', 13, 'Normal', 12, 4, 5, 1, 'Description Grape', 100, 0, 2, '2024-03-18 08:00:33', '2024-05-23 04:42:54'),
(10, 'Apple', 14, 'Normal', 12, 4, 5, 1, 'Description Apple', 100, 0, 2, '2024-03-18 08:00:56', '2024-05-23 04:42:54'),
(11, 'Banana', 15, 'Normal', 12, 4, 5, 1, 'Description Banana', 100, 0, 2, '2024-03-18 08:01:16', '2024-05-23 04:42:54');

INSERT INTO `users` (`id`, `name`, `user_name`, `password`, `salt`, `phone_number`, `email`, `status`, `address`, `role`, `avatar_id`, `created_at`, `updated_at`) VALUES
(1, 'User 1 updated', 'register1', 'e47cc3f5a5d88a7719f6a06408dc37f1', 'JlMOCUofftIAiGnqnGUIpGrEOWHCfmVQKbEOaKdSTlfqraxdCv', '0987654321', 'register1@gmail.com', 1, '123 demo user 11', 1, 18, '2024-03-05 13:53:22', '2024-05-06 04:45:59'),
(2, 'USER 2', 'register2', '8eb9ce668ac514f363f2edb071ebdda4', 'QIfqXoiQsVsTiPkFvGIwypdFEVFCAytxMnTvoZcwILDpBIJTgr', '0987654311', 'register2@gmail.com', 1, '123 demo demo', 1, 0, '2024-05-02 11:07:53', '2024-05-07 01:47:59');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;