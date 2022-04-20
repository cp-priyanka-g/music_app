# music_app
Music app using Golang, Gin Framework with mysql database

Requirements Go 1.17

1.Run below command to install go dependencies -go mod tidy

2.Run go server -go run main.go

Start the docker to connect with database -sudo docker start local-mysql -sudo docker exec -it local-mysql bash -mysql -uroot -p [Enter Password]



Album CRUD API
Method : Post

Endpoint :/album

Description :API to create Album

Authorization:
                username:priyanka@gmail.com 
                Password:123

Headers :
Body : { "name":"Bollywood", "description":"refreshing bollywoord music", "image_url":"http://www.google/justin-bieber", "is_published":1, "artist_id":2 }
Response :

If any server error is there then,

Status Code: 500 Internal server error

Data: no response data

If request will success ,

Status Code: 200 Ok

Headers : none

Body :none

Response : Message": "Album Created Successfully"

4.1

Method : PUT

Endpoint :/album/edit

Description :API to update Album

Authorization:
                username:priyanka@gmail.com 
                Password:123

Body : { "name":"Falguni Pathak", "description":"rendition of fresh new songs", "is_published":1, "image_url":"http://www.google/justin-bieber", "id":2 }
Response :

If any server error is there then,

Status Code: 500 Internal server error

Data: no response data

If request will success ,

Status Code: 200 Ok

Headers : none

Body :none

Response : Message": "Album Updated Successfully"

4.2

Method : DELETE

Endpoint :/album/remove

Description :API to Delete Album

Authorization:
                username:priyanka@gmail.com 
                Password:123


Body :{ "id":2 }
Response :

If any server error is there then,

Status Code: 500 Internal server error

Data: no response data

If request will success ,

Status Code: 200 Ok

Headers : none

Body :none

Response : Message": "Album Deleted Successfully"

4.3

Method : READ

Endpoint :/album/show

Description :API to READ Album

Authorization:
                username:priyanka@gmail.com 
                Password:123
                or
      
                username:anisha@gmail.com 
                Password:1234

Body : none
Response :

If any server error is there then,

Status Code: 500 Internal server error

Data: no response data

If request will success ,

Status Code: 200 Ok

Headers : none

Body :none

Response : [ { "album_id": 0, "name": "Sanam", "description": "Recreated old songs ", "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/Justin_Bieber_in_2015.jpg/800px-Justin_Bieber_in_2015.jpg", "is_published": 0, "created_at": "2022-03-21 09:42:41", "updated_at": "2022-03-21 09:42:41", "artist_id": 0 }, { "album_id": 0, "name": "Justin bieber", "description": "Recreated old songs,retro creation", "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/Justin_Bieber_in_2015.jpg/800px-sanampuri.jpg", "is_published": 1, "created_at": "2022-03-22 03:43:25", "updated_at": "2022-03-22 03:43:25", "artist_id": 0 }, { "album_id": 0, "name": "Shreya Ghosal", "description": "melidious song og bollywood", "image_url": "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcR8JE6WblIyHLrTnUFBRhYcy4LfIkVHhvq1CNRErv0ZXozW1n6v", "is_published": 0, "created_at": "2022-04-14 06:23:28", "updated_at": "2022-04-14 06:23:28", "artist_id": 0 } ]

