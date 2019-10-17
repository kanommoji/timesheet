-- MariaDB dump 10.17  Distrib 10.4.7-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: timesheet
-- ------------------------------------------------------
-- Server version	10.4.7-MariaDB-1:10.4.7+maria~bionic

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
  `start_time_am_hours` int(11) DEFAULT NULL,
  `start_time_am_minutes` int(11) DEFAULT NULL,
  `start_time_am_seconds` int(11) DEFAULT NULL,
  `end_time_am_hours` int(11) DEFAULT NULL,
  `end_time_am_minutes` int(11) DEFAULT NULL,
  `end_time_am_seconds` int(11) DEFAULT NULL,
  `start_time_pm_hours` int(11) DEFAULT NULL,
  `start_time_pm_minutes` int(11) DEFAULT NULL,
  `start_time_pm_seconds` int(11) DEFAULT NULL,
  `end_time_pm_hours` int(11) DEFAULT NULL,
  `end_time_pm_minutes` int(11) DEFAULT NULL,
  `end_time_pm_seconds` int(11) DEFAULT NULL,
  `overtime` int(11) DEFAULT NULL,
  `total_hours_hours` int(11) DEFAULT NULL,
  `total_hours_minutes` int(11) DEFAULT NULL,
  `total_hours_seconds` int(11) DEFAULT NULL,
  `coaching_customer_charging` float DEFAULT NULL,
  `coaching_payment_rate` float DEFAULT NULL,
  `training_wage` float DEFAULT NULL,
  `other_wage` float DEFAULT NULL,
  `company` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `description` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incomes`
--

LOCK TABLES `incomes` WRITE;
/*!40000 ALTER TABLE `incomes` DISABLE KEYS */;
INSERT INTO `incomes` VALUES (1,'001',12,2018,3,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,15000,10000,0,0,'siam_chamnankit','KBTG Coaching'),(2,'001',12,2018,4,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,10000,'siam_chamnankit','siam_chamnankit and SHR operation'),(3,'001',12,2018,6,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'shuhari','[IMC]GSB: Agile Project Mgmt'),(4,'001',12,2018,7,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'shuhari','[IMC]GSB: Agile Project Mgmt'),(5,'001',12,2018,11,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','[PTT-GC] 2 Days Agile Project Management'),(6,'001',12,2018,12,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','[PTT-GC] 2 Days Agile Project Management'),(7,'001',12,2018,13,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'shuhari','[IMC]GSB: Agile Project Mgmt'),(8,'001',12,2018,14,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'shuhari','[IMC]GSB: Agile Project Mgmt'),(9,'001',12,2018,15,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,10000,'siam_chamnankit','Meeting siam_chamnankit'),(10,'001',12,2018,16,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','Day 1 of 6 days Agile for Software Development - Fujitsu'),(11,'001',12,2018,17,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,10000,0,0,'siam_chamnankit','KBTG Coaching'),(12,'001',12,2018,18,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,10000,'siam_chamnankit','Meeting with SW Park team + Meeting with Mobilfe on Bank Agent Platform'),(13,'001',12,2018,19,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,10000,0,0,'siam_chamnankit','[KBTG] 3 Days Agile Testing in Action'),(14,'001',12,2018,20,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,10000,0,0,'siam_chamnankit','[KBTG] 3 Days Agile Testing in Action'),(15,'001',12,2018,21,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,10000,0,0,'siam_chamnankit','[KBTG] 3 Days Agile Testing in Action'),(16,'001',12,2018,24,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,15000,0,0,'siam_chamnankit','TDEM - Coaching and feedback Fujitsu'),(17,'001',12,2018,26,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,10000,'siam_chamnankit','Meeting with Mobilfe on Bank Agent Platform + Meeting with Fujitsu MD'),(18,'001',12,2018,27,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,10000,0,0,'siam_chamnankit','[KBTG] 2 Days Agile Project Management'),(19,'002',12,2018,11,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at siam_chamnankit Dojo'),(20,'002',12,2018,12,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at siam_chamnankit Dojo'),(21,'002',12,2018,13,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(22,'002',12,2018,14,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(23,'002',12,2018,15,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,5000,'shuhari','ประชุมสรุปประจำปี 2018'),(24,'002',12,2018,17,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(25,'002',12,2018,18,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(26,'002',12,2018,19,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(27,'002',12,2018,20,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(28,'002',12,2018,21,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(29,'002',12,2018,24,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(30,'002',12,2018,25,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(31,'002',12,2018,26,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(32,'002',12,2018,27,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(33,'003',12,2018,1,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,15000,0,'siam_chamnankit','Technical Excellence at Khonkean'),(34,'003',12,2018,2,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,15000,0,'siam_chamnankit','Technical Excellence at Khonkean'),(35,'003',12,2018,4,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','KBTG'),(36,'003',12,2018,6,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','TDD with JAVA'),(37,'003',12,2018,7,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','TDD with JAVA'),(38,'003',12,2018,11,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','KBTG'),(39,'003',12,2018,12,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','KBTG'),(40,'003',12,2018,13,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','Docker and Kubernetes'),(41,'003',12,2018,14,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','Docker and Kubernetes'),(42,'003',12,2018,17,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,10000,0,'siam_chamnankit','TDD with JAVA'),(43,'005',12,2018,11,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at siam_chamnankit Dojo'),(44,'005',12,2018,12,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at siam_chamnankit Dojo'),(45,'005',12,2018,13,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(46,'005',12,2018,14,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(47,'005',12,2018,15,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,5000,'shuhari','ประชุมสรุปประจำปี 2018'),(48,'005',12,2018,17,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(49,'005',12,2018,18,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(50,'005',12,2018,19,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(51,'005',12,2018,20,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(52,'005',12,2018,21,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(53,'005',12,2018,24,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(54,'005',12,2018,25,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(55,'005',12,2018,26,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(56,'005',12,2018,27,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(57,'005',12,2018,28,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(58,'006',12,2019,11,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN'),(59,'006',12,2019,12,9,0,0,12,0,0,13,0,0,18,0,0,0,8,0,0,0,0,0,0,'shuhari','work at TN');
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
INSERT INTO `members` VALUES (1,'001','siam_chamnankit','ประธาน ด่านสกุลเจริญกิจ','Prathan Dansakulcharoenkit','prathan@scrum123.com',0,15000,1875,80000,5000,10,0,NULL,0),(2,'001','shuhari','ประธาน ด่านสกุลเจริญกิจ','Prathan Dansakulcharoenkit','prathan@scrum123.com',0,15000,1875,0,0,10,0,NULL,0),(3,'002','shuhari','นารีนารถ เนรัญชร','Nareenart Narunchon','nareenart@scrum123.com',0,0,0,25000,0,5,750,'salary',1500),(4,'003','siam_chamnankit','สมเกียรติ ปุ๋ยสูงเนิน','Somkiat Puisungnoen','somkiat@scrum123.com',0,15000,1875,15000,0,10,750,NULL,0),(5,'003','shuhari','สมเกียรติ ปุ๋ยสูงเนิน','Somkiat Puisungnoen','somkiat@scrum123.com',0,15000,1875,40000,5000,10,0,NULL,0),(6,'004','siam_chamnankit','ธวัชชัย จงสุวรรณไพศาล','Thawatchai Jongsuwanpaisan','thawatchai@scrum123.com',0,15000,1875,40000,5000,10,0,NULL,0),(7,'004','shuhari','ธวัชชัย จงสุวรรณไพศาล','Thawatchai Jongsuwanpaisan','thawatchai@scrum123.com',0,15000,1875,0,0,10,0,NULL,0),(8,'005','shuhari','อภิพล สุขเกลอ','Apipol sukgler','golf.apipol@scrum123.com',0,0,0,40000,1200,5,750,'salary',1500),(9,'006','shuhari','ภาณุมาศ แสนโท','Panumars Seanto','panumars@scrum123.com',0,0,0,25000,0,5,750,'salary',1500);
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `timesheets`
--

DROP TABLE IF EXISTS `timesheets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `timesheets` (
  `id` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `member_id` varchar(3) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `month` int(11) DEFAULT NULL,
  `year` int(11) DEFAULT NULL,
  `total_hours_hours` int(11) DEFAULT NULL,
  `total_hours_minutes` int(11) DEFAULT NULL,
  `total_hours_seconds` int(11) DEFAULT NULL,
  `total_coaching_customer_charging` float DEFAULT NULL,
  `total_coaching_payment_rate` float DEFAULT NULL,
  `total_training_wage` float DEFAULT NULL,
  `total_other_wage` float DEFAULT NULL,
  `payment_wage` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `timesheets`
--

LOCK TABLES `timesheets` WRITE;
/*!40000 ALTER TABLE `timesheets` DISABLE KEYS */;
INSERT INTO `timesheets` VALUES ('003201712','003',12,2017,88,0,0,0,0,120000,0,120000),('006201812','006',12,2018,0,0,0,0,0,0,0,0),('007201911','007',12,2019,0,0,0,0,0,0,0,0),('007201912','007',12,2019,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `timesheets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactions` (
  `id` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
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
  `date_transfer` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `comment` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES ('001201712siam_chamnankit','001',12,2017,'siam_chamnankit','ประธาน ด่านสกุลเจริญกิจ',85000,30000,40000,155000,80000,5000,0,75000,75000,10,7500,67500,142500,'รอการตรวจสอบ',NULL,NULL),('001201812shuhari','001',12,2018,'shuhari','ประธาน ด่านสกุลเจริญกิจ',0,40000,0,40000,0,0,0,0,40000,10,4000,36000,36000,'รอการตรวจสอบ',NULL,NULL),('001201812siam_chamnankit','001',12,2018,'siam_chamnankit','ประธาน ด่านสกุลเจริญกิจ',75000,30000,0,145000,80000,5000,0,75000,65000,10,6500,58500,133500,'รอการตรวจสอบ',NULL,NULL),('001201912shuhari','001',12,2019,'shuhari','ประธาน ด่านสกุลเจริญกิจ',0,0,0,0,0,0,0,0,0,0,0,0,0,'รอการตรวจสอบ',NULL,NULL),('001201912siam_chamnankit','001',12,2019,'siam_chamnankit','ประธาน ด่านสกุลเจริญกิจ',0,0,0,0,0,0,0,0,0,0,0,0,0,'รอการตรวจสอบ',NULL,NULL),('002201812shuhari','002',12,2018,'shuhari','นารีนารถ เนรัญชร',0,0,6500,6500,25000,0,750,24250,6500,5,325,6175,30425,'รอการตรวจสอบ',NULL,NULL),('004201718siam_chamnankit','004',12,2018,'siam_chamnankit','ธวัชชัย จงสุวรรณไพศาล',50000,70000,10000,130000,40000,5000,0,35000,90000,10,9000,81000,116000,'รอการตรวจสอบ',NULL,NULL),('006201912shuhari','006',12,2019,'shuhari','ภาณุมาศ แสนโท',0,0,6500,6500,25000,0,750,24250,6500,5,325,6175,30425,'รอการตรวจสอบ','','');
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

-- Dump completed on 2019-10-17 11:17:18
