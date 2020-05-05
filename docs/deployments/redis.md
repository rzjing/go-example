
## Redis 部署文档

1、docker

```bash
docker run --name redis -p 6379:6379 -d redis:5.0-alpine
```

2、docker swarm

```yaml
version: "3.7"
services:
  redis:
    image: redis:5.0-alpine
    ports:
      - target: 6379
        published: 6379
        protocol: tcp
        mode: host
    volumes:
      - /data/docker/etc/redis/redis.conf:/etc/redis.conf:ro
      - /data/docker/redis_data:/data:rw
    command: redis-server /etc/redis.conf
    networks:
      - stack
    deploy:
      replicas: ${REPLICAS:-1}
      placement:
        constraints:
          - node.labels.redis == true
      update_config:
        parallelism: 1
        delay: 10s
        order: stop-first
      restart_policy:
        condition: on-failure
        delay: 5s
        max_attempts: 3
        window: 120s

networks:
  stack:
    driver: overlay
```

> 需要创建配置文件、数据挂载目录。
