# go-mailservice

go-mailservice is a simple http server is used to send the emails. It features Rest endpoint to send an email. It is written using <a href="https://github.com/gin-gonic/gin"><strong>Gin Web Framework</strong></a> to make server high performant


## Contents

- [MailService](#go-mailservice)
  - [Usage](#usage)
  - [1. Design Pattern](#1-design-pattern)
  - [2. Middleware](#2-middleware)
  - [3. Frameworks](#3-frameworks)
  - [4. Future Email Providers Support](#4-future-email-providers-support)
  - [5. API Token authentication](#5-api-token-authentication)
  - [Performance Metrics](#performance-metrics)
  - [Limitations](#limitations)

## Usage

To install go-mailservice, you need to install [Go](https://golang.org/)(**version 1.12+ is required**) and set your Go workspace.

1. This project uses go modules and provides a make file. You should be able to simply install and start:

```sh
$ git clone https://github.com/anil-appface/go-mailservice.git
$ cd go-mailservice
$ make
$ ./go-mailservice
```


## 1. Design Pattern

Since we are using the gin web framework, we made the application rest api routes very easy and following the standard way of versioning the api's (/api/v1/)

## 2. Middleware

Out of the box Gin web framework allows you to add the middleware function. In this project, Middleware functionality has been added for logging & very basic <code>x-api-key</code> header authentication to validate the requests incoming.

## 3. Frameworks

Using the below frameworks in this project:

1. <a href="https://github.com/gin-gonic/gin"><strong>Gin Web Framework</strong></a>
2. <a href="https://github.com/mailjet/mailjet-apiv3-go"><strong>Mailjet API</strong></a>

## 4. Future Email Providers Support

Added mailProvider as interface in the app.

If everwant to add the new email-provider then we should

1. Implement mailProvider interface.
2. Add the additional case inside switch block in <a href="https://github.com/anil-appface/go-mailservice/blob/238d130f07fbe18a0936b16052b821ffddbdbda4/mailproviders/mailProvider.go#L17">GetEmailProvider(mailProvider string) MailProvider </a> Method


## 5. API Token authentication

Added the <code>x-api-key</code> value as <code>test</code> in the middleware to validate every request. If the request is not passed with header x-api-key as test then you will get unauthorised with status code 401 as error.

## Rest API 


<code>Endpoint  <strong> https://localhost:8080/api/v1/sendEmail  </strong> </code>

<code>HTTP Method:  <strong>POST </strong></code>

<code>Request Body  <strong> 
  {
    "emailProvider": "MAILJET",
    "emailParams": {
        "recepients": ["anilamilineni01@gmail.com"],
        "subject": "some subject",
        "body": "some body",
        "cc": ["anilamilineni01@gmail.com"],
        "bcc": ["anilamilineni01@gmail.com"]
      }
  } 
  </strong> </code>

<code>HTTP Headers  <strong> x-api-key : test  </strong> </code>


NOTE: cc,bcc fields are made optional. 

## Performance Metrics

Benchmarking for this application is not done.

<p align="justify"><i>"As this application uses Gin web framework, the default logs of gin server shows how much time is consumed by each request to send response back."</i></p>

## Limitations

1. Added very basic level of authentication on api routes and the server is running on http mode.
2. There is no field validation in api.
4. There is no much test cases written in the interest of time.
