# Ryde Back-end Developer Test

## About the Project
   The repo contains a simple CRUD server to get, add, update and delete users. It is 
   build using Golang with GIN for the API framework. The user data is stored in MongoDB
   
   

## Getting Started
1. Install [Go(v1.19.3)](https://go.dev/doc/install/) and [MongoDB(v6.0.4)](https://www.mongodb.com/docs/manual/installation/)
2. Clone this repo in your system path
3. Run command to setup MongoDB collection
    ```
    make db-setup
    ``` 
4. Run command to start server on localhost:8080
    ```
   make start-server
   ```
## Running Tests
## Advanced Requirements
### User Auth Strategy
    We could use a JWT based strategy where the client would send the JWT in the request
    headers and each endpoint on the server side will be protected. It would first validate
    the JWT for authentication and redirect the user to the login if she is not authenticated
    For authorization, we can store the role of each user in database and have restrict
    permission to specific endpoints
    
    
### Logging Strategy
    We can log and to stdout and redirect it to a log file. We can use tools like ELK
    (ElasticSearch, Logstash and Kibana) to collect, index and visualise logs. It will be 
     useful for logging API errors and significant event. Logs will include information like 
     the loglevel, origin information, error etc. Along with this, we could think about storing
     all request and responses in a store like Hadoop for analytics purposes.
### Followers / Following Feature
    We can add two new fields "followers" and "following" for each user which will 
    hold the _id of other users. We can create a PATCH endpoint to modify those fields.
    If User A starts following UserB, we will modify the "following" of User A and 
    "follower" of User B.
   ```
   {
      "id" : ObjectID("1abd4e"),
      .
      .
      .
      "followers":[
         ObjectID("1aee4e")
         ObjectID("1a344e")
      ],
      "following":[
         ObjectID("1a334e")
         ObjectID("1a23de")
      ]
      
   }
   ```
### Query nearby friends
    We can make the address field nested and include the latitude and longitude
    of the place. We can get the coordinate information by querying using _id of 
    followers and following for a user. Finally, we can set a radius threshold 
    and filter the friends who fall in that radius

   