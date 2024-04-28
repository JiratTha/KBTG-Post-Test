### Software Feature

1. **swagger**
2. **adjustable tax level**
3. **deduct value with version**
4. **MVC Style**
5. **Some Testing**
6. **Transaction Update**

---

### How to Run the Program

#### Setup Instructions

1. **Start the Database**:
   Ensure Docker is installed on your machine. Begin by starting the database using Docker Compose:
   ```
   docker-compose up -d
   ```

2. **Initialize the Database**:
   Once the database service is running, run the following Go script to create tables based on the SQL files: **By uncomment file initDB.go**
   ```
   go run initDB.go
   ```
   ****After successfully initializing the database, remember to comment it back with expectation for first file in initDB.go.****
   This script uses SQL files `allowance.sql` and `personal_deduction` to set up the necessary database tables.


   
3. **Build the Docker Image**:
   Build a Docker image for the application:
   ```
   docker build -t jirat-tax .
   ```

#### Running the Application


- **Directly on PC**:
  To run directly on your PC, you can use Go to execute the main application file. This method requires setting
  environment variables either in your system or within your IDE:
  ```
  go run main.go
  ```
#### FAQ

1. If error during DOcker build pleas make sure that initDb.go is comment except package main line

---

#### Ready to Use

After following these steps, your application should be up and running and ready to use. Ensure that you have set all
the required environment variables before starting the service, especially when running directly on your PC.

---
