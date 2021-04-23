# Mercari Office Hour at Go Conference 2021 Spring

Let's explore the Mercari-ish microservices consist of [Go](https://golang.org/), [gRPC](https://grpc.io/), [Kubernetes](https://kubernetes.io/) and [Istio](https://istio.io/) **on your laptop**!

![image](https://user-images.githubusercontent.com/2134196/115878694-ca802980-a483-11eb-80cb-fd56e941f168.png)

## Prerequisites

-   Go
-   Docker

## Usage

```console
$ make cluster
```

### Signup

```console
$ curl -s -XPOST -d '{"name":"gopher"}' localhost:30000/auth/signup | jq .
{
  "customer": {
    "id": "ec1fcc77-b565-4477-b609-62bf0c403903",
    "name": "gopher"
  }
}
```

### Signin

```console
$ TOKEN=$(curl -s -XPOST -d '{"name":"gopher"}' localhost:30000/auth/signin | jq .access_token -r)
```

### Create a item

```console
$ curl -s -XPOST -d '{"title":"Keyboard","price":30000}' -H "authorization: bearer $TOKEN" localhost:30000/catalog/items | jq .
{
  "item": {
    "id": "bda92da6-3270-4255-a756-dbe7d0aa333e",
    "customer_id": "ec1fcc77-b565-4477-b609-62bf0c403903",
    "title": "Keyboard",
    "price": "30000"
  }
}
```

### List items

```console
$ curl -s -XGET -d "{}" -H "authorization: bearer $TOKEN" localhost:30000/catalog/items | jq .
{
  "items": [
    {
      "title": "Laptop",
      "price": "20000"
    },
    {
      "id": "bda92da6-3270-4255-a756-dbe7d0aa333e",
      "title": "Keyboard",
      "price": "30000"
    },
    {
      "title": "Mobile Phone",
      "price": "10000"
    }
  ]
}

```

### Get a item detail

```console
$ curl -s -XGET -d "{}" -H "authorization: bearer $TOKEN" localhost:30000/catalog/items/bda92da6-3270-4255-a756-dbe7d0aa333e | jq .
{
  "item": {
    "id": "bda92da6-3270-4255-a756-dbe7d0aa333e",
    "customer_id": "ec1fcc77-b565-4477-b609-62bf0c403903",
    "customer_name": "gopher",
    "title": "Keyboard",
    "price": "30000"
  }
}
```

## Cleanup

```console
$ make clean
```
