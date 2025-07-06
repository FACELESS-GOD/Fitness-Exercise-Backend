CREATE database AuthData ; 

CREATE TABLE SEC_USER (
	UserID 			long NOT NULL auto_increment,
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
    Primary KEY (UserID)
);