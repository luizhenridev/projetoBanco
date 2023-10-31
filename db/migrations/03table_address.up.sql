CREATE table address (
    FOREIGN KEY (account_id) references client(account_id)
    street varchar(255),
    number int
);
