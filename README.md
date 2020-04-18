# Hash Code Challenge

## 1 - Introduction

The purpose of this file is to present information about the work developed to solve the code challenge prepared by the company **Hash Pagamentos** that can be founded in the following link: 

*Website*: https://github.com/hashlab/hiring/blob/master/challenges/en-us/backend-challenge.md

In order to summarise, the project comprehends the development of two back-end applications for the management of users, products, promotions and discounted dates, in addition to the evaluation of discounts applied to products based on users and discounted dates settings. It is composed by a **microservice 1** and a **microservice 2** developed using **NodeJS** and **Go** programming languages, respectively, which communicate together through **RPC** (*Remote Procedure Call*), in addition to a **Mongo** database.

The **microservice 1** is responsible for evaluating discounts applied to products, while the **microservice 2** is responsible for managing users, products, promotions and discounted dates.

In addition to meeting the original requirements of the technical challenge, as indicated above the additional functionalities of the **microservice 2** were developed as an **extended approach** in order to **facilitate** the handling of the project and, thus, the evaluation of discounts when obtaining one or more products.

Throughout this documentation, a few aspects will be highlighted, such as, the configuration of environment variables of **Mongo** database and the procedures adopted to run the project with **Docker** containers.

Finally, the last section named **Project Dynamics** illustrates a brief report of how the solution works in practice.

## 2 - API Documentation

The documentation of the API implemented in **Go** programming language that refers to the **microservice 2** was developed following the **OpenAPI 3.0** specification. Inside the directory **api-docs** there is a script named **swagger-json-to-html.py**. When running it using the **openapi-ms-2.json** file, it is generated a HTML page named **index.html** within the current directory that illustrates details about the API *endpoints*.

## 3 - Project Organization

The developed solution is organized according to the structure of directories and files summarized below:

### 3.1 - Back-end

The following directories contain the APIs implementation using **NodeJS** and **Go** programming language, respectively.

#### 3.1.1 - NodeJS

The **microservice 1** is responsible for the evaluation of discounts applied to one or more products.

**nodejs/cmd/server**: it contains the configuration and implementation of the gRPC server.

**nodejs/internal/grpc/services/impl**: it contains the implementation of the services related to the handling of gRPC requests, as well as the elaboration of gRPC responses.

**nodejs/internal/grpc/services/impl_test**: it contains the tests of the implementation of the services using the **JavaScript** testing framework named **jest**.

**nodejs/internal/grpc/services/server**: it contains an abstraction of the server that allows to "attach" some resources to make them available during the API requests. Here, it's used to store a structure that holds attributes to manage the data.

**nodejs/internal/models**: it contains the definition of the data entities used by both the API and the database.

**nodejs/internal/mongodb**: it contains the implementation directed to the database configuration along with **CRUD** operations.

**nodejs/internal/proto/entities**: it contains the specification of the **protocol buffers** entities.

**nodejs/internal/proto/services**: it contains the specification of the **protocol buffers** services.

**nodejs/internal/services/promotion**: it contains the implementation of the evaluation of promotions.

**nodejs/internal/services/promotion/discounted-date**: it contains the implementation of the examination of discounted dates.

**nodejs/internal/tests**: it contains the configuration of the test cases using the **JavaScript** testing framework named **jest**.

**nodejs/internal/utils**: it contains supporting functions, such as, to generate and format dates, as well as to evaluate the equivalence of JSON objects used during the tests.

**nodejs/.env**: it contains the variables for the configuration of the **development** environment.

**nodejs/internal/tests/.env**: it contains variables for the configuration of the **test** environment.

The **nodejs/.env** file contains the environment variables referring to the connection to **Mongo** database, as well as the exposure of the access address for gRPC communication, indicated below:

```
DB_USERNAME=user
DB_PASSWORD=password
DB_HOST=db
DB_PORT=27017
DB_NAME=db
```

```
GRPC_SERVER_HOST=0.0.0.0
GRPC_SERVER_PORT=50051
```

#### 3.1.2 - Go

The **microservice 2** is responsible for the management of users, products, promotions and discounted dates.

**go/cmd/server**: it contains the configuration and implementation of the gRPC server, as well as the exposure of a HTTP server to accept API requests and send them to the corresponding server services.

**go/internal/grpc/entities**: it contains the outputs related to **protocol buffers** entities resulting from the execution of **proto-gen.sh** script.

**go/internal/grpc/services**: it contains the outputs related to **protocol buffers** services resulting from the execution of **proto-gen.sh** script.

**go/internal/grpc/services/impl**: it contains the implementation of the services related to the handling of API requests, as well as the elaboration of its responses.

**go/internal/grpc/services/impl_test**: it contains the tests of the implementation of the services using the **Go** language test package.

**go/internal/grpc/services/server**: it contains an abstraction of the server that allows to "attach" some resources to make them available during the API requests. Here, it's used to store a structure that holds attributes to manage the data.

**go/internal/middlewares**: it contains intermediate validations of parameters transmitted through API requests.

**go/internal/models**: it contains the definition of the data entities used by both the API and the database.

**go/internal/mongodb**: it contains the implementation directed to the database configuration along with **CRUD** operations (*create*, *read*, *update* and *delete*).

**go/internal/mongodb_test**: it contains the tests of the implementation of **CRUD** operations using the **Go** language test package.

**go/internal/proto/entities**: it contains the specification of the **protocol buffers** entities.

**go/internal/proto/services**: it contains the specification of the **protocol buffers** services.

**go/internal/utils**: it contains supporting functions, such as, to generate random data used during the tests.

**go/internal/yaml/services**: it contains the specification of the API routes generated through a *reverse proxy* using **grpc-gateway** that acts as a *proxy* between a REST client and the gRPC server, publishing the API *endpoints* based on **proto** files. The result of this procedure is similar to the generation of API routes following the process of "annotated" proto files.

**go/scripts**: it contains the script file named **proto-gen.sh** used to generate the outputs from the **proto** files related to the entities and services necessary when configurating the gRPC server.

**go/.env**: it contains the variables for the configuration of the **development** environment.

**go/.test.env**: it contains the variables for the configuration of the **test** environment.

The **go/.env** file contains the environment variables referring to the connection to **Mongo** database, as well as the exposure of the access address for gRPC and HTTP communication, as indicated below:

```
DB_USERNAME=user
DB_PASSWORD=password
DB_HOST=db
DB_PORT=27017
DB_NAME=db
```

```
GRPC_SERVER_HOST=0.0.0.0
GRPC_SERVER_PORT=50052
```

```
HTTP_SERVER_HOST=0.0.0.0
HTTP_SERVER_PORT=8082
```

In addition, so that the **microservice 2** can communicate with the **microservice 1** to obtain one or more products with the possibility of discounts, the access address of the **microservice 1** must be also defined in the file.

```
GRPC_SERVER_HOST_MS_1=0.0.0.0
GRPC_SERVER_PORT_MS_1=50051
```

In order to not compromise the integrity of the database used by the project in terms of data generated from the execution of the test cases, two Mongo databases will be used.

In this sense, to facilitate future explanations regarding the details of the databases, consider that the database used for the storage of data in a "normal" actions is the **development** database and the one used for the storage of data resulting from the test cases is the **test** database.

These databases are named **db** and **test-db** by the environment variable **DB_NAME** of the **nodejs/.env** and **go/.env** files; and **TEST_DB_NAME** of the **nodejs/internal/tests/.env** and **go/.test.env** files, respectively.

(P.S. It is necessary to pay special attention to the database environment variables defined in these two previous files in case they are changed.)

### 3.2 - Mongo

The **mongodb/scripts/1-create_collections.js** file contains instructions for creating the **users**, **products**, **promotions** and **discountedDates** collections, as detailed below:

#### 3.2.1 - Collections

**Users**

In the **users** collection each document contains the data of a user.

This way, the **_id** field refers to the unique identifier of the user and the **first_name**, **last_name** and **date_of_birth** (**year**, **month** and **day**) fields refer to the first name, last name and date of user's birth (with year, month and day), respectively.

| Fields        | Data type |
|:--------------|:----------|
| _id           | Object ID |
| first_name    | String    |
| last_name     | String    |
| date_of_birth | Field     |

The date_of_birth field is configured as follows:

| Fields | Data type |
|:-------|:----------|
| year   | Integer   |
| month  | Integer   |
| day    | Integer   |

```
*application/json*

"date_of_birth": {
        "year": <Integer between 1-9999>,
        "month": <Integer between 1-12>,
        "day": <Integer between 1-31>
    }
```

**Products**

In the **products** collection each document contains data of a product.

This way, the **_id** field refers to the unique identifier of the product and the **price_in_cents**, **title** and **description** fields refer to the price in cents, title and description, respectively.

| Fields         | Data type |
|:---------------|:----------|
| _id            | Object ID |
| price_in_cents | Integer   |
| title          | String    |
| description    | String    |

In addition to the previous fields, there is other field named **discount** that refers to the discount applied to the product, displayed only when it is applicable. In this case, the pct and value_in_cents fields refer to the percentage and the value of the discount in cents, respectively.

| Fields         | Data type |
|:---------------|:----------|
| pct            | Float     |
| value_in_cents | Integer   |

**Promotions**

In the **promotions** collection each document contains the data of a promotion.

This way, the **_id** field refers to the unique identifier of the promotion and the **code**, **title**, **description**, **max_discount_pct** and **products** fields refer to the code, title, description, maximum discount percentage and a list of ids of all its products, respectively.

| Fields           | Data type |
|:-----------------|:----------|
| _id              | Object ID |
| code             | String    |
| title            | String    |
| description      | String    |
| max_discount_pct | Float     |
| products         | Array     |

The list of all ids of its products is configured as follows:

```
*application/json*

"products": [
        <The id of a product>
    ]
```

Additionally, each promotion must have a unique value for the **code** field and **in this context of the challenge**, to set up the promotion of discounted dates, its code field **must** be filled out with the value **DISCOUNTEDDATES**. Otherwise, the **microservice 1** will not evaluate any discount based on the criteria of the discounted dates.

**Discounted Dates**

In the **discountedDates** collection each document contains the data of a discounted date.

This way, the **_id** field refers to the unique identifier of the discounted date and **title**, **description**, **discount_pct** and **date** (**year**, **month** and **day**) fields refer to the title, description, percentage of discount, and date in which the discount will be applicable to one or more products, respectively.

| Fields       | Data type |
|:-------------|:----------|
| _id          | Object ID |
| title        | String    |
| description  | String    |
| discount_pct | Float     |
| date         | Field     |

The date field is configured as follows:

| Fields | Data type |
|:-------|:----------|
| year   | Integer   |
| month  | Integer   |
| day    | Integer   |

```
"date": {
        "year": <Integer between 0-9999>,
        "month": <Integer between 0-12>,
        "day": <Integer between 0-31>
    }
```

#### 3.2.2 - Configurations of Docker database containers

To execute the solution through **Docker** containers, it is necessary to relate the environment variables of the **mongodb/.env** and **mongodb/.test.env** files with the corresponding environment variables directed to the development and test databases defined in both **back-end** applications settings.

To do this, the environment variables of the **mongodb/.env** and **mongodb/.test.env** files must be associated with the environment variables of the **back-end/nodejs/.env** and **back-end/go/.env**; and **back-end/nodejs/internal/tests/.env** and **back-end/go/.test.env** files, respectively.

Additionally, it is necessary to indicate that the environment variables **DB_HOST** of the **back-end/nodejs/.env** and **back-end/go/.env** files, and **TEST_DB_HOST** of the **back-end/nodejs/internal/tests/.env** and **back-end/go/.test.env** files must be related to the database **services** defined in the **docker-compose.yml** file.

The **docker-compose.yml** file contains the database services:

```
services:
  ...

  db:
    container_name: db
    build:
      context: ./mongodb
      dockerfile: Dockerfile
    env_file:
      - ./mongodb/.env
    ...

  test-db:
    container_name: test-db
    build:
      context: ./mongodb
      dockerfile: Dockerfile
    env_file:
      - ./mongodb/.test.env
    ...
```

**Development**

The **mongodb/.env** file contains the database environment variables:

```
MONGO_INITDB_ROOT_USERNAME=user
MONGO_INITDB_ROOT_PASSWORD=password
MONGO_INITDB_DATABASE=db
```

The **back-end/nodejs/.env** file contains the database environment variables:

```
DB_USERNAME=user
DB_PASSWORD=password
DB_HOST=db
DB_PORT=27017
DB_NAME=db
```

The **back-end/go/.env** file contains the database environment variables:

```
DB_USERNAME=user
DB_PASSWORD=password
DB_HOST=db
DB_PORT=27017
DB_NAME=db
```

**Test**

The **mongodb/.test.env** file contains the database environment variables:

```
MONGO_INITDB_ROOT_USERNAME=user
MONGO_INITDB_ROOT_PASSWORD=password
MONGO_INITDB_DATABASE=test-db
```

The **back-end/nodejs/internal/tests/.env** file contains the database environment variables:

```
TEST_DB_USERNAME=user
TEST_DB_PASSWORD=password
TEST_DB_HOST=test-db
TEST_DB_PORT=27017
TEST_DB_NAME=test-db
```

The **back-end/go/.test.env** file contains the database environment variables:

```
TEST_DB_USERNAME=user
TEST_DB_PASSWORD=password
TEST_DB_HOST=test-db
TEST_DB_PORT=27017
TEST_DB_NAME=test-db
```

**Important note**

After the project has been successfully executed, it is possible to check the data of the development and test databases resulting from the operations carried out at a command prompt with access to instructions directed to Docker:

```
$ docker exec -it <The id of the container of the corresponding database> /bin/bash
```

```
$ mongo
```

To do this, we must **always** inform the username and password that were previously defined by the database environment variables prior to accessing data.

```
$ use admin
```

```
$ db.auth(<Username>, <Password>)
```

```
$ show dbs;
```

```
$ use <Database name>;
```

In the case of the envinronment variables are kept as they were delivered, if the **id** of the container corresponds to the service named **db**, the data are obtained from the development database:

```
$ use db;
```

On the other hand, if the **id** of the container corresponds to the service named **test-db**, the data are obtained from the test database:

```
$ use test-db;
```

## 4 - How to execute the project?

**In the case of the environment variables of all the .env and .test.env files from all directories are kept as they were delivered I strongly believe that it will not be necessary any change before executing the project.**.

**Prior** to run the project, it is first necessary to configure the IP address (*host*) configured by Docker so that the **microservice 2** can communicate with the **microservice 1** using gRPC.

The *host* corresponds to the value informed when executing a command at a command prompt with access to instructions directed to Docker:

```
$ docker-machine ip
```

(P.S. Because of the dependencies related to the back-end services, it may take some time to them attach to other services properly. Then, to confirm if everything is up and running ok execute the command *docker container ls -a*)

After that, navigate to the project's root directory where the **docker-compose.yml** file is, and assign the *host* to the **GRPC_SERVER_HOST_MS_1** variable in the **back-end_2** service:

```
  ...
  back-end_2:
    container_name: back-end_2
    ...
    environment:
      - GRPC_SERVER_HOST_MS_1=<IP address configured by Docker>
    ...
```

Still at a command prompt with access to instructions directed to Docker where the docker-compose.yml file is, run the command:

```
$ docker-compose up -d
```

If there are no errors, the API *endpoints* will be accessed using the address composed by the *host* and the HTTP server port **8082** (P.S. The **microservice 2** will communicate with the **microservice 1** through the gRPC server *port* **50051**). For example:

```
http://{host}:8082
```

In continuity, suppose the *host* is: 192.168.99.100. As a result, the API requests can be performed through a front-end client or test tool like Postman using the address:

```
http://192.168.99.100:8082
```

In addition, it is also worth emphasizing that the entire configuration related to **Docker** was evaluated in this documentation based on the **DockerToolbox** tool for Windows.

## 5 - How to use the API *endpoints*?

The API requests related to the **microservice 2** are performed through the HTTP server port **8082** and the API responses can be viewed by means of a **front-end** client or test tool, for example **Postman**.

In what follows, there is a guide that includes API requests for creating, obtaining, updating and deleting data from the database.

(P.S. Before checking the following examples, consider that no data is recorded prior to this explanation.)

### Management of Users

#### Creation of a User

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8082/user
```

```
*application/json*

Body: {
    "first_name": "User1",
    "last_name": "User1",
    "date_of_birth": {
        "year": 1990,
        "month": 1,
        "day": 2
    }
}
```

Response:

```
Code: 200 OK - In the case of the user is successfully created.
```

```
*application/json*

Body: {
    "id": "5e54be952c06e1046199cbf8",
    "first_name": "User1",
    "last_name": "User1",
    "date_of_birth": {
        "year": 1990,
        "month": 1,
        "day": 2
    }
}
```

#### Listing of Users

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/users
```

Response:

```
Code: 200 OK - In the case of the list of all users is successfully obtained.
```

```
*application/json*

Body: {
    "users": [
        {
            "id": "5e54be952c06e1046199cbf8",
            "first_name": "User1",
            "last_name": "User1",
            "date_of_birth": {
                "year": 1990,
                "month": 1,
                "day": 2
            }
        }
    ]
}
```

#### Obtainment of a User by its id

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/users/5e54be952c06e1046199cbf8
```

Response:

```
Code: 200 OK - In the case of the user is successfully obtained.
```

```
*application/json*

Body: {
    "id": "5e54be952c06e1046199cbf8",
    "first_name": "User1",
    "last_name": "User1",
    "date_of_birth": {
        "year": 1990,
        "month": 1,
        "day": 2
    }
}
```

#### Updating of a User by its id

Request:

```
Method: HTTP PUT
```

```
URL: http://{host}:8082/users/5e54be952c06e1046199cbf8
```

```
*application/json*

Body: {
    "first_name": "User2",
    "last_name": "User2",
    "date_of_birth": {
        "year": 1990,
        "month": 11,
        "day": 27
    }
}
```

Response:

```
Code: 200 OK - In the case of the user is successfully updated.
```

```
*application/json*

Body: {
    "id": "5e54be952c06e1046199cbf8",
    "first_name": "User2",
    "last_name": "User2",
    "date_of_birth": {
        "year": 1990,
        "month": 11,
        "day": 27
    }
}
```

#### Deletion of a User by its id

Request:

```
Method: HTTP DELETE
```

```
URL: http://{host}:8082/users/5e54be952c06e1046199cbf8
```

Response:

```
Code: 200 OK - In the case of the user is successfully deleted.
```

```
*application/json*

Body: {
    "id": "5e54be952c06e1046199cbf8",
    "first_name": "User2",
    "last_name": "User2",
    "date_of_birth": {
        "year": 1990,
        "month": 11,
        "day": 27
    }
}
```

### Management of Products

#### Creation of a Product

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8082/product
```

```
*application/json*

Body: {
    "price_in_cents": 100,
    "title": "Blue Pen",
    "description": "A pen with blue ink"
}
```

Response:

```
Code: 200 OK - In the case of the product is successfully created.
```

```
*application/json*

Body: {
    "id": "5e5584244007b9649b6837c7",
    "price_in_cents": 100,
    "title": "Blue Pen",
    "description": "A pen with blue ink"
}
```

#### Listing of Products

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/products
```

```
Header: X-USER-ID <The id of a user> (Optional)
```

Response:

```
Code: 200 OK - In the case of the list of all products is successfully obtained.
```

```
*application/json*

Body: {
    "products": [
        {
            "id": "5e5584244007b9649b6837c7",
            "price_in_cents": 100,
            "title": "Blue Pen",
            "description": "A pen with blue ink"
        }
    ]
}
```

or even, when there is a discount:

```
*application/json*

Body: {
    "products": [
        {
            "id": "5e5584244007b9649b6837c7",
            "price_in_cents": 95,
            "title": "Blue Pen",
            "description": "A pen with blue ink",
            "discount": {
                "pct": 5,
                "value_in_cents": 5
            }    
        }
    ]
}
```

(P.S. As indicated above, the scheme of changing API request bodies when there is a discount that **is greater than zero** is performed both when obtaining a single product and the list of all products.)

#### Obtainment of a Product by its id

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/products/5e5584244007b9649b6837c7
```

```
Header: X-USER-ID <The id of a user> (Optional)
```

Response:

```
Code: 200 OK - In the case of the product is successfully obtained.
```

```
*application/json*

Body: {
    "id": "5e5584244007b9649b6837c7",
    "price_in_cents": 100,
    "title": "Blue Pen",
    "description": "A pen with blue ink"
}
```

or even, when there is a discount:

```
*application/json*

Body: {
    "id": "5e5584244007b9649b6837c7",
    "price_in_cents": 95,
    "title": "Blue Pen",
    "description": "A pen with blue ink",
    "discount": {
        "pct": 5,
        "value_in_cents": 5
    }
}
```

**Important note**

Whenever the API request to obtain one product is performed and there is a discount that **is greater than zero**, the API response body of the related product is modified.

First of all, the *price_in_cents* field must be adjusted since a reduction of value is necessary. After that, a new field named **discount** is elaborated along with values for the **pct** and **value_in_cents** fields that will be also presented.

In that example, since the original value of a product is 100 and the discount is 5, the new *price_in_cents* field is 95.

#### Updating of a Product by its id

Request:

```
Method: HTTP PUT
```

```
URL: http://{host}:8082/products/5e5584244007b9649b6837c7
```

```
*application/json*

Body: {
    "price_in_cents": 200,
    "title": "Red Pen",
    "description": "A pen with red ink"
}
```

Response:

```
Code: 200 OK - In the case of the product is successfully updated.
```

```
*application/json*

Body: {
    "id": "5e5584244007b9649b6837c7",
    "price_in_cents": 200,
    "title": "Red Pen",
    "description": "A pen with red ink"
}
```

#### Deletion of a Product by its id

Request:

```
Method: HTTP DELETE
```

```
URL: http://{host}:8082/products/5e5584244007b9649b6837c7
```

Response:

```
Code: 200 OK - In the case of the product is successfully deleted.
```

```
*application/json*

Body: {
    "id": "5e5584244007b9649b6837c7",
    "price_in_cents": 200,
    "title": "Red Pen",
    "description": "A pen with red ink"
}
```

### Management of Promotions

#### Creation of a Promotion

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8082/promotion
```

```
*application/json*

Body: {
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 10
}
```

or even, if there is one or more products to be associated with the promotion:

```
*application/json*

Body: {
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 10
    "products": [
        "5e5584244007b9649b6837c7"
    ]
}
```

Response:

```
Code: 200 OK - In the case of the promotion is successfully created.
```

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d62b",
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 10
}
```

or even, if there is one or more products associated with the promotion:

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d62b",
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 10,
    "products": [
        "5e5584244007b9649b6837c7"
    ]
}
```

**Important note**

The inclusion of promotional products is a procedure that can be performed when creating, or even updating, a promotion.

With this in mind, a promotion is only valid for products that are associated with it. That is, if there is a discount related to a particular promotion, it is **only** applicable to products that are associated with it. Because of that, all other products that are also registered, but are not related to any promotion, will not have any reduction of values.

#### Listing of Promotions

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/promotions
```

Response:

```
Code: 200 OK - In the case of the list of all promotions is successfully obtained.
```

```
Body: {
    "promotions": [
        {
            "id": "5e6383a641261e8e7b49d62b",
            "code": "DISCOUNTEDDATES",
            "title": "Discounted Dates",
            "description": "The promotion of discounted dates",
            "max_discount_pct": 10
        }
    ]
}
```

or even, if there is one or more products associated with the promotion:

```
Body: {
    "promotions": [
        {
            "id": "5e6383a641261e8e7b49d62b",
            "code": "DISCOUNTEDDATES",
            "title": "Discounted Dates",
            "description": "The promotion of discounted dates",
            "max_discount_pct": 10,
            "products": [
                "5e5584244007b9649b6837c7"
            ]
        }
    ]
}
```

#### Obtainment of a Promotion by its id

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/promotions/5e6383a641261e8e7b49d62b
```

Response:

```
Code: 200 OK - In the case of the promotion is successfully obtained.
```

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d62b",
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 10
}
```

or even, if there is one or more products associated with the promotion:

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d62b",
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 10,
    "products": [
        "5e5584244007b9649b6837c7"
    ]
}
```

#### Updating of a Promotion by its id

Request:

```
Method: HTTP PUT
```

```
URL: http://{host}:8082/promotions/5e6383a641261e8e7b49d62b
```

```
*application/json*

Body: {
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 25
}
```

or even, if there is one or more products to be associated with the promotion:

```
*application/json*

Body: {
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 25
    "products": [
        "5e5584244007b9649b6837c7"
    ]
}
```

Response:

```
Code: 200 OK - In the case of the promotion is successfully updated.
```

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d62b",
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 25
}
```

or even, if there is one or more products associated with the promotion:

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d62b",
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 25
    "products": [
        "5e5584244007b9649b6837c7"
    ]
}
```

#### Deletion of a Promotion by its id

Request:

```
Method: HTTP DELETE
```

```
URL: http://{host}:8082/promotions/5e6383a641261e8e7b49d62b
```

Response:

```
Code: 200 OK - In the case of the promotion is successfully deleted.
```

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d62b",
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 25
}
```

or even, if there is one or more products associated with the promotion:

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d62b",
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 25
    "products": [
        "5e5584244007b9649b6837c7"
    ]
}
```

### Management of Discounted Dates

#### Creation of a Discounted Date

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8082/discountedDate
```

```
*application/json*

Body: {
	"title": "Black Friday 2020",
	"description": "The discount of Black Friday 2020",
	"discount_pct": 10,
	"date": {
		"year": 2020,
		"month": 11,
		"day": 27
	}
}
```

Response:

```
Code: 200 OK - In the case of the discounted date is successfully created.
```

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d61c",
    "title": "Black Friday 2020",
    "description": "The discount of Black Friday 2020",
    "discount_pct": 10,
    "date": {
        "year": 2020,
        "month": 11,
        "day": 27
    }
}
```

**Important note**

A discounted date must be configured with a unique combination of values for the **year**, **month** and **day** fields.

In order to create the discount date related to **user's birthday**, the **date** field **must** be configured with the value **0** for the **year**, **month** and **day** fields:

```
*application/json*

Body: {
	"title": "User's Birthday",
	"description": "The discount of user's birthday",
	"discount_pct": 5,
	"date": {
		"year": 0,
		"month": 0,
		"day": 0
	}
}
```

#### Listing of Discounted Dates

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/discountedDates
```

Response:

```
Code: 200 OK - In the case of the list of all discounted dates is successfully obtained.
```

```
Body: {
    "discountedDates": [
        {
            "id": "5e6383a641261e8e7b49d61c",
            "title": "Black Friday 2020",
            "description": "The discount of Black Friday 2020",
            "discount_pct": 10,
            "date": {
                "year": 2020,
                "month": 11,
                "day": 27
            }
        }
    ]
}
```

#### Obtainment of a Discounted Date by its id

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/discountedDates/5e6383a641261e8e7b49d61c
```

Response:

```
Code: 200 OK - In the case of the discounted date is successfully obtained.
```

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d61c",
    "title": "Black Friday 2020",
    "description": "The discount of Black Friday 2020",
    "discount_pct": 10,
    "date": {
        "year": 2020,
        "month": 11,
        "day": 27
    }
}
```

#### Updating of a Discounted Date by its id

Request:

```
Method: HTTP PUT
```

```
URL: http://{host}:8082/discountedDates/5e6383a641261e8e7b49d61c
```

```
*application/json*

Body: {
	"title": "Black Friday 2020",
	"description": "The discount of Black Friday 2020",
	"discount_pct": 35,
	"date": {
		"year": 2020,
		"month": 11,
		"day": 27
	}
}
```

Response:

```
Code: 200 OK - In the case of the discounted date is successfully updated.
```

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d61c",
    "title": "Black Friday 2020",
    "description": "The discount of Black Friday 2020",
    "discount_pct": 35,
    "date": {
        "year": 2020,
        "month": 11,
        "day": 27
    }
}
```

#### Deletion of a Discounted Date by its id

Request:

```
Method: HTTP DELETE
```

```
URL: http://{host}:8082/discountedDates/5e6383a641261e8e7b49d61c
```

Response:

```
Code: 200 OK - In the case of the discounted date is successfully deleted.
```

```
*application/json*

Body: {
    "id": "5e6383a641261e8e7b49d61c",
    "title": "Black Friday 2020",
    "description": "The discount of Black Friday 2020",
    "discount_pct": 35,
    "date": {
        "year": 2020,
        "month": 11,
        "day": 27
    }
}
```

## 6 - Tests

In order to test the solution a few **test sets** were developed.

The tests will be executed on the running **back-end** containers. 

To do this, at a command prompt with access to instructions directed to Docker, launch a bash terminal within the related **back-end** container:

```
$ docker exec -it <The id of the container of the corresponding back-end application> /bin/bash
```

### 6.1 Microservice 1

(P.S. These tests involve removing the record related to the promotion of discounted dates from the test database if it already exists.)

### 6.1.1 Services

These tests are related to the implementation of services in order to obtain one or more products with the possibility of discounts.

To execute them on the running **back-end_1** container, navigate to the **app/internal/grpc/services/impl_test** directory.

So, if you prefer to evaluate all tests at once, run the command:

```
$ npm run test
```

However, it is also possible to run each test separately using the commands:

**Tests of the implementation of services directed to Products**

```
$ npm test -- -t "TestGetAllProducts.WithoutAnyDiscountOfDates"
```

```
$ npm test -- -t "TestGetAllProducts.WithOnlyTheDiscountOfUser\'sBirthday"
```

```
$ npm test -- -t "TestGetAllProducts.WithOnlyTheDiscountOfOtherDiscountedDate"
```

```
$ npm test -- -t "TestGetAllProducts.WithTheMaximumDiscountOfDates"
```

```
$ npm test -- -t "TestGetProduct.WithoutAnyDiscountOfDates"
```

```
$ npm test -- -t "TestGetProduct.WithOnlyTheDiscountOfUser\'sBirthday"
```

```
$ npm test -- -t "TestGetProduct.WithOnlyTheDiscountOfOtherDiscountedDate"
```

```
$ npm test -- -t "TestGetProduct.WithTheMaximumDiscountOfDates"
```

### 6.2 Microservice 2

(P.S. These tests involve creating, editing and removing records from the test database.)

### 6.2.1 Database

The tests that were developed are related to **CRUD** operations (*create*, *read*, *update* and *delete*) in the test database.

To execute them on the running **back-end_2** container, navigate to the **app/internal/mongodb_test** directory.

So, if you prefer to evaluate all tests at once, run the command:

```
$ go test -v
```

However, it is also possible to run each test separately using the commands:

**Tests of the CRUD operations directed to Users**

```
$ go test -v -run=TestCreateUser
```

```
$ go test -v -run=TestGetAllUsers
```

```
$ go test -v -run=TestGetUser
```

```
$ go test -v -run=TestUpdateUser
```

```
$ go test -v -run=TestDeleteUser
```

**Tests of the CRUD operations directed to Products**

```
$ go test -v -run=TestCreateProduct
```

```
$ go test -v -run=TestGetAllProducts
```

```
$ go test -v -run=TestGetProduct
```

```
$ go test -v -run=TestUpdateProduct
```

```
$ go test -v -run=TestDeleteProduct
```

**Tests of the CRUD operations directed to Promotions**

```
$ go test -v -run=TestCreatePromotion
```

```
$ go test -v -run=TestGetAllPromotions
```

```
$ go test -v -run=TestGetPromotion
```

```
$ go test -v -run=TestUpdatePromotion
```

```
$ go test -v -run=TestDeletePromotion
```

**Tests of the CRUD operations directed to Discounted Dates**

```
$ go test -v -run=TestCreateDiscountedDate
```

```
$ go test -v -run=TestGetAllDiscountedDates
```

```
$ go test -v -run=TestGetDiscountedDate
```

```
$ go test -v -run=TestUpdateDiscountedDate
```

```
$ go test -v -run=TestDeleteDiscountedDate
```

### 6.2.2 Services

These tests are related to the implementation of services to create, read, update and delete users, products, promotions and discounted dates.

To execute them on the running **back-end_2** container, navigate to the **app/internal/grpc/services/impl_test** directory.

So, if you prefer to evaluate all tests at once, run the command:

```
$ go test -v
```

Nevertheless, it is also possible to run each test separately using the commands:

**Tests of the implementation of services directed to Users**

```
$ go test -v -run=TestCreateUser
```

```
$ go test -v -run=TestGetAllUsers
```

```
$ go test -v -run=TestGetUser
```

```
$ go test -v -run=TestUpdateUser
```

```
$ go test -v -run=TestDeleteUser
```

**Tests of the implementation of services directed to Products**

```
$ go test -v -run=TestCreateProduct
```

```
$ go test -v -run=TestGetAllProducts
```

```
$ go test -v -run=TestGetProduct
```

```
$ go test -v -run=TestUpdateProduct
```

```
$ go test -v -run=TestDeleteProduct
```


**Tests of the implementation of services directed to Promotions**

```
$ go test -v -run=TestCreatePromotion
```

```
$ go test -v -run=TestGetAllPromotions
```

```
$ go test -v -run=TestGetPromotion
```

```
$ go test -v -run=TestUpdatePromotion
```

```
$ go test -v -run=TestDeletePromotion
```

**Tests of the implementation of services directed to Discounted Dates**

```
$ go test -v -run=TestCreateDiscountedDate
```

```
$ go test -v -run=TestGetAllDiscountedDates
```

```
$ go test -v -run=TestGetDiscountedDate
```

```
$ go test -v -run=TestUpdateDiscountedDate
```

```
$ go test -v -run=TestDeleteDiscountedDate
```

## 7 - Project Dynamics

In what follows, there is a brief account of how the solution works in practice meeting the requirements specified in the comments of the code challenge.

The listing of products, or even the obtainment of a single product, with discounts provided by the promotion of discounted dates are dependent on the following factors:

1. Firstly, a promotion **must** be created with the code **DISCOUNTEDDATES**, a maximum discount percentage and a list of promotional products.

2. Secondly, the discounted dates **must** be accordingly configured with the discount amount and the date in which the discount is applicable based on the year, month and day.

3. Lastly, when an API request is performed to obtain one or more products, the discount will be evaluated **always** based on the data of the current date and the discounted dates registered. Additionally, if it is also informed the **X-USER-ID** in the *header* of the corresponding request and it contains the *id* of a given user, the discount will be analyzed based on the data of the corresponding user as well.

As previou explained, the discount is **always** limited by the value of the **max_discount_pct** field of the related promotion. This way, if the calculated discount exceeds the maximum discount percentage of the related promotion, it will be adjusted to meet this criteria.

### Data Record

First, it is necessary to illustrate the data to be registered before obtaining one or more products with the possibility of discounts.

(P.S. Consider that no data is recorded prior to this explanation.)

#### Creation of Users

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8082/user
```

```
*application/json*

Body: {
    "first_name" : "User1",
    "last_name" : "User1",
    "date_of_birth": {
        "year": 1990,
        "month": 1,
        "day": 2
    }
}
```

Response:

```
Code: 200 OK - In the case of the user is successfully created.
```

```
*application/json*

Body: {
    "id": "5e5583684007b9649b6837c5",
    "first_name": "User1",
    "last_name": "User1",
    "date_of_birth": {
        "year": 1990,
        "month": 1,
        "day": 2
    }
}
```

Database:

```
> db.users.find({"_id": ObjectId("5e5583684007b9649b6837c5")}).pretty();
{
    "_id" : ObjectId("5e5583684007b9649b6837c5"),
    "first_name" : "User1",
    "last_name" : "User1",
    "date_of_birth" : {
        "year" : 1990,
        "month" : 1,
        "day" : 2
    }
}
```

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8082/user
```

```
*application/json*

Body: {
    "first_name": "User2",
    "last_name": "User2",
    "date_of_birth": {
        "year": 1990,
        "month": 11,
        "day": 27
    }
}
```

Response:

```
Code: 200 OK - In the case of the user is successfully created.
```

```
*application/json*

Body: {
    "id": "5e5583e84007b9649b6837c6",
    "first_name": "User2",
    "last_name": "User2",
    "date_of_birth": {
        "year": 1990,
        "month": 11,
        "day": 27
    }
}
```

Database:

```
> db.users.find({"_id": ObjectId("5e5583e84007b9649b6837c6")}).pretty();
{
    "_id" : ObjectId("5e5583e84007b9649b6837c6"),
    "first_name" : "User2",
    "last_name" : "User2",
    "date_of_birth" : {
        "year" : 1990,
        "month" : 11,
        "day" : 27
    }
}
```

#### Creation of a Product

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8082/product
```

```
*application/json*

Body: {
    "price_in_cents": 100,
    "title": "Blue Pen",
    "description": "A pen with blue ink"
}
```

Response:

```
Code: 200 OK - In the case of the product is successfully created.
```

```
*application/json*

Body: {
    "id": "5e5584244007b9649b6837c7",
    "price_in_cents": 100,
    "title": "Blue Pen",
    "description": "A pen with blue ink"
}
```

Database:

```
> db.products.find({"_id": ObjectId("5e5584244007b9649b6837c7")}).pretty();
{
    "_id" : ObjectId("5e5584244007b9649b6837c7"),
    "price_in_cents" : 100,
    "title" : "Blue Pen",
    "description" : "A pen with blue ink"
}
```

#### Creation of the Promotion of Discounted Dates

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8082/promotion
```

```
*application/json*

Body: {
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 10,
    "products": [
        "5e5584244007b9649b6837c7"
    ]
}
```

Response:

```
Code: 200 OK - In the case of the promotion is successfully created.
```

```
*application/json*

Body: {
    "id": "5e5599374007b9649b6837ca",
    "code": "DISCOUNTEDDATES",
    "title": "Discounted Dates",
    "description": "The promotion of discounted dates",
    "max_discount_pct": 10,
    "products": [
        "5e5584244007b9649b6837c7"
    ]
}
```

Database:

```
> db.promotions.find({"_id": ObjectId("5e5599374007b9649b6837ca")}).pretty();
{
    "_id" : ObjectId("5e5599374007b9649b6837ca"),
    "code" : "DISCOUNTEDDATES",
    "title" : "Discounted Dates",
    "description" : "The promotion of discounted dates",
    "max_discount_pct" : 10,
    "products" : [
        "5e5584244007b9649b6837c7"
    ]
}
```

#### Creation of Discounted Dates

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8082/discountedDate
```

1st Discounted date

```
*application/json*

Body: {
	"title": "Black Friday 2020",
	"description": "The discount of Black Friday 2020",
	"discount_pct": 10,
	"date": {
		"year": 2020,
		"month": 11,
		"day": 27
	}
}
```

Response:

```
Code: 200 OK - In the case of the discounted date is successfully created.
```

```
*application/json*

Body: {
    "id": "5e55984c4007b9649b6837c9",
    "title": "Black Friday 2020",
    "description": "The discount of Black Friday 2020",
    "discount_pct": 10,
    "date": {
        "year": 2020,
        "month": 11,
        "day": 27
    }
}
```

Database:

```
> db.discountedDates.find({"_id": ObjectId("5e55984c4007b9649b6837c9")}).pretty();
{
    "_id" : ObjectId("5e55984c4007b9649b6837c9"),
    "title" : "Black Friday 2020",
    "description" : "The discount of Black Friday 2020",
    "discount_pct" : 10,
    "date" : {
        "year" : 2020,
        "month" : 11,
        "day" : 27
    }
}
```

2nd Discounted date

```
*application/json*

Body: {
	"title": "User's Birthday",
	"description": "The discount of user's birthday",
	"discount_pct": 5,
	"date": {
		"year": 0,
		"month": 0,
		"day": 0
	}
}
```

Response:

```
Code: 200 OK - In the case of the discounted date is successfully created.
```

```
*application/json*

Body: {
    "id": "5e5596d04007b9649b6837c8",
    "title": "User's Birthday",
    "description": "The discount of user's birthday",
    "discount_pct": 5,
    "date": {
        "year": 0,
        "month": 0,
        "day": 0
    }
}
```

Database:

```
> db.discountedDates.find({"_id": ObjectId("5e5596d04007b9649b6837c8")}).pretty();
{
    "_id" : ObjectId("5e5596d04007b9649b6837c8"),
    "title" : "User's Birthday",
    "description" : "The discount of user's birthday",
    "discount_pct" : 5,
    "date" : {
        "year" : 0,
        "month" : 0,
        "day" : 0
    }
}
```

Below, there a simple explanation of a few possible scenarios without and with discounts when obtaining one or more products taking into consideration the previous data registered.

### 7.2.1 Products without any Discount of Dates

It may happen due to some circumstances, such as the three described below:

1. The **microservice 1** is not active.

2. The **microservice 1** is active, but the **X-USER-ID** is not informed. In this case, there will be no evaluation of the discount of user's birthday and, therefore, the discount is only evaluated considering other discounted dates.

As an example, suppose that the current date is **1-1-2020** and it is not a discounted date. (The previously discount dates are **0-0-0** and **27-11-2020**.) This way, no discount is applicable.

3. The **microservice 1** is active and the **X-USER-ID** is informed. Because of this, consider two situations:

     3.1. If the **X-USER-ID** does not contain an **id** of a user registered, then the second circumstance described above will be investigated.

     3.2. On the other hand, if the **X-USER-ID** contains an **id** of a registered user, for example the **id** of the user named **User1**, the discount of user's birthday will be evaluated. As an example, suppose that the current date is **1-1-2020** and has no relation to the date of user's birth. (The date of User1's birth is **2-1-1990**.) As before, the second circumstance described above will be investigated.

### 7.2.2 Products with the Discount of User's Birthday

It may happen whenever the **microservice 1** is active and the API request is performed in a date that is related to the birth date of some user registered.

As an example, suppose that the current date is **2-1-2020** and, thus, it is related to the date of User1's birth that is **2-1-1990**. Because of that, if the **id** of the **User1** is assigned to **X-USER-ID** in the *Header* of the API request, the discount of user's birthday is applicable. In this case, the discount is **5** because it is the discount percentage of the discounted date **0-0-0**.

Database:

```
> db.products.find({"_id": ObjectId("5e5584244007b9649b6837c7")}).pretty();
{
    "_id" : ObjectId("5e5584244007b9649b6837c7"),
    "price_in_cents" : 100,
    "title" : "Blue Pen",
    "description" : "A pen with blue ink"
}
```

```
> db.discountedDates.find({"_id": ObjectId("5e5596d04007b9649b6837c8")}).pretty();
{
    "_id" : ObjectId("5e5596d04007b9649b6837c8"),
    "title" : "User's Birthday",
    "description" : "The discount of user's birthday",
    "discount_pct" : 5,
    "date" : {
        "year" : 0,
        "month" : 0,
        "day" : 0
    }
}
```

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/products
```

```
Header: X-USER-ID 5e5583684007b9649b6837c5
```

Response:

```
Code: 200 OK - In the case of the list of all products is successfully obtained.
```

```
*application/json*

Body: {
    "products": [
        {
            "id": "5e5584244007b9649b6837c7",
            "price_in_cents": 95,
            "title": "Blue Pen",
            "description": "A pen with blue ink",
            "discount": {
                "pct": 5,
                "value_in_cents": 5
            }
        }
    ]
}
```

### 7.2.3 Products with only the Discount of Other Discounted Dates

It may happen whenever the **microservice 1** is active and the API request is performed without the **X-USER-ID**.

Following this idea, even if the API request is performed in a date related to the birth date of some user registered, the discount of user's birthday is not evaluated since no **X-USER-ID** is not informed.

As an example, suppose that the current date is **27-11-2020** and it is related to the discounted date of *Black Friday 2020*. In this case, the discount is **10** because it is the discount percentage of the discounted date **27-11-2020**.

Database:

```
> db.discountedDates.find({"_id": ObjectId("5e55984c4007b9649b6837c9")}).pretty();
{
    "_id" : ObjectId("5e55984c4007b9649b6837c9"),
    "title" : "Black Friday 2020",
    "description" : "The discount of Black Friday 2020",
    "discount_pct" : 10,
    "date" : {
        "year" : 2020,
        "month" : 11,
        "day" : 27
    }
}
```

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/products
```

Response:

```
Code: 200 OK - In the case of the list of all products is successfully obtained.
```

```
*application/json*

Body: {
    "products": [
        {
            "id": "5e5584244007b9649b6837c7",
            "price_in_cents": 90,
            "title": "Blue Pen",
            "description": "A pen with blue ink",
            "discount": {
                "pct": 10,
                "value_in_cents": 10
            }
        }
    ]
}
```

### 7.2.4 Products with Maximum Discount of the Discounted Dates

It may happen whenever the **microservice 1** is active and the API request is performed in a date that, at tha same time, is related to the birth date of some user registered, as well as other discounted date that is different from **0-0-0**.

As an example, suppose that the current date is **27-11-2020**. This date is related to the User2's birth that is **27-11-1990**, as well as it is related to the discounted date of *Black Friday 2020* that is **27-11-2020**.

This way, if the **id** of the **User2** is assigned to **X-USER-ID** in the *Header* of the API request, the discount is the sum of the amounts defined for both the discounted dates of user's birthday (**0-0-0**) and *Black Friday 2020* (**27-11-2020**), that is **15**.

Nevertheless, the maximum discount percentage that is applied is **always** limited by the **max_discount_pct** field of the promotion of discounted dates. In this case, the **max_discount_pct** is **10**.

Database:

```
> db.promotions.find({"_id": ObjectId("5e5599374007b9649b6837ca")}).pretty();
{
    "_id" : ObjectId("5e5599374007b9649b6837ca"),
    "code" : "DISCOUNTEDDATES",
    "title" : "Discounted Dates",
    "description" : "The promotion of discounted dates",
    "max_discount_pct" : 10,
    "products" : [
        "5e5584244007b9649b6837c7"
    ]
}
```

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8082/products
```

```
Header: X-USER-ID 5e5583e84007b9649b6837c6
```

Response:

```
Code: 200 OK - In the case of the list of all products is successfully obtained.
```

```
*application/json*

Body: {
    "products": [
        {
            "id": "5e5584244007b9649b6837c7",
            "price_in_cents": 90,
            "title": "Blue Pen",
            "description": "A pen with blue ink",
            "discount": {
                "pct": 10,
                "value_in_cents": 10
            }
        }
    ]
}
```