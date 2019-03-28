# golang-springboot-archetype

golang-springboot-archetype is a starter project for any project that wants to integrate with spring-cloud-config, eureka and actuator.
This includes golang-springboot for the actuator, config and eureka implementation

The project parts are:
* actuator - exposes some methods like springboot actuator (/health, /metrics, /env, /info)
* config - connects to "spring cloud config server" and loads the configuration
* eureka - register, deregister and Heartbeat with eureka service
* handler - http manager. Allows to expose new endpoints
* service - Your application implementation

### How run it?

we have two options for starting:
* using env vars
* using command line params

#### env vars

we must export the nexts env vars:

| env vars | description | example |
| ---------- | ---------- | ---------- |
| spring_profiles_active | spring cloud config profile | test
| spring_cloud_config_uri | uri of spring cloud config server | http://127.0.0.1:30606
| spring_cloud_config_label | spring cloud config label | test
| server_port | exposed port | 8080
| eureka_instance_ip_address | IP of this application, accesible by eureka | 10.0.1.5
| spring_application_name | name of application | my-golang-application

#### command line params

we report the following parameters in the same order:
* spring_profiles_active
* spring_cloud_config_uri
* spring_cloud_config_label
* server_port
* eureka_instance_ip_address
* spring_application_name

example:

```golang
go run main.go pre http://localhost:8080 pre 8080 10.5.5.10 my-golang-app
```

this base project exposes:
    http://127.0.0.1:8080/

with the import of actuator, also exposes:
    http://127.0.0.1:8080/info
    http://127.0.0.1:8080/health
    http://127.0.0.1:8080/env
    http://127.0.0.1:8080/metrics

