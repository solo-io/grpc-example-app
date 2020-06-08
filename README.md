# gRPC Example App

A simple app composed of two gRPC microservices, `books` and `records`. Currently each microservice only implements a single endpoint.

The purpose of this app is to demonstrate the API Management features of the Solo.io Developer Portal working with gRPC.  

To build the app: `make all`

To deploy to k8s: `kubectl apply -f install/grpc-example-app-default.yaml`

To consume the app as a client, see `test/e2e_test.go`
