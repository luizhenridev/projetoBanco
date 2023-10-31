CREATE TABLE client (
    account_id varchar(255) PRIMARY KEY NOT NULL,
    cpf varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    FOREIGN KEY (address) REFERENCES address
    );