# Golang-Web-Crawler
This repository creates a web crawler using Graph data structure

## Pre-requirements
**IMPORTANT:**

The application requires Golang version 1.19, godotenv and zap logger library

Install requirements in go.mod file using the command below:
```
go mod tidy
```

## Usage
- Change directory to golang-web-crawler/cmd/web
- Build the application with the command below
```
go build .
```
- Create ".env" file and set environment variables: BASE_URL and CRAWL_EXTERNAL_LINKS
- Run application using command below
```
./web
```
