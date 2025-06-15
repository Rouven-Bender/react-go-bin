# Setup
## generate database
```bash
sqlite3 db.sqlite
# in the sqlite3 repl
.read schema.sql
.quit
```
## compile
```bash
./createRandomSecrets.sh # this will generate the secret files needed for compiling
go build .
```
## setting up the first user to start using the app
```bash
./reactgobin add-user
```
this will print the secret key you need for this user on the login page
# Usage
```bash
go run . serve # this will launch the server
go run . add-user # this will add a new user and print the secret key needed for logging in as that user
```
