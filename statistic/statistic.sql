drop database statDB;
CREATE database statDB;
use statDB;
CREATE table Stat (
    id int auto_increment primary key,
    dir_path varchar(200),
    total_size varchar(20), 
    loading_time int,
    date_time_of_load time
)