# SAFE TRAVEL - Backedend

Web service handling RESTful requests.

## Database

Data is stored in a document oriented manner. The databaser is a JSON file storing information for each country. Example row looks as follows:
```
{
    "country": "Bulgaria",
    "others": [
        {
            "code": "Montenegro", "color": "orange"
        },
        {
            "code": "Serbia", "color": "red"
        },
        {
            "code": "Macedonia", "color": "green"
        },
        {
            "code": "Greece", "color": "green"
        },
        {
            "code": "Romania", "color": "red"
        }        
    ]
}
```

## REST API

### GET
For a given country returns json with countries and corresponding color. Example of GET request with curl:
```
curl GET http://localhost:8080/countries/Poland
```

### POST
Allows to add entry to database with given json file. Example of POST request with curl:
```
curl http://localhost:8080/countries \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"country": "Bulgaria",
                "others": [
                    {
                        "code": "Montenegro", "color": "orange"
                    },
                    {
                        "code": "Serbia", "color": "red"
                    },
                    {
                        "code": "Macedonia", "color": "green"
                    },
                    {
                        "code": "Greece", "color": "green"
                    },
                    {
                        "code": "Romania", "color": "red"
                    }        
                ]
            }'
```

### PUT
Update existing entry with given json file. Entry must exists. Example of POST request with curl:
```
curl http://localhost:8080/countries \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"country": "Bulgaria",
                "others": [
                    {
                        "code": "Montenegro", "color": "orange"
                    },
                    {
                        "code": "Serbia", "color": "red"
                    },
                    {
                        "code": "Macedonia", "color": "green"
                    },
                    {
                        "code": "Greece", "color": "green"
                    },
                    {
                        "code": "Romania", "color": "red"
                    }        
                ]
            }'
```
### DELETE
Deletes given country from database. Example of DELETE request with curl:
```
curl DELETE http://localhost:8080/countries/Poland
```
