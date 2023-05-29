# Get It Started -

- Create a database ada-api in your Postgres local instance
- Rename `.env.exmaple` to `.env`.
- Add in the database user and password for your postgres instance containing the database ada-api.
- In the root folder run `go run main.go`.
- Get the API docs at http://localhost:3000/swagger/index.html

## Tasks

### Endpoints 

- [x] /api/student 
    - [x] / [GET] returns all students
    - [x] /:school_id [GET] returns a specific student
    - [x] / [POST] creates a new student
    - [x] /:school_id [PUT] updates the student data
    - [x] /:school_id [DELETE] deletes the specified student

- [x] /api/instructor
    - [x] / [GET] returns all instructors
    - [x] /:school_id [GET] returns a specific instructor
    - [x] / [POST] creates a new instructor
    - [x] /:school_id [PUT] updates the instructor data
    - [x] /:school_id [DELETE] deletes the specified instructor

- [x] /api/building
    - [x] / [GET] returns all buildings
    - [x] /:name [GET] returns a specific building
    - [x] / [POST] creates a new building
    - [x] /:name [PUT] updates the building data
    - [x] /:name [DELETE] deletes the specified building

- [x] /api/room
    - [x] / [GET] returns all rooms
    - [x] /:building_name [GET] returns all rooms on a specified building_name
    - [ ] /:name [GET] returns a specific room
    - [x] / [POST] creates a new room
    - [x] /:name [PUT] updates the room data
    - [x] /:name [DELETE] deletes the specified room

- [x] /api/key
  - [x] /:bulding_name [GET] get all keys in a building
  - [x] / [POST] creates a new key
  - [x] /:rfid [PUT] updates a key
  - [x] /:rfid [DELETE] deletes a key

- [x] /api/record
  - [x] / [GET] get all records
  - [x] / [POST] creates a new record
    - [x] should also update the key status

## ENDPOINTS POTENTIAL PROBLEMS/BUGS

- [ ] implement a limit or range of records
- [ ] fix /api/room/:name [GET] endpoint

- [ ] create unit tests for all handlers 