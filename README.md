Fiber API Architecture:

1. Directory Structure:

```bash
/api
    /controllers
    /middlewares
    /models
    /routes
    /utils
/config
/database
    /migrations
    /seeds
/public
    /uploads
    /logs
main.go
```

2. Components:

   - **/api/controllers**: This directory contains the logic for handling API
     endpoints. Typically, each resource (e.g., users, posts) would have its own
     controller.

   - **/api/middlewares**: Middlewares are functions that run before hitting the
     main controllers. This directory will contain custom middleware functions
     such as authentication, logging, CORS, etc.

   - **/api/models**: It contains data structures and operations related to the
     database. Here you'd define your structs that map to database tables and
     any related methods or functions.

   - **/api/routes**: Define and organize your API routes here. It's a good
     practice to have a separate file for each resource's routes.

   - **/api/utils**: Utility functions that can be used across the application,
     e.g., password hashing, JWT token management, error handling functions,
     etc.

   - **/config**: Configuration-related things like setting up environment
     variables, application constants, etc.

   - **/database**: /migrations: For managing database changes and versioning.

   - **/seeds**: For initial data seeding of the database.

   - **/public/uploads**: If your application supports file uploads, storing
     them temporarily or permanently in this directory could be an option.
     Remember to have security measures in place if allowing uploads.

   - **/public/logs**: For saving logs if you opt to save them as files.

   - **main.go**: This is the entry point for your application. It sets up the
     Fiber app instance, connects to the database, loads middlewares,
     initializes routes, and starts the server.

3. Local Environment:

   - **GO**: Ensure go is installed by running `go --version`.

   - **MOD**: run `go mod download` for dependencies.

   - **.env**: Environment file which should include the following environment
     variables ['DB_HOST', 'DB_USER', 'DB_PASSWORD', 'DB_NAME'].

4. Docker Environment:

   - **TEST**: run `docker compose -f docker-compose.test.yml`, to launch the
     environment. This includes a mariadb service which can be used to test
     reads and writes to the database.

   - **DEVELOPMENT**: TODO

   - **PRODUCTION**: TODO
