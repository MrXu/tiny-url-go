#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

apt-get update

echo "installing golang and setting up GOPATH"
bash ${DIR}/goinstall.sh --64


echo "installing mysql ..."

sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password password root'
sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password root'
sudo apt-get -y install mysql-server



echo "creating database and tables ..."

DB_NAME="tinyurl_db"
USER_NAME="tinyurl"
PASSWORD="tinyurl"
mysql -uroot -proot -e "CREATE DATABASE ${DB_NAME} /*\!40100 DEFAULT CHARACTER SET utf8 */;"
mysql -uroot -proot -e "CREATE USER ${USER_NAME}@localhost IDENTIFIED BY '${PASSWORD}';"
mysql -uroot -proot -e "GRANT ALL PRIVILEGES ON ${DB_NAME}.* TO '${USER_NAME}'@'localhost';"
mysql -uroot -proot -e "FLUSH PRIVILEGES;"
mysql -u${USER_NAME} -p${PASSWORD} ${DB_NAME} < ${DIR}/create_tables.sql



echo "building go app ..."

cp -R ${DIR}/../../tiny-url-go $HOME/go/src/
mv $HOME/go/src/tiny-url-go $HOME/go/src/tinyUrl
cd $HOME/go/src/tinyUrl

go get github.com/go-sql-driver/mysql
go get github.com/gin-gonic/gin

go run main.go

echo "serving go app"