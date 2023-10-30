CREATE TABLE bank(
    account_id varchar(255) PRIMARY KEY NOT NULL, 
    cpf varchar(255),
    name varchar(255), 
    email varchar(255),
    address varchar(255),
    transactiontype varchar(255),
    value int,
    currency varchar(255),
    transactiondate varchar(255),
    description varchar(255), 
    balance int, 
    status varchar(255), 
    message varchar(255));