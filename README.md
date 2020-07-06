#### Dockerizing a GoLang Web API

Login to Docker Hub

```zsh
$ docker login
```

create a GoLang Web API container

```zsh
$ docker build -t {yourDockerUsername}/golang-api:1.0.0 .
```

Test the GoLang Web API container by running it. It should be visible at localhost:8080

```zsh
$ docker run -p 8080:8080 {yourDockerUsername}/golang-api:1.0.0
```

Push the container to your Docker Hub account repository

```zsh
$ docker push {yourDockerUsername}/golang-api:1.0.0
```
