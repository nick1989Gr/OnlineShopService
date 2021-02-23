CREATE DATABASE `OnlineShop` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE `OnlineShop`.`items` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `manufacturer` VARCHAR(255) NOT NULL,
  `itemType` VARCHAR(60) NOT NULL,
  `price` DECIMAL(13,2) NOT NULL,
  `quantity` INT NOT NULL,
  PRIMARY KEY (`productId`));

INSERT INTO `OnlineShop`.`items` (`manufacturer`, `itemType`, `price`, `quantity`) VALUES ("Levis", "Jeans", 70, 4);  
INSERT INTO `OnlineShop`.`items` (`manufacturer`, `itemType`, `price`, `quantity`) VALUES ("Timberland", "hat", 20, 5);  
INSERT INTO `OnlineShop`.`items` (`manufacturer`, `itemType`, `price`, `quantity`) VALUES ("Diesel", "Jeans", 40, 5);    
INSERT INTO `OnlineShop`.`items` (`manufacturer`, `itemType`, `price`, `quantity`) VALUES ("Nike", "Shoes", 120, 4);  
INSERT INTO `OnlineShop`.`items` (`manufacturer`, `itemType`, `price`, `quantity`) VALUES ("Addidas", "Shoes", 110, 3);  
INSERT INTO `OnlineShop`.`items` (`manufacturer`, `itemType`, `price`, `quantity`) VALUES ("Rogue", "Sport Trousers", 50, 2);  
INSERT INTO `OnlineShop`.`items` (`manufacturer`, `itemType`, `price`, `quantity`) VALUES ("Lacoste", "T-Shirt", 50, 1);  
