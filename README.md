# CI/CD Orchestrator

This is a demo CI/CD orchestrator with three microservices: `build-service`,
`test-service`, and `deployment-service`. It exists purely as an example
microservices and Kubernetes deployment—it has no persistence and does not
actually run CI/CD jobs.

## Kubernetes deployment

To use the existing public images from Docker Hub:

1. Run `kubectl apply -f kubernetes/deployment.yaml`.

To create your own images and push to a different container registry:

1. Build and push Docker images for each service.
2. Update `kubernetes/deployment.yaml` with your image URLs.
3. Run `kubectl apply -f kubernetes/deployment.yaml`.

### Add ngrok as your API gateway

1. Edit the `kubernetes/ngrok-api-gateway.yaml` file and replace
   `{YOUR_NGROK_DOMAIN}` with a [branded
domain](https://ngrok.com/docs/guides/how-to-set-up-a-custom-domain/), or using
an ngrok-controlled domain like `{YOUR_SUBDOMAIN}.ngrok.dev`.
2. Run `kubectl apply -f kubernetes/ngrok-api-gateway.yaml`.

## Project structure

- `services/`: Contains the source code for each microservice
- `kubernetes/`: Contains Kubernetes deployment configurations

## Routes

- `build-service`:
  - `POST /builds/trigger`: Trigger a new build
  - `GET /builds/{id}`: Retrieve details about a specific build
  - `GET /builds`: List all builds
- `test-service`:
  - `POST /tests/run`: Initiate a new test run
  - `GET /tests/{id}`: Retrieve details about a specific test run
  - `GET /tests`: List all test runs
- `deployment-service`:
  - `POST /deployments/create`: Create a new deployment
  - `GET /deployments/{id}`: Retrieve details about a specific deployment
  - `GET /deployments`: List all deployments`

