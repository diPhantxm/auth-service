# About
This project is test of creating microservices in golang with [go-kit package](https://github.com/go-kit/kit) and [CLI](https://github.com/chaseSpace/kit).
# Architecture

Repository Pattern is used to perform CRUD operations to database. If you want to use another driver for db, implement [repository](internal/storage/Repos/userRepo.go) and add it to [factory](internal/storage/Repos/userRepoFactory.go). To add repository implementation to storage instance according to driver, it uses Factory Pattern.

In [service](cmd/service/service.go) change connection string or use it as flag.

Auth service is token based. If you want to implement your own token controller, you can implement [interface](internal/tokenController/controller.go) and use your controller in [service](cmd/service/service.go).

![Architecture](https://user-images.githubusercontent.com/43153608/169884233-95b5e2fa-fcf6-4531-b704-9354fad21e36.jpg)
