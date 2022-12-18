# simple-rest-api-server-fan-in
Simple rest-api server with fan-in collector/combiner.

Our application consists of three layers of abstractions:
1. Handler - transport layer
2. Service - business logic layer
3. Repository - data source (db, external resources or third party data providers)


### Project structure

```
.
├── cmd
│   └── main.go
├── k8s/
│   └── deplyment.yaml
├── internal/
│    └── api/
│       ├── v1/
│       │   │
│       │   ├── errs/
│       │   │   └── errs.go
│       │   ├── handlers/
│       │   │   ├── profile.go
│       │   │   └── user.go
│       │   ├── models/
│       │   │   ├── config.go
│       │   │   └── data.go
│       │   ├── repository/
│       │   │   ├── data/        
│       │   │   │   └── data.go
│       │   │   └── repisitory.go
│       │   └── services/
│       │       ├── data   
│       │       │   ├── data.go
│       │       │   ├── sorter.go
│       │       │   └── sorter_test.go
│       │       └── services.go
│       └── api.go
├── pkg/
│    ├── render/
│    │  └── render.go
│    └── rest/ 
│       ├── cllient.go
│       └── retry.go
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
 
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
- sortKey (String). Can be empty
- limit   (Integer). Greater than 1, less than 200, 20 by default


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

### Local development

#### Mock Server
You can use mock server to makes things easier in local development
```
$ go run cmd/mock_server/main.go
```
or 
```
$ make mock-server
```
Mock server runs on port 3000.

It provides rest-api endpoints:
- GET http://localhost:3000/assignment132/assignment/main/duckduckgo
- GET http://localhost:3000/assignment132/assignment/main/google
- GET http://localhost:3000/assignment132/assignment/main/wikipedia

#### Linter
To run linter use
```
$ make lint 
```

#### Run tests
```
$ make test
```

#### Run server locally
To run go application on host machine use commands
```
$ go run ./cmd/.
```
or
```
$ make run
```

#### Run in docker container
```
$ make docker-run
```
See Makefile to get more details.

### Kube deployment

Building and pushing the docker image to docker hub
```
# Login to docker with your docker Id
$ docker login
```
Image:
`docker.io/mamanegoryu1/simple-rest-server`

#### Starting a local Kubernetes cluster
```
$ minikube start
```

Now deploy our app to the cluster
```
$ kubectl apply -f k8s/deployment.yml
```

The deployment is created. You can get the deployments like this:
```
$ kubectl get deployments
```

You can type the following command to get the pods in the cluster:
```
$ kubectl get pods
```

Pods are allocated a private IP address by default and cannot be 
reached outside of the cluster. You can use the kubectl port-forward 
command to map a local port to a port inside the pod like this:
```
$ kubectl port-forward simple-rest-server-69b45499fb-7fh87 8080:8080
```



---
## What we can add or improve
- Use health checker library, smth like https://github.com/nelkinda/health-go
- Use Viper to init config
- Implement Exponential backoff strategy for HTTP requests
- ZeroLog for logs
