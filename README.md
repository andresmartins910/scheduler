# Scheduler

Start server:
``` shell
go run .
```

Curls:

```shell
# To test the GET / route
curl http://localhost:1323/

# To test the GET /:id route:
curl http://localhost:1323/1

# To test the POST / route:
curl -X POST -H "Content-Type: application/json" -d '{"title":"Task 1","status":"open"}' http://localhost:1323/

# To test the PUT /:id route:
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Task","status":"in progress"}' http://localhost:1323/2

# To test the DELETE /:id route:
curl -X DELETE http://localhost:1323/1
```