CREATE TABLE IF NOT EXISTS contact (
    id SERIAL PRIMARY KEY,
    name varchar(50),
    phone varchar(15)
);

INSERT INTO contact (name, phone) VALUES ('test', '+5511944445555');