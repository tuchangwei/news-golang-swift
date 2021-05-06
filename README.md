# Brief intro
It is a side project that is used to practise my golang api skill, I am planning to develop 
a SwiftUI client project to interact with it.

Note: If you want to run the app, besides cofig the settings under `config/config.ini`, you need to put your aws credential file under `~/.aws/credentials` , please check [this](https://asanchez.dev/blog/upload-to-amazon-s3-with-golang/) for more information.

# API

## Auth

### register

request:
```shell
curl --location --request POST 'localhost:7777/api/v1/register' \
--data-raw '{
    "email": "changweitu@gmail.com",
    "password": "123456"
}'
```
response:
```shell
{"message":"OK","result":1,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg0NjQyLCJpc3MiOiJnby1uZXdzIn0.Q_3qBVN2nTYpZ50YPrs32ciNQDnjusyPbt5REOOqzD8"}
```
### login
request:
```shell
curl --location --request POST 'localhost:7777/api/v1/login' \
--data-raw '{
    "email": "changweitu@gmail.com",
    "password": "123456"
}'
```
response:
```shell

{"message":"OK","result":1,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg0OTI1LCJpc3MiOiJnby1uZXdzIn0.BovWURxcA0lhzV6dYlSxamUiJDgVx58YNJtIwp4iNCk"}%  
```

### change password
request:
```shell
curl --location --request POST 'localhost:7777/api/v1/changePassword' \
--data-raw '{
    "password":"000000"
}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg2NDQxLCJpc3MiOiJnby1uZXdzIn0.skk1eRMpeFM21q1e2VfjgAaesRja65ej6gF_VpbZs0A"  
```
response:
```shell

{"message":"OK","result":1}
```
## User

### get user
request: 
```shell
curl --location --request GET 'localhost:7777/api/v1/users/1' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg1MTU3LCJpc3MiOiJnby1uZXdzIn0.6Tmy1O0pLj5UTyiljx7IWco20jS1laiT4EmAD6pJtKY"  
```
response:
```shell
{"data":{"username":"","avatar":"","role":1,"id":1},"message":"OK","result":1}
```

### get users
request:
```shell
curl --location --request GET 'localhost:7777/api/v1/users?pageSize=3&pageNum=0' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg2NDQxLCJpc3MiOiJnby1uZXdzIn0.skk1eRMpeFM21q1e2VfjgAaesRja65ej6gF_VpbZs0A"
```
response:
```shell

{"data":[{"username":"tu_tu","avatar":"url","role":2,"id":1}],"message":"OK","result":1,"total":1}
```

### edit user
request:
```shell
curl --location --request PUT 'localhost:7777/api/v1/users/1' \
--data-raw '{
    "role": 2,
    "username": "tu_tu",
    "avatar": "url"
}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg1MTU3LCJpc3MiOiJnby1uZXdzIn0.6Tmy1O0pLj5UTyiljx7IWco20jS1laiT4EmAD6pJtKY"
```
response:
```shell

{"message":"OK","result":1} 
```
## Post

### create post (post_type: 1 text, 2 image, 3 url)
request:
```shell
curl --location --request POST 'localhost:7777/api/v1/posts' \
--data-raw '{
    "title": "my first post",
    "post_type": 1,
    "content": "Hello everyone, it is my first post"
}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoYW5nd2VpdHVAZ21haWwuY29tIiwiZXhwIjoxNjE5MTg3OTYzLCJpc3MiOiJnby1uZXdzIn0.5llD3LSSFBS282070dEPhbBNWjW0FdlYaZuDzRJ5nSs"
```
response:
```shell
{"message":"OK","result":1}
```
## Friends
### get followers
request:
```shell
 curl --location --request GET 'localhost:7777/api/v1/followers' \ 
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjFAMS5jb20iLCJleHAiOjE2MTkzODc0MTksImlzcyI6ImdvLW5ld3MifQ.8wiry_aCUODCzpDzIj5fhtci474wA9pcN8CKsLprWJQ"
```
response:
```shell
{"data":[{"username":"","avatar":"","role":1,"id":439},{"username":"","avatar":"","role":1,"id":440},{"username":"","avatar":"","role":1,"id":441},{"username":"","avatar":"","role":1,"id":442},{"username":"","avatar":"","role":1,"id":443},{"username":"","avatar":"","role":1,"id":444},{"username":"","avatar":"","role":1,"id":445},{"username":"","avatar":"","role":1,"id":446},{"username":"","avatar":"","role":1,"id":447},{"username":"","avatar":"","role":1,"id":448},{"username":"","avatar":"","role":1,"id":449},{"username":"","avatar":"","role":1,"id":450},{"username":"","avatar":"","role":1,"id":451},{"username":"","avatar":"","role":1,"id":452},{"username":"","avatar":"","role":1,"id":453},{"username":"","avatar":"","role":1,"id":454},{"username":"","avatar":"","role":1,"id":455},{"username":"","avatar":"","role":1,"id":456},{"username":"","avatar":"","role":1,"id":457},{"username":"","avatar":"","role":1,"id":458}],"message":"OK","result":1,"total":20}
```
### get followings
request:
```shell
curl --location --request GET 'localhost:7777/api/v1/followings' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjFAMS5jb20iLCJleHAiOjE2MTkzODc0MTksImlzcyI6ImdvLW5ld3MifQ.8wiry_aCUODCzpDzIj5fhtci474wA9pcN8CKsLprWJQ"
```
response:
```shell
{"data":[{"username":"","avatar":"","role":1,"id":417},{"username":"","avatar":"","role":1,"id":418},{"username":"","avatar":"","role":1,"id":419},{"username":"","avatar":"","role":1,"id":420},{"username":"","avatar":"","role":1,"id":421},{"username":"","avatar":"","role":1,"id":422},{"username":"","avatar":"","role":1,"id":423},{"username":"","avatar":"","role":1,"id":424},{"username":"","avatar":"","role":1,"id":425},{"username":"","avatar":"","role":1,"id":426},{"username":"","avatar":"","role":1,"id":427},{"username":"","avatar":"","role":1,"id":428},{"username":"","avatar":"","role":1,"id":429},{"username":"","avatar":"","role":1,"id":430},{"username":"","avatar":"","role":1,"id":431},{"username":"","avatar":"","role":1,"id":432},{"username":"","avatar":"","role":1,"id":433},{"username":"","avatar":"","role":1,"id":434},{"username":"","avatar":"","role":1,"id":435},{"username":"","avatar":"","role":1,"id":436}],"message":"OK","result":1,"total":20}
```
### follow user
request:
```shell
curl --location --request POST 'localhost:7777/api/v1/follow' \
--data-raw '{
"id": 439
}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjFAMS5jb20iLCJleHAiOjE2MTkzOTQ5NDMsImlzcyI6ImdvLW5ld3MifQ.vr43aOzy3ttE0J8C-SlrH5stzy95hhhs4oyZvCrRq6E"
```
response:
```shell
{"message":"OK","result":1}
```

### unfollow user
request:
```shell
curl --location --request POST 'localhost:7777/api/v1/unfollow' \
--data-raw '{
"id": 439
}' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjFAMS5jb20iLCJleHAiOjE2MTkzOTQ5NDMsImlzcyI6ImdvLW5ld3MifQ.vr43aOzy3ttE0J8C-SlrH5stzy95hhhs4oyZvCrRq6E"
```
response:
```shell
{"message":"OK","result":1}
```
### upload image
request:
```shell
curl --location --request POST 'localhost:7777/api/v1/upload' \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjFAMS5jb20iLCJleHAiOjE2MjAzMTA1OTIsImlzcyI6ImdvLW5ld3MifQ.QNIyS3r5MNNyx_hQPb2ZpqZfzw0CO5wkQSLB64cPC4E" \
-H "Content-Type: multipart/form-data" \
-F "file=@WX20210428-164204.png"

```