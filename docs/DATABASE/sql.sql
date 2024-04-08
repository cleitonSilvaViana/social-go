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
    number INT NOT NULL, /* número do imóvel */
    ZIP_CODE INT NOT NULL, /* CEP */

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

CREATE TABLE User (
    uid CHAR(36) ,
    nick VARCHAR(20) UNIQUE NOT NULL,
    img_perfil BLOB,
    First_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    gender VARCHAR(50),
    birth_date DATE NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    contactID INT NOT NULL,
    cityID INT,
    password VARCHAR(100) NOT NULL,

    PRIMARY KEY(uid),
    FOREIGN KEY(cityID) REFERENCES City(ID),
    FOREIGN KEY(ContactID) REFERENCES Contact(ID)
);

CREATE TABLE Company (
    uid CHAR(32) NOT NULL,
    nick varchar(20) UNIQUE NOT NULL,
    img_perfil BLOB,
    fundation DATE NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    contactID INT NOT NULL,
    AddressID INT NOT NULL,

    PRIMARY KEY (uid),
    FOREIGN KEY (ContactID) REFERENCES Contact(ID),
    FOREIGN KEY (AddressID) REFERENCES Address(ID)
);
