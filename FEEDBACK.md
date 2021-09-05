
Devops / SRE TakeHome Hiring Process
# Feedback

I enjoyed working on this challenge a lot. 

Had some hurdles in the beginning, but I'm glad that I was able to pull through.

It took me roughly 6 hours, over the span of 3 days.

# Brief Overview

The application is a simple user-api to add/view users. Once API has started, you can access it on 8080 port where its running.

## Flow
The flow will look like this:

Start Cluster:
```bash
minikube start
```

Apply kubernetes config files:
```bash
#Need this so that HPA can collect metrics
kubectl apply -f kubernetes/metrics-server/

#Prometheus Setup
kubectl apply -f kubernetes/monitoring/custom-resources/

#Prometheus Setup
kubectl apply -f kubernetes/monitoring/prometheus

#Prometheus Setup
kubectl apply -f kubernetes/monitoring/prometheus-operator/

#Grafana Setup
kubectl apply -f kubernetes/monitoring/grafana/

#MongoDB Setup - Along with user-api
kubectl apply -f kubernetes/mongo/

#User-api setup as seperate deployment
kubectl apply -f kubernetes/app/
```

On kubernetes we can use `port-forward` to access our services.
In our case, we can view prometheus, grafana dashboard and test our user api.

## Prometheus

```bash
kubectl -n monitoring port-forward svc/prometheus-operated 9090:9090
```
You can view prometheus dashboard to check metrics and debug if targets are up or not.
http://localhost:9090/targets

---

## Grafana

```bash
kubectl -n monitoring port-forward svc/grafana 3000:3000
```
Login to Grafana http://localhost:3000/login, `user: admin, pass: admin.`

- http://localhost:3000/datasources
  - Provide Datasource: http://prometheus-operated:9090

- http://localhost:3000/dashboard/import
  - Import Dashboard: https://grafana.com/grafana/dashboards/12594.

You'll be able to monitor mongodb metrics on that Dashboard.

---

## App - UserAPI
Using port-forward again to access our user-api:

I have deployed user-api in two methods:

- In its own Deployment
```bash
kubectl port-forward service/user-api-service 8080:8080
```

- Alongside MongoDB Container in a Statefulset of MongoDB
```bash
kubectl port-forward service/mongodb-user-api-service 8080:8080
```


Now we can access app on localhost

Insert User: localhost:8080/users/
```bash
$ curl -i -H "Content-Type: application/json" -X POST localhost:8080/users/ -d '{"name":"Foo Barrington","id":1}'

HTTP/1.1 200 OK
```

View All Users
```bash
$ curl localhost:8080/users/

[{"Id":1,"Name":"Foo Barrington"},{"Id":2,"Name":"Jane Doerty"}]
```

View User By ID
```bash
$ curl localhost:8080/users/1

{"Id":1,"Name":"Foo Barrington"}
```

---
## Kubectl

You can use kubectl cli to check other resources and debug issues
```bash
#to view HPAs
kubectl get hpa

#view monitoring pods
kubectl get pods -n monitoring
```


---

## Components
### 1. Simple Service Webapp
- `Dockerfile` present in `users` directory with `code`, you can use it to build image
- `muhammadabdulraheem/user-api-go` is image repo address where image is pushed on `Docker-Hub`

---

### 2. Setup Kubernetes

- Installed `minikube` in my local workstation to deploy image build in `[1]`.
- Initially set up app differently in deployment along with service and hpa.
- Statefulset deployed with persistent volume for mongo database, along with metrics container for prometheus, also with app container as well
---

### 3. Monitoring (Prometheus/Grafana)

- Configured Prometheus with conditions to monitor apps using `ServiceMonitors`
- Deployed Grafana and configured correct prometheus data source
- Using this dashboard to view mongodb metrics
  - https://grafana.com/grafana/dashboards/12594


