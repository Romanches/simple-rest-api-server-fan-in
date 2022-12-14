# simple-rest-api-server-fan-in
Simple rest-api server with fan-in collector/combiner.



### Project structure

```
.
├── cmd
│   └──  main.go
├── helm
│       └── chart.yaml
└── internal
    └── app
        ├── handlers
        │   ├── course.go
        │   ├── lecture.go
        │   ├── profile.go
        │   └── user.go
        ├── service
        │   ├── models        
        │   ├── course.go
        │   ├── lecture.go
        │   └── user.go
        └── repository
            ├── course.go
            └── user.go
    
```

### Available routes

* GET `/health` 

Health check endpoint. \
Response body example:
```
{
    "status": "success",
    "data": "Ok!"
}
```

###

### What we can add or improve
- Use health checker library, smth like https://github.com/nelkinda/health-go
- Load config from file

