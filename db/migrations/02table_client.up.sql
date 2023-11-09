CREATE TABLE db_client (
    account_id varchar(255) PRIMARY KEY NOT NULL,
    cpf varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    address_id varchar(255),
    CONSTRAINT fk_address FOREIGN KEY (address_id) REFERENCES db_address(address_id)
);