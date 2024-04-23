# Микросервисный проект по курсу «Промышленная разработка на C#/Go»
[![Actions Status](https://github.com/appleboy/ssh-action/workflows/remote%20ssh%20command/badge.svg)](https://github.com/qveex/beer-api-itmo/actions)

# Тема проекта - магазин пива 🍺
## [Предварительное ТЗ в Notion](https://flicker-jobaria-33d.notion.site/3Heads-5024c35e69bd4e5abc06d9f3c93e3713?pvs=74)

## Команда: [@qveex](https://github.com/xincas), [@xincas](https://github.com/qveex), [@Krutov777](https://github.com/Krutov777)

### В проекте используется
- Go
- C#
- gRPC
- PostgreSQL
- GORM
- EF
- Github Actions

## Микросервисы

### Авторизация [@Krutov777](https://github.com/Krutov777)
- `Register`
- `Login`
- `Logout`
- `IsAdmin`
- `GetUserInfo`
> [proto](https://github.com/Krutov777/protos/blob/main/proto/sso/sso.proto)

### Каталог [@qveex](https://github.com/xincas)
- `GetBeers`
- `GetBeer`
- `CreateBeer`
- `UpdateBeer`
- `DeleteBeer`
- `SetFavorite`
- `GetFavorites`
- `DeleteFavorite`
> [proto](https://github.com/qveex/beer-api-itmo/blob/main/go_beer_catalog/api/proto/catalog.proto)

### Заказы [@xincas](https://github.com/qveex)
- `GetOrder`
- `ListOrders`
- `CreateOrder`
- `UpdateOrder`
- `DeleteOrder`
> [proto](https://github.com/qveex/beer-api-itmo/blob/main/protos/order/order_service.proto)