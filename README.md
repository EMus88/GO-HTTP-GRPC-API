# To use project:

Clone git repository from https://github.com/EMus88/Go-HTTP-GRPC-API

Execute docker Commands:

    docker build --tag=grpc-api .

    docker run -it -p 8080:8080 -p 8081:8081 grpc-api


# About project:
This project performs the calculation of Fibonacci numbers.This is implemented through HTTP REST API and GRPC.

For use HTTP API, you need to use the URL   http://localhost:8080/numbers?start=5&end=7
where "start" is the first number of fibonacci and "end" is the last number.
In response, you will get Fibonacci numbers between first and last number thet you entered.

For GRPC you need use port 8081.