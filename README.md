# Donut RPC
![donut](./pixel-donut.webp)


a fake donut store built using [grpc](https://grpc.io) and sqlite


## Setup 
pull the docker container
```shell
docker pull ghcr.io/s1ntaxe770r/donut-store:latest
```
start the service 
```shell
docker run -p 5050:5050 9052:9052 -d donut-store:latest
```

## Testing 
to test the service install [grpcurl](https://github.com/fullstorydev/grpcurl)
```shell
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

## List the methods
```shell
grpcurl -plaintext  localhost:5050 DonutShop
```
sample response
```shell
DonutShop.GetDonut
DonutShop.GetDonuts
DonutShop.GetVersion
DonutShop.MakeDonut
```

### Method (DonutShop.GetDonut)
```shell
grpcurl -plaintext localhost:5050 DonutShop/GetDonuts 
```

sample response 
```json
{
  "donuts": [
    {
      "name": "bagel",
      "price": 32,
      "image": "https://none.com/bagel",
      "id": "1"
    },
    {
      "name": "glazed donut",
      "price": 32,
      "image": "https://none.com/bagel",
      "id": "3"
    },
  ]
}
```
### Method (DonutShop.GetDonuts)
```shell
grpcurl -d '{"name":"bagel"}'  -plaintext  localhost:5050 DonutShop/GetDonut
```

sample response 
```json
{
  "name": "bagel",
  "price": 32,
  "image": "https://none.com/bagel",
  "id": "1"
}
```
### Method (DonutShop.GetVersion)
```shell 
grpcurl -d '{}'  -plaintext  localhost:5050 DonutShop/GetVersion
```

sample response
```json
{
  "number": "v0.1"
}
```
### Method (DonutShop.MakeDonut)
```shell
grpcurl -d '{"name":"bagel","price":32,"image":"https://none.com/bagel","id":"001"}'  -plaintext  localhost:5050 DonutShop/MakeDonut
```

```json
{
  "name": "bagel"
}
```

## Accessing metrics 
```shell
curl http://localhost:9092/metrics
```




