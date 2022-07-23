## Что это?

Микросервис, работающий с очередью NATS. Он читает данные из очереди и записывает их
в кэш и базу данных. В случае попадания невалидных данных в канал сервис записывает данные
в отдельную таблицу в базе данных

## Пример модели данных

```json
{
"order_uid": "b563feb7b2b84b6test",
"track_number": "WBILMTESTTRACK",
"entry": "WBIL",
"delivery": {
"name": "Test Testov",
"phone": "+9720000000",
"zip": "2639809",
"city": "Kiryat Mozkin",
"address": "Ploshad Mira 15",
"region": "Kraiot",
"email": "test@gmail.com"
},
"payment": {
"transaction": "b563feb7b2b84b6test",
"request_id": "",
"currency": "USD",
"provider": "wbpay",
"amount": 1817,
"payment_dt": 1637907727,
"bank": "alpha",
"delivery_cost": 1500,
"goods_total": 317,
"custom_fee": 0
},
"items": [
{
"chrt_id": 9934930,
"track_number": "WBILMTESTTRACK",
"price": 453,
"rid": "ab4219087a764ae0btest",
"name": "Mascaras",
"sale": 30,
"size": "0",
"total_price": 317,
"nm_id": 2389212,
"brand": "Vivienne Sabo",
"status": 202
}
],
"locale": "en",
"internal_signature": "",
"customer_id": "test",
"delivery_service": "meest",
"shardkey": "9",
"sm_id": 99,
"date_created": "2021-11-26T06:22:19Z",
"oof_shard": "1"
}
```

## База данных

В качестве базы данных был использован Postgres. Схема базы данных выглядит следующим образом

<img src="https://github.com/infamax/nats-streaming-server/blob/main/%D1%81%D1%85%D0%B5%D0%BC%D0%B0%20%D0%B1%D0%B4.jpg" alt="Схема бд"/>

## АPI

 Поддерживаются следующие методы

1) Создание заказа в базе данных

   http://localhost:8080/create_order_db. HTTP method: POST.

   Обязательные параметры:
   
   order - структура следующего ввида

   Пример запроса:

   ```
   POST /create_order_db HTTP/1.1
   Host: localhost:8080
   Content-Type: text/plain
   Content-Length: 1149

   {
   "order_uid": "b563feb7b2b84b6test",
   "track_number": "WBILMTESTTRACK",
   "entry": "WBIL",
   "delivery": {
   "name": "Test Testov",
   "phone": "+9720000000",
   "zip": "2639809",
   "city": "Kiryat Mozkin",
   "address": "Ploshad Mira 15",
   "region": "Kraiot",
   "email": "test@gmail.com"
   },
   "payment": {
   "transaction": "b563feb7b2b84b6test",
   "request_id": "",
   "currency": "USD",
   "provider": "wbpay",
   "amount": 1817,
   "payment_dt": 1637907727,
   "bank": "alpha",
   "delivery_cost": 1500,
   "goods_total": 317,
   "custom_fee": 0
   },
   "items": [
   {
   "chrt_id": 9934930,
   "track_number": "WBILMTESTTRACK",
   "price": 453,
   "rid": "ab4219087a764ae0btest",
   "name": "Mascaras",
   "sale": 30,
   "size": "0",
   "total_price": 317,
   "nm_id": 2389212,
   "brand": "Vivienne Sabo",
   "status": 202
   }
   ],
   "locale": "en",
   "internal_signature": "",
   "customer_id": "test",
   "delivery_service": "meest",
   "shardkey": "9",
   "sm_id": 99,
   "date_created": "2021-11-26T06:22:19Z",
   "oof_shard": "1"
   }
   ```
   
2) Создание заказа в кэшэ

   http://localhost:8080/create_order_cache. HTTP method: POST.
   
   Запрос такого же вида как при создании заказа в бд.

3) Получение заказа из бд

   http://localhost:8080/get_order_db. HTTP method: GET.

   Обязательный параметр order-uid - id заказа

   Пример:
   
   ```
   GET /get_order_db/b563feb7b2b84b6test HTTP/1.1
   Host: localhost:8080
   ```

   
4) Получение заказа из кэша 

   http://localhost:8080/get_order_cache. HTTP method: GET.

   Обязательный параметр order-uid - id заказа.
   Запрос такой же вида как при получение заказа из базы данных ответ

5) Удаление заказа

   Удаляет заказ из базы данных

   http://localhost:8080/delete_order. HTTP method: DELETE.

   Обязательный параметр order-uid - id заказа.

   Пример:
   ```
   DELETE /delete_order/b563feb7b2b84b6test HTTP/1.1
   Host: localhost:8080
   ```

6) Обновление заказа
   
   Обновляет заказ в базе данных

   http://localhost:8080/update_order. HTTP method: PATCH

   Обязательный параметр - json модель заказа

   Пример:

```
PATCH /update_order HTTP/1.1
Host: localhost:8080
Content-Type: text/plain
Content-Length: 1152

{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testovich",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 320,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
```

