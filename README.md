# Dashboards as Code

This repository serves as a demonstration, sample implementation and comparison of the different open-source tools allowing a completely declarative and fully-functional dashboards-as-code setup. Currently only [Grafana](https://grafana.com/grafana/) and [Perses](https://perses.dev/) are the supported dashboard visualization solutions with [Prometheus](https://prometheus.io/) metrics as the data source.

The code requires only Docker Compose to be installed on your machine. It is also very easy to translate the setup to Kubernetes manifests with [ConfigMaps and Volumes](https://kubernetes.io/docs/concepts/configuration/configmap/#using-configmaps-as-files-from-a-pod).

> &#x26a0;&#xfe0f; *Important: Make sure you have Docker Compose installed before you try running this demo. You can find installation instructions at: <https://docs.docker.com/compose/install/>*

## Overview

### Grafana

Grafana is the leading tool for dashboard visualizations of observability data.
It was created in 2014 and with the years it has grown to provide not only dashboards but also advanced alerting, the ability to explore metrics and logs,
and combine different data sources in the same chart (with more than a hundred data sources officially supported).
On the downside, its data model (with only JSON supported) is pretty complex and editing a dashboard's code is difficult and impractical. Additionally, the embedding of Grafana charts and dashboards is limited.

In order to run Grafana with a pre-configured dashboard it needs to have some configuration files in the appropriate folders:

* datasource.yaml - defines the location of the data source for the dashboard
* provider.yaml - an auxiliary config file instructing Grafana to read dashboard definitions from the file system
* dashboard.json - the JSON representation (YAML is not supported) of a dashboard visualizing some standard Golang application metrics

For this demo those files plus a Docker Compose file are located in the `grafana` folder.

### Perses

Perses is a new open-source solution for observability data visualization.
Although it was released pretty recently (in the end of 2022), it is already supported by [The Linux Foundation](https://www.linuxfoundation.org/).
Compared to Grafana, Perses's data model is [clear and flexible](https://perses.dev/docs/perses/v0.45.0/api/dashboard.md/) (with support for JSON and YAML), and editing a dashboard's code is easy.
It also provides npm packages so that anyone can embed charts and dashboards in their own UI.
On the downside, being a very young project means that it lacks some features: the currently supported data sources are only Prometheus and Grafana Tempo, and the choice of chart types is limited.

In order to run Perses with a pre-configured dashboard it needs to have some configuration files in the appropriate folders:

* config.yaml - sets up the basic Perses configuration
* globaldatasource.yaml - defines the location of the data source for all dashboards (you can define data sources per dashboard too)
* project.yaml - dashboards in Perses are grouped into projects
* dashboard.yaml - the YAML representation (JSON is also supported) of a dashboard visualizing some standard Golang application metrics

For this demo those files plus a Docker Compose file are located in the `perses` folder.

### Prometheus

In this demo Prometheus is not in the same Docker Compose stack as the data visualization tools so we can demonstrate how to connect to any Prometheus-compatible data source exposed on your local machine.

The Prometheus Docker Compose file, the custom Prometheus configuration file and an auxiliary demo application are located in the `demo-prometheus` folder.

#### Demo application

We are using a minimal Golang demo application using the [Prometheus Go client library](https://github.com/prometheus/client_golang) to populate the Prometheus server with standard metrics.

The demo application is located in the `demo-prometheus/demo-app` folder.

## Running the demo

First start the demo Prometheus stack.

```shell
docker compose -f demo-prometheus/docker-compose.yml up -d
```

You can explore the UI of the demo Prometheus at <http://localhost:9090>.\
You can explore the demo app metrics output at <http://localhost:8080/metrics>.

Once the demo application is serving metrics and Prometheus is collecting data from it (you see a line chart [here](http://localhost:9090/graph?g0.expr=go_goroutines&g0.tab=0)),
you can start the pre-configured Grafana or Perses.

### Running Grafana

To start Grafana you just need to run its Docker Compose file:

```shell
docker compose -f grafana/docker-compose.yml up -d
```

Then, you can explore the demo dashboard at <http://localhost:3000/d/demo/demo-dashboard>.\
For this demo Grafana is configured so that it does not require a username & password.

### Running Perses

To start Perses you just need to run its Docker Compose file:

```shell
docker compose -f perses/docker-compose.yml up -d
```

Then, you can explore the demo dashboard at <http://localhost:8888/projects/Demo/dashboards/Demo>.\
For this demo Perses is configured so that it does not require a username & password.

### Cleanup

Once you are done with your work, you can clean up all used resources by executing the following commands:

```shell
docker compose -f demo-prometheus/docker-compose.yml down
docker compose -f grafana/docker-compose.yml down
docker compose -f perses/docker-compose.yml down
```
