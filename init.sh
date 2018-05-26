#!/bin/sh

set -e

REDIS_CMD="redis-server"
[ -z "$REDIS_MODE" ] && REDIS_MODE="redis"
[ "$REDIS_MODE" = "sentinel" ] && REDIS_CMD="redis-sentinel"

cp /app/${REDIS_MODE}.conf /etc/redis.conf
/app/rmc
exec "$REDIS_CMD" /etc/redis.conf
