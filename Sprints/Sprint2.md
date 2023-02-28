## **Frontend**

### **User Stories**

1. create an angular dev environment
   Install: Node.js, npm, Angular CLI & Git CLI (if needed)
   This task is done when the developer is able to code via Angular with VSCode and push changes to GitHub
   ~1h
2. install some kind of component library
   Research component libraries and install the one that matches our design
   This task is done when a component library has been installed and tested on the main page.
   ~1h
3. implement hello world on the main page, a page for viewing posts, and a page for search results
   Create pages on angular and implement navigation by url
   This task is done when the user can switch between different pages by the url
   ~1h
4. Implement website header and navigation bar
   Pick a color scheme, create a header, use tailwind and daisyUI to design buttons
   This task is done when we have buttons that link to different pages/search bars on the Classical site
   ~1h
5. Display posts in the table with test data
   Use mock data to display a table of user posts on the results page
   This task is done when the user is able to see a list of posts that are formatted
   ~1h
6. Access posts with getClassPosts API and display them in the posts table ui
   Grab all posts matching course code, rank by number of upvotes, display in table in ranked order
   This task will be done when posts can successfully be accessed from backend and displayed in a ranked fashion in UI
   ~1h
7. implement upvote/downvote ui
   Design a ranking system, sort posts by votes
   This task will be done when users are able to use the up/downvote button
   ~1h
8. Implement addVote to upvote/downvote UI
   Connect upvote/downvote ui to addVote API request
   This task will be done when users CAN up/downvote posts by clicking a button
   ~1h
9. create post posting component, and make http POST
   Implement post component to add message and link, checked with regex to match group chat links (ex: Discord, GroupMe)
   The user is able to add posts to a class page
   ~1h
10. Polish class posts UI
    implement asthetic links with names chosen by user, sort posts by upvotes, etc.
    This task will be done when the class post UI is asthetically pleasing and functional
    ~1h
11. Create searchable box on main page with searchClass API
    The user is able to search through classes and see the results on the result page, when a class is clicked on the result page it opens the class page with all the comments.
    This task will be done when users can see/access a search box and type in said search box
    ~1h
12. Create search results page displaying searchClass API data
    Grab links/data from backend function call (ranked in order of #upvotes), create numbered list of size: #search results, insert links into list
    This task is done when the user’s search action successfully displays all posts linked to matching course codes, ranked & displayed by number of upvotes
    ~1h
13. Add trending posts on main page with getTrendingClasses
    Grab trending posts from backend, auto-rank posts by popularity in ranked list, display ranked list on main page
    This task is done when we have a section of the main page with functional hyperlinks to trending/popular classes
    ~1h

### **Issues your team planned to address**

Create an angular dev environment

- Install: Node.js, npm, Angular CLI & Git CLI (if needed)
- This task is done when the developer is able to code via Angular with VSCode and push changes to GitHub
- ~1h

Install some kind of component library

- Research component libraries and install the one that matches our design
- This task is done when a component library has been installed and tested on the main page.
- ~1h

Implement hello world on the main page, a page for viewing posts, and a page for search results

- Create pages on angular and implement navigation by url
- This task is done when the user can switch between different pages by the url
- ~1h

Implement website header and navigation bar

- Pick a color scheme, create a header, use tailwind and daisyUI to design buttons
- This task is done when we have buttons that link to different pages/search bars on the Classical site
- ~1h

### **Issues successfully completed**

- Create an angular dev environment
- Install some kind of component library
- Implement hello world on the main page, a page for viewing posts, and a page for search results

### **Which ones didn't and why?**

- Implement website header and navigation bar
  - This issue was not completed becuase our team needs to meet to decide on a color scheme and UI inspiration.

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

Create a database

- Create mySql database with mySQL CLI
- Set up sql code
- Create tables for classes with variable ids
- Create a mock class to test on
- 1 hour

Setup go dev environment

- This task is done when the developer can run a simple Go program through VSCode
- Download VSCode, GO, and make a simple "Hello World" program
- ~1 hour

Establish database schema

- This task is done when the developer can use SQL queries to navigate through the relational tables created
- Create tables (w/ foreign keys, etc…)
- Test queries to make sure relational tables working as intended
- ~1 hours

### **Issues successfully completed**

- Create a database
- Setup go dev environment
- Establish a database schema

### **Which ones didn't and why?**

-N/A
