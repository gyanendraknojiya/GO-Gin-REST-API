# GO-Gin-REST-API

## Features

- Auto reload using AIR.
- All REST APIs (GET, POST, PUT, DELETE)

## Start server

```bash
air run main.go
```

## API

### Get All Users

URL: GET localhost:8080/

**Response:**

```json
[
  {
    "id": "1",
    "name": "Leanne Graham",
    "userName": "Bret",
    "email": "Sincere@april.biz",
    "phone": "1-770-736-8031 x56442",
    "website": "hildegard.org"
  },
  {
    "id": "2",
    "name": "Ervin Howell",
    "userName": "Antonette",
    "email": "Shanna@melissa.tv",
    "phone": "010-692-6593 x09125",
    "website": "anastasia.net"
  },
  {
    "id": "3",
    "name": "Clementine Bauch",
    "userName": "Samantha",
    "email": "Nathan@yesenia.net",
    "phone": "1-463-123-4447",
    "website": "ramiro.info"
  },
  {
    "id": "4",
    "name": "Patricia Lebsack",
    "userName": "Karianne",
    "email": "Julianne.OConner@kory.org",
    "phone": "493-170-9623 x156",
    "website": "kale.biz"
  },
  {
    "id": "5",
    "name": "Chelsey Dietrich",
    "userName": "Kamren",
    "email": "Lucio_Hettinger@annie.ca",
    "phone": "(254)954-1289",
    "website": "demarco.info"
  }
]
```

### Get user by ID

**URL:** GET localhost:8080/user/1

**Response:**

```json
{
  "id": "1",
  "name": "Leanne Graham",
  "userName": "Bret",
  "email": "Sincere@april.biz",
  "phone": "1-770-736-8031 x56442",
  "website": "hildegard.org"
}
```

### Add new user

**URL:** POST localhost:8080/addUser

Body

```json
{
  "id": "7",
  "name": "Gyanendra Knojiya",
  "userName": "Gyan",
  "email": "gyanendrak064@gmail.com",
  "phone": "+918802879231",
  "website": "gyanendra.tech"
}
```

**Response:**

```json
[
...,
  {
    "id": "7",
    "name": "Gyanendra Knojiya",
    "userName": "Gyan",
    "email": "gyanendrak064@gmail.com",
    "phone": "+918802879231",
    "website": "gyanendra.tech"
  }
]
```

### Update a user

**URL:** PUT localhost:8080/updateUser

Body

```json
{
  "id": "7",
  "name": "Gyanendra Verma"
}
```

**Response:**

```json
[
...,
  {
    "id": "7",
    "name": "Gyanendra Verma",
    "userName": "Gyan",
    "email": "gyanendrak064@gmail.com",
    "phone": "+918802879231",
    "website": "gyanendra.tech"
  }
]
```

### Delete a user

**URL:** DELETE localhost:8080/deleteUser

Body

```json
{
  "id": "7"
}
```

**Response:**

```json
[
  {
    "id": "1",
    "name": "Leanne Graham",
    "userName": "Bret",
    "email": "Sincere@april.biz",
    "phone": "1-770-736-8031 x56442",
    "website": "hildegard.org"
  },
  {
    "id": "2",
    "name": "Ervin Howell",
    "userName": "Antonette",
    "email": "Shanna@melissa.tv",
    "phone": "010-692-6593 x09125",
    "website": "anastasia.net"
  },
  {
    "id": "3",
    "name": "Clementine Bauch",
    "userName": "Samantha",
    "email": "Nathan@yesenia.net",
    "phone": "1-463-123-4447",
    "website": "ramiro.info"
  },
  {
    "id": "4",
    "name": "Patricia Lebsack",
    "userName": "Karianne",
    "email": "Julianne.OConner@kory.org",
    "phone": "493-170-9623 x156",
    "website": "kale.biz"
  },
  {
    "id": "5",
    "name": "Chelsey Dietrich",
    "userName": "Kamren",
    "email": "Lucio_Hettinger@annie.ca",
    "phone": "(254)954-1289",
    "website": "demarco.info"
  }
]
```
