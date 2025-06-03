# AstroGo API Gateway

## Overview
This project represents a microservices-based system deployed on Kubernetes, using an API Gateway (KrakenD), NGINX ingress, Kafka for event streaming, and persistent volumes for storage. It is designed for scalability, modular service management, and robust event processing.

![Ekran görüntüsü 2025-06-03 041326](https://github.com/user-attachments/assets/591cccfa-dc0f-4d3d-b548-c501b0a87f2c)


## Architecture
- **API Gateway**: KrakenD
- **Backend Services**:
  - Astronaut Service
  - Mission Service
  - Fuel Service

## Kubernetes Deployment

### Prerequisites
- Kubernetes cluster
- kubectl configured to communicate with the cluster

### Deployment Steps
1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd astrogo
   ```

2. **Deploy KrakenD API Gateway**
   ```bash
   kubectl apply -k k8s/krakend
   ```

3. **Verify Deployment**
   ```bash
   kubectl get pods -n astrogo -l app=krakend
   ```

4. **Access the API Gateway**
   The API Gateway is exposed via Ingress at `api.astrogo.local`.

## Configuration
- **KrakenD Configuration**: Located in `k8s/krakend/configmap.yaml`
- **Deployment Configuration**: Located in `k8s/krakend/deployment.yaml`
- **Ingress Configuration**: Located in `k8s/krakend/ingress.yaml`

## Kubernetes Setup Details
- **ConfigMap**: Contains the KrakenD configuration, specifying endpoints and backend services.
- **Deployment**: Manages the KrakenD pods, ensuring they are running and configured correctly.
- **Ingress**: Routes external traffic to the KrakenD service, allowing access to the API Gateway.

## Kubernetes Layers
- **Pods**: The smallest deployable units in Kubernetes, running the KrakenD container.
- **Services**: Expose the KrakenD pods to the network, allowing internal communication.
- **Ingress**: Manages external access to the services, providing load balancing and SSL termination.
- **ConfigMaps**: Store configuration data in key-value pairs, used by the KrakenD pods.
- **Deployments**: Manage the desired state for pods and replica sets, ensuring high availability.

## Detailed Kubernetes Setup in `k8s/krakend`
- **Deployment**: Defines the KrakenD deployment with resource limits and requests, ensuring efficient resource usage.
- **Service**: Exposes the KrakenD pods internally, allowing other services to communicate with it.
- **Ingress**: Configures external access to the KrakenD service, with specific host and path rules.

## Troubleshooting
- Check pod logs for errors:
  ```bash
  kubectl logs -n astrogo -l app=krakend
  ```
