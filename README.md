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

#### GET `/health` 

Health check endpoint. \
Response body example:
```
{
    "status": "success",
    "data": "Ok!"
}
```

#### GET `/data`

Possible parameters: \
- sortKey (String), no srting bu default
- limit   (Integer), Greater than 1, less than 200, 20 by default


Response body example:
```
{
    "status": "success",
    "data": [
        {
            "url": "www.example.com/abc1",
            "views": 1000,
            "relevanceScore": 0.1
        },
        {
            "url": "www.example.com/abc2",
            "views": 2000,
            "relevanceScore": 0.2
        },
        {
            "url": "www.example.com/abc3",
            "views": 3000,
            "relevanceScore": 0.3
        },
        {
            "url": "www.example.com/abc4",
            "views": 4000,
            "relevanceScore": 0.4
        },
        {
            "url": "www.example.com/abc5",
            "views": 5000,
            "relevanceScore": 0.5
        }
    ]
}
```

###

---
### What we can add or improve
- Use health checker library, smth like https://github.com/nelkinda/health-go
- Load config from file
- Implement Exponential backoff strategy for HTTP requests
- ZeroLog
