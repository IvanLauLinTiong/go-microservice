# Commands

1. Docker run redis:

```cmd
docker run --rm -p 6379:6379 redis:5.0
```

Intereact with container:
```cmd
docker exec -it <myrediscontainer> redis-cli
```

2. Create Order:

```cmd
curl -X POST -d '{"customer_id":"'$(uuidgen)'", "line_items":[{"item_id":"'$(uuidgen)'", "quantity":5, "price":1999}]}' localhost:3000/orders
```


3. Redis commands

GET keys:
```cmd
GET "order:your_key"
```

View set members `orders`:
```cmd
SMEMBERS orders
```
