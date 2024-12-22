# calc-api

## prereqs

make sure that you have git and go installed on your computer.

## installation

1. clone the repo

    ```sh
    git clone https://github.com/pelumbum/calc-api.git
    ```

## run the proj

1. navigate to the project dir

    ```sh
    cd calc-api
    ```

2. run the proj

    ```sh
    go run cmd/main/main.go
    ```

## testing the Project

you can get a result by creating an http request to the api with json data. here an example using `curl`:

```sh
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'
```