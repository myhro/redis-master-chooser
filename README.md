Redis Master Chooser
====================

Redis Master Chooser, or `rmc`, is a tool designed help [Redis Sentinel][sentinel] to behave properly on [dynamic environments like when running under Docker and/or Kubernetes][sentinel-docker]. It decides:

- Whether a new Redis instance should be the master.
- Whether a new Redis instance should be a slave and to which instance is the current master that it has to connect to;
- And if a new Redis Sentinel instance should connect to the default master or to the current one (after a previous failover).

## Usage

There is a Docker Compose configuration that can be used to simulate a multi-node setup, together with [the `docker-compose scale` command][dc-scale]. The tool itself is written in Go and can be built with:

    $ make build

## Previous works

This project was inspired by the Redis Sentinel example setup described on the book [Kubernetes: Up and Running][book] (Kelsey Hightower, Brendan Burns and Joe Beda), which is unfortunately not complete.

The setup provided in the repository [Smile-SA/redis-ha][smile-redis] offers a better approach, but is a little overcomplicated and doesn't deal with all the use cases that can happen during a failover. It also requires custom Docker images that can be avoided.


[book]: http://shop.oreilly.com/product/0636920043874.do
[dc-scale]: https://docs.docker.com/compose/reference/scale/
[sentinel-docker]: https://redis.io/topics/sentinel#sentinel-docker-nat-and-possible-issues
[sentinel]: https://redis.io/topics/sentinel
[smile-redis]: https://github.com/Smile-SA/redis-ha
