# Traffic toll Microservice
Project state -> In development

## Services
- OBU -> data sender for coords
- Receiver -> receive data from OBU and sit on Apache Kafka queue
- Distance calculator -> calculate the distance travelled from queue data
- Invoice generator  -> invoice generator for the customer
- Invoice calculator -> invoice calculator using the distance travelled during delivery windows
- Gateway -> API Gateway so data can be passed between the invoicer and invoice calculator
- Promtetheus -> metrics
- Grafana -> timeseries 

## Installation

1. Clone the repository:
```bash
git clone https://github.com/lucaliebenberg/traffic-toll-microservice.git
```

2. Build the project:
  ```bash
  go build - o traffic-toll-microservice/bin
  ```

3. Run the project:
```
docker run --name kafka -p 9092:9092 -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true bitnami/kafka:latest # docker container
```
```
make run receiver # terminal 1

make run calculator # terminal 2

make run agg # terminal 3

make obu # terminal 4
```

4. To run executable bin file:
```bash
./traffic-toll-microservice/bin
```

## Contributing
If you would like to contribute to the project, please follow these steps:
  1. Fork the repository.
  2. Create a new branch (git checkout -b feature-branch).
  3. Make your changes.
  4. Commit your changes (git commit -am 'Add new feature').
  5. Push to the branch (git push origin feature-branch).
  6. Create a new Pull Request.
