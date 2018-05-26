Redis Master Chooser
====================

Redis Master Chooser, or `rmc`, is a tool designed help [Redis Sentinel][sentinel] to behave properly on [dynamic environments like when running under Docker and/or Kubernetes][sentinel-docker]. It decides:

- Whether a new Redis instance should be the master.
- Whether a new Redis instance should be a slave and to which instance is the current master that it has to connect to;
- And if a new Redis Sentinel instance should connect to the default master or to the current one (after a previous failover).

## Usage

There is a Docker Compose configuration that can be used to simulate a multi-node setup, together with [the `docker-compose scale` command][dc-scale]. The tool itself is written in Go and can be built with:

    $ make build

## Configuration

Different configuration parameters can be defined via environment variables. The supported settings are:

Environment Variable | Description | Default value
-------------------- | ----------- | -------------
`REDIS_CONF` | Configuration file that will be updated. It expects both Redis and Redis Sentinel to use a file with the same name. | `/etc/redis.conf`
`REDIS_DEFAULT_MASTER` | Redis host which all the other instances will connect to, if Redis Sentinel is not available to point to the current master. | `localhost`
`REDIS_HOST` | Redis server that will be used to test cluster connectivity. Can be a load balancer like a Docker Compose/Kubernetes Service. | `localhost`
`REDIS_MASTER_NAME` | [Name of the cluster][sentinel-config] configured in Redis Sentinel. | `ha-master`
`REDIS_MODE` | Defines if the Redis instance is running on either `redis` or `sentinel` mode. | `redis`
`REDIS_PORT` | Port that will be used to connect to the `REDIS_HOST`. | `6379`
`REDIS_TIMEOUT` | Maximum amount of time to wait when connecting to either Redis or Redis Sentinel instances before considering a failure. Should be specified in [Go ParseDuration format][parse-duration]. | `100ms`
`SENTINEL_HOST` | Redis Sentinel server that will be used to ask for current master. Can be a load balancer like `REDIS_HOST`. | `localhost`
`SENTINEL_PORT` | Port that will be used to connect to the `SENTINEL_HOST`. | `26379`

## Previous works

This project was inspired by the Redis Sentinel example setup described on the book [Kubernetes: Up and Running][book] (Kelsey Hightower, Brendan Burns and Joe Beda), which is unfortunately not complete.

The setup provided in the repository [Smile-SA/redis-ha][smile-redis] offers a better approach, but is a little overcomplicated and doesn't deal with all the use cases that can happen during a failover. It also requires custom Docker images that can be avoided.


[book]: http://shop.oreilly.com/product/0636920043874.do
[dc-scale]: https://docs.docker.com/compose/reference/scale/
[parse-duration]: https://golang.org/pkg/time/#ParseDuration
[sentinel-config]: https://redis.io/topics/sentinel#configuring-sentinel
[sentinel-docker]: https://redis.io/topics/sentinel#sentinel-docker-nat-and-possible-issues
[sentinel]: https://redis.io/topics/sentinel
[smile-redis]: https://github.com/Smile-SA/redis-ha
