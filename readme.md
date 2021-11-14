# Simple GO Lang REST API

> Simple RestFul API to create, read, update and delete function. No database implementation.

## Start


``` sh
# Install mux router
go get -u github.com/gorilla/mux
```


## Endpoints

### Get All 
``` sh
GET /books
GET /members
```
### Get Single Item
```sh
GET api/books/{id}
GET api/members/{id}
```

### Delete 
``` sh
DELETE api/books/{id}
DELETE api/members/{id}
```

### Create Book
``` sh
POST /books

# Request sample
# {
#   "isbn":"1231245",
#   "title":"Book X",
#   "author":{"firstname":"Albert",  "lastname":"Camus"}
# }
```

### Update Book
``` sh
PUT api/books/{id}

# Request sample
# {
#   "isbn":"1231245",
#   "title":"Updated Title",
#   "author":{"firstname":"Albert",  "lastname":"Camus"}
# }

```


```sh

## App Info

### Version

1.0.0

### License

This project is licensed under the MIT License
