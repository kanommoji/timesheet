-- MariaDB dump 10.17  Distrib 10.4.8-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: timesheet
-- ------------------------------------------------------
-- Server version	10.4.8-MariaDB-1:10.4.8+maria~bionic

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `incomes`
--

DROP TABLE IF EXISTS `incomes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `incomes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `member_id` varchar(3) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `month` int(11) DEFAULT NULL,
  `year` int(11) DEFAULT NULL,
  `day` int(11) DEFAULT NULL,
  `start_time_am` time DEFAULT NULL,
  `end_time_am` time DEFAULT NULL,
  `start_time_pm` time DEFAULT NULL,
  `end_time_pm` time DEFAULT NULL,
  `overtime` int(11) DEFAULT NULL,
  `total_hours` time DEFAULT NULL,
  `coaching_customer_charging` float DEFAULT NULL,
  `coaching_payment_rate` float DEFAULT NULL,
  `training_wage` float DEFAULT NULL,
  `other_wage` float DEFAULT NULL,
  `company` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `description` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incomes`
--

LOCK TABLES `incomes` WRITE;
/*!40000 ALTER TABLE `incomes` DISABLE KEYS */;
INSERT INTO `incomes` VALUES (1,'001',12,2018,3,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',15000,10000,0,0,'Siam Chamnankit','KBTG Coaching'),
(2,'001',12,2018,4,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,10000,'Siam Chamnankit','Siam Chamnankit and SHR operation'),
(3,'001',12,2018,6,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Shuhari','[IMC]GSB: Agile Project Mgmt'),
(4,'001',12,2018,7,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Shuhari','[IMC]GSB: Agile Project Mgmt'),
(5,'001',12,2018,11,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','[PTT-GC] 2 Days Agile Project Management'),
(6,'001',12,2018,12,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','[PTT-GC] 2 Days Agile Project Management'),
(7,'001',12,2018,13,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Shuhari','[IMC]GSB: Agile Project Mgmt'),
(8,'001',12,2018,14,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Shuhari','[IMC]GSB: Agile Project Mgmt'),
(9,'001',12,2018,15,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,10000,'Siam Chamnankit','Meeting Siam Chamnankit'),
(10,'001',12,2018,16,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','Day 1 of 6 days Agile for Software Development - Fujitsu'),
(11,'001',12,2018,17,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,10000,0,0,'Siam Chamnankit','KBTG Coaching'),
(12,'001',12,2018,18,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,10000,'Siam Chamnankit','Meeting with SW Park team + Meeting with Mobilfe on Bank Agent Platform'),
(13,'001',12,2018,19,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,10000,0,0,'Siam Chamnankit','[KBTG] 3 Days Agile Testing in Action'),
(14,'001',12,2018,20,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,10000,0,0,'Siam Chamnankit','[KBTG] 3 Days Agile Testing in Action'),
(15,'001',12,2018,21,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,10000,0,0,'Siam Chamnankit','[KBTG] 3 Days Agile Testing in Action'),
(16,'001',12,2018,24,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,15000,0,0,'Siam Chamnankit','TDEM - Coaching and feedback Fujitsu'),
(17,'001',12,2018,26,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,10000,'Siam Chamnankit','Meeting with Mobilfe on Bank Agent Platform + Meeting with Fujitsu MD'),
(18,'001',12,2018,27,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,10000,0,0,'Siam Chamnankit','[KBTG] 2 Days Agile Project Management'),
(19,'002',12,2018,11,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at Siam Chamnankit Dojo'),
(20,'002',12,2018,12,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at Siam Chamnankit Dojo'),
(21,'002',12,2018,13,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(22,'002',12,2018,14,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(23,'002',12,2018,15,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,5000,'Shuhari','ประชุมสรุปประจำปี 2018'),
(24,'002',12,2018,17,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(25,'002',12,2018,18,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(26,'002',12,2018,19,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(27,'002',12,2018,20,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(28,'002',12,2018,21,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(29,'002',12,2018,24,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(30,'002',12,2018,25,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(31,'002',12,2018,26,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(32,'002',12,2018,27,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(33,'003',12,2018,1,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,15000,0,'Siam Chamnankit','Technical Excellence at Khonkean'),
(34,'003',12,2018,2,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,15000,0,'Siam Chamnankit','Technical Excellence at Khonkean'),
(35,'003',12,2018,4,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','KBTG'),
(36,'003',12,2018,6,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','TDD with JAVA'),
(37,'003',12,2018,7,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','TDD with JAVA'),
(38,'003',12,2018,11,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','KBTG'),
(39,'003',12,2018,12,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','KBTG'),
(40,'003',12,2018,13,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','Docker and Kubernetes'),
(41,'003',12,2018,14,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','Docker and Kubernetes'),
(42,'003',12,2018,17,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,10000,0,'Siam Chamnankit','TDD with JAVA'),
(43,'005',12,2018,11,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at Siam Chamnankit Dojo'),
(44,'005',12,2018,12,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at Siam Chamnankit Dojo'),
(45,'005',12,2018,13,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(46,'005',12,2018,14,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(47,'005',12,2018,15,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,5000,'Shuhari','ประชุมสรุปประจำปี 2018'),
(48,'005',12,2018,17,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(49,'005',12,2018,18,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(50,'005',12,2018,19,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(51,'005',12,2018,20,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(52,'005',12,2018,21,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(53,'005',12,2018,24,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(54,'005',12,2018,25,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(55,'005',12,2018,26,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(56,'005',12,2018,27,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN'),
(57,'005',12,2018,28,'09:00:00','12:00:00','13:00:00','18:00:00',0,'08:00:00',0,0,0,0,'Shuhari','work at TN');
/*!40000 ALTER TABLE `incomes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `members`
--

DROP TABLE IF EXISTS `members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `members` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `member_id` varchar(3) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `company` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `member_name_th` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `member_name_eng` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `overtime_rate` float DEFAULT NULL,
  `rate_per_day` float DEFAULT NULL,
  `rate_per_hour` float DEFAULT NULL,
  `salary` float DEFAULT NULL,
  `income_tax_1` float DEFAULT NULL,
  `income_tax_53_percentage` int(11) DEFAULT NULL,
  `social_security` float DEFAULT NULL,
  `status` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `travel_expense` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `members`
--

LOCK TABLES `members` WRITE;
/*!40000 ALTER TABLE `members` DISABLE KEYS */;
INSERT INTO `members` VALUES (1,'001','Siam Chamnankit','ประธาน ด่านสกุลเจริญกิจ','Prathan Dansakulcharoenkit','prathan@scrum123.com',0,15000,1875,80000,5000,10,0,NULL,0),
(2,'001','Shuhari','ประธาน ด่านสกุลเจริญกิจ','Prathan Dansakulcharoenkit','prathan@scrum123.com',0,15000,1875,0,0,10,0,NULL,0),
(3,'002','Shuhari','นารีนารถ เนรัญชร','Nareenart Narunchon','nareenart@scrum123.com',0,0,0,25000,0,5,750,'salary',1500),
(4,'003','Siam Chamnankit','สมเกียรติ ปุ๋ยสูงเนิน','Somkiat Puisungnoen','somkiat@scrum123.com',0,15000,1875,15000,0,10,750,NULL,0),
(5,'003','Shuhari','สมเกียรติ ปุ๋ยสูงเนิน','Somkiat Puisungnoen','somkiat@scrum123.com',0,15000,1875,40000,5000,10,0,NULL,0),
(6,'004','Siam Chamnankit','ธวัชชัย จงสุวรรณไพศาล','Thawatchai Jongsuwanpaisan','thawatchai@scrum123.com',0,15000,1875,40000,5000,10,0,NULL,0),
(7,'004','Shuhari','ธวัชชัย จงสุวรรณไพศาล','Thawatchai Jongsuwanpaisan','thawatchai@scrum123.com',0,15000,1875,0,0,10,0,NULL,0),
(8,'005','Shuhari','อภิพล สุขเกลอ','Apipol sukgler','golf.apipol@scrum123.com',0,0,0,40000,1200,5,750,'salary',1500),
(9,'006','Shuhari','ภาณุมาศ แสนโท','Panumars Seanto','panumars@scrum123.com',0,0,0,25000,0,5,750,'salary',1500);
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `timesheets`
--

DROP TABLE IF EXISTS `timesheets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `timesheets` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `timesheets_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `member_id` varchar(3) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `month` int(11) DEFAULT NULL,
  `year` int(11) DEFAULT NULL,
  `total_hours` time DEFAULT NULL,
  `total_coaching_customer_charging` float DEFAULT NULL,
  `total_coaching_payment_rate` float DEFAULT NULL,
  `total_training_wage` float DEFAULT NULL,
  `total_other_wage` float DEFAULT NULL,
  `payment_wage` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `timesheets`
--

LOCK TABLES `timesheets` WRITE;
/*!40000 ALTER TABLE `timesheets` DISABLE KEYS */;
INSERT INTO `timesheets` VALUES (1,'003201712','003',12,2017,'88:00:00',0,0,120000,0,120000),
(2,'006201812','006',12,2018,'00:00:00',0,0,0,0,0);
/*!40000 ALTER TABLE `timesheets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `member_id` varchar(3) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `month` int(11) DEFAULT NULL,
  `year` int(11) DEFAULT NULL,
  `company` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `member_name_th` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `coaching` float DEFAULT NULL,
  `training` float DEFAULT NULL,
  `other` float DEFAULT NULL,
  `total_incomes` float DEFAULT NULL,
  `salary` float DEFAULT NULL,
  `income_tax_1` float DEFAULT NULL,
  `social_security` float DEFAULT NULL,
  `net_salary` float DEFAULT NULL,
  `wage` float DEFAULT NULL,
  `income_tax_53_percentage` int(11) DEFAULT NULL,
  `income_tax_53` float DEFAULT NULL,
  `net_wage` float DEFAULT NULL,
  `net_transfer` float DEFAULT NULL,
  `status_checking_transfer` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `date_transfer` date DEFAULT NULL,
  `comment` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (1,'001',12,2017,'Siam Chamnankit','ประธาน ด่านสกุลเจริญกิจ',85000,30000,40000,155000,80000,5000,0,75000,75000,10,7500,67500,142500,'รอการตรวจสอบ',NULL,NULL),
(2,'004',12,2018,'Siam Chamnankit','ธวัชชัย จงสุวรรณไพศาล',50000,70000,10000,130000,40000,5000,0,35000,90000,10,9000,81000,116000,'รอการตรวจสอบ',NULL,NULL);
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-10-12 15:26:57
