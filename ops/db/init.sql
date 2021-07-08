-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Hôte : localhost
-- Généré le : mar. 06 juil. 2021 à 15:47
-- Version du serveur : 10.4.19-MariaDB
-- Version de PHP : 7.4.20

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `test`
--

-- --------------------------------------------------------

--
-- Structure de la table `publicity`
--

CREATE TABLE `publicity` (
                             `id` varchar(50) NOT NULL,
                             `ua` varchar(255) NOT NULL,
                             `ip` varchar(255) NOT NULL,
                             `os` varchar(255) NOT NULL,
                             `created_at` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Structure de la table `publicity_type`
--

CREATE TABLE `publicity_type` (
                                  `publicity_id` varchar(50) NOT NULL,
                                  `type_id` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- --------------------------------------------------------

--
-- Structure de la table `type`
--

CREATE TABLE `type` (
                        `id` varchar(50) NOT NULL,
                        `type` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Déchargement des données de la table `type`
--

INSERT INTO `type` (`id`, `type`) VALUES
('3d3f859a-de60-11eb-ba80-0242ac130004', 'visible'),
('4472b29c-de60-11eb-ba80-0242ac130004', 'click'),
('49993e76-de60-11eb-ba80-0242ac130004', 'enum_impression');

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `publicity`
--
ALTER TABLE `publicity`
    ADD PRIMARY KEY (`id`);

--
-- Index pour la table `type`
--
ALTER TABLE `type`
    ADD PRIMARY KEY (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
