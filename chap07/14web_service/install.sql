drop database gwp;
create database gwp;
drop user gwp;
create user gwp@localhost identified by gwp
GRANT ALL PRIVILEGES ON gwp TO 'gwp'@'localhost'
