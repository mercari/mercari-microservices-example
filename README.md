# Mercari Office Hour at Go Conference 2021 Spring

Let's explore the Mercari-ish microservices consist of [Go](https://golang.org/), [gRPC](https://grpc.io/), [Kubernetes](https://kubernetes.io/) and [Istio](https://istio.io/) **on your laptop**!

![image](https://user-images.githubusercontent.com/2134196/115878694-ca802980-a483-11eb-80cb-fd56e941f168.png)

### Microservices

#### Gateway

-   This is the only microservice which is facing the out side of the Kubernetes cluster and acts as just a proxy.
-   This microservice is responsible for authenticate the request by verifying the access token (JWT) with public keys wehich can be fetched from the Authority microservice.
-   This microservice also transcodes the JSON requests to the gRPC Protocol Buffers.

#### Authority

-   This microservice is responsible for issuing the access token (JWT) for the customer app.
-   This microservice also provides public keys as a gRPC endpoint to make other microservices be albe to verify the signature of the access token.

#### Catalog

-   This microservice is responsible for aggragating data from the Customer and the Item microservices to make a API caller easily consume it.
-   This microserivce acts like a Backend For Frontend (BFF).

#### Customer

-   This microservice is responsible for storing the customer information to the database and providing it as APIs.

#### Item

-   This microservice is responsible for storing the item information to the database and providing it as APIs.

## Usage

### Prerequisites

-   Go
-   Docker

### Lanuch the Kubernetes cluster

```console
$ make cluster
```

This make target does following tasks:

-   Launch the [Kubernetes](https://kubernetes.io/) cluster on your laptop by using [kind](https://github.com/kubernetes-sigs/kind).
-   Install [Istio](https://istio.io/) to the Kubernetes cluster.
-   Build Docker images of all microservices placed under the `/services` directory.
-   Deploy all microservices to the Kubernetes cluster.

After this make target will have been finished, the Mercari-ish service will listen on port `30000`.\
Now you can play with the Mercari-ish service like below!

### API

#### Sign up

```console
$ curl -s -XPOST -d '{"name":"gopher"}' localhost:30000/auth/signup | jq .
{
  "customer": {
    "id": "ec1fcc77-b565-4477-b609-62bf0c403903",
    "name": "gopher"
  }
}
```

#### Sign in

```console
$ TOKEN=$(curl -s -XPOST -d '{"name":"gopher"}' localhost:30000/auth/signin | jq .access_token -r)
```

#### Create a item

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

#### List items

```console
$ curl -s -XGET -H "authorization: bearer $TOKEN" localhost:30000/catalog/items | jq .
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

#### Get a item detail

```console
$ curl -s -XGET -H "authorization: bearer $TOKEN" localhost:30000/catalog/items/bda92da6-3270-4255-a756-dbe7d0aa333e | jq .
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

### Clean up

```console
$ make clean
```
