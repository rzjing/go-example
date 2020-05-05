
## 账号管理模块接口文档

**Version: 1.0**

1、获取账号列表

* URL `/v1/account/`
* Method `GET`
* Header `token: tokenString`
* Params

|   field    |  type  | required | commit |
| ---------- | ------ | -------- | ------ |
| name       | string | N | 账号名称 |
| email      | string | N | 邮件地址 |
| page       | int    | N | 页码, 默认 1 |
| page_size  | int    | N | 分页数量, 默认 10, 最大 100 |
| sort       | string | N | 排序字段+方式, e.g. updated_at desc |

* Response

```json5
{
    "code": 200,  // 状态码
    "data": {     // 消息体
        "count": 1, // 账号总数
        "list": [   // 账号列表
            {
                "id": 1,
                "name": "Go Example",
                "email": "go@example.com",
                "status": 0,  // 启停状态, 0 启用 1 停用
                "created_at": 1588477044,
                "updated_at": 1588477044
            }
        ]
    }
}
```

2、创建账号

* URL `/v1/account/`
* Method `POST`
* Header `token: tokenString`
* Params

|   field    |  type  | required | commit |
| ---------- | ------ | -------- | ------ |
| name       | string | Y | 账号名称 |
| email      | string | Y | 邮件地址 |
| password   | string | Y | 登录密码 |
| status     | string | Y | 启停状态, 0 启用 1 停用 |
| remark     | string | N | 备注 ｜

* Response

```json5
{
    "code": 200,
    "msg": "created"
}
```

3、修改账号

* URL `/v1/account/:id`
* Method `PUT`
* Header `token: tokenString`
* Params

|   field    |  type  | required | commit |
| ---------- | ------ | -------- | ------ |
| name       | string | Y | 账号名称 |
| password   | string | N | 登录密码 |
| status     | string | Y | 启停状态, 0 启用 1 停用 |
| remark     | string | N | 备注 ｜

* Response

```json5
{
    "code": 200,
    "msg": "changed"
}
```

4、删除账号

* URL `/v1/account/:id`
* Method `DELETE`
* Header `token: tokenString`
* Response

```json5
{
    "code": 200,
    "msg": "deleted"
}
```
