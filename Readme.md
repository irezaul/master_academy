CouchBase deafult port 8093

MYSQL   COUCHBASE
Database   Bucket
Table   none/type
Row   Json Document
3306   8093
DataCapacity  20MB

REST API-php

A= Application
P= Programming

Rest API to Couchbase

Resquest semd by via Post-

Post_Url http://localhost:8093/query/service
KEY - Content-Type & VALUE- application/json
BODY - 
{
   "statement":"SELECT * FROM Database/Bucket WHERE type='request';"
}