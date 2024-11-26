-- Active: 1725978681310@@127.0.0.1@5432@shop@public
CREATE TABLE Products (
  id int PRIMARY KEY,
  name varchar,
  cost decimal,
  description varchar(255),
  image_path varchar(255),
  сount int,
  manufacturer_id int,
  review_id int
);

CREATE TABLE Categories (
  id int PRIMARY KEY,
  name varchar(30)
);

CREATE TABLE ProductsCaterogries (
  product_id int REFERENCES Products(id) ON DELETE CASCADE,
  category_id int REFERENCES Categories(id) ON DELETE CASCADE
);

CREATE TABLE Users (
  id int PRIMARY KEY,
  username varchar(30),
  password varchar(30)  
);

CREATE TABLE ProductReviews (
  id int PRIMARY KEY,
  text varchar(100),
  user_id int REFERENCES Users(id) ON DELETE CASCADE
);

CREATE TABLE Manifacturers (
  id int PRIMARY KEY,
  name varchar(30),
  description varchar(255),
  address varchar(100),
  contacts varchar(100)
);

CREATE TABLE AccessLevels (
  id int PRIMARY KEY,
  access_level int DEFAULT 0 NOT NULL
);

CREATE TABLE Workers (
  id int PRIMARY KEY,
  login varchar(20) NOT NULL,
  password varchar(20) NOT NULL, 
  access_id int DEFAULT 0 REFERENCES AccessLevels(id) ON DELETE SET DEFAULT
);



CREATE TABLE DeliveryAddresses (
  id int PRIMARY KEY,
  user_id INT REFERENCES Users(id) ON DELETE CASCADE,
  address varchar(100) NOT NULL
);

CREATE TABLE OrderStatus (
  id int PRIMARY KEY,
  text varchar(100)
);

CREATE TABLE Orders (
  id int PRIMARY KEY,
  create_date date NOT NULL,
  delivery_date date NOT NULL,
  total_cost decimal NOT NULL,
  user_id int REFERENCES Users(id) ON DELETE SET NULL,
  delivery_address_id int REFERENCES DeliveryAddresses(id) ON DELETE SET NULL,
  status_id int REFERENCES OrderStatus(id) ON DELETE SET NULL
);

CREATE TABLE OrdersProducts (
  order_id int REFERENCES Orders(id) ON DELETE CASCADE ,
  product_id int REFERENCES Products(id) ON DELETE CASCADE
);

CREATE TABLE Cart (
  user_id int REFERENCES Users(id),
  product_id int REFERENCES Products(id)
);

-- Категории (3 категории)
INSERT INTO Categories (id, name) VALUES 
(1, 'Процессоры'),
(2, 'Видеокарты'),
(3, 'Материнские платы');

-- Производители (3 производителя)
INSERT INTO Manifacturers (id, name, description, address, contacts) VALUES 
(1, 'Intel', 'Производитель процессоров', 'США', 'intel@example.com'),
(2, 'NVIDIA', 'Производитель видеокарт', 'США', 'nvidia@example.com'),
(3, 'ASUS', 'Производитель материнских плат', 'Тайвань', 'asus@example.com');

-- Товары (5 товаров в каждой категории)
-- Категория: Процессоры
INSERT INTO Products (id, name, cost, description, image_path, сount, manufacturer_id) VALUES 
(1, 'Intel Core i5-12400F', 15000.00, '6 ядер, 12 потоков, 2.5 ГГц', 'path/to/image1.jpg', 100, 1),
(2, 'Intel Core i7-13700K', 35000.00, '16 ядер, 24 потока, 3.4 ГГц', 'path/to/image2.jpg', 50, 1),
(3, 'Intel Core i9-13900K', 55000.00, '24 ядра, 32 потока, 3.0 ГГц', 'path/to/image3.jpg', 30, 1),
(4, 'AMD Ryzen 5 5600X', 12000.00, '6 ядер, 12 потоков, 3.7 ГГц', 'path/to/image4.jpg', 75, 1),
(5, 'AMD Ryzen 7 5800X', 20000.00, '8 ядер, 16 потоков, 3.8 ГГц', 'path/to/image5.jpg', 60, 1);

-- Категория: Видеокарты
INSERT INTO Products (id, name, cost, description, image_path, сount, manufacturer_id) VALUES 
(6, 'NVIDIA RTX 3060', 30000.00, '12 ГБ GDDR6, 3584 CUDA', 'path/to/image6.jpg', 40, 2),
(7, 'NVIDIA RTX 3070', 45000.00, '8 ГБ GDDR6, 5888 CUDA', 'path/to/image7.jpg', 30, 2),
(8, 'NVIDIA RTX 3080', 70000.00, '10 ГБ GDDR6X, 8704 CUDA', 'path/to/image8.jpg', 20, 2),
(9, 'AMD RX 6700 XT', 35000.00, '12 ГБ GDDR6', 'path/to/image9.jpg', 35, 2),
(10, 'AMD RX 6800 XT', 55000.00, '16 ГБ GDDR6', 'path/to/image10.jpg', 25, 2);

-- Категория: Материнские платы
INSERT INTO Products (id, name, cost, description, image_path, сount, manufacturer_id) VALUES 
(11, 'ASUS ROG Strix B550-F', 15000.00, 'ATX, AM4, PCIe 4.0', 'path/to/image11.jpg', 50, 3),
(12, 'ASUS TUF Gaming X570', 20000.00, 'ATX, AM4, PCIe 4.0', 'path/to/image12.jpg', 40, 3),
(13, 'MSI MPG Z590 Gaming Edge', 18000.00, 'ATX, LGA1200, PCIe 4.0', 'path/to/image13.jpg', 35, 3),
(14, 'Gigabyte B450M DS3H', 8000.00, 'Micro-ATX, AM4', 'path/to/image14.jpg', 60, 3),
(15, 'ASRock B550 Pro4', 9000.00, 'ATX, AM4', 'path/to/image15.jpg', 55, 3);

-- Связь товаров и категорий
INSERT INTO ProductsCaterogries (product_id, category_id) VALUES 
(1, 1), (2, 1), (3, 1), (4, 1), (5, 1),  -- Процессоры
(6, 2), (7, 2), (8, 2), (9, 2), (10, 2), -- Видеокарты
(11, 3), (12, 3), (13, 3), (14, 3), (15, 3); -- Материнские платы

-- Пользователи
INSERT INTO Users (id, username, password) VALUES 
(1, 'user1', 'password1'),
(2, 'user2', 'password2'),
(3, 'user3', 'password3');

-- Адреса доставки
INSERT INTO DeliveryAddresses (id, user_id, address) VALUES 
(1, 1, 'ул. Ленина, д. 10, кв. 5'),
(2, 2, 'ул. Пушкина, д. 20, кв. 15'),
(3, 3, 'ул. Советская, д. 30, кв. 25');

-- Статусы заказов
INSERT INTO OrderStatus (id, text) VALUES 
(1, 'Оформлен'),
(2, 'Отправлен'),
(3, 'Доставлен');

-- Заказы
INSERT INTO Orders (id, create_date, delivery_date, total_cost, user_id, delivery_address_id, status_id) VALUES 
(1, '2024-11-01', '2024-11-05', 75000.00, 1, 1, 1),
(2, '2024-11-10', '2024-11-15', 35000.00, 2, 2, 2);

-- Продукты в заказах
INSERT INTO OrdersProducts (order_id, product_id) VALUES 
(1, 3), (1, 8),  -- Заказ 1: процессор и видеокарта
(2, 7), (2, 12); -- Заказ 2: видеокарта и материнская плата

-- Корзина
INSERT INTO Cart (user_id, product_id) VALUES 
(1, 5),
(2, 9),
(3, 15);

