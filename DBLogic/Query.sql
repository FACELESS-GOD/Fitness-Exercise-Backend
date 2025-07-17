CREATE database AuthData_Test ; 
use AuthData_Test ; 

CREATE TABLE SEC_USER (
    UserName        varchar(255) NOT NULL,
	Password        varchar(255) NOT NULL,
	Email           varchar(255),
	FirstName       varchar(255),
	MiddleName      varchar(255),
	LastName        varchar(255),
	Designation     long NOT NULL,
	AuthorizationId  long NOT NULL,
    IsValid boolean NOT NULL,
    Last_Modified_Date varchar(255) Not null,
    Primary KEY (UserName)
);

CREATE TABLE Token_Store (	
    UserName        varchar(255) NOT NULL,	
	Token        varchar(255) NOT NULL,	
    Last_Modified_Date varchar(255) Not null,
    Primary KEY (UserName)
);

