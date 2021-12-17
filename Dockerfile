version: "3.9"

services:
  postgres:
    image: postgres:13.3
    environment:
      POSTGRES_DB: "postgres"
      POSTGRES_USER: "postgres"
      POSTGRES_PASSWORD: "qwerty"
    ports:
      - "5432:5432"