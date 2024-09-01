To run the service locally, you need to have the following installed:
go - preferably 1.18 and above
To get all dependencies simply run `go mod tidy`

After cloning the service, go to main directory and run the command
 ```go run .```

There will be link printed in terminal 
```http://localhost:8080/graphiql```

Click on it and a graphiql interface will open up. You can run the queries and mutations there.

### **Sample Mutations And Queries**

- I have added 5 Users already with Ids from `user1` to `user5`

#### Creating a User :
```
mutation{
  createUser(name : "parag"){
   	Id
    name
  }
}
```

#### Follow Someone
```
mutation{
  followUser(myId: "user1", targetId: "user5")
}
```
if the target user exist and is not already followed it will return true else false with message

#### UnFollow Someone
```
mutation{
  unfollowUser(myId: "user1", targetId: "user5")
}
```

#### Get All Followers
```
query{
  followers(Id : "user1"){
    Id
    name
  }
}
```

#### Get All Following
```
query{
  followings(Id : "user1"){
    Id
    name
  }
}
```