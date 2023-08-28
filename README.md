<div align='center' style="text-align: center;">

<h1 style="border:0;margin:1rem">Tickitz Backend</h1>

Backend for Tickitz

[Suggestion](mailto:adjiedewantara24@gmail.com)

<hr>
<br>

</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Contributors](#contributors)
- [Related Project](#related-projects)
- [License](#license)
- [Suggestion](#suggestion)

## Overview

Tickitz is a movie ticket booking platform that provides an easy, fast, and convenient booking experience for film enthusiasts. With Tickitz, users can effortlessly explore the latest movie schedules, select preferred seats, and make secure payments, all in one place.

## Features

- Authentication
- Search Movie (search, sort, filter, create)
- User Role, Profile (Update)
- Error Handling
- Manage Movie (CRUD)
- etc.

## Technologies Used

- [Golang](https://go.dev/learn/)
- [JsonWebToken](https://github.com/golang-jwt/jwt)
- [Postgresql](https://www.postgresql.org/)
- [Cloudinary](https://cloudinary.com/)
- etc.

## Getting Started

### Installation

1. Clone this repo

   ```bash
   git clone https://github.com/Ravictation/tickitzgolang.git
   ```



2. Install all dependencies

   ```bash
   go build 
   ```

4. Create .env file

   ```env
    DB_HOST = [YOUR DATABASE HOST]
    DB_USER = [YOUR DATABASE USER]
    DB_PASSWORD = [YOUR DATABASE PASSWORD]
    DB_DATABASE = [YOUR DATABASE NAME]
   DB_PORT = [YOUR DATABASE PORT]
    KEY = [YOUR KEY]
    PORT = [YOUR PORT]
    CLOUDINARY_URL = [YOUR CLOUDINARY API]


   ```

5. Start the local server

   ```bash
   go run ./cmd/main.go
   ```



## Contributors

Currently, there are no contributors to this project. If you would like to contribute, you can submit a pull request.

## Related Projects

- [Zwallet Front End](https://github.com/zakifrhn/tickitz-fe) - Front End

## License

This project is licensed under the ISC License

## Suggestion

If you find bugs / find better ways / suggestions you can pull request.
