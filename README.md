# todo-go

A simple Todo application built with Go and hosted on Google Cloud Run.

## API Endpoints

For your convenience, use the postman collection stored at the root of this project `todo-go.postman_collection.json`

### Task

#### Create Task

Create a new task.

```bash
curl --request POST \
     --url https://todo-go-6rxr265wza-uc.a.run.app/tasks \
     --header 'Content-Type: application/json' \
     --data-raw '{
         "title": "Do Laundry",
         "description": "Do Laundry Every Friday",
         "user_id": "1d960ab7-fbe4-4301-a4cd-2b8777f64317"
     }'
```

#### Read All Tasks

Get a list of all tasks.

```bash
curl --request GET \
     --url https://todo-go-6rxr265wza-uc.a.run.app/tasks
```

#### Read Task By ID

Get details of a specific task by providing its ID.

```bash
curl --request GET \
     --url https://todo-go-6rxr265wza-uc.a.run.app/tasks/{{taskID}}
```

#### Update Task

Update an existing task by providing its ID.

```bash
curl --request PUT \
     --url https://todo-go-6rxr265wza-uc.a.run.app/tasks/{{taskID}} \
     --header 'Content-Type: application/json' \
     --data-raw '{
         "title": "Mow the yard",
         "description": "Honeydo list"
     }'
```

#### Delete Task

Delete a task by providing its ID.

```bash
curl --request DELETE \
     --url https://todo-go-6rxr265wza-uc.a.run.app/tasks/{{taskID}}
```

### User

#### Create User

Create a new user.

```bash
curl --request POST \
     --url https://todo-go-6rxr265wza-uc.a.run.app/users \
     --header 'Content-Type: application/json' \
     --data-raw '{
         "username": "UserTwo",
         "email": "2@gmail.com"
     }'
```

#### Read All Users

Get a list of all users.

```bash
curl --request GET \
     --url https://todo-go-6rxr265wza-uc.a.run.app/users
```

#### Read User By ID

Get details of a specific user by providing their ID.

```bash
curl --request GET \
     --url https://todo-go-6rxr265wza-uc.a.run.app/users/{{userID}}
```

#### Update User

Update an existing user by providing their ID.

```bash
curl --request PUT \
     --url https://todo-go-6rxr265wza-uc.a.run.app/users/{{userID}} \
     --header 'Content-Type: application/json' \
     --data-raw '{
         "username": "User2",
         "email": "user2@yahoo.com"
     }'
```

#### Delete User

Delete a user by providing their ID.

```bash
curl --request DELETE \
     --url https://todo-go-6rxr265wza-uc.a.run.app/users/{{userID}}
```

---

Remember to replace `{{taskID}}` and `{{userID}}` with actual IDs when making requests.
