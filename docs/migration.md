# Migration
#### Update package list
```bash
sudo apt update
```

#### Install MySQL
```bash
sudo apt install mysql-server
```

#### Start MySQL service
```bash
sudo systemctl start mysql
```

#### Enable MySQL to start on boot
```bash
sudo systemctl enable mysql
```

#### Run security script (recommended)
```bash
sudo mysql_secure_installation
```

#### Check MySQL version
```bash
mysql --version
```

#### Connect to MySQL (if password is set)
```bash
mysql -u root -p
```

#### Create a new user
```bash
mysql -u root -p
```
```sql
CREATE USER 'barn_manager'@'localhost' IDENTIFIED BY 'barn_manager';
GRANT ALL PRIVILEGES ON *.* TO 'barn_manager'@'localhost';
FLUSH PRIVILEGES;
```

#### Run migration
```bash 
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -database "mysql://barn_manager:barn_manager@tcp(localhost:3306)/barn_test" -path /Users/doblakov/magistratura/hw/go/havchik_podbirator/barn_manager/migrations up
migrate -database "mysql://barn_manager:barn_manager@tcp(localhost:3306)/barn_test" -path /Users/doblakov/magistratura/hw/go/havchik_podbirator/barn_manager/migrations down
```
