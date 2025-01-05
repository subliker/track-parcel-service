# Track parcel service

This documentation is also available in Russian [there](./README_RU.md).

<img src="./assets/track-parcel-service.svg">

## Project Idea
This service was implemented to gain experience working with a project structure based on microservice architecture, describing gRPC interaction between services, using a message broker, and deploying the project in a clustered environment.

The **Parcel Tracking Service** was a perfect fit for these requirements. The service was initially split into 6 independent microservices with their own areas of responsibility:

- **Account Service**: Works with user data in the system for display, authentication, and authorization.
- **User Client Service**: Works with the Telegram API, allows users to receive information about parcels, subscribe to updates via the Telegram bot interface. It is used for registering users in the system by their *Telegram user id*.
- **Parcel Service for Users**: Used for retrieving parcel data and managing user subscriptions.
- **Manager Client Service**: Works with the Telegram API, allows managers to create, modify, and update parcels in the system. It is used for registering managers in the system by their *Telegram user id*.
- **Parcel Service for Managers**: Used for managing parcel data by managers. When parcel data is updated, an event is sent to the notification service.
- **Notification Service**: Works with other services via a message broker. It receives events from the queue for notifications and then sends the collected notifications to services responsible for sending them.

## Project Setup
The setup instructions use task commands. You need to install [Taskfile](https://taskfile.dev/) or refer to the commands in the `Taskfile.yml` and manually enter them 😈.

### Test Environment
#### Service Configuration
If you don't want to dig into the configuration of each microservice, you can run the project using the **default** configuration. To do this, use the `task set-example` command. The environment variable and configuration files will be copied to the actual configs.

For **manual** configuration, you can also use the default configurations (which describe all the fields) and modify them according to your needs, or, by reviewing the config structure or documentation, create your own.

Service configs are located in `configs/config.toml`. Configuration of databases, brokers, and other non-Golang services happens using environment variables in `docker-compose`.

To make the client-bots work, it is **important** to specify tokens in the configuration of the `internal/services/manager_bot_service` and `internal/services/user_bot_service` services.

### Cluster Environment
You should already have `kubectl` configured for this project.

In the `k8s_example` directory, there are example service configurations. To make the service work, you only need to add base64-encoded bot tokens to `k8s_example/services/user_bot_service.yml` and `k8s_example/services/manager_bot_service.yml`.

To encrypt the token, run the command `echo -n "<your-token>" | base64` and copy the resulting string into the Secret in the above-mentioned configs.

Ensure that all services are configured properly. Now you can apply the configurations:
```sh
// start broker
kubectl apply -f k8s_example/broker

// start databases, migrations and admin panel
kubectl apply -f k8s_example/database
kubectl apply -f k8s_example/database/migration
kubectl apply -f k8s_example/database/admin

// Start services
kubectl apply -f k8s_example/services
```

## Selected Tools
### Service Development
#### Programming Language
Go (Golang) was chosen for service development, which is my main programming language. There are several reasons why I prefer this language for this type of service:

- Go allows writing code in a consistent style, focusing more on the logic of the code than on its implementation, which is important.
- The language's rich toolset and the large community with libraries provide many ready-made solutions.
- Built-in code generation support makes it easy to work with gRPC, localizations, documentation, and other tools.
- The module and package hierarchy allows for flexible code reuse and abstraction of service functions from one another.
- The final image of the containerized application is only a few dozen megabytes when properly containerized.
- Built-in tools for context handling allow easy integration of graceful shutdowns, request rollbacks, etc.

#### Interface-Based Approach
Interfaces are used for the tools and components of the services. These interfaces interact with the application layers. This approach was chosen to simplify testing (mocking) and to allow using multiple implementations of tools and components.

For example, integrating a different logger into a microservice is as simple as defining and importing a new implementation. The interface implementations return a ready-to-use interface after calling the `New` function.

In the `main` package, the choice of implementations is made (via flags or configuration), which are then passed to the application layers as interfaces.

#### Logging
The logging tool in all services is [uber-go/zap](https://github.com/uber-go/zap).

Zap was chosen for its rich set of logging tools, flexible log formatting, speed, and the ability to use remote log aggregators (for ELK and similar systems).

The interface and implementation are described in `pkg/logger` and `pkg/logger/zap`, respectively. In `main.go`, the `New` function from the `zap` package is called to create the logger, which is then passed to the application layers.

The logger has methods for formatted (Infof, Warnf, etc.) and unformatted (Info, Warn, etc.) logging, as well as the ability to create a logger with a field that can store important data (e.g., the current layer or handler).

For ELK system integration, the zap implementation supports a remote source for log forwarding. The target address is provided as an argument when creating the logger.

#### Configuration
The tool used for reading service configurations is [spf13/viper](https://github.com/spf13/viper), with the configuration files written in TOML format.

Viper was chosen for its ability to read configurations from both environment variables and files, support for default values, and automatic configuration structure assembly.

The configuration initializer is located in the `pkg/config` module, which should be imported to configure Viper by default. The configuration structure, reading, default value setting, and validation are provided in each service's package in `internal/config`.

Configuration validation is performed using a global validator.

#### Containerization
The microservices containerization happens in two stages. In the first stage, the project is built into a Go-based image. In the second stage, a minimal Alpine image is used, with the built application copied from the first stage. Thus, the final image contains only the application build and an empty Alpine Linux system.

Alpine was chosen for its small size and the availability of standard utilities for command-line access.

#### REST API Server
The **Parcel Service for Managers** uses a REST API server to provide access to parcel management via external automated systems.

The router from [gorilla/mux](https://github.com/gorilla/mux) is used for routing. It was chosen for its simplicity, speed, and interaction with standard (from the internal package) HTTP structures and interfaces.

For automatic documentation generation from annotations, [swaggo/swag](https://github.com/swaggo/swag) is used.

#### Stores and Repositories
For working with data storage, we use stores and repositories. This approach also uses interfaces and implementations, so services are not dependent on the choice of database or storage.

SQL queries are generated using [Masterminds/squirrel](https://github.com/Masterminds/squirrel), which simplifies and types query creation.

### Instruction Overview
For quick and easy interaction with services, we use the [Taskfile](https://taskfile.dev/) (an alternative to Makefile). It is used in the project to define build commands, apply standard configurations, run services, generate code, and more.

### Test Environment
To set up a test or development environment, we use `docker-compose` and the standard microservice configuration, which allows you to run services locally without conflicts.

In `docker-compose`, databases, the message broker, and ELK services are started. The remaining services can be started via `go run`.

### Cluster Environment
The orchestration tool used in this project is Kubernetes. In the `k8s_example` directory, microservice configurations, migration jobs, database configurations, and message broker settings are provided.

Microservices were successfully tested in the cluster environment.
