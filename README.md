# Inventory API

## 1. Getting started

An inventory application

### 1.1 View Application and Documentation

- Application live url - 
- Documentation - 

### 1.2 Requirements

Before starting, make sure you have at least those components on your workstation:

- An up-to-date release of [Go](https://go.dev/doc/install) 
- A database (MongoDB).

### 1.3 Project configuration

Start by cloning this project on your workstation.

```sh
git clone https://github.com/emmabraboke/inventory-api 
```

The next thing will be to install all the dependencies of the project.

```sh
cd inventory-api
go get 
```

Once the dependencies are installed, you can now configure your project by creating a new `.env` file containing your environment variables used for development.

```
cp .env.example .env
vi .env
```

### 1.4 Running the app

You are now ready to launch the NestJS application using the command below.

```sh
# Run the
$ go run main.go

```

You can now head to `http://localhost:5000/documentation` and see your API Swagger docs.

## 2. Project structure

This template was made with a well-defined directory structure.

```sh
    ├───cmd
│   ├───handlers
│   │   ├───customerHandler      
│   │   ├───invoiceHandler       
│   │   ├───productHandler       
│   │   ├───saleHandler
│   │   ├───transactionHandler
│   │   └───userHandler
│   ├───middlewares
│   └───routes
├───docs
├───internals
│   ├───database
│   │   └───mongo
│   ├───entity
│   │   ├───customerEntity
│   │   ├───invoiceEntity
│   │   ├───productEntity
│   │   ├───responseEntity
│   │   ├───saleEntity
│   │   ├───transactionEntity
│   │   └───userEntity
│   ├───enum
│   ├───repository
│   │   ├───customerRepo
│   │   │   └───mongoRepo
│   │   ├───invoiceRepo
│   │   │   └───mongoRepo
│   │   ├───productRepo
│   │   │   └───mongoRepo
│   │   ├───saleRepo
│   │   │   └───mongoRepo
│   │   ├───transactionRepo
│   │   │   └───mongoRepo
│   │   └───userRepo
│   │       └───mongoRepo
│   └───service
│       ├───cryptoService
│       ├───customerService
│       ├───invoiceService
│       ├───paymentService
│       ├───productService
│       ├───saleService
│       ├───tokenService
│       ├───transactionService
│       ├───userService
│       └───validationService
└───utils
```

## 3. Some NPM commands

```sh
# development
$ npm run start

# watch mode
$ npm run start:dev

# Lint the project files using TSLint
$ npm run lint

# Format document
$ npm run format

# Run the migrations
$ npm run migration

# Revert the migrations
$ npm run migration:down
```

## 4. Technologies/Tools

- NestJS
- ObjectionJs
- KnexJs
- TypeScript
