# Сервис управления складом

Часть экосистемы Havchik Podbirator - микросервис, ответственный за управление хранением продуктов и создание списков покупок.

## Обязанности сервиса

- Отслеживание продуктов в хранилище пользователя (холодильник, кладовая и т.д.)
- Мониторинг сроков годности продуктов
- Расчет необходимых количеств ингредиентов

## Архитектура

Этот сервис является частью системы Havchik Podbirator:
- Обменивается данными с сервисом управления меню для получения планов питания
- Предоставляет данные об инвентаризации сервису управления меню
- Создает списки покупок на основе планов питания и текущих запасов


# Curl examples

```bash 
curl http://localhost:8080/api/v1/products

curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Apple", "ID": "Apple", "weight_per_pkg": 100, "expiration_date": "2026-01-02T15:04:05Z"}'

curl -X POST http://localhost:8080/api/v1/products/check-availability \
  -H "Content-Type: application/json" \
  -d '{
    "steps": ["step1"],
    "ingredients": [
      {"product_id": "Aplle", "amount": 5}
    ]
  }'
```