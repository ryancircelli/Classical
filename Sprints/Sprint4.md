## **Backend**

### **User Stories**

1. Create a database
   This task is done when the developer is able to see a defined database structure within mySql
   ~1h
2. Setup Go dev environment
   This task is done when the developer can run a simple Go program through VSCode
   ~1h
3. Establish database schema
   In mySql, create neccessary key relationships between tables and make sure each table has relevant information
   This task is done when the developer can use SQL queries to navigate through the relational tables created
   ~1h
4. Access the database with go
   Set up Go code to connect to database
   This task is done when connection is working between Go backend and the mySql database
   ~1h
5. Create getClassPosts endpoint on go without logic
   Implement Go logic, setup endpoint, and test that endpoint can send data
   The task is done when within main the get function is set up and linked to a getClassPosts function
   ~1h
6. Add getClassPosts() logic with necessary queries and test cases (sorted by upvote column)
   Create an SQL query to get all posts from a given class ID
   Sort posts by most upvotes
   Send sorted posts object
   Make sure object is running correctly by running tests in command line
   ~1 hour
7. Create addPost() endpoint with logic (make sure to add logic to delete after x downvotes)
   Bind the JSON recevied to a new JSON that will bea dded to the database
   Task is done when there is a specific endpoint that can be called that will add the post to the database
   ~1 hour
8. Create addVote() endpoint with logic
   Task is done when there is a specific endpoint that can be called that will add one vote to a specific post's upvote number via post's ID
   Create an SQL query to get specific post ID and add 1 to the current upvotes of the post
   Send messages back
   ~1 hour
9. Create searchClass() (sorted based on upvotes) endpoint with logic
   Create an SQL query to get all classes and return correct class based on class ID
   Task is done when there is a specific endpoint that can be called that will lists posts based on getClassPoints()
   ~1 hour
10. Create getTrendingClasses endpoint with logic (ranked by total upvotes for all posts)
    Task is done when there is a specific endpoint that can be called to return a sorted list(based on most post upvotes)
    Create a SQL query to get the classes
    Sort post by most upvotes
    Send sorted posts object
    ~1 hour

### **Issues your team planned to address**

Update all functions to saistify updated database schema

- Add in timePosted (Unix timestamp) for post, add in lastUpdated for class (Unix timestamp), add total_votes for class, add postClassName for post.
- postClassName references className now. Deprecate classId (it was never used by frontend).
- Update createClass, createClassPost, decreasePostVotes, increasePostVotes, getPostsByClassName, getClassByName, getClasses
- ~3 hours

Change getClassByName to accept prefix or postfix and return all corresponding classes

- Change SELECT to accept varying number of characters
- change to match on prefix or post fix
- Make case insensitive
- ~1.5 hours

### **Issues successfully completed**

- Update all functions to saistify updated database schema
- Change getClassByName to accept prefix or postfix and return all corresponding classes

### **Which ones didn't and why?**

- None

### Backend Unit Test Cases

- TestGetClasses
  Tests that the http request can be sent across to the database and the correct json output is received

- TestCreateClass
  Should send a json object and a class will be created and stored in the database. Json object of the new class is sent back

- TestDeleteClass
  Should send a request with the class name that is to be deleted attached and that class will be deleted

- TestCreatePost
  Should send a json object with the post information inside and the new post will be stored in the database

- TestGetPostsByClassID
  Should send a request with the class ID as a variable and the posts will be returned as a json oject

- TestGetClassesByName
  Should send a request with the class name as a variable and the class will be returned as a json object

- TestGetTrendingClass
  Should send a request and a json object with the classes listed in descending order based on total_votes will be sent back

- TestIncreasePostVote
  Should send a post request with json object that contains postId and postClassName and return post with id = ? was updated.

- TestDecreasePostVote
  Should send a post request with json object that contains postId and postClassName and return post with id = ? was updated.

### Updated Backend API Documentation

## _Prequisites_

1. Install MySQL
2. Startup MySQL
3. Connect to SQL localhost at port 3306 in MySQL workbench
   -username: root
   -password: password123
   WARNING: may have to edit above credentials this based on SQL database computer settings
   -This can be done by editing the file "db/db.go" and the statement _DB, err = sql.Open("mysql", "{username}:{password}@tcp(localhost:{port})/{schemaName}")_

4. Create SQL schema named "classical"
5. Run SQL queries in "data-access/create-tables.sql" to create the necessary tables by copy pasting the commands
6. Install Go

---

_Running the server_

1. Run the command "go run main.go"
   - Make sure that "Connected"
2. Go to Postman and follow API requests below to populate database

_Create Class_
![Alt text](/Backend/assets/createClass.png?raw=true "Create Class")

_Delete Class_
![Alt text](/Backend/assets/deleteClass.png?raw=true "Delete Class")

_Get Classes_
![Alt text](/Backend/assets/getClasses.png?raw=true "Get Classes")

_Create Post_
![Alt text](/Backend/assets/createPost.png?raw=true "Create Post")

_Get Posts By Class Id_
![Alt text](/Backend/assets/getPostByClassID.png?raw=true "Get Posts By Class Id")

_Increase Post Votes_
![Alt text](/Backend/assets/increasePostVotes.png?raw=true "increase post Votes")

_Decrease Post Votes_
![Alt text](/Backend/assets/increasePostVotes.png?raw=true "Decrease post votes is the same")

_Get Trending Classes_
![Alt text](/Backend/assets/getTrendingClasses.png?raw=true "Get Trending Classes")

_Get Classes By Name_
![Alt text](/Backend/assets/searchClassesByName.png?raw=true "Get Classes By Name")
