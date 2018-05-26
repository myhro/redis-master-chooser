package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ExecCmd(name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), GetRedisTimeout())
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	text := string(out)

	if ctx.Err() != nil {
		return text, ctx.Err()
	} else if err != nil {
		return text, err
	}

	return text, nil
}

func GetEnv(key string, defValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defValue
	}
	return v
}

func GetRedisConf() string {
	return GetEnv("REDIS_CONF", "/etc/redis.conf")
}

func GetRedisHost() string {
	return GetEnv("REDIS_HOST", "localhost")
}

func GetRedisPort() string {
	return GetEnv("REDIS_PORT", "6379")
}

func GetRedisTimeout() time.Duration {
	s := GetEnv("REDIS_TIMEOUT", "100ms")
	t, err := time.ParseDuration(s)
	if err != nil {
		log.Fatalf("REDIS_TIMEOUT error: %v\n", err)
	}
	return t
}

func GetSentinelHost() string {
	return GetEnv("SENTINEL_HOST", "localhost")
}

func GetSentinelMaster() string {
	defaultMaster := GetEnv("REDIS_DEFAULT_MASTER", "localhost")
	masterName := GetEnv("REDIS_MASTER_NAME", "ha-master")

	out, err := ExecCmd("redis-cli", "-h", GetSentinelHost(), "-p", GetSentinelPort(), "SENTINEL", "get-master-addr-by-name", masterName)
	if err != nil {
		return defaultMaster
	}

	addr := strings.Split(out, "\n")[0]
	return addr
}

func GetSentinelPort() string {
	return GetEnv("SENTINEL_PORT", "26379")
}

func UpdateConfigSentinel(masterAddr string) {
	log.Printf("Setting Redis master to: %v\n", masterAddr)
	placeholder := "{{MASTER_ADDR}}"

	input, err := ioutil.ReadFile(GetRedisConf())
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	for i, l := range lines {
		if strings.Contains(l, placeholder) {
			lines[i] = strings.Replace(l, placeholder, masterAddr, 1)
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(GetRedisConf(), []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateConfigSlave(masterAddr string) {
	log.Printf("Setting Redis master to: %v\n", masterAddr)

	f, err := os.OpenFile(GetRedisConf(), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(fmt.Sprintf("slaveof %v %v\n", masterAddr, GetRedisPort()))
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
}

func main() {
	mode := GetEnv("REDIS_MODE", "redis")
	if mode == "redis" {
		_, err := ExecCmd("redis-cli", "-h", GetRedisHost(), "-p", GetRedisPort(), "PING")
		if err == nil {
			log.Print("Redis mode: Slave")
			m := GetSentinelMaster()
			UpdateConfigSlave(m)
		} else {
			log.Print("Redis mode: Master")
		}
	} else if mode == "sentinel" {
		log.Print("Redis mode: Sentinel")
		m := GetSentinelMaster()
		UpdateConfigSentinel(m)
	}
}
