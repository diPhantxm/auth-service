# About
This project is test of creating microservices in golang with [go-kit package](https://github.com/go-kit/kit) and [CLI](https://github.com/chaseSpace/kit).
# Architecture

Storage uses an interface with methods for performing database operations since different database drivers work differently. If you want to use another driver, implement storage that inherits Abstract Storage and UserRepository interface, then change package in [service](pkg/service/service.go). 

In [service](cmd/service/service.go) change connection string or use it as flag.

Auth service is token based. If you want to implement your own token controller, you can implement [interface](internal/tokenController/controller.go) and use your controller in [service](cmd/service/service.go).

![Architecture](https://user-images.githubusercontent.com/43153608/168493608-875f3191-7687-4905-b038-2faf5b6dcb9e.jpg)
