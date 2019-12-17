-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Dec 17, 2019 at 02:58 PM
-- Server version: 10.1.38-MariaDB
-- PHP Version: 5.6.40

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `stock_management`
--

-- --------------------------------------------------------

--
-- Table structure for table `customer`
--

CREATE TABLE `customer` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `address` text NOT NULL,
  `added_at` datetime NOT NULL,
  `update_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `customer`
--

INSERT INTO `customer` (`id`, `name`, `phone`, `address`, `added_at`, `update_at`) VALUES
(2, 'nana', '12345678890', 'Jl. Jalan no 10', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(3, 'nanang', '12345678890', 'Jl. Jalan no 10', '2019-12-10 15:33:27', '0000-00-00 00:00:00'),
(4, 'Abdul Hafiz', '081317384592', '', '2019-12-10 15:41:20', '0000-00-00 00:00:00'),
(9, 'Abdul Hafiz', '081317384592', 'Jl. Benda Timur 8B, Blok E7 No. 9, Pamulang 2', '2019-12-11 14:53:54', '0000-00-00 00:00:00'),
(10, 'Hafid', '0812344567890', 'Jl. Jalan no 10', '2019-12-11 15:46:01', '0000-00-00 00:00:00'),
(13, 'Budi', '0812312313', 'Jl. Jalan no 10', '2019-12-12 15:50:45', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `master_order`
--

CREATE TABLE `master_order` (
  `id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `unit_price` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `payment_status` enum('UNPAID','PAID') NOT NULL DEFAULT 'UNPAID',
  `added_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `master_order`
--

INSERT INTO `master_order` (`id`, `product_id`, `customer_id`, `quantity`, `unit_price`, `total_price`, `payment_status`, `added_at`, `updated_at`) VALUES
(6, 5, 8, 10, 2000, 20000, 'UNPAID', '2019-12-11 14:53:23', '0000-00-00 00:00:00'),
(7, 11, 9, 15, 5000, 75000, 'UNPAID', '2019-12-11 14:53:54', '0000-00-00 00:00:00'),
(8, 8, 10, 21, 20000, 420000, 'UNPAID', '2019-12-11 15:46:01', '0000-00-00 00:00:00'),
(9, 8, 11, 20, 20000, 400000, 'UNPAID', '2019-12-12 10:08:49', '0000-00-00 00:00:00'),
(10, 3, 12, 1, 1234, 1234, 'UNPAID', '2019-12-12 14:24:39', '0000-00-00 00:00:00'),
(11, 3, 2, 1, 1234, 1234, 'UNPAID', '2019-12-12 15:49:51', '0000-00-00 00:00:00'),
(12, 3, 13, 2, 1234, 2468, 'UNPAID', '2019-12-12 15:50:45', '0000-00-00 00:00:00'),
(13, 8, 4, 15, 20000, 300000, 'UNPAID', '2019-12-12 17:09:50', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `code` varchar(20) NOT NULL,
  `price` int(11) NOT NULL,
  `qty` int(11) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `added_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`id`, `name`, `code`, `price`, `qty`, `is_deleted`, `added_at`, `updated_at`) VALUES
(1, 'Batu update 1', 'BT01', 10000, 80, 1, '2019-12-12 03:01:36', '2019-11-15 00:00:00'),
(3, 'Batu Kali U', 'BT002', 1234, 56, 0, '2019-12-12 08:50:45', '2019-12-12 15:50:45'),
(5, 'Batu', 'BT123', 2000, 50, 0, '2019-12-10 08:14:49', '2019-12-10 15:14:49'),
(7, 'Tissue a', 'TS01', 10000, 100, 1, '2019-11-27 13:38:33', '0000-00-00 00:00:00'),
(8, 'Kerupuk Update', 'KR01', 20000, 115, 0, '2019-12-12 10:09:50', '2019-12-12 17:09:50'),
(9, '', '', 0, 0, 1, '2019-11-27 13:03:49', '0000-00-00 00:00:00'),
(10, '', '', 0, 0, 1, '2019-11-27 13:06:50', '0000-00-00 00:00:00'),
(11, 'Aqua . edit', 'AQ01', 5000, 200, 0, '2019-11-27 13:28:57', '0000-00-00 00:00:00'),
(12, 'tess aaa', '22', 2222, 3333, 1, '2019-11-27 13:14:27', '0000-00-00 00:00:00'),
(13, 'tes 222', 'tes', 1, 2, 1, '2019-11-27 13:28:46', '0000-00-00 00:00:00'),
(14, 'werwr update', '35', 35, 3, 1, '2019-11-27 13:31:20', '0000-00-00 00:00:00'),
(15, 'Tissue', 'TS01', 4000, 33, 1, '2019-11-29 08:26:29', '0000-00-00 00:00:00'),
(16, 'aasd ee', 'we', 2, 2, 1, '2019-11-27 13:32:33', '0000-00-00 00:00:00'),
(17, 'Test', 'test', 100, 10, 1, '2019-11-29 04:18:28', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `stock`
--

CREATE TABLE `stock` (
  `id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `supply`
--

CREATE TABLE `supply` (
  `id` int(11) NOT NULL,
  `suplier_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `added_at` date NOT NULL,
  `updated_at` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`) VALUES
(1, 'hafiz', '$2a$04$doB2VvibSgPK8U1dDmJMq.8devn6r7IeDQ7r/PEwC8qePozXv3Lmi'),
(2, 'hafiz2', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `master_order`
--
ALTER TABLE `master_order`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `stock`
--
ALTER TABLE `stock`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `supply`
--
ALTER TABLE `supply`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customer`
--
ALTER TABLE `customer`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `master_order`
--
ALTER TABLE `master_order`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `product`
--
ALTER TABLE `product`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `stock`
--
ALTER TABLE `stock`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `supply`
--
ALTER TABLE `supply`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
