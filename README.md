# music_app
Music app using Golang, Gin Framework with mysql database

Requirements
Go 1.17

1.Run below command to install go dependencies
    -go mod tidy


2.Run go server
    -go run main.go

3. Start the docker to connect with database
    -sudo docker start local-mysql
    -sudo docker exec -it local-mysql bash 
        -mysql -uroot -p [Enter Password]

Music-App API Documentation

1. User registeration Api
- Method : Post
- Endpoint :/api/v1/register
- Description : API for new user to register in app
- Request:
    - Headers : none
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :{
        "name":"priya",
        "email":"priya@gmail.com"
    }
    - Response :
     "message": "Register Successfully"

2. Admin registeration Api
- Method : Post
- Endpoint :/api/v1/register/admin-register
- Description : API for Admin to register in app
- Request:
    - Headers : none
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :{
        "name":"priyanka",
        "email":"priyanka@gmail.com"
    }
    - Response :
     "message": "Admin Register Successfully"

2.1 Admin login API
- Method : Post
- Endpoint :/api/login
- Description : API for Admin to login in app
- Request:
    - Headers : none
    - Body : {
  "email":"priyanka@gmail.com"   
}
    

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :
    {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJpeWFua2FAZ21haWwuY29tIiwidXNlciI6dHJ1ZSwiZXhwIjoxNjQ5MjQyNjA3LCJpYXQiOjE2NDkwNjk4MDcsImlzcyI6IlByaXlhIn0.WjIZ043AM0UBL4yHfQHuJgx4ACJ2gGjTcdmMD7PrhdY"
}


3. Artist CRUD API

- Method : Post
- Endpoint :/api/v1/artist
- Description :API to create artist
- Request:
    - Headers :Token of admin
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :{
        "name":"Justin",
        "image_url":"http://www.google/justin-bieber"
    }
    - Response :
    Message": "Artist Added Successfully"

3.1 
- Method : PUT
- Endpoint :/api/v1/artist/:id
- Description :API to Update Artist
- Request:
    - Headers :Token of admin
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :{
        "name":"Justin",
        "image":"bieber.jpg",
        "id":2
    }
    - Response :
    Message": "Artist Updated Successfully"

3.2
- Method : Get
- Endpoint :/api/v1/artist/show
- Description :API to Get Artist
- Request:
    - Headers :Token of admin
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
[
    {
        "id": 0,
        "name": "justin jack",
        "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/Justin_Bieber_in_2015.jpg/800px-Justin_Bieber_in_2015.jpg",
        "created_at": "2022-03-19 04:40:01",
        "updated_at": "2022-03-19 04:40:01"
    },
    {
        "id": 0,
        "name": "Taylor swift",
        "image_url": "https://api.time.com/wp-content/uploads/2019/04/taylor-swift-time-100-2019-082.jpg?quality=85&zoom=2",
        "created_at": "2022-03-19 04:43:30",
        "updated_at": "2022-03-19 04:43:30"
    },
    {
        "id": 0,
        "name": "Justin Bieber",
        "image_url": "https://lh3.googleusercontent.com/pW7Jv2o8g0bkXFi11hrumm_N0e7KAf5pc5bawoSdD44uTLAYQi-Eeh1t1HileeiMx-9pXN6hQROW-OBEzWQWcEs2",
        "created_at": "2022-03-19 05:26:07",
        "updated_at": "2022-03-19 05:26:07"
    },
    {
        "id": 0,
        "name": "Priyanka Gour",
        "image_url": "http://localhots.com",
        "created_at": "2022-03-21 03:55:41",
        "updated_at": "2022-03-21 09:24:36"
    },
    {
        "id": 0,
        "name": "",
        "image_url": "",
        "created_at": "2022-03-24 11:48:31",
        "updated_at": "2022-03-24 11:48:31"
    },
    {
        "id": 0,
        "name": "",
        "image_url": "",
        "created_at": "2022-03-25 04:18:08",
        "updated_at": "2022-03-25 04:18:08"
    },
    {
        "id": 0,
        "name": "",
        "image_url": "",
        "created_at": "2022-03-25 08:58:23",
        "updated_at": "2022-03-25 08:58:23"
    },
    {
        "id": 0,
        "name": "",
        "image_url": "",
        "created_at": "2022-03-25 08:59:06",
        "updated_at": "2022-03-25 08:59:06"
    },
    {
        "id": 0,
        "name": "priya",
        "image_url": "",
        "created_at": "2022-03-25 09:15:43",
        "updated_at": "2022-03-25 09:15:43"
    },
    {
        "id": 0,
        "name": "priya",
        "image_url": "",
        "created_at": "2022-03-25 09:15:54",
        "updated_at": "2022-03-25 09:15:54"
    },
    {
        "id": 0,
        "name": "priya",
        "image_url": "",
        "created_at": "2022-03-25 09:29:35",
        "updated_at": "2022-03-25 09:29:35"
    },
    {
        "id": 0,
        "name": "Priyanka",
        "image_url": "",
        "created_at": "2022-03-25 09:37:37",
        "updated_at": "2022-03-25 09:37:37"
    },
    {
        "id": 0,
        "name": "",
        "image_url": "",
        "created_at": "2022-03-25 09:38:11",
        "updated_at": "2022-03-25 09:38:11"
    },
    {
        "id": 0,
        "name": "",
        "image_url": "",
        "created_at": "2022-03-25 09:38:32",
        "updated_at": "2022-03-25 09:38:32"
    },
    {
        "id": 0,
        "name": "Falguni Pathak",
        "image_url": "http://t0.gstatic.com/licensed-image?q=tbn:ANd9GcSscnxk4KfPRmEmgPZJLf5DKh26D0shiDYPtiMG2wvUWwujxbZJemdsJGan-RhK",
        "created_at": "2022-03-29 04:45:00",
        "updated_at": "2022-03-29 04:45:00"
    }
]

3.3
- Method : Delete
- Endpoint :/api/v1/artist/remove
- Description :API to Delete Artist
- Request:
    - Headers :Token of admin
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :{
       
        "id":2
    }
    - Response :
    Message": "Artist Deleted Successfully"

4. Album CRUD API

- Method : Post
- Endpoint :/api/v1/artist
- Description :API to create Album
- Request:
    - Headers :Token of admin
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :{
        "name":"Bollywood",
        "description":"refreshing bollywoord music",
        "image_url":"http://www.google/justin-bieber",
        "is_published":1,
        "artist_id":2
    }
    - Response :
    Message": "Album Created Successfully"

4.1 

- Method : PUT
- Endpoint :/api/v1/album/edit
- Description :API to update Album
- Request:
    - Headers :Token of admin
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :{
        "name":"Falguni Pathak",
        "description":"rendition of fresh new songs",
        "is_published":1,
        "image_url":"http://www.google/justin-bieber",
         "id":2
    }
    - Response :
    Message": "Album Updated Successfully"


4.2 
- Method : DELETE
- Endpoint :/api/v1/album/remove
- Description :API to Delete Album
- Request:
    - Headers :Token of admin
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :{
                "id":2
    }
    - Response :
    Message": "Album Deleted Successfully"

4.3 
- Method : READ
- Endpoint :/api/v1/album/show
- Description :API to READ Album
- Request:
    - Headers :Token of admin
    - Body : none

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
    [{
        "name": "Sanam",
        "description":"Recreated old songs",
        "image_url":"https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/Justin_Bieber_in_2015.jpg/800px-Justin_Bieber_in_2015.jpg,
        "is_published": 0,
        "created_at": "2022-03-21 09:42:41",
        "updated_at:"2022-03-21 09:42:41" 
    },
     {
        "name": "Justin bieber",
        "description":"Recreated old songs,retro creation ",
        "image_url":"https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/Justin_Bieber_in_2015.jpg/800px-sanampuri.jpg ",
        "is_published": 1,
        "created_at": "2022-03-22 03:43:25",
        "updated_at:"2022-03-22 03:43:25" 
    }
    ]
   
		
4.4
- Method : POST
- Endpoint :/api/v1/album/add
- Description :API to ADD Album(multiple insert)
- Request:
    - Headers :Token of admin
    - Body : {
        "album_id":"2",
        "track_id":[1,2,3,4]
    }

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
4.5
- Method : DELETE
- Endpoint :/api/v1/album/remove-track
- Description :API to Remove Album(multiple delete)
- Request:
    - Headers :Token of admin
    - Body : {
        "album_id":2,
        "track_id":[1,2]
    }

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
	
5. API of CRUD operation on Track
- Method : POST
- Endpoint :/api/v1/track
- Description :API to create Track
- Request:
    - Headers :Token of admin
    - Body : 
    {
        "name":"willow",
        "track_index":2,
        "track_url":"www.google.com",
        "image_url":"http://bedje-jwh",
        "is_published":1,
        "artist_id":1
     }

- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
    { "Message": "Track Created Successfully"}

5.1 
- Method : PUT
- Endpoint :/api/v1/track/edit
- Description :API to Update Track
- Request:
    - Headers :Token of admin
    - Body : 
    {
        "name":" Samne wali",
        "track_index":2,
        "track_url":"www.google.com",
        "image_url":"http://bedje-jwh",
        "is_published":1,
        "track_id":1
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
    { "Message": "Track Updated Successfully"}

5.2
- Method : DELETE
- Endpoint :/api/v1/track/remove
- Description :API to Delete Track
- Request:
    - Headers :Token of admin
    - Body : 
    {
       "track_id":1
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
    { "Message": "Track Deleted Successfully"}
    		
5.3
- Method : GET
- Endpoint :/api/v1/track/show
- Description :API to GET Track
- Request:
    - Headers :Token of admin
    - Body : 
    {
       "track_id":1
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
    [
    {
        "track_id": 0,
        "name": "Samne wali",
        "track_index": 1,
        "track_url": "http://google.com/sanamPuri",
        "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/in_2015.jpg/800px-Justin_Bieber_in_2015.jpg",
        "is_published": 1,
        "created_at": "2022-03-21 10:49:14",
        "updated_at": "2022-03-21 10:52:45",
        "artist_id": 1
    },
    {
        "track_id": 0,
        "name": "One time",
        "track_index": 1,
        "track_url": "http://google.com/justine-bieber",
        "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/in_2015.jpg/800px-Justin_Bieber_in_2015.jpg",
        "is_published": 1,
        "created_at": "2022-03-21 10:54:26",
        "updated_at": "2022-03-21 10:54:26",
        "artist_id": 2
    },
    {
        "track_id": 0,
        "name": "baby",
        "track_index": 2,
        "track_url": "http://google.com/justine-bieber",
        "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/in_2015.jpg/800px-Justin_Bieber_in_2015.jpg",
        "is_published": 1,
        "created_at": "2022-03-21 10:55:02",
        "updated_at": "2022-03-21 10:55:02",
        "artist_id": 2
    },
    {
        "track_id": 0,
        "name": "somebody",
        "track_index": 3,
        "track_url": "http://google.com/justine-bieber",
        "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/in_2015.jpg/800px-Justin_Bieber_in_2015.jpg",
        "is_published": 1,
        "created_at": "2022-03-21 10:55:30",
        "updated_at": "2022-03-21 10:55:30",
        "artist_id": 2
    },
    {
        "track_id": 0,
        "name": "samne wali ",
        "track_index": 2,
        "track_url": "http://sanampuri/samnewali/",
        "image_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/sanampuri.jpg",
        "is_published": 1,
        "created_at": "2022-03-21 10:56:16",
        "updated_at": "2022-03-21 10:56:16",
        "artist_id": 1
    },
    {
        "track_id": 0,
        "name": "Falguni",
        "track_index": 2,
        "track_url": "http://www.google.com",
        "image_url": "hhfeiuf ",
        "is_published": 1,
        "created_at": "2022-04-04 08:30:49",
        "updated_at": "2022-04-04 08:30:49",
        "artist_id": 2
    },
    {
        "track_id": 0,
        "name": "willow",
        "track_index": 2,
        "track_url": "www.google.com",
        "image_url": "http://bedje-jwh",
        "is_published": 1,
        "created_at": "2022-04-04 10:59:48",
        "updated_at": "2022-04-04 10:59:48",
        "artist_id": 1
    }
]
6. API for CRUD operation on Playlist
- Method : POST
- Endpoint :/api/v1/playlist
- Description :API to Create playlist
- Request:
    - Headers :Token of admin
    - Body : 
    {
       "name":"Bollywood ",
       "description":"favorite Bollywood songs",
       "image_url": https://media.istockphoto.com/photos/mountain-landscape-picture-id517188688?k=20&m=517188688&s=612x612&w=0&h=i38qBm2P-6V4vZVEaMy_TaTEaoCMkYhvLCysE7yJQ5Q,
       "is_published":1
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
    { "Message": "Playlist Created Successfully"}

6.1
- Method : PUT
- Endpoint :/api/v1/playlist/edit
- Description :API to Update playlist
- Request:
    - Headers :Token of admin
    - Body : 
    {
       "name":"Bollywood ",
       "description":"favorite Bollywood songs",
       "image_url": https://media.istockphoto.com/photos/mountain-landscape-picture-id517188688?k=20&m=517188688&s=612x612&w=0&h=i38qBm2P-6V4vZVEaMy_TaTEaoCMkYhvLCysE7yJQ5Q,
       "is_published":1,
       "playlist_id":1
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
    { "Message": "Playlist Updated Successfully"}
	
6.2
- Method : DELETE
- Endpoint :/api/v1/playlist/remove
- Description :API to Delete playlist
- Request:
    - Headers :Token of admin
    - Body : 
    {
       
       "playlist_id":1
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
    { "Message": "Playlist Deleted Successfully"}

6.3
- Method : GET
- Endpoint :/api/v1/playlist/show
- Description :API to Display playlist
- Request:
    - Headers :Token of admin
    - Body : 
   
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :[
    {
        "playlist_id": 0,
        "name": "Bollywood",
        "description": "favorite Bollywood songs",
        "image_url": "https://media.istockphoto.com/photos/mountain-landscape-picture-id517188688?k=20&m=517188688&s=612x612&w=0&h=i38qBm2P-6V4vZVEaMy_TaTEaoCMkYhvLCysE7yJQ5Q=",
        "is_published": 0,
        "created_at": "",
        "updated_at": ""
    },
    {
        "playlist_id": 0,
        "name": "hollywood",
        "description": "favorite mind refreshing",
        "image_url": "https://images.unsplash.com/photo-1471879832106-c7ab9e0cee23?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=1080&fit=max",
        "is_published": 0,
        "created_at": "",
        "updated_at": ""
    },
    {
        "playlist_id": 0,
        "name": "hollywcfgdghood",
        "description": "favorite mooddegwet refreshing",
        "image_url": "https://images.unsplash.com/photo-1471879832106-c7ab9e0cee23?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=1080&fit=max",
        "is_published": 0,
        "created_at": "",
        "updated_at": ""
    },
    {
        "playlist_id": 0,
        "name": "Priya",
        "description": "heidjwqh diwqhd iqhdoiq",
        "image_url": "jkeh ehjeio",
        "is_published": 0,
        "created_at": "",
        "updated_at": ""
    }
]

6.4

- Method : POST
- Endpoint :/api/v1/playlist/add-track-playlist
- Description :API to add playlist
- Request:
    - Headers :Token of admin
    - Body : 
    {
       
       "playlist_id":1,
       "track_id":2
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
    {"Message": "Playlist Added in Track Successfully"}
 
6.5
- Method : DELETE
- Endpoint :/api/v1/playlist/remove-track-playlist
- Description :API to DELETE playlist from track
- Request:
    - Headers :Token of admin
    - Body : 
    {
       
       "playlist_id":1,
       "track_id":2
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :{
        "Message": "Playlist Removed from Track  Successfully"
    }

6.6

- Method : GET
- Endpoint :/api/v1/playlist/get"
- Description :API to GET playlist with track information 
- Request:
    - Headers :Token of admin
    - Body : none

   - Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :[
    {
        "id": 1,
        "playlist_id": 1,
        "track_id": 1,
        "pname": "",
        "description": "",
        "name": "",
        "image_url": ""
    },
    {
        "id": 2,
        "playlist_id": 1,
        "track_id": 2,
        "pname": "",
        "description": "",
        "name": "",
        "image_url": ""
    },
    {
        "id": 3,
        "playlist_id": 1,
        "track_id": 3,
        "pname": "",
        "description": "",
        "name": "",
        "image_url": ""
    },
    {
        "id": 4,
        "playlist_id": 1,
        "track_id": 4,
        "pname": "",
        "description": "",
        "name": "",
        "image_url": ""
    },
    {
        "id": 5,
        "playlist_id": 1,
        "track_id": 5,
        "pname": "",
        "description": "",
        "name": "",
        "image_url": ""
    },
    {
        "id": 8,
        "playlist_id": 2,
        "track_id": 1,
        "pname": "",
        "description": "",
        "name": "",
        "image_url": ""
    }
]

7. API for crud Operation on Favourite track

- Method : DELETE
- Endpoint :/api/v1/favourite-track/create
- Description :API to Add favourite trcak by user
- Request:
    - Headers :none
    - Body : 
    {
       
       "track_id":1,
       "user_id":2,
       "fav_track_index":2
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :{
        "Message": "Favourite Track created Successfully"
    }
		

7.1

- Method : DELETE
- Endpoint :/api/v1/unfavourite-track
- Description :API to DELETE or unfavourite the track
- Request:
    - Headers :none
    - Body : 
    {       
       "fav_track_id":2
     }
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :{
      "Message": "Track Removed from Favourite Successfully"
    }
		
7.2

- Method :GET
- Endpoint :/api/v1/favourite-track
- Description :API to GET favourite track
- Request:
    - Headers :none
    - Body : none
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Body :none
    - Response :
   [
    {
        "fav_track_id": 1,
        "name": "Priyanka",
        "track_url": "http://google.com/sanamPuri"
    },
    {
        "fav_track_id": 3,
        "name": "Priyanka",
        "track_url": "http://google.com/justine-bieber"
    },
    {
        "fav_track_id": 4,
        "name": "Priyanka",
        "track_url": "http://google.com/justine-bieber"
    },
    {
        "fav_track_id": 5,
        "name": "Ridhaan",
        "track_url": "http://sanampuri/samnewali/"
    }
]
     
    
   
 7.3 API to get favourite track by passing id in param 

- Method :GET
- Endpoint :/api/v1/favourite-track/:id
- Description :API to GET favourite track by id
- Request:
    - Headers : id 
    - Body : none
   
- Response :

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : param (pass id of your favourite track)
    - Body :none
    - Response :
  {
    "fav_track_id": 1,
    "name": "Priyanka",
    "track_url": "http://google.com/sanamPuri"
}

