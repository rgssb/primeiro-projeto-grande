CREATE DATABASE IF NOT EXISTS devebook;
USE devebook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    ID int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(555) not null,
    criadoEm timestamp default current_timestamp()
)ENGINE=INNODB;