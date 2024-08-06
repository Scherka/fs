drop database statDB;
CREATE database statDB;
use statDB;
CREATE table Stat (
    id int auto_increment primary key not null,
    dir_path varchar(200)not null,
    total_size varchar(100) not null, 
    loading_time float not null,
    date_of_request date not null,
    time_of_request time not null
)