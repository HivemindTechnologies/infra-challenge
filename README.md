# Hivemind's Infrastructure Challenge

Welcome to Hivemind's Infrastructure Challenge! This project is a coding challenge that aims to test your skills in deploying a web service in a highly available environment using Kubernetes. The service is a simple Go-based web application that greets the user with a message and a unique tag.

**Here is what we would like you to do:**

* Deploy a web service to AWS using Terraform.
* Set the HELLO_TAG environment variable to a unique value.
* Add a URL parameter to the greeter.

**To complete the challenge, you will also need to:**

* Build a pipeline for continuous integration and deployment.
* Create documentation that explains how to set up and use the service.
* Ensure that your solution is production-ready by implementing best practices for security, performance, and reliability.
* Submit your solution by sending us a zip file via email.

It's important to note that while we expect the solution to be production-ready, we understand that there may be certain limitations or constraints based on the environment in which the solution is deployed. If there are any areas of production readiness that you are unable to achieve due to environmental limitations, please provide an explanation in your documentation of how you would typically address those concerns in a production environment.

It is also important to keep in mind that the goal of the challenge is not to impose unrealistic requirements or for you to submit a perfect solution, but rather to evaluate the your ability to think critically about the problem and your skills in implementing a solution. Make sure to document any tradeoffs or compromises made and be ready to showcase your solution if invited to a feedback round.

## Building the application

You can build and run a go app in many ways, easiest is the following:

```shell
go build -o greeter greeter.go
./greeter
```
