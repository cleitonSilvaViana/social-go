/* CREATE DATABASE */

/* logic model: */
CREATE TABLE Country (
    ID INT AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL UNIQUE,
    sigla VARCHAR(3) NOT NULL UNIQUE,
    code INT NOT NULL UNIQUE,
    language VARCHAR(20) NOT NULL,

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

CREATE TABLE Contact (
    ID INT AUTO_INCREMENT,
    email VARCHAR(50) NOT NULL,
    phone CHAR(32) NOT NULL,
    site VARCHAR(100),

    PRIMARY KEY(ID)
);



CREATE TABLE User (
    uid CHAR(32) ,
    nick VARCHAR(20) UNIQUE NOT NULL,
    img_perfil BLOB,
    First_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
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
    uid CHAR(32),
    nick varchar(20) UNIQUE NOT NULL,
    img_perfil BLOB,
    document VARCHAR(50),
    slogan VARCHAR(100),
    description TEXT NOT NULL,
    fundation DATE,
    public_place VARCHAR(50),
    contactID INT,
    cityID INT,
    createdAt DATE,

    PRIMARY KEY (uid),
    FOREIGN KEY(ContactID) REFERENCES Contact(ID)
);

/*****  FALTA ESTA TABELA ******/
CREATE TABLE Group (
    ID INT,
    creatorID CHAR(32) NOT NULL,
    description VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,

    PRIMARY KEY(ID),
    FOREIGN KEY(creatorID) REFERENCES User(UID)
);


CREATE TABLE UserGroup (
    userID CHAR(32) NOT NULL UNIQUE,
    groupID INT NOT NULL UNIQUE,

    CONSTRAINT PK_UserGroup PRIMARY KEY (userID, groupID)
);
