<!--
	Style guide:
	https://github.com/tailscale/tailscale/blob/main/api.md
-->

# Example API

The Example API is a (mostly) RESTful API.<br/>
Typically, POST bodies should be JSON encoded and responses will be JSON encoded.

Ref `server/routes.go`.

# Authentication

None.

# APIs

* **[Users](#Users)**
  - [GET users](#users-get)

* **[User](#User)**
  - [GET user](#user-get)
  - [POST user address](#user-move)


* **[Top](#Top)**
  - [GET top rides](#top-get)


# Users

<a name=users-get></a>
#### `GET /users` - lists all users

Supports pagination.

##### Parameters
None

##### Query Parameters
* `city` - filter by city
* `limit`
* `offset`

##### Example
```sh
http GET :9000/users city==boston limit==1
```

Response
```json
[
    {
        "address": "31118 Allen Gateway Apt. 60",
        "city": "boston",
        "credit_card": "6464362441",
        "edges": {},
        "id": "2e147ae1-47ae-4400-8000-000000000009",
        "name": "Cindy Medina"
    }
]
```

# User

<a name=user-get></a>
#### `GET /user/:id` - get one user

Self-explanatory

##### Parameters
* `id`

##### Query Parameters
None

##### Example
```sh
http GET :9000/user/c28f5c28-f5c2-4000-8000-000000000026
```

Response
```json
[
    {
        "address": "31118 Allen Gateway Apt. 60",
        "city": "boston",
        "credit_card": "6464362441",
        "edges": {},
        "id": "2e147ae1-47ae-4400-8000-000000000009",
        "name": "Cindy Medina"
    }
]
```

<a name=user-move></a>
#### `POST /user/move` - change a user address

Self-explanatory

##### Parameters
None

##### Query Parameters
None

###### POST Body
```json
{
    "address": "nowhere",
    "id": "c28f5c28-f5c2-4000-8000-000000000026"
}
```

##### Example
```sh
http POST :9000/user/move id=c28f5c28-f5c2-4000-8000-000000000026 address=nowhere
```

Response
```json
{
    "address": "nowhere",
    "city": "amsterdam",
    "credit_card": "5844236997",
    "edges": {},
    "id": "c28f5c28-f5c2-4000-8000-000000000026",
    "name": "Maria Weber"
}
```


# Top

<a name=top-get></a>
#### `GET /top` - lists top rides

Also shows the rides' user and vehicle.

##### Parameters
None

##### Query Parameters
None

##### Example
```sh
http GET :9000/top
```

Response
```json
[
    {
        "city": "amsterdam",
        "edges": {
            "user": {
                "address": "nowhere",
                "city": "amsterdam",
                "credit_card": "5844236997",
                "edges": {},
                "id": "c28f5c28-f5c2-4000-8000-000000000026",
                "name": "Maria Weber"
            },
            "vehicle": {
                "city": "amsterdam",
                "creation_time": "2019-01-02T09:04:05Z",
                "current_location": "62609 Stephanie Route",
                "edges": {},
                "ext": {
                    "color": "red"
                },
                "id": "aaaaaaaa-aaaa-4800-8000-00000000000a",
                "owner_id": "c28f5c28-f5c2-4000-8000-000000000026",
                "status": "in_use",
                "type": "scooter"
            }
        },
        "end_address": "66037 Belinda Plaza Apt. 93",
        "end_time": "2018-12-14T14:04:05Z",
        "id": "ab020c49-ba5e-4800-8000-00000000014e",
        "revenue": 77,
        "rider_id": "c28f5c28-f5c2-4000-8000-000000000026",
        "start_address": "1905 Christopher Locks Apt. 77",
        "start_time": "2018-12-13T09:04:05Z",
        "vehicle_city": "amsterdam",
        "vehicle_id": "aaaaaaaa-aaaa-4800-8000-00000000000a"
    }
]

```
