<div align='center'>

<img src="./icon.png" width="121" style="border-radius:25px"/><br>

[![Tech Stack](https://skillicons.dev/icons?i=golang,mysql)](#tech-stack)

<h2>Gadgetify</h2>
<h3 align="center">Powered by Scalingo ⚡</h3>

[Demo](https://gadgetify-go.osc-fr1.scalingo.io/)

</div><br>

## Table of Contents

- [Overview](#overview)
- [Tech Stack](#tech-stack)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Instalation \& Environments](#instalation--environments)
- [Postman Documentation](#postman-documentation)
- [Table Structure](#table-structure)
- [Related Projects](#related-projects)
- [Contributors](#contributors)
- [License](#license)
- [Report](#report)

## Overview

This project provides a RESTful API for managing products in an online gadget shop. It supports basic CRUD (Create, Read, Update, Delete) operations for products.

The REST API follows the principles of Representational State Transfer (REST), which enables easy integration with various clients, including web and mobile applications. It utilizes the HTTP protocol for communication, allowing clients to perform operations such as retrieving, creating, updating, and deleting resources.

### Features

- Authorization & Authentication
- Upload Images
- CRUD
- Error Handling & Validation

## Tech Stack

- [Go-lang](https://go.dev/)
- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/) (for storing data)
- [Cloudinary](https://cloudinary.com/) (for storing images)
- [JSON Web Token](https://jwt.io/) (authorization)
- [Scalingo](https://scalingo.com/) (for deploying)
- and other packages (you can see in go.mod)

## Getting Started

### Prerequisites

- [Go-lang](https://go.dev/)
- [MySQL](https://www.mysql.com/)

### Instalation & Environments

1. Clone this repository to your local

   ```bash
   git clone https://github.com/nyannss/gadgetify-go.git
   ```

2. Install dependencies

   ```bash
   cd gadgetify-go && go mod download
   ```

3. Setup environments (you can see in `.env.example`)

   - Database server using MySQL

     ```env
     APP_PORT = (port as you want)
     DB_HOST = (put your db host)
     DB_PORT = (put your port of db host)
     DB_USER = (put your db username)
     DB_PASS = (put your db password)
     DB_NAME = (put your db  name)
     ```

   - JSON Web Token Secret Key (prefer using random string) [[see more information]](<https://jwt.io/introduction>)

     ```env
     JWT_SECRET_KEY = (put your secret key)
     ```

   - Image server using Cloudinary [[you can create account in here]](<https://cloudinary.com/>)

     ```env
     CLOUDINARY_NAME = (put your cloudinary name)
     CLOUDINARY_KEY = (put your cloudinary key)
     CLOUDINARY_SECRET = (put your cloudinary secret)
     ```

4. Last, run the app

   ```bash
   go run main.go
   ```

## Postman Documentation

You can see the documentation from [Postman](https://elements.getpostman.com/redirect?entityId=26209677-03d497ed-7164-41b8-b466-70405ef30f5e&entityType=collection).

## Table Structure

You can download table structure (ddl) from [this link](/ddl.sql).


## Contributors

- [nyannss](https://github.com/nyannss)

## License

This project using ISC License

## Report

Any error report you can pull request
or contact: <nyannss@proton.me>
