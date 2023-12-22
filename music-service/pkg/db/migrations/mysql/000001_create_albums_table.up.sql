CREATE TABLE IF NOT EXISTS Albums(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Name VARCHAR (127) NOT NULL,
    Artist VARCHAR (127) NOT NULL,
    Genre VARCHAR (127) NOT NULL,
    UploaderID INT NOT NULL,
    FOREIGN KEY (UploaderID) REFERENCES Users(ID),
    PRIMARY KEY (ID)
)