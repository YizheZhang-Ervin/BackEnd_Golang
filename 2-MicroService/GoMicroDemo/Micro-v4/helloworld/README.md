# Hello World

## Run
```
micro run . --name helloworld
```

## Query Service
```
micro call helloworld Greeter.Hello '{"name": "John"}'
```

## List Services
```shell
micro services
```