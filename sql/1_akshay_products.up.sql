CREATE TABLE products.products (
	id BIGINT auto_increment NULL,
	name varchar(100) NULL,
	description varchar(100) NULL,
	price DOUBLE NULL,
	CONSTRAINT products_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci
AUTO_INCREMENT=1;
