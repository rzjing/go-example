
## 示例

**Version: 1.0**

1、健康检测

* URL `/`
* Method `HEAD`
* Response `200`

2、Hello World

* URL `/trying`
* Method `GET`
* Response

```json5
{
    "code": 200,
    "info": "hello world.",
    "time": "2020-04-28 15:26:36"
}
```

3、Login

* URL `/login`
* Method `POST`
* Params

|   field    |  type  | required | commit |
| ---------- | ------ | -------- | ------ |
| email      | string | Y | 邮件地址 |
| password   | string | Y | 登录密码 |

* Response

```json5
{
    "code": 200,
    "token": "3bc728b2d9feafd17c3e2550eerfd942"
}
```

> 请求频率每分钟 10 次。

4、Logout

* URL `/logout`
* Method `POST`
* Header `token: tokenString`
* Response

```json5
{
    "code": 200,
}
```
