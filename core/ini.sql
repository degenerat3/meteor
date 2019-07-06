CREATE USER meteoruser with password "dbpassword";
CREATE DATABASE meteor;
GRANT ALL PRIVILEGES ON DATABASE meteor TO meteoruser;