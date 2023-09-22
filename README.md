# Todo App in one Go binary
## Go + HTMX
*Binary works solo, but I will make a docker image for hosting*

`docker build -t go-todo-app .`

### run locally to test  

`docker run --rm -p 8000:8000 go-todo-app`  

## gcp crap 👍🏼
> project: todo-app  
> project id: todo-app-go-htmx-id  
> Service URL: https://service-name-rdve5qkcjq-uw.a.run.app

`gcloud auth login`

`gcloud projects create todo-app-go-htmx-id --name="todo-app"` 

`gcloud config set project todo-app-go-htmx-id`

`gcloud services enable containerregistry.googleapis.com`

`docker tag todo-app gcr.io/todo-app-go-htmx-id/todo-app`

`docker push gcr.io/todo-app-go-htmx-id/todo-app`

## the juice

`gcloud run deploy service-name --image gcr.io/todo-app-go-htmx-id/todo-app --platform managed --port=8000` <-- default port is 8080  Consider updating app to use `8080`  
*takes a while*

https://console.cloud.google.com/run/domains *manage custom domains*
## aws crap 👎🏼
### login to ecr  

`$(aws ecr get-login --region region-name --no-include-email)`  

### create ecr repo  

`aws ecr create-repository --repository-name go-todo-app`  

### tag image
>Replace account-id with your AWS account id and region-name with your AWS region (like us-west-1).  

`docker tag my-todo-app:latest account-id.dkr.ecr.region-name.amazonaws.com/go-todo-app:latest`  

### push image  

`docker push account-id.dkr.ecr.region-name.amazonaws.com/go-todo-app:latest`