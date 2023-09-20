-- 
CREATE TABLE User (
    id              BIGINT  NOT NULL AUTO_INCREMENT,
    username        NVARCHAR(50)     NOT NULL,
    email           NVARCHAR(100)    NOT NULL,
    name            NVARCHAR(50),     
    family          NVARCHAR(50),
    password        NVARCHAR(100),
    role            NVARCHAR(20),
    image           TEXT,
    phone_number    INTEGER,
    created_at      DATETIME,
    PRIMARY KEY (id),
);
/************************************************/
CREATE TABLE Post (
    id              BIGINT NOT NULL AUTO_INCREMENT,
    title           NVARCHAR(200),    
	content            TEXT,    
	is_published    BOOLEAN,      
	created_at      DATETIME,
	published_at    DATETIME,
	userID          BIGINT,
    PRIMARY KEY (id),
);
CREATE TABLE Comment (
    id              BIGINT NOT NULL AUTO_INCREMENT,
    name            NVARCHAR(50),  
	email           NVARCHAR(100),   
	content         TEXT,    
	is_published    BOOLEAN,      
	published_at    DATETIME, 
	created_at      DATETIME, 
	PostID          BIGINT,
    PRIMARY KEY (id),
    FOREIGN KEY (PostID) REFERENCES Post(id),
);
/***********************************************/
CREATE TABLE Portfolio (
    id              BIGINT NOT NULL AUTO_INCREMENT,
	UserID          BIGINT NOT NULL,
	content         TEXT,      
    PRIMARY KEY (id),
    FOREIGN KEY (UserID) REFERENCES User(UserID),
);
CREATE TABLE Project (
    id              BIGINT NOT NULL AUTO_INCREMENT,
    title           NVARCHAR(200),
	contect         TEXT,  
	image           TEXT,  
	is_done         BOOLEAN, 
	created_at      DateTime,
    PRIMARY KEY (id),
);
CREATE TABLE Technology (
    id              BIGINT NOT NULL AUTO_INCREMENT,
    title           NVARCHAR(100),
	content         TEXT, 
	image           TEXT,
    PRIMARY KEY (id),
);
CREATE TABLE portfolio_projects (
    id              BIGINT NOT NULL AUTO_INCREMENT,
    PortfolioID     BIGINT,
    ProjectID       BIGINT,
    PRIMARY KEY (id),
    FOREIGN KEY (PortfolioID) REFERENCES Portfolio(id),
    FOREIGN KEY (ProjectID) REFERENCES Project(id),
);
CREATE TABLE portfolio_techs (
    id              BIGINT NOT NULL AUTO_INCREMENT,
    PortfolioID     BIGINT,
    TechnologyID    BIGINT,
    skill_level     INTEGER,
    PRIMARY KEY (id),
    FOREIGN KEY (PortfolioID) REFERENCES Portfolio(id),
    FOREIGN KEY (Technology) REFERENCES Technology(id),
);
CREATE TABLE project_techs (
    id              BIGINT NOT NULL AUTO_INCREMENT,
    ProjectID       BIGINT,
    TechnologyID    BIGINT,
    PRIMARY KEY (id),
    FOREIGN KEY (PortfolioID) REFERENCES Portfolio(id),
    FOREIGN KEY (Technology) REFERENCES Technology(id),
);