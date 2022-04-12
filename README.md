# AlthCartV2

## Introduction

AlthCartV2 is mainly focused on handling multi-user's cart. AlthCartV2 use JWT to sign and verify authentication, JWT is stored at cookie and will expired within 1 hours (Can be changed at controller/auth.go). AlthCartV2 use UIKit for styling, the reason is to make it looks modern, simple, clean and Jquery for lifecycle.

## Installation
- Golang 1.18
- Iris V12
```
go get github.com/kataras/iris/v12@master
```
- JWT
```
go get github.com/kataras/iris/v12/middleware/jwt
```
- SQLX
```
go get github.com/jmoiron/sqlx
```
- PostgreSQL 14
```
go get github.com/lib/pq
```

## Preview
- Login
![Login Page](https://i.imgur.com/52RKuGb.png "Login Page")

- Product List
![Product List Page](https://i.imgur.com/ppBx4to.png "Product List Page")

- Carts
![Carts Page](https://i.imgur.com/z6bKmCQ.png "Carts Page")
