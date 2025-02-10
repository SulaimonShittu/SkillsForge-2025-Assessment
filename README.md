# SkillsForge-2025-Assessment
SkillsForge Technical assessment for 2025 Internship role
Track: Backend Development

### File Structure

├── section-one/ │ ├── non-repeating-char.go │ └── string-compression.go ├── task-management/
│ ├── endpoints.go │ ├── helper.go │ └── task.go ├── go.mode ├── main.go └── README.md

├── section-one/
│   ├── non-repeating-char.go
│   └── string-compression.go
├── task-management/      
│   ├── endpoints.go 
│   ├── helper.go
│   └── task.go
├── go.mode
├── main.go
└── README.md


## Section one

**Question One**
Implement a string compression. For example, aaaabbbccddddddee would become a4b3c2d6e2. If the
length of the string is not reduced, return the original string.

Test Cases.
---
section_one.StringCompress("bbcceeee") == ‘b2c2e4’
section_one.StringCompress("aaabbbcccaaa") == ‘a3b3c3a3’
section_one.StringCompress("a") = a

**Solution** : In file string-compression.go - Function StringCompress 
<img width="960" alt="s1" src="https://github.com/user-attachments/assets/637023d1-5ef9-41ce-ae5b-92ea6b9fec46" />


**Question Two**
Given a string, write a function to find the first non-repeating character. If all characters
repeat, return -1.

Test Cases.
---
section_one.FirstNonRepeating("swiss") == 'w'

**Solution** : In file non-repeating-char.go - Function FirstNonRepeating
<img width="960" alt="s2" src="https://github.com/user-attachments/assets/a2c94da8-dd5a-49dc-8ba0-8f1794f3ccac" />



## Section Two (Task Management API)

# Task Management API Documentation

## Base URL
```
http://localhost:8080
```

## Data Models

## Endpoints

### Create Task
`POST /tasks`

Creates a new task in the system.

**Sample Request Body:**
```json
{
  "title": "Complete Project Proposal",
  "description": "Write and review Q1 project proposal",
  "category": "work",
  "completed": false,
  "due_date": "2025-02-20T15:00:00Z"
}
```

**Sample Response:** `200 OK`
```json
{
  "data": {
    "id": 1,
    "title": "Complete Project Proposal",
    "description": "Write and review Q1 project proposal",
    "category": "work",
    "completed": false,
    "due_date": "2025-02-20T15:00:00Z",
    "created_at": "2025-02-10T12:00:00Z"
  },
  "message": "Task created successfully",
  "error": null
}
```

### List Tasks
`GET /tasks`

Retrieves a paginated list of tasks, with optional category filtering.

**Query Parameters:**
- `page` (integer, default: 1): Page number
- `pageSize` (integer, default: 10): Number of items per page
- `category` (string, optional): Filter tasks by category

**Sample Response:** `200 OK`
```json
{
  "data": [
    {
      "id": 1,
      "title": "Complete Project Proposal",
      "description": "Write and review Q1 project proposal",
      "category": "work",
      "completed": false,
      "due_date": "2025-02-20T15:00:00Z",
      "created_at": "2025-02-10T12:00:00Z"
    }
  ],
  "message": "Showing page 1 of tasks",
  "error": null
}
```

### Get Task
`GET /tasks/{id}`

Retrieves a specific task by ID.

**Sample Response:** `200 OK`
```json
{
  "data": {
    "id": 1,
    "title": "Complete Project Proposal",
    "description": "Write and review Q1 project proposal",
    "category": "work",
    "completed": false,
    "due_date": "2025-02-20T15:00:00Z",
    "created_at": "2025-02-10T12:00:00Z"
  },
  "message": null,
  "error": null
}
```

### Update Task
`PUT /tasks/{id}`

Updates an existing task. Supports partial updates.

**Sample Request Body:**
```json
{
  "title": "Complete Project Proposal - Updated",
  "completed": true
}
```

**Sample Response:** `200 OK`
```json
{
  "data": {
    "id": 1,
    "title": "Complete Project Proposal - Updated",
    "description": "Write and review Q1 project proposal",
    "category": "work",
    "completed": true,
    "due_date": "2025-02-20T15:00:00Z",
    "created_at": "2025-02-10T12:00:00Z"
  },
  "message": "Task updated successfully",
  "error": null
}
```

### Delete Task
`DELETE /tasks/{id}`

Deletes a specific task.

**Sample Response:** `200 OK`
```json
{
  "data": null,
  "message": "Task deleted successfully",
  "error": null
}
```

## Error Responses

### 400 Bad Request
```json
{
  "data": null,
  "message": null,
  "error": "Invalid request body"
}
```

### 404 Not Found
```json
{
  "data": null,
  "message": null,
  "error": "Task not found"
}
```

### 405 Method Not Allowed
```json
{
  "data": null,
  "message": null,
  "error": "Method not allowed"
}
```
