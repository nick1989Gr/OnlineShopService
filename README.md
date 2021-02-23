# OnlineShopService

This is an example of a REST web server in Go.

For the puprose of illustrating how to create a web service in Go we use a simple MySql database with an Items table.

You can find more information on how to setup your database in one of the following sections.

## Features implemented

We implement REST API calls for the Item endpoint for the following functionalities:

- Add an item
- Get all items
- Get item with id
- Remove an item
- Update an item

## Packages used

- For logging we use sirupsen/logrus
- gocraft/web is used as go mux and middleware package

## Setting up your database

You can use a docker image for your mysql database.
The following command will run locally the docker image. Make sure that the port
you expose is the same as in github.com/nick1989Gr/OnlineShopService/database package
in SetupDatabase method.

```
docker run -it -d -v %CD%/data:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password123 mysql:latest
```

You can use any mysql client like MySQL workbench to initialize your container with our database by executing the db_init.sql file.
