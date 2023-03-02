*API Cookbook*

*Prequisites*
---------------
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
----------------

*Running the server*

1. Run the command "go run main.go"
    - Make sure that "Connected"
2. Go to Postman and follow API requests below to populate database

*Create Class*
![Alt text](/Backend/assets/createClass.png?raw=true "Create Class")

*Delete Class*
![Alt text](/Backend/assets/deleteClass.png?raw=true "Delete Class")

*Get Classes*
![Alt text](/Backend/assets/getClasses.png?raw=true "Get Classes")

*Create Post*
![Alt text](/Backend/assets/createPost.png?raw=true "Create Post")

*Get Posts By Class Id*
![Alt text](/Backend/assets/getPostByClassID.png?raw=true "Get Posts By Class Id")
