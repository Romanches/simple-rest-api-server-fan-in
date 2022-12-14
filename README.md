# simple-rest-api-server-fan-in
Simple rest-api server with fan-in collector/combiner.



### Project structure

```
.
├── cmd
│   └──  main.go
└── internal
    └── app
        ├── handlers
        │   ├── health.go
        │   ├── data.go
        ├── service
        │   ├── models        
        │   ├── data.go
        ├── repository
        │   ├── duckduck.go
        │   ├── google.go
        │   └── wikipedia.go
        └── app.go
```

