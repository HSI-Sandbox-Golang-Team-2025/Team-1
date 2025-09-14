# HSI SANDBOX LMS

This is a LMS for the HSI Sandbox, there will be user for administrator, students (santri), and instructors (mentor).

## TABLE OF CONTENTS

1. [Documentation](#documentation)
2. [Installation](#installation)
3. [Team](#team)

## Documentation

Click link below to see the documentation

<!-- 1. [USE CASE](/docs/USE_CASE.md) -->

1. [USE CASE](https://drive.google.com/file/d/10l_BOmjo5rqUMowp0NO_n_da5NYnkXBp/view?usp=sharing)
2. [ERD (Entity Relationship Diagram)](/docs/ERD.md)
3. [SYSTEM ARCHITECTURE](/docs/SYSTEM_ARCHITECTURE.md)

## Installation

1. Clone this repository

## TEAM

1. [RAGIL BURHANUDIN PAMUNGKAS](https://github.com/sipamungkas)
2. [AAS ASRINI](https://github.com/asrini07)
3. [AZIZ RIDHWAN PRATAMA](https://github.com/ziprawan)

## project structure

```
project/
├── cmd/
│   └── app/
│       └── main.go             # Main application logic
├── cmd/
│   └── docs/                   # project documentation like usecase, erd, system architecture
├── internal/
│   ├── user/                   # Feature: User
│   │   ├── handler/            # HTTP handlers (Fiber)
│   │   │   └── http.go
│   │   ├── service/            # Business logic
│   │   │   └── service.go
│   │   ├── repository/         # Data access (GORM)
│   │   │   └── repository.go
│   │   ├── user.go             # User model
│   │   └── refresh_token.go    # Refresh token model
├── pkg/                        # Shared utilities or helpers
│   ├── config/
│   │   └── config.go           # Env loader
│   ├── db/
│   │   └── db.go               # GORM init (Neon)
│   ├── jwt/
│   │   └── jwt.go              # JWT helpers
│   ├── logger/
│   │   └── logger.go           # Fiber logger setup (optional)
│   ├── middleware/
│   │   └── auth.go             # JWT auth middleware
│   ├── password/
│   │   └── hash.go             # bcrypt helpers
│   ├── response/
│   │   └── response.go         # unified JSON response helpers
│   └── validator/
│       └── validator.go        # go-playground/validator glue
├── configs/                    # Configuration files
│   └── .env.example
├── go.mod                      # Go module definition
└── go.sum                      # Go module checksum file

```
