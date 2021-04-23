# Brief intro
It is a side project that is used to practise my golang api skill, I am planning to develop 
a SwiftUI client project to interact with it.

# API

## Auth

### register

```shell
curl --location --request POST 'localhost:7777/api/v1/register' \
--data-raw '{
    "email": "changweitu@gmail.com",
    "password": "123456"
}'
{"message":"OK","result":1,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg0NjQyLCJpc3MiOiJnby1uZXdzIn0.Q_3qBVN2nTYpZ50YPrs32ciNQDnjusyPbt5REOOqzD8"}
```
### login

```shell
curl --location --request POST 'localhost:7777/api/v1/login' \
--data-raw '{
    "email": "changweitu@gmail.com",
    "password": "123456"
}'
{"message":"OK","result":1,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg0OTI1LCJpc3MiOiJnby1uZXdzIn0.BovWURxcA0lhzV6dYlSxamUiJDgVx58YNJtIwp4iNCk"}%  
```

### change password

```shell
curl --location --request POST 'localhost:7777/api/v1/changePassword' \
--data-raw '{
    "password":"000000"
}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg2NDQxLCJpc3MiOiJnby1uZXdzIn0.skk1eRMpeFM21q1e2VfjgAaesRja65ej6gF_VpbZs0A"  
{"message":"OK","result":1}
```

## User

### get user
```shell
curl --location --request GET 'localhost:7777/api/v1/users/1' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg1MTU3LCJpc3MiOiJnby1uZXdzIn0.6Tmy1O0pLj5UTyiljx7IWco20jS1laiT4EmAD6pJtKY"  
{"data":{"username":"","avatar":"","role":1,"id":1},"message":"OK","result":1}
```

### get users

```shell
curl --location --request GET 'localhost:7777/api/v1/users?pageSize=3&pageNum=0' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg2NDQxLCJpc3MiOiJnby1uZXdzIn0.skk1eRMpeFM21q1e2VfjgAaesRja65ej6gF_VpbZs0A"
{"data":[{"username":"tu_tu","avatar":"url","role":2,"id":1}],"message":"OK","result":1,"total":1}
```

### edit user

```shell
curl --location --request PUT 'localhost:7777/api/v1/users/1' \
--data-raw '{
    "role": 2,
    "username": "tu_tu",
    "avatar": "url"
}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg1MTU3LCJpc3MiOiJnby1uZXdzIn0.6Tmy1O0pLj5UTyiljx7IWco20jS1laiT4EmAD6pJtKY"
{"message":"OK","result":1} 
```
## Post

### create post (post_type: 1 text, 2 image, 3 url)

```shell
curl --location --request POST 'localhost:7777/api/v1/posts' \
--data-raw '{
    "title": "my first post",
    "post_type": 1,
    "content": "Hello everyone, it is my first post"
}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg3OTYzLCJpc3MiOiJnby1uZXdzIn0.5llD3LSSFBS282070dEPhbBNWjW0FdlYaZuDzRJ5nSs"
{"message":"OK","result":1}
```
