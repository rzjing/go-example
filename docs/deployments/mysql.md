
## MySQL 部署文档

1、docker

```bash
docker run --name mysql -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 -d mysql:5.7
```

2、docker swarm

```yaml
version: "3.7"
services:
  mysql:
    image: mysql:${TAG:-latest}
    ports:
      - target: 3306
        published: 3306
        protocol: tcp
        mode: host
    volumes:
      - /data/docker/etc/mysql/my.cnf:/etc/mysql/my.cnf:ro
      - /data/docker/mysql_data:/data/mysql_data:rw
    environment:
      MYSQL_ROOT_PASSWORD: "123456"
      # MYSQL_RANDOM_ROOT_PASSWORD: "yes"
    command:
      - /bin/sh
      - -c
      - cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && docker-entrypoint.sh mysqld
    networks:
      - stack
    deploy:
      replicas: ${REPLICAS:-1}
      placement:
        constraints:
          - node.labels.mysql == true
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

> 需要创建配置文件、数据目录。
