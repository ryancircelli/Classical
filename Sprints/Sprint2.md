## **Frontend**

### **Issues your team planned to address**

Implement website header and navigation bar
   Pick a color scheme, create a header, use tailwind and daisyUI to design buttons
   This task is done when we have buttons that link to different pages/search bars on the Classical site
   ~1h

Create searchable box on main page that opens page using search value
   The user is able to search in the search bar, and open the results page of the desired search input.
   This task will be done when users can see/access a search box and type in said search box. A cypress unit test is implemented to test the page opening.
   ~2h

Display posts in the table with test data
   Use mock data to display a table of user posts on the results page
   This task is done when the user is able to see a list of posts that are formatted
   ~2h

Add trending posts on main page with getTrendingClasses
   Grab trending posts from backend, auto-rank posts by popularity in ranked list, display ranked list on main page
   This task is done when we have a section of the main page with functional hyperlinks to trending/popular classes
   ~1h

### **Issues successfully completed**

- Implement website header and navigation bar
- Create searchable box on main page that opens page using search value
- Display posts in the table with test data

### **Which ones didn't and why?**

- Add trending posts on main page with getTrendingClasses
  - This was not completed because we could not figure out how to mock the api using cypress.

---

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

Access the database with go

- Set up code in main.go to connect to the sql database using a mysql driver
- Connection between go backend and mySql database is successful upon executing go main.go
- ~1 hr

Create getClassPosts endpoint on go with logic (5 and 6 combined)

- Implemented go logic, setup endpoint using mux to create an API between mysql and go
- Get function is linked between getClassPosts in go code and classController.go and postController.go
- ~1 hour

  Test cases:

  - query classes and response with a JSON object with classes that are in the database at the current time

Create addPost() endpoint with logic (make sure to add logic to delete after x downvotes)

- Setup in postController.go, bind the JSON received to a new JSON that is added to the database
- endpoint is called and adds post to the database
- Logic for votes not yet implemented
- ~ 1 hour

  Test cases:

  - AddPost shown with json passed in with a classId and the postName along with the postContent

### **Issues successfully completed**

- Access the database with go
- Create getClassPosts endpoint on go with logic (5 and 6 combined)
- Create addPost() endpoint with logic (make sure to add logic to delete after x downvotes)

### **Which ones didn't and why?**

- Logic for upvotes and downvotes is not yet implemented because of the complexity of adding and querying posts and classes thus far, will be added next sprint.

- Test cases for backend using go test -v is difficult to use because of connecting to the mysql database and resources online are limited there are test cases shown using postman to demonstrate that the functions work correctly and also test cases shown that are almost complete

## **Backend API Documentation **##

_API Cookbook_

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
