# tiny-url-go

### Required environment
Clean linux server, assuming neither Go or Mysql installed

### Deployment:
1. login to linux server and clone the project: `git clone https://github.com/MrXu/tiny-url-go.git`
2. go to `deploy` folder and execute `bash deploy.sh`and the application should be running at `localhost:8080`
 1. the script will install Go and configure GOPATH
 2. the script will install Mysql and create database and tables
 3. the script will deploy project into go/src folder and run the application
 
 
### Test
1. Shorten url: 
 1. `curl -sX POST -H 'Content-Type: application/json' 'localhost:8000/api/v1/shorten' -d '{"url":"http://averylongurl.com/verylong"}'`
 2. should return `{"short":"localhost:8080/GuNWmTOi"}`
2. Get original url: 
 1. `curl -sX POST -H 'Content-Type: application/json' 'localhost:8000/api/v1/original' -d '{"short":"http://localhost:8080/GuNWmTOi"}'`
 2. should return `{"url":"http://averylongurl.com/verylong"}`
