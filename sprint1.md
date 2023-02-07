## **Backend**

### **User Stories**

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
