/* CREATE DATABASE */
CREATE DATABASE social;


USE social;
/*
https://countrycode.org/
*/

/* logic model: */
CREATE TABLE Country (
    ID INT AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL UNIQUE,
    country_code VARCHAR(9) NOT NULL UNIQUE,
    iso_code CHAR(2) NOT NULL UNIQUE,

    PRIMARY KEY (ID)
);

CREATE TABLE State (
    ID INT AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL ,
    countryID INT NOT NULL,

    PRIMARY KEY (ID),
    FOREIGN KEY (CountryID) REFERENCES Country(ID)
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
    email_personal VARCHAR(50) NOT NULL,
    email_work VARCHAR(50) NOT NULL,
    phone CHAR(13) NOT NULL,
    site VARCHAR(100),

    PRIMARY KEY(ID)
);


CREATE TABLE User (
    uid CHAR(32) ,
    nick VARCHAR(20) UNIQUE NOT NULL,
    img_perfil BLOB,
    First_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    gender VARCHAR(50),
    birth_date DATE NOT NULL,
    createdAt DATE,
    contactID INT,
    cityID INT,

    PRIMARY KEY(uid),
    FOREIGN KEY(cityID) REFERENCES City(ID),
    FOREIGN KEY(ContactID) REFERENCES Contact(ID)
);

CREATE TABLE Company (
    uid CHAR(32) NOT NULL,
    nick varchar(20) UNIQUE NOT NULL,
    img_perfil BLOB,
    fundation DATE NOT NULL,
    createdAt DATE,
    contactID INT,
    AddressID INT,

    PRIMARY KEY (uid),
    FOREIGN KEY (ContactID) REFERENCES Contact(ID),
    FOREIGN KEY (AddressID) REFERENCES Address(ID)
);
