CREATE TABLE bank(
    account_id varchar(255) NOT NULL,
    FOREIGN KEY (account_id) references client(account_id),
    account_id_to varchar(255) NOT NULL
    FOREIGN KEY (account_id_to) references client(account_id), 
    transactiontype varchar(255),
    transactiondate varchar(255),
    value int,
    balance int);