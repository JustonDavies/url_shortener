## Go URL Shortener
Example project for a Go webservice that acts a web URL shortening service


## Application Server Installation

---

This README outlines the process for (re)installing / updating an instance of the URL Shortener web service.

---
#### Requirements

Use of this application requires GoLang 1.9.3 and PostgreSQL with user in models.go

---
#### Compile

Compiling the application

    go get
    go build

---

### Use

You can test the application via the commands below

       curl -X POST \
         http://localhost:8000/v1/short \
         -H 'cache-control: no-cache' \
         -H 'content-type: application/json' \
         -d '{
         "url": "https://www.forbes.com/forbes/welcome/?toURL=https://www.forbes.com/sites/karstenstrauss/2017/04/20/the-highest-paying-jobs-in-tech-in-2017/&refURL=https://www.google.co.in/&referrer=https://www.google.co.in/"
       }'

Based on the return value you can then

        curl -X GET \
          http://localhost:8000/v1/short/RESULT_HERE \
          -H 'cache-control: no-cache' \

...and you will get the original URL back.

### TODO
* Refactor inserts and lookups into the model