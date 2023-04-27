-- Database: proj2

-- DROP DATABASE IF EXISTS proj2;

CREATE DATABASE proj2
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_India.1252'
    LC_CTYPE = 'English_India.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
	
create table moviestest(
    movieID int NOT NULL,
    movieName varchar(50) NOT NULL
)

insert into moviestest (
    movieID,
    movieName
)
values
	(1, 'John Wick'),
    (2, 'Tangled'),
    (3, 'Andhadhun'),
	(4, 'DDLJ'),
    (5, 'Singham'),
    (6, 'RRR');
	
	

