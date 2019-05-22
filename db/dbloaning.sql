-- phpMyAdmin SQL Dump
-- version 4.8.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 22, 2019 at 12:09 PM
-- Server version: 10.1.33-MariaDB
-- PHP Version: 7.2.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `dbloaning`
--

-- --------------------------------------------------------

--
-- Table structure for table `dates`
--

CREATE TABLE `dates` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `date_value` varchar(255) DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `ktp_details`
--

CREATE TABLE `ktp_details` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `provincial_id` int(10) UNSIGNED DEFAULT NULL,
  `date_id` int(10) UNSIGNED DEFAULT NULL,
  `month_id` int(10) UNSIGNED DEFAULT NULL,
  `sub_district_id` int(10) UNSIGNED DEFAULT NULL,
  `year_id` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `ktp_details`
--

INSERT INTO `ktp_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `provincial_id`, `date_id`, `month_id`, `sub_district_id`, `year_id`) VALUES
(1, '2019-05-21 10:02:15', '2019-05-21 10:02:15', NULL, 0, 0, 0, 0, 0),
(1112345678, '2019-05-22 08:32:45', '2019-05-22 08:32:45', NULL, 11, 34, 56, 12, 78),
(1234567890, '2019-05-21 10:03:43', '2019-05-21 10:03:43', NULL, 90, 56, 34, 78, 12),
(4294967295, '2019-05-22 07:28:49', '2019-05-22 07:28:49', NULL, 98, 54, 32, 76, 33);

-- --------------------------------------------------------

--
-- Table structure for table `loanings`
--

CREATE TABLE `loanings` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `date` varchar(255) DEFAULT NULL,
  `ktp_detail_id` int(10) UNSIGNED DEFAULT NULL,
  `birth_date` varchar(255) DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `period` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `loanings`
--

INSERT INTO `loanings` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `date`, `ktp_detail_id`, `birth_date`, `gender`, `amount`, `period`) VALUES
(1, '2019-05-21 10:02:15', '2019-05-21 10:02:15', '2019-05-21 10:31:30', 'Bayu', '2019-05-08', 0, '', 'Pria', 20000, 12),
(2, '2019-05-21 10:03:42', '2019-05-22 08:32:04', NULL, 'Bayu', '2019-05-08', 1111111106, '2012-09-10', 'Pria', 20000, 12),
(3, '2019-05-21 15:54:43', '2019-05-21 15:54:43', '2019-05-21 15:55:51', 'bayu', '2019-05-22', 1234567890, '2019-05-25', 'laki', 123455, 1234),
(4, '2019-05-22 07:28:49', '2019-05-22 07:34:45', '2019-05-22 08:33:38', 'Ruins', '2019-05-30', 4294967200, '2019-05-30', 'Laki', 1234567, 3),
(5, '2019-05-22 08:32:44', '2019-05-22 08:32:44', '2019-05-22 08:33:26', 'Haha', '2019-05-23', 1112345678, '2019-05-31', 'Laki', 9, 2);

-- --------------------------------------------------------

--
-- Table structure for table `months`
--

CREATE TABLE `months` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name_month` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `provincials`
--

CREATE TABLE `provincials` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name_provincial` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `sub_districts`
--

CREATE TABLE `sub_districts` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name_sub_district` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `years`
--

CREATE TABLE `years` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name_year` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `dates`
--
ALTER TABLE `dates`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_dates_deleted_at` (`deleted_at`);

--
-- Indexes for table `ktp_details`
--
ALTER TABLE `ktp_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_ktp_details_deleted_at` (`deleted_at`);

--
-- Indexes for table `loanings`
--
ALTER TABLE `loanings`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_loanings_deleted_at` (`deleted_at`);

--
-- Indexes for table `months`
--
ALTER TABLE `months`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_months_deleted_at` (`deleted_at`);

--
-- Indexes for table `provincials`
--
ALTER TABLE `provincials`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_provincials_deleted_at` (`deleted_at`);

--
-- Indexes for table `sub_districts`
--
ALTER TABLE `sub_districts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_sub_districts_deleted_at` (`deleted_at`);

--
-- Indexes for table `years`
--
ALTER TABLE `years`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_years_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `dates`
--
ALTER TABLE `dates`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `ktp_details`
--
ALTER TABLE `ktp_details`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2147483647;

--
-- AUTO_INCREMENT for table `loanings`
--
ALTER TABLE `loanings`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `months`
--
ALTER TABLE `months`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `provincials`
--
ALTER TABLE `provincials`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `sub_districts`
--
ALTER TABLE `sub_districts`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `years`
--
ALTER TABLE `years`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
