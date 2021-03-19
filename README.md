# microservice-architecture-01
Kubernetes fundamentals

### Deployment

## Create namespace
    kubectl create namespace healthapp
    kubectl config set-context --current --namespace=healthapp

## Run for base version:
    kubectl -f apply ./manifests_base

## Test:
    curl -H 'Host: arch.homework' http://192.168.49.2/health

## Run for star version:
    kubectl -f apply ./manifests_with_star

## Test:
    curl -H 'Host: arch.homework' http://192.168.49.2/otusapp/possokhov/health
