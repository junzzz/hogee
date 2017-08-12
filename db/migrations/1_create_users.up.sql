CREATE TABLE users (
    id SERIAL,
    name text,
    address text,
    rank integer
);

INSERT INTO users(name, address, rank) VALUES('docker太郎', '北海道', 99);