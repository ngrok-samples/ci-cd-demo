# CI/CD Orchestrator

This project implements a simple CI/CD orchestrator with three microservices: Build, Test, and Deployment.

## Services

1. Build Service: Manages build processes
2. Test Service: Manages test runs
3. Deployment Service: Manages deployments

## Local Development

To run the services locally:

1. Install Docker and Docker Compose
2. Run `docker-compose up` in the project root

## Kubernetes Deployment

To deploy to Kubernetes:

1. Build and push Docker images for each service
2. Update `kubernetes/deployment.yaml` with your image URLs
3. Run `kubectl apply -f kubernetes/deployment.yaml`

## Testing

Use the `scripts/test_services.sh` script to test the services. Make sure to update the ngrok URL in the script before running.

## Project Structure

- `services/`: Contains the source code for each microservice
- `kubernetes/`: Contains Kubernetes deployment configurations
- `scripts/`: Contains utility scripts

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
