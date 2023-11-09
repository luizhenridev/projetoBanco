CREATE TABLE db_bank(
    transaction_id varchar(255) PRIMARY KEY NOT NULL,
    account_id varchar(255),
    CONSTRAINT fk_account_id FOREIGN KEY (account_id) references db_client(account_id),
    account_id_to varchar(255),
    CONSTRAINT fk_account_id_to FOREIGN KEY (account_id_to) references db_client(account_id), 
    transactiontype varchar(255),
    transactiondate varchar(255),
    value int,
    balance int);