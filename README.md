# SMS

School Management System built in Go

SMS is a School Management System API which provides admin and regular endpoints
to perform actions required to efficiently manage schools (universities).

## Features

- Create, Read, Update and Delete student accounts
- Create, Read, Update and Delete student courses
- Update the courses a student takes
- List the courses a student currently takes
- List the number of students currently taking a course

## Setup 

### docker and docker-compose

- Clone this project using
  ```shell
    git clone github.com/Emmrys-Jay/altschool-sms.git
  ```

- Create a `config.env` file in the project root directory, and populate it with
  the project configuration.

  Example `config.env` file:

  ```shell
    SECRET_KEY=mySecretKey
    ADMIN_MATRIC_NUMBER=admin
    ADMIN_PASSWORD=admin
  ```

- Start up and build containers using 
  ```shell
    docker-compose up --build
  ```


### local setup 

- Clone this project using 
  ```shell
    git clone github.com/Emmrys-Jay/altschool-sms.git
  ```

- Create a `config.env` file in the project root directory, and populate it with 
the project configuration.

  Example `config.env`:

  ```shell
    PORT=3000
    SECRET_KEY=mySecretKey
    DB_HOST=127.0.0.1
    DB_USERNAME=default
    DB_PASSWORD=password
    DB_PORT=3306
    DB_NAME=sms
    ADMIN_MATRIC_NUMBER=admin
    ADMIN_PASSWORD=admin
  ```

- Run project using
  ```shell
    go run main.go
  ```
  OR 
  ```shell
    go build . && ./altschool-sms
  ```

## Test API

After starting up the application, test API via swagger UI at:

[http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)

`NB:` Ensure the port `:3000` from the swagger UI URL above matches the PORT from your 
`config.env` file when deploying locally.


## Disclaimer

Thanks to [Altschool Africa](https://altschoolafrica.com) for issuing this project.

