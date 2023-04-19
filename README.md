# Classical
A hub for finding group chats for your classes.

Post your groupchats on Classical to collaborate with other classmates.
Upvote to increase a groupchat's ranking.
Downvote to decrease a groupchat's rating

## Front-End
* Chandler Carnes
* Ryan Circelli

## Back-End
* Travis Maddox
* Victor Polisetty

# Setup and Running

## Run Front-End
### Setup
- cd ./frontend
- npm i

### Angular
- cd ./frontend
- ng serve --open

### Cypress
- cd ./frontend
- npx cypress open

---------------
## Run Back-End
### Setup
1. Install MySQL
2. Startup MySQL
3. Connect to SQL localhost at port 3306 in MySQL workbench
    -username: root
    -password: password123
    WARNING: may have to edit above credentials this based on SQL database computer settings
        -This can be done by editing the file "db/db.go" and the statement 	*DB, err = sql.Open("mysql", "{username}:{password}@tcp(localhost:{port})/{schemaName}")*

4. Create SQL schema named "classical"
5. Run SQL queries in "data-access/create-tables.sql" to create the necessary tables by copy pasting the commands
6. Install Go

### Run Server
1. cd ./backend
2. Run the command "go run main.go"
    - Make sure that "Connected"