# Introduction
   This is a RESTful API that perform CRUD to a social network database.   
(Currently includes user and email data, but will add more data and user relation in the future.)

# Set Up

- Run mysql db in container, map to host port 8081, password is "password".  

	`docker run --name mysql -p 8081:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql`

# Test

`cd ./src/social_network`  
`go test `

# API call

## Create User

**URL:**  
`api/vi/users/  `

**Method:**  
POST  

**URL Params:**  
None  

**Data Params:**  
`{"name":"name1", "email":"email1@mail.com"}  `

**Success Response:**  
- Code: 201  
- Content:`{"message": "User created successfully!", "resourceId": UserID} `

**Error Reponse:**  
- Code: 500  
- Content:`{"message": "User not created!"} ` 

## Get Single User

**URL:**  
`api/vi/users/:id/  `

**Method:**  
GET  

**URL Params:**  
- Required: 
`id=[integer]` 

**Data Params:**  
None

**Success Response:**  
- Code: 200  
- Content:`{"data": {"ID": "1", "Name": "name1", "Email": "email1@mail.com"}} `

**Error Reponse:**  
- Code: 500  
- Content:`{"message": "User not read!"} ` 

## Update User

**URL:**  
`api/vi/users/:id/  `

**Method:**  
PATCH  

**URL Params:**  
- Required: 
`id=[integer]`  

**Data Params:**  
`{"name":"updateName1", "email":"update1@mail.com"}  `

**Success Response:**  
- Code: 200  
- Content:`{"message": "User updated successfully!", "resourceId": UserID} `

**Error Reponse:**  
- Code: 500  
- Content:`{"message": "User not update!"} ` 

## Delete User

**URL:**  
`api/vi/users/:id/  `

**Method:**  
DELETE  

**URL Params:**  
- Required: 
`id=[integer]`  

**Data Params:**  
None

**Success Response:**  
- Code: 200  
- Content:`{"message": "User deleted successfully!"} `

**Error Reponse:**  
- Code: 500  
- Content:`{"message": "User not delete!"} ` 