drop database statDB;
CREATE database statDB;
use statDB;
CREATE table Stat (
    id int auto_increment primary key,
    dir_path varchar(200),
    total_size varchar(100), 
    loading_time varchar(15),
    date_of_request date,
    time_of_request time
)