# K8s kubernetes

> Documentation: https://kubernetes.io/

## Overview

Kubernetes is a portable, extensible, open source platform for managing containerized workloads and services, that facilitates both declarative configuration and automation. It has a large, rapidly growing ecosystem. Kubernetes services, support, and tools are widely available.

## About

This project has a lot of ordered yaml files to exemplify the main features and advantages of using Kubernetes in your project.

### Files

- **server.go:** it´s the application main file

---

#### YAML Files

- **01-Kind:** it´s to create a new cluster and its nodes
  - https://kind.sigs.k8s.io/docs/user/configuration/#nodes
- **02_pod:** it´s to create a now POD, unit that contains the provisioned containers
- **03_replicaset:** it´s to scale the application, create multiple Pods and split application traffic
- **04_deployment:** it´s to deploy the application; If the Deployment sees that the build version has changed, it will kill all ReplicaSets and Pods and run again
- **05*service*[nodeport|loadbalancer]:** it´s the gateway to the application
  - ClusterIP
  - NodePort
  - LoadBalancer
  - HeadlerService
- **06_configmap-env:** it´s a configuration map to use in different ways in the application
- **07_secret:** it´s a sensitive data configuration map to use in different ways in the application
- **08_metrics-server:** it´s to collect real-time metrics of how many resources each Pod (each part of Kubernetes) is consuming
- **09_hpa:** it´s to automatically allocate the instances (pods) we need according to resource consumption
- **10_pvc:** Persistent Volume Claim - it´s to be used to make the persistence request of a certain volume
- **11_statefulset:** it´s to create volumes that we want to be created whenever we scale a new database replica
- **12_ingress:** a Loadbalance service generates an external IP when we are running managed services (Azure, Google, Amazon...); Ingress becomes a single entry point and routes the request to the service we want, like an API Gateway
- **13_cluster-issuer:** it´s to crceate TLS certificates
