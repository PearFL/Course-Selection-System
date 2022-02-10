-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: localhost    Database: course_select
-- ------------------------------------------------------
-- Server version	8.0.28-0ubuntu0.20.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `bind`
--

DROP TABLE IF EXISTS `bind`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `bind` (
  `teacher_id` varchar(255) NOT NULL,
  `course_id` varchar(255) NOT NULL,
  PRIMARY KEY (`teacher_id`,`course_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bind`
--

LOCK TABLES `bind` WRITE;
/*!40000 ALTER TABLE `bind` DISABLE KEYS */;
/*!40000 ALTER TABLE `bind` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `choice`
--

DROP TABLE IF EXISTS `choice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `choice` (
  `student_id` varchar(255) NOT NULL,
  `course_id` varchar(255) NOT NULL,
  PRIMARY KEY (`student_id`,`course_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `choice`
--

LOCK TABLES `choice` WRITE;
/*!40000 ALTER TABLE `choice` DISABLE KEYS */;
/*!40000 ALTER TABLE `choice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course`
--

DROP TABLE IF EXISTS `course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `course` (
  `course_id` varchar(255) NOT NULL,
  `teacher_id` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `capacity` int DEFAULT NULL,
  `cap_selected` int DEFAULT NULL,
  PRIMARY KEY (`course_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course`
--

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;
INSERT INTO `course` VALUES ('0ae16eff-f791-4872-8a1d-dd614f3e541e','9eb72094-69ab-4407-9caf-6b39b657704f','线性代数',20,0),('130b494c-3f3f-41df-9cc0-c53baa6b5a47','1bd82c07-71ee-4dc0-bc7b-36d4638eaaa5','高等数学',30,0),('3db62ab9-840d-496c-9fcb-ebccb7442a0f','9057eee5-f1ba-4ff7-a9a4-ceec62a0ca1d','低等数学',30,0),('74d65040-a2c8-4ad9-ad7c-f204c1cce158','e982ba88-e858-47ca-8a65-3263e73119b6','复变函数',120,0),('c9c088a9-2337-4f8e-bbe0-57f5437c56cd','ee33e324-3ec4-42fd-9e0f-3a79323a7768','信号系统',20,0),('d45a8627-f0f3-432c-8270-3d592efc47f8','89180547-a847-4541-95b1-2aa703d482e9','中等数学',20,0);
/*!40000 ALTER TABLE `course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `member`
--

DROP TABLE IF EXISTS `member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8 */;
CREATE TABLE `member` (
  `user_id` varchar(255) NOT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `user_type` int DEFAULT NULL,
  `is_deleted` tinyint(2) NOT NULL DEFAULT 0,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member`
--

LOCK TABLES `member` WRITE;
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
INSERT INTO `member` VALUES ('1bd82c07-71ee-4dc0-bc7b-36d4638eaaa5','yeye','hulu','fdf44dd58209f4f21104e92b850886ae',2,0),('2a4a0b95-dd0e-41eb-b56b-b3d5d50e539e','Alice','Alice','25d55ad283aa400af464c76d713c07ad',3,0),('81a47efe-65cd-4f1b-a789-68a538d1c39a','JudgeAdmin','JudgeAdmin','3f83fdd7e989c398ea157c58e3241dcd',1,0),('89180547-a847-4541-95b1-2aa703d482e9','icu','ui','5967ee71b80c6e3761b5414fd5b22dd7',2,0),('9057eee5-f1ba-4ff7-a9a4-ceec62a0ca1d','twt','twt','2ba38c72f45a830f2b8243a3a22236dd',2,0),('9eb72094-69ab-4407-9caf-6b39b657704f','wz','wz','36eed342177e32c235203aca93accde3',2,0),('e3551ad5-196d-42f7-8326-3687abda5f28','Tom','Tom','ddb14e5d1f19779c276e8d7b89ebd103',3,0),('e982ba88-e858-47ca-8a65-3263e73119b6','Bob','Bob','25d55ad283aa400af464c76d713c07ad',2,0),('ee33e324-3ec4-42fd-9e0f-3a79323a7768','mxh','mxh','4c3e4195b649599434c7ee3f1ed0d8d9',2,0);
/*!40000 ALTER TABLE `member` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-02-08  0:54:39
