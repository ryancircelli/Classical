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
![Alt text](/Backend/assets/createClassPost.png?raw=true "Create Post")

*Get Posts By Class Name*
![Alt text](/Backend/assets/getPostByClassName.png?raw=true "Get Posts By Class Id")

*Increase Post Votes*
![Alt text](/Backend/assets/increasePostVotes.png?raw=true "Increase Post Votes")

*Decrease Post Votes*
![Alt text](/Backend/assets/decreasePostVotes.png?raw=true "Decrease Post Votes")

*Search Class By Name*
![Alt text](/Backend/assets/searchClasses.png?raw=true "Search Class By Name")

*Get Trending Classes*
![Alt text](/Backend/assets/getTrendingClasses.png?raw=true "Get Trending Classes")
