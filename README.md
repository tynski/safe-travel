# SAFE TRAVEL - Backedend

Web service handling RESTful requests.

## GET
```
GET /countries/:country 
```
For a given country returns json with countries and corresponding color.

## POST
```
POST /countries
```
Allows to add entry to database with given json file. Example of json file:
```
{"country":"Poland","Neighbors":{"Germany":"Red","Slovakia":"Green"}}
```

## PUT
```
PUT /countries
```
Update existing entry with given json file. Entry must exists.

## DELETE
```
DELETE /countries/:country
```
Deletes given country from database.