
# Grit Space

Grit Space is a web application that allows users to register, create posts, comment on posts, and like/dislike posts and comments. The application includes user authentication, session management, and various features for filtering and categorizing posts.

## Features

1. **User Authentication & Sessions**
   - User registration with email, username, and password
   - Password encryption (bcrypt)
   - User login and session management
   - Session expiration after a set time

2. **User Communication (Posts & Comments)**
   - Users can create posts
   - Users can comment on posts
   - Posts must belong to at least one category
   - Non-registered users can only view posts & comments

3. **Likes & Dislikes System**
   - Users can like/dislike posts and comments
   - Users cannot like & dislike the same post simultaneously
   - Like/dislike counts must be visible

4. **Filtering & Categorization**
   - Users can filter posts by category
   - Users can view only their own posts
   - Users can filter posts they have liked

5. **Database Queries (SQLite)**
   - Persistent storage for users, posts, and comments
   - Prevent SQL injection and invalid data storage

6. **Dockerization**
   - Project includes a Dockerfile
   - Application runs inside a Docker container

7. **Permissions & Access Control**
   - Only registered users can create posts/comments
   - Only registered users can like/dislike posts
   - Empty posts/comments are rejected

8. **Error Handling & HTTP Responses**
   - Handle 400 Bad Requests for invalid inputs
   - Handle 500 Internal Server Errors
   - Use correct HTTP methods for actions
   - Prevent 404 Not Found errors

9. **Performance & Code Quality**
   - Optimize database queries
   - Follow clean coding practices

## Bonus Features (Optional)
   - Password hashing using bcrypt
   - Use UUIDs for user IDs
   - Clear README & documentation
