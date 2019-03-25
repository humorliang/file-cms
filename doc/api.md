## 文件操作系统的API
#### 登陆
url:http://localhost:8080/v1/user/login
method:post
- 请求参数：
```bash
form
user_name:user
pass_word:123456
```
- 响应
```json
{
    code:0,
    msg:"success",
    data:{
        "user_id":1
    }
}

```
#### 注册
url:http://localhost:8080/v1/user/register
method:post
- 请求参数：
```
user_name:user
pass_word:123456
```
- 响应
```json
{
    code:0,
    msg:"success",
    data:{
        "user_id":1
    }
}
```
#### 上传文件
url:http://localhost:8080/v1/auth/file/upload
method:post
- 请求参数：
```
file_name:file1
file_data:xxxx
```
- 响应
```json
{
    code:0,
    msg:"success",
    data:{
        "file_id":1
    }
}
```
#### 下载文件
url:http://localhost:8080/v1/auth/file/download/:file_id
method:get
- 请求参数：
```
```
- 响应
```json
{
    code:0,
    msg:"success",
    data:{}
}
```
#### 获取文件列表
url:http://localhost:8080/v1/auth/files/:page_num
method:get
- 请求参数：
```
```
- 响应
```json
{
    code:0,
    msg:"success",
    data:{
        "file_list":[{
            "file_id":"1",
            "file_name":"文件1"
        },{
            "file_id":"2",
            "file_name":"文件2"
        }]
    }
}
```