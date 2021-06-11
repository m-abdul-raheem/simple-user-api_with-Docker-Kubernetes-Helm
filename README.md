# Simple User API

## Overview

The User API is a simple api to add/view users.

Once the API have started, you can access it through the following link: <http://localhost:443/users/>.

## Examples of functionalities

You can use Postman API Tool or other tools, on windows/linux/mac.
I'm using curl on linux as basic examples.

Insert User
```bash
$ curl -i -H "Content-Type: application/json" -X POST localhost:443/users/ -d '{"name":"Foo Barrington","id":1}

HTTP/1.1 200 OK
```

View All Users
```bash
$ curl localhost:443/users/

[{"Id":1,"Name":"Foo Barrington"},{"Id":2,"Name":"Jane Doerty"}]
```

View User By ID
```bash
$ curl localhost:443/users/1

{"Id":1,"Name":"Foo Barrington"}
```


## Starting services

### Deploying using Binary

Make sure you have Go setup on your system.
Clone this repository.
Dependent packages are defined inside go.mod.
You can run the basic command to download those dependencies for your code.
```bash
go mod download
```
After that, you can create binary from users/code directory
```bash
go build
```
This will create a binary which you can simply run.
### Deploying using docker and docker-compose

* Docker Engine
* Docker Compose 
  
I have create a docker-compose file which will automatically start database and API server.
  
You can also use docker only to start a mongo container seperately and API seperately.

Use the following command to deploy all services in your local environment.
```bash
$ docker-compose up
```

Once the services have started, you can access the web through the following link: <http://localhost:443/users/>.
 
You can stop and delete the containers, but database will still remain consistent.

### Deploying using docker only
If you want to deploy API and DB using docker only, you can use following commands:
```bash
$ docker run --name mongodb_host --rm mongo:4.2.11
```
Above will start db container
Yon can then build user-api container via Dockerfile inside user directory
```bash
$ docker build . -t user-api
$ docker run --link mongodb_host -p 127.0.0.1:443:443/tcp user-api
```
This will link our user api container internally with mongodb host, and we can access our api at 127.0.0.1:443.

### Deploying using Kubernetes and Helm

* Kubernetes Cluster
* Helm and kubectl configured with cluster config
  
I have also added kubernetes config files in helm templates so we can deploy our API on kubernetes as well if want.

We can use below command.
```bash
$ helm --namespace main-namespace upgrade/create -f ./helm/users-api/values/dev.yml --set "runtimevar=TestValue" users-api-dev ./helm/users-api
```

### Automating Whole Process
I can further use CI/CD pipelines (Github Actions, Gitlab CI, Jenkins, Concourse, Codefresh, etc) to automate each and every step mentioned above.
Normally we would automate only building and deployment steps.
In this case, we can go with Docker Image Build, and deployment on kubernetes cluster using helm.

We can add other steps to our pipeline as well like testing, linting, pre-deploys, post-deploys, etc
