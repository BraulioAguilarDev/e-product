# Ecommerce example

The goal in this example is to use Suite by testify.

# Test code

Generating the pg container

```sh
$ docker run -d -p 5432:5432 --name ecommerce -e POSTGRES_PASSWORD=password postgres
```

Creating the product table. Make sure to have the database previously.

```sh
$ docker exec -i ecommerce /bin/bash -c "PGPASSWORD=password psql --username postgres <db>" < schema.sql
```
