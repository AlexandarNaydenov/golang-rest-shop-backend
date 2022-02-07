# Golang backend for online shop

## Local setup:

### Prerequisites:

- GOLang
- MySQL

### Steps:

1. Logging with your credentials using mysql cli.
```mysql
mysql -u root -p
Enter password:
```

2. Create a new database
```mysql
create database online_shop;
use online_shop;
```

3. Use initial sql file to create the tables and insert initial data
```mysql
source sql/initialSetup/tableProducts.sql;
source sql/initialSetup/tableOrders.sql;
source sql/initialSetup/tableOrderedProducts.sql;
```

4. Start main.go with those environment variables
```
MYSQL_USER = <your mysql username>
MYSQL_PASSWORD = <your mysql password>
MYSQL_IP_ADDRESS = <to ip address of the mysql> default: localhost:3306
```

### Swagger

Go to\
**_localhost:8080/swagger/index.html_**\
to see the available endpoint and try example call.