# Coffee Store

This is my pet project, the goal of which is to develop a solution for a small coffee company.

Basic requirements that the solution must meet:
- Authorization of organization employees
- Role-based access control
- CRUD products
- Order formation
- Order tracking
- Forwarding the order to the kitchen/barista
- Availability of a loyalty program
- Services observability

# Project architecture at the initial stage
![alt text](docs/architecture.png "zxc")

## Microservices

Depending on the business logic, the following microservices are planned

| Service        | Host         |
|----------------|--------------|
| ProductService | grpc:5000    |
| AuthService    | grpc:5001    |
| RoyaltyService | grpc:5002    |
| OrderService   | grpc:5003    |
| StorageService | S3Minio:5004 |
| BaristaService | http:5005    |
| KitchenService | http:5006    |
| API Gateway    | http:8000    |

## Database

As shown in the diagram, a database-per-service approach is used.

## Message Brokers

To ensure fault tolerance, an asynchronous approach to communication of some microservices is used, in particular, the SAGA and transactional outbox patterns are used.

## Observability

- Metrics:
  - Prometheus
  - Grafana
- Tracing:
  - Jaeger
- Logging:
  - Graylog

# Build 

```sh
make build
```

# RoadMap
- ProductService
- AuthService
- RoyaltyService
- OrderService
- StorageService
- BaristaService
- KitchenService
- API Gateway
- Integration Tests
- UI (Dart/Flutter)