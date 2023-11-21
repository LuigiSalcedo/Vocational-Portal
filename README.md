# Vocational-Portal
This respository contains the source code about a Vocational-Portal. Here You can find information about profesional and technical careers in Colombia.
## Preparing the server
### Dependencies
To run vocaportal-server you have to get the next requeriments:
- Go: [Download Go here](https://go.dev/dl/)
- PostgreSQL: [Download PostgrsSQL here](https://www.postgresql.org/download/)
### Preparing database
Once installed go and PostgreSQL you have to create a database. In this example It will named **"vocadb"**.
Now You have to install [goose](http://pressly.github.io/goose/) to manage the database migrations.
You can install It writing the next command on your command line console (CLC):
```
go install github.com/pressly/goose/v3/cmd/goose@lastest
```
### Loading database data
With the database created and goose installed, create a enviroment variable (EV) named **vocaportal_db_data**. The variable's value
must be like this:
```
"user=<your postgres user> password=<your postgres password> dbname=vocadb host=localhost sslmode=disable"
```
With the EV saved open your CLC and go to the project root folder (where go.mod file is located) and type the next command:
```
goose -dir migrations postgres %vocaportal_db_data% up
```
The database should be updated.
### Running
The program is compiled on Windows 64 bits. If You are on a Windows 64 bits You can execute server.exe file. If You are on a different
operative system You can execute It with the next command:
```
go run server.go
```
The server should run on localhost:8088.
### Documentation
The server documentation can be accessed on http://localhost:8088/documentation
