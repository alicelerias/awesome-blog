# awesome-blog

AWESOME! is a photography blog and texts where members can follow each other and interact through comments and likes.

## Index

- [Backend - GO lang](#backend---go)
- [Frontend - TYPESCRIPT](#frontend---typescript)
- [Running the application](#running-the-application)

## Backend - GO

Dependencies:

- `gin` - web framework
- `dgrijalva/jwt-go` - tokens jwt generator
- `postgres` - db

Folders:
- `config/` - shortcuts to the application's private contents;

- `api/` - controllers, auth and middlewares;

- `cache/` - redis configs to cache some endpoints;

- `database/` - database connection, application queries and mutations; 

- `migrations/` - queries recipes in SQL language;

- `models/` - database models;

- `scripts/` - scripts to build;

- `seed/` - seed to populate db;

- `types/` - application types;

- `main.go` - endpoints;

## Frontend - TYPESCRIPT

Dependencies:
- `react`
- `typescript`
- `tailwind` - styling
- `prettier` - format

Folders:
- `api/` queries and mutations using axios;

- `components/` tsx components; 

- `configs/` shortcuts to the application's private contents;

- `types/` application types;

- `context/` current user context;

- `main.tsx` Application routes;

## Running the application

- Postgres:
  - `Set your env dependencies`
- Backend:
  - `cd backend`
  - `go mod tidy` 
  - `go test ./...` 
  - Start
    `go install github.com/pilu/fresh@latest`
      - `fresh`
- Frontend
  - `cd frontend/`
  - `npm i`
  - `npm run format`
  - `npm test`
  - `npm start`
