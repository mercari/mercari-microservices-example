# Mercari Office Hour at Go Conference 2021 Spring

Let's explore the Mercari-ish microservices consist of [Go](https://golang.org/), [gRPC](https://grpc.io/), [Kubernetes](https://kubernetes.io/) and [Istio](https://istio.io/) **on your laptop**!

![image](https://user-images.githubusercontent.com/2134196/115878694-ca802980-a483-11eb-80cb-fd56e941f168.png)

### Microservices

You can find each microservice's implementation and gRPC API definitions under the `/services` directory.

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

After this make target will have been finished, you can check the status of microservices with `./script/kubectl` which is just a tiny wrapper for `kubectl` like below:\

```console
$ ./script/kubectl get pods --all-namespaces | grep -P '^(gateway|authority|catalog|customer|item)'
authority   app-7b559dfd9f-dcr2v   2/2     Running     0   44s
authority   app-7b559dfd9f-z8c54   2/2     Running     0   44s
catalog     app-67cc897d9c-dhcv7   2/2     Running     0   36s
catalog     app-67cc897d9c-nfk7x   2/2     Running     0   36s
customer    app-565bfc5884-bgb8r   2/2     Running     0   28s
customer    app-565bfc5884-lt6q2   2/2     Running     0   28s
gateway     app-cc456cf4d-nsghg    2/2     Running     0   51s
gateway     app-cc456cf4d-wq47s    2/2     Running     0   51s
item        app-84db48bdf-h7q7b    2/2     Running     0   19s
item        app-84db48bdf-l5mnl    2/2     Running     0   19s
```

Now, the Mercari-ish service is listening on port `30000`, and you can explore it like below!

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
