# Hivemind's Infrastructure Challenge

Welcome to Hivemind's Infrastructure Challenge! This project is a coding challenge that aims to test your skills in deploying a web service using Kubernetes.

## Overview

You are joining an existing client project aimed at migrating a Kubernetes setup. They currently have an ancient Kubernetes cluster running with a single HTTP service called the `greeting-provider` (located in `services/greeting-provider/`). This service is a JavaScript application, and the existing code, Dockerfile, and Kubernetes manifest for deployment are provided. The client wishes to transition to a new infrastructure managed as code.

Your task is to develop the infrastructure code required to set up a new Kubernetes cluster, initially for a development environment, but with the capability to expand to multiple environments like staging and production in the future. For the scope of this project, you should set up one or more Kubernetes clusters locally using a tool such as Minikube.

Additionally, you're tasked with reviewing the existing `greeting-provider` service and its deployment strategy.

Furthermore, a new HTTP service named `greeter` needs to be added, as specified in `services/greeter`. This service, written in Go, requires a Dockerfile and deployment configuration. It should be externally accessible and will interact with the `greeting-provider` service.

For all services, the client would like that logs are sent to their `OpenSearch` instance. You should have received the `URL` as well as the credentials to access it.

## Expectations:

* **Infrastructure-as-code**: Clean and extensible infrastructure code to easily spin up environments with the above mentioned features.
* **Documentation**: Provide clear documentation for reproducibly setting up, using, and deploying the services.
* **Production-Readiness**: Ensure that your solution adheres to best practices concerning security, performance, and reliability.
* **Submission**: Submit your solution via email as a zip file.

While we expect your solution to be production-ready, we understand there may be limitations due to the deployment environment. Please include in your documentation any production-readiness concerns that cannot be addressed in this environment, along with explanations of how you would resolve these in a production setup.

Remember, the goal of this challenge is to assess your problem-solving skills and your approach to implementing a solution, rather than aiming for absolute perfection. Document any trade-offs or compromises made, and be prepared to discuss your solution if selected for a feedback session.
