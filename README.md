# Service Template

This is a project for me to figure out what is the best program layout for a business Microservice project in GRPC and Go. This is not the final version, which will supports transaction. Adding transaction will bring more complex into the project. In case there is no need for transaction, you can use it as the project template. It also doesn't expose the service through GRPC interface. For most people, I do recommend the full featured version [here](link).

##How to use this project
This project is best to be used as a starting point when creating a business gRPC Microservice project. It already has rich functionality to build in and is working, so there is no need to start from scratch. The purpose of it is not including many functions into it, but to build a flexible foundation, which can be extended easily. When I design and build the project, I followed SOLID design in object-oriented programing and Go's concise coding style, so it can be used as a living example of application design and coding style when you try to enforce it in your organization or your code. You may have different different design and coding style with what I have here, it is easy to change it to fit into your needs. 

## Use as a template to start a service project
### Functional Feature:
1. Support different implementation of database by changing configuration file ( Currently it supports MySQL and CouchDB) (The database can be SQL and NoSQL)
2. Support different implementation of logging lib by changing configuration file ( Currently it supports ZAP and Logrus)( The logger lib need to support common interface similar to ZAP and Logrus)
3. Support business level transaction inside a service without database transaction code in business layer( it doesn't suppose nested transaction or transaction across multiple Microservice)  
4. Using Dependency Injection to create concrete types and wire the whole application together.
5. Application configurations are saved in yaml file and can be changed easily. 

### Design Feature:
##### 1. Programming on interface 
Access outside functions through interface
Has three layer: use case, model and persistence. Each layer access other layer through interface ( Except for model layer, which doesn't need interface)
##### 2. Create concrete type through Dependency Inject by using factory method pattern
##### 3. Minimize Dependency
Dependency between different layers is only on interface ( except for model, which doesn't have interface)
Interface is defined in top level package and separated from concrete type. 
Each concrete type is defined in a separate sub-package and file 
##### 4. Function Isolation
Isolate different layer by package and file
Isolate each use case by package 
Isolate each implementation ( for example database implementation) by package
##### 5. Open-closed principle
whenever a new feature is added, instead modify existing code, try to add new code
  

### Coding Style:
1. No use of package level variable except in "appcontainer" package
2. Minimize use of constant
Constants are read-only global variables, even though they are better than mutable ones, still should be restricted.
Basically, functions should encapsulate itself. 
3. Log full stack trace for error
4. Errors are only handled on top level ( All other levels should only add information and propagate error to upper level)
5. separation of concerns
business logic, functional requirements (For example, retry, timeout, transaction) and technical implementation ( for example, database, logger ) are different concerns and shouldn't be mixed in one piece of code. 
6. Naming Convention 
function or block level variable should be named according to Go's concise naming convention, but type or interfaces shouldn't. For them, readability overweight concise, you should tell what it for from it's name.   

## Getting Started

### Installation and Setting Up

Don't need to follow all steps in this section to get the code to work. The simplest way is to get the code from github and run it. However, it will throw exception when accesses database. So, I'd recommend you install at least one database ( MySQL is better), then most part of the code will work. 

#### Download Code

```
go get github.com/jfeng45/sericeconfig
```

#### Set Up MySQL

There are two database implementations, MySQL and CouchDB, but most functions are implemented in MySQL. You'd better install at least one of them. 
```
Install MySQL
run SQL script in script folder to create database and table
```
#### Install CouchDB

The code works fine without it. This part just shows the feature of switching database by changing configuration.
 
Installation on [Windows](https://docs.couchdb.org/en/2.2.0/install/windows.html)

Installation on [Linux](https://docs.couchdb.org/en/2.2.0/install/unix.html)

Installation on [Mac](https://docs.couchdb.org/en/2.2.0/install/mac.html)

CouchDB [Example](https://github.com/go-kivik/kivik/wiki/Usage-Examples)

#### Set up CouchDB

```
Access Fauxton through broswer: http://localhost:5984/_utils/# (login with: admin/admin)
Create new database "service_config" in Fauxton
Add the following document to database ( "_id" and "_rev" are generated by database, no need to change it)
{
  "_id": "80a9134c7dfa53f67f6be214e1000fa7",
  "_rev": "4-f45fb8bdd454a71e6ae88bdeea8a0b4c",
  "uid": 10,
  "username": "Tony",
  "department": "IT",
  "created": "2018-02-17T15:04:05-03:00"
}
```
#### Install Cache Service (Another Microservice)

Without it, only calling another Microservice piece won't work, the rest of code works just fine. Please follow instructions in [reservegrpc](https://github.com/jfeng45/reservegrpc) to set up the service.

###S tart Application

#### Start MySQL Server
```
cd [MySQLroot]/bin
mysqld
```

#### Start CouchDB Server
```
It should already been started
```
#### Start Cache Service

Please follow instructions in [reservegrpc](https://github.com/jfeng45/reservegrpc) to start the server.

#### Run main
In "main.go", under main() function, there are two functions "testMySql()" ( which reads configurations from "configs/appConifgDev.yaml") and "testCouchDB()"  which reads from "configs/appConifgProd.yaml") to test MySQL and CouchDB separately.
```
cd [rootOfProject]/cmd
go run main.go
```
## License

[MIT](LICENSE.txt) License



