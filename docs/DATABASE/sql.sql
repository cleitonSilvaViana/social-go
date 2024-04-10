/* CREATE DATABASE */
CREATE DATABASE social;


USE social;
/*
https://countrycode.org/
*/

/* logic model: */
CREATE TABLE Country (
    CCA3 CHAR(3) NOT NULL UNIQUE,
    name VARCHAR(20) NOT NULL UNIQUE,

    PRIMARY KEY (cca3)
);

CREATE TABLE State (
    ID INT AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL ,
    countryID CHAR(3) NOT NULL,

    PRIMARY KEY (ID),
    FOREIGN KEY (CountryID) REFERENCES Country(CCA3)
);

CREATE TABLE City (
    ID INT AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    stateID INT NOT NULL,

    PRIMARY KEY(ID),
    FOREIGN KEY (stateID) REFERENCES State(ID)
);

CREATE TABLE Address (
    ID INT NOT NULL AUTO_INCREMENT,
    cityID INT NOT NULL,
    public_place VARCHAR(50), /* logradouro */
    number INT, /* número do imóvel */
    ZIP_CODE INT, /* CEP */

    PRIMARY KEY(ID),
    FOREIGN KEY (cityID) REFERENCES City(ID)
);

CREATE TABLE Contact (
    ID INT AUTO_INCREMENT,
    email VARCHAR(50) NOT NULL UNIQUE,
    phone CHAR(13) UNIQUE,
    site VARCHAR(100) UNIQUE,

    PRIMARY KEY(ID)
);

CREATE TABLE Profile {
    uid char(36),
    img_perfil BLOB,
    nick VARCHAR(20) UNIQUE NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    password VARCHAR(100) NOT NULL,
    type char /* 1 - user, 0 - company */

    contactID INT NOT NULL,
    AddressID INT NOT NULL,

    PRIMARY KEY(uid)
    FOREIGN KEY(cityID) REFERENCES City(ID),
    FOREIGN KEY(ContactID) REFERENCES Contact(ID)
}

CREATE TABLE User (
    uid char(36) NOT NULL,
    First_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    gender VARCHAR(50),
    birth_date DATE NOT NULL,
    
    PRIMARY KEY(uid)
);

CREATE TABLE Company (
    uid char(36) NOT NULL,
    fundation DATE NOT NULL,

    PRIMARY KEY(uid)
);
