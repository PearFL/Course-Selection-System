# 字节后端寒假camp

[TOC]

## 一、大作业 - 选排课系统

### 1 作业说明

- 模块：需要实现 4 个模块，包括登录、成员、排课、抢课。

- 评分说明：实现每个接口会得到相应的分数，最后以总分进行排名。



### 2 模块及考察点

系统包括以下模块

- 登录模块
  - 考察登录的设计与实现，对 HTTP 协议的理解。
    - 账密登录
    - Cookie Session

- 成员模块
  - 考察工程实现能力。
    - CURD 及对数据库的操作
    - 参数校验
      - 参数长度
      - 弱密码校验
    - 权限判断

- 排课模块
  - 主要考察算法（二分图匹配）的实现。

- 抢课模块
  - 主要考察简单秒杀场景的设计。



### 3 物料说明

- zip 包，里面包括 path 和接口的定义
  下面第二节来具体说

- 云主机的账号和密码



### 4 评分

- 课程结束后会对每个小组的系统（除抢课部分）进行跑分（提供服务部署地址，会有自动化的判分程序对大家的系统进行评分），隔天公布成绩。

#### 4.1 评分细则

##### 4.2 功能

完成每个功能模块会获得相应的分数。

##### 4.3 性能

必须在答案正确的情况下，性能才计分。



### 5 作业提交

- 部署：提供服务地址，每天将进行评分。

- 代码：上传至 GitHub 并提供地址，最终将进行代码评审。

- 文档：提供设计文档地址。



### 6 附录

#### 6.1 Go 相关

- https://go.dev/tour/list

#### 6.2 MySQL 相关

- [高性能 ](https://book.douban.com/subject/10443458/)[MySQL](https://book.douban.com/subject/10443458/)

#### 6.3 HTTP 相关

- [HTTP 权威指南](https://book.douban.com/subject/10746113/)
- [图解 HTTP](https://book.douban.com/subject/25863515/)



## 二、开发需求

### 1 总体信息

#### 1.1 服务器信息

ip：
账号：
密码：
SQL密码：



#### 1.3 接口信息

| 模块         | 接口名                      | 接口                      |
| ------------ | --------------------------- | ------------------------- |
| 登录         | 登录                        | POST /auth/login          |
| 登出         | POST /auth/logout           |                           |
| 获取个人信息 | GET /auth/whoami            |                           |
| 成员         | 创建成员                    | POST /member/create       |
| 获取单个成员 | GET /member                 |                           |
| 批量获取成员 | GET /member/list            |                           |
| 更新成员     | POST /member/update         |                           |
| 删除成员     | POST /member/delete         |                           |
| 排课         | 创建课程                    | POST /course/create       |
| 获取课程     | GET /course/get             |                           |
| 绑定课程     | POST /teacher/bind_course   |                           |
| 解绑课程     | POST /teacher/unbind_course |                           |
| 获取老师课程 | GET /teacher/get_course     |                           |
| 排课         | POST /course/schedule       |                           |
| 抢课         | 学生抢课                    | POST /student/book_course |
| 获取学生课表 | GET /student/course         |                           |



#### 1.4 得分标准

![image-20220207163731734](C:\Users\pearfl\AppData\Roaming\Typora\typora-user-images\image-20220207163731734.png)



### 2 types.go信息

```go
package types

// 说明：
// 1. 所提到的「位数」均以字节长度为准
// 2. 所有的 ID 均为 int64（以 string 方式表现）

// 通用结构

type ErrNo int

const (
	OK                 ErrNo = 0
	ParamInvalid       ErrNo = 1   // 参数不合法
	UserHasExisted     ErrNo = 2   // 该 Username 已存在
	UserHasDeleted     ErrNo = 3   // 用户已删除
	UserNotExisted     ErrNo = 4   // 用户不存在
	WrongPassword      ErrNo = 5   // 密码错误
	LoginRequired      ErrNo = 6   // 用户未登录
	CourseNotAvailable ErrNo = 7   // 课程已满
	CourseHasBound     ErrNo = 8   // 课程已绑定过
	CourseNotBind      ErrNo = 9   // 课程未绑定过
	PermDenied         ErrNo = 10   // 没有操作权限
	StudentNotExisted  ErrNo = 11   // 学生不存在
	CourseNotExisted   ErrNo = 12   // 课程不存在
	StudentHasNoCourse ErrNo = 13  // 学生没有课程
	StudentHasCourse   ErrNo = 14  // 学生有课程

	UnknownError       ErrNo = 255 // 未知错误
)


type ResponseMeta struct {
	Code ErrNo
}

type TMember struct {
	UserID   string
	Nickname string
	Username string
	UserType UserType
}

Password

type TCourse struct {
	CourseID string
	Name     string
	TeacherID string
}

Cap
Selected

// -----------------------------------

// 成员管理

type UserType int

const (
	Admin   UserType = 1
	Student UserType = 2
	Teacher UserType = 3
)

// 系统内置管理员账号
// 账号名：JudgeAdmin 密码：JudgePassword2022

// 创建成员
// 参数不合法返回 ParamInvalid

// 只有管理员才能添加

type CreateMemberRequest struct {
	Nickname string   // required，不小于 4 位 不超过 20 位
	Username string   // required，只支持大小写，长度不小于 8 位 不超过 20 位
	Password string   // required，同时包括大小写、数字，长度不少于 8 位 不超过 20 位
	UserType UserType // required, 枚举值
}

type CreateMemberResponse struct {
	Code ErrNo
	Data struct {
		UserID string // int64 范围
	}
}

// 获取成员信息

type GetMemberRequest struct {
	UserID string
}

// 如果用户已删除请返回已删除状态码，不存在请返回不存在状态码

type GetMemberResponse struct {
	Code ErrNo
	Data TMember
}

// 批量获取成员信息

type GetMemberListRequest struct {
	Offset int
	Limit  int
}

type GetMemberListResponse struct {
	Code ErrNo
	Data struct {
		MemberList []TMember
	}
}

// 更新成员信息

type UpdateMemberRequest struct {
	UserID   string
	Nickname string
}

type UpdateMemberResponse struct {
	Code ErrNo
}

// 删除成员信息
// 成员删除后，该成员不能够被登录且不应该不可见，ID 不可复用

type DeleteMemberRequest struct {
	UserID string
}

type DeleteMemberResponse struct {
	Code ErrNo
}

// ----------------------------------------
// 登录

type LoginRequest struct {
	Username string
	Password string
}

// 登录成功后需要 Set-Cookie("camp-session", ${value})
// 密码错误范围密码错误状态码

type LoginResponse struct {
	Code ErrNo
	Data struct {
		UserID string
	}
}

// 登出

type LogoutRequest struct{}

// 登出成功需要删除 Cookie

type LogoutResponse struct {
	Code ErrNo
}

// WhoAmI 接口，用来测试是否登录成功，只有此接口需要带上 Cookie

type WhoAmIRequest struct {
}

// 用户未登录请返回用户未登录状态码

type WhoAmIResponse struct {
	Code ErrNo
	Data TMember
}

// -------------------------------------
// 排课

// 创建课程
// Method: Post
type CreateCourseRequest struct {
	Name string
	Cap  int
}

type CreateCourseResponse struct {
	Code ErrNo
	Data struct {
		CourseID string
	}
}

// 获取课程
// Method: Get
type GetCourseRequest struct {
	CourseID string
}

type GetCourseResponse struct {
	Code ErrNo
	Data TCourse
}

// 老师绑定课程
// Method： Post
// 注：这里的 teacherID 不需要做已落库校验
// 一个老师可以绑定多个课程 , 不过，一个课程只能绑定在一个老师下面
type BindCourseRequest struct {
	CourseID  string
	TeacherID string
}

type BindCourseResponse struct {
	Code ErrNo
}

// 老师解绑课程
// Method： Post
type UnbindCourseRequest struct {
	CourseID  string
	TeacherID string
}

type UnbindCourseResponse struct {
	Code ErrNo
}

// 获取老师下所有课程
// Method：Get
type GetTeacherCourseRequest struct {
	TeacherID string
}

type GetTeacherCourseResponse struct {
	Code ErrNo
	Data struct {
		CourseList []*TCourse
	}
}

// 排课求解器，使老师绑定课程的最优解， 老师有且只能绑定一个课程
// Method： Post
type ScheduleCourseRequest struct {
	TeacherCourseRelationShip map[string][]string // key 为 teacherID , val 为老师期望绑定的课程 courseID 数组
}

type ScheduleCourseResponse struct {
	Code ErrNo
	Data map[string]string   // key 为 teacherID , val 为老师最终绑定的课程 courseID
}

type BookCourseRequest struct {
	StudentID string
	CourseID  string
}

// 课程已满返回 CourseNotAvailable

type BookCourseResponse struct {
	Code ErrNo
}

type GetStudentCourseRequest struct {
	StudentID string
}

type GetStudentCourseResponse struct {
	Code ErrNo
	Data struct {
		CourseList []TCourse
	}
}
```


#### 2.1 数据表设计及参数解析

##### 2.1.1 member表

| 参数     | 字段类型 | 字段说明 |
| -------- | -------- | -------- |
| user_id   | varchar(255)   | 用户id   |
| nickname | varchar(255)   | 昵称     |
| username | varchar(255)   | 用户名   |
| user_type | int      | 用户类型 |
| password | varchar(255)   | 用户密码 |

**Nickname可以修改，Username不可修改，因为Username是登录账号名**



##### 2.1.2 course表
| 参数      | 字段类型 | 字段说明 |
| --------- | -------- | -------- |
| course_id  | varchar(255)   | 课程id   |
| name      | varchar(255)   | 课程名     |
| teacher_id | varchar(255)   | 老师用户id   |
| capacity      | int     | 课程最大容量 |
| cap_selected | int      | 已选择课程的学生数 |

**TCourse表的TeacherID同TMember表的UserID做关联**



##### 2.1.3 choice表
| 参数      | 字段类型 | 字段说明 |
| --------- | -------- | -------- |
| student_id | varchar(255)   | 学生用户id   |
| course_id  | varchar(255)   | 课程id   |


##### 2.1.4 UserType表（不用创建，或者创建个固定表作为参考）

| 参数     | 数值类型 | 数值 |说明 |
| -------- | -------- | -------- | -------- |
| Admin    | int   | 1 | 管理员 |
| Student  | int   | 2 |学生 |
| Teacher  | int   | 3 |老师 |



##### 2.1.5 ErrNo表（不用创建，或者创建个固定表作为参考）

| 参数     | 数值类型 | 数值 |说明 |
| -------- | -------- | -------- | -------- |
| OK    | int   | 0 | 请求成功 |
| ParamInvalid  | int   | 1 | 参数不合法 |
| UserHasExisted  | int   | 2 | 该 Username 已存在 |
| UserHasDeleted    | int   | 3 | 用户已删除 |
| UserNotExisted  | int   | 4 | 用户不存在 |
| WrongPassword  | int   | 5 | 密码错误 |
| LoginRequired    | int   | 6 | 用户未登录 |
| CourseNotAvailable  | int   | 7 |课程已满 |
| CourseHasBound  | int   | 8 |课程已绑定过 |
| CourseNotBind    | int   | 9 | 课程未绑定过 |
| PermDenied  | int   | 10 |没有操作权限 |
| StudentNotExisted  | int   | 11 |学生不存在 |
| CourseNotExisted    | int   | 12 | 课程不存在 |
| StudentHasNoCourse  | int   | 13 |学生没有课程 |
| StudentHasCourse  | int   | 14 |学生有课程 |
| UnknownError  | int   | 255 |未知错误 |



### 3 router.go信息

```go
package types

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	// 成员管理
	g.POST("/member/create")
	g.GET("/member")
	g.GET("/member/list")
	g.POST("/member/update")
	g.POST("/member/delete")

	// 登录

	g.POST("/auth/login")
	g.POST("/auth/logout")
	g.GET("/auth/whoami")

	// 排课
	g.POST("/course/create")
	g.GET("/course/get")

	g.POST("/teacher/bind_course")
	g.POST("/teacher/unbind_course")
	g.GET("/teacher/get_course")
	g.POST("/course/schedule")

	// 抢课
	g.POST("/student/book_course")
	g.GET("/student/course")

}
```



#### 3.1 成员管理接口设计

##### 3.1.1 /member/create 创建成员

URL：180.184.74.1/api/v1/member/create

支持请求：JSON

HTTP请求方式：POSTC:\Users\pearfl\Downloads\camp_v1_0\types\types.go

**请求参数**

| 参数 | 是否必选 | 类型 | 范围 | 说明     |
| ---- | -------- | ---------- |---- |-------- |
| Nickname   | true   | string  |不小于4位，不超过20位 | 昵称 |
| Username   | true  | string  | 只支持大小写，长度不小于8位，不超过20位       | 用户名 |
| Password   | true  | string |同时包括大小写、数字，长度不少于8位，不超过 20 位         | 用户密码 |
| UserType   | true  | int |枚举值（1：管理员；2：学生；3：教师）         | 用户类型 |



**返回字段说明：**

| 参数 | 字段类型 | 字段说明                    |
| ---- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |



**测试数据：**

```json
Nickname:xwy天下第一
Username:xwytxdytxdy
Password:1q2w3e4R
UserType:1
```



**返回结果：**

正确

```json
{
    "Code": 0
}
```



错误（用户名已存在）0,1,2

```json
{
    "Code": 2
}
```



##### 3.1.2 /member 获取成员信息

URL：180.184.74.1/api/v1/member

支持请求：JSON

HTTP请求方式：GET

**请求参数**

| 参数 | 是否必选 | 类型  | 说明     |
| ---- | -------- | ---------- |-------- |
| UserID | true   | string  | 用户id |


**返回字段说明：**

| 参数  | 字段类型 | 字段说明                    |
| ----- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |
| Data | TMember  | 当前UserID在Tmember表的那一组数据 |


**测试数据：**

```json
UserID:6
```

**返回结果：**

正确

```json
{
    "Code": 0,
    "Data": {
		"UserID":"6",
        "Nickname":"xwy天下第一",
        "Username":"xwytxdytxdy",
        "UserType": 1
    }
}
```



错误（用户不存在）

```json
{
    "Code": 4
}
```



##### 3.1.3 /member/list 获取所有成员信息

URL：180.184.74.1/api/v1/member/list

支持请求：JSON

HTTP请求方式：GET

**请求参数**

| 参数 | 是否必选 | 类型  | 说明     |
| ---- | -------- | ----------  |-------- |
| Offset | false  | int  | Offset:n意思是跳过n个数据 |
| Limit | false  | int | Limit:n意思是选取n个数据 |

**如果Offset未传入参数则默认从第一个数据开始**

**如果Limit未传入参数则默认选取到最后一个数据结束**

**返回字段说明：**

| 参数  | 字段类型 | 字段说明                    |
| ----- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |
| MemberList | TMember数组 | 输出整个Tmember表 |


**测试数据：**

```json
Offset:5
Limit:2
```

**返回结果：**

正确

```json
{
    "Code": 0,
    "Data":{
        "MemberList":[
            {
                "UserID":"6",
                "Nickname":"xwy天下第一",
                "Username":"xwytxdytxdy",
                "UserType": 1
            }
            ,
            {
                "UserID":"7",
                "Nickname":"MX天下第一",
                "Username":"mxtxdytxdy",
                "UserType": 1
            }
    	]
    }
}
```



错误（参数不合法）

```json
{
    "Code": 1
}
```



##### 3.1.4 /member/update 更新成员信息

URL：180.184.74.1/api/v1/member/update

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数 | 是否必选 | 类型  | 说明     |
| ---- | -------- | ----------  |-------- |
| UserID | true     | string | 用户id |
| Nickname | true | string | 昵称 |

**返回字段说明：**

| 参数 | 字段类型 | 字段说明                    |
| ---- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |


**测试数据：**

```json
UserID:6
Nickname:xwy永远的神
```

**返回结果：**

正确

```json
{
    "Code": 0
}
```


错误（用户不存在）

```json
{
    "Code": 4
}
```



##### 3.1.5 /member/delete 更新成员信息

URL：180.184.74.1/api/v1/member/delete

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数 | 是否必选 | 类型  | 说明     |
| ---- | -------- | ----------  |-------- |
| UserID | true     | string | 用户id |

**返回字段说明：**

| 参数 | 字段类型 | 字段说明                    |
| ---- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |


**测试数据：**

```json
UserID:6
```

**返回结果：**

正确

```json
{
    "Code": 0
}
```


错误（用户不存在）

```json
{
    "Code": 4
}
```


#### 3.2 登录接口设计

##### 3.2.1 /auth/login 登录账号

URL：180.184.74.1/api/v1/auth/login

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数     | 是否必选 | 类型   |说明     |
| -------- | -------- | ------ | -------- |
| Username | true     | string | 用户名   |
| Password | true     | string | 用户密码 |

**登录成功后需要 Set-Cookie("camp-session", ${value})**

**返回字段说明：**

| 参数  | 字段类型 | 字段说明                    |
| ----- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |
| Data | UserID | 返回对应的UserID数据 |


**测试数据：**

```json
Username:xwytxdytxdy
Password:1q2w3e4R
```

**返回结果：**

正确

```json
{
    "Code": 0,
    "Data":{
        "UserID": "6"
    }
    
}
```

错误（密码错误）

```json
{
    "Code": 5
}
```

##### 3.2.2 /auth/logout 登出账号

URL：180.184.74.1/api/v1/auth/logout

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数     | 是否必选 | 类型   |说明     |
| -------- | -------- | ------ | -------- |
| 无 |      |  |    |

**登出成功后需要删除 Cookie **

**返回字段说明：**

| 参数 | 字段类型 | 字段说明                    |
| ---- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |


**测试数据：**

```json
无
```

**返回结果：**

正确

```json
{
    "Code": 0
}
```

错误（用户未登录）

```json
{
    "Code": 6
}
```

##### 3.2.3 /auth/whoami 测试是否登录成功，只有此接口需要带上 Cookie

URL：180.184.74.1/api/v1/auth/whoami

支持请求：JSON

HTTP请求方式：GET

**请求参数**

| 参数     | 是否必选 | 类型   |说明     |
| -------- | -------- | ------ | -------- |
| 无 |      |  |    |


**返回字段说明：**

| 参数  | 字段类型 | 字段说明                    |
| ----- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |
| Data | Tmember      | 返回当前用户的Tmember数据 |

**测试数据：**

```json
无
```

**返回结果：**

正确

```json
{
    "Code": 0,
    "Data": {
		"UserID":"6",
        "Nickname":"xwy天下第一",
        "Username":"xwytxdytxdy",
        "UserType": 1
    }
}
```

错误（用户未登录）

```json
{
    "Code": 6
}
```

#### 3.3 排课接口设计

##### 3.3.1 /course/create 创建课程

URL：180.184.74.1/api/v1/course/create

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数     | 是否必选 | 类型   |说明     |
| -------- | -------- | ------ | -------- |
| Name | true     | string | 课程名   |
| Cap | true     | int | 用户密码 |

**返回字段说明：**

| 参数  | 字段类型 | 字段说明                    |
| ----- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |
| Data | CourseID | 返回对应的CourseID数据 |


**测试数据：**

```json
Name:高等数学
Cap:500
```

**返回结果：**

正确

```json
{
    "Code": 0,
    "Data":{
        "CourseID": "8"
    }
    
}
```

错误（没有操作权限）

```json
{
    "Code": 10
}
```

##### 3.3.2 /course/get 获取课程

URL：180.184.74.1/api/v1/course/get

支持请求：JSON

HTTP请求方式：GET

**请求参数**

| 参数     | 是否必选 | 类型   |说明     |
| -------- | -------- | ------ | -------- |
| CourseID | true     | string | 课程id   |


**返回字段说明：**

| 参数  | 字段类型 | 字段说明                    |
| ----- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |
| Data | TCourse | 返回对应的TCourse数据 |


**测试数据：**

```json
CourseID:8
```

**返回结果：**

正确

```json
{
    "Code": 0,
    "Data":{
    	"CourseID": "8",
        "Name": "高等数学",
        "TeacherID": "张三"
    }
}
```

错误（没有操作权限）

```json
{
    "Code": 10
}
```

##### 3.3.3 /teacher/bind_course 老师绑定课程

URL：180.184.74.1/api/v1/teacher/bind_course

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数     | 是否必选 | 类型   |说明     |
| -------- | -------- | ------ | -------- |
| CourseID | true     | string | 课程id   |
| TeacherID | true     | string | 教师id   |

**返回字段说明：**

| 参数 | 字段类型 | 字段说明                    |
| ---- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |


**测试数据：**

```json
CourseID:8
TeacherID:10
```

**返回结果：**

正确

```json
{
    "Code": 0
}
```

错误（没有操作权限）

```json
{
    "Code": 10
}
```



##### 3.3.4 /teacher/unbind_course 老师解绑课程

URL：180.184.74.1/api/v1/teacher/unbind_course

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数     | 是否必选 | 类型   |说明     |
| -------- | -------- | ------ | -------- |
| CourseID | true     | string | 课程id   |
| TeacherID | true     | string | 教师id   |

**返回字段说明：**

| 参数 | 字段类型 | 字段说明                    |
| ---- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |


**测试数据：**

```json
CourseID:8
TeacherID:10
```

**返回结果：**

正确

```json
{
    "Code": 0
}
```

错误（没有操作权限）

```json
{
    "Code": 10
}
```

##### 3.3.5 /teacher/get_course 获取老师下所有课程

URL：180.184.74.1/api/v1/teacher/get_course

支持请求：JSON

HTTP请求方式：GET

**请求参数**

| 参数     | 是否必选 | 类型   |说明     |
| -------- | -------- | ------ | -------- |
| TeacherID | true     | string | 教师id   |

**返回字段说明：**

| 参数  | 字段类型 | 字段说明                    |
| ----- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |
| Data | TCourse数组 | 返回对应的TCourse数组 |

**测试数据：**

```json
TeacherID:10
```

**返回结果：**

正确

```json
{
    "Code": 0,
    "Data":{
        "CourseList":[
            {
                "CourseID": "8",
                "Name": "高等数学",
                "TeacherID": "张三"
            }
            ,
            {
                "CourseID": "9",
                "Name": "线性代数",
                "TeacherID": "张三"
            }
        ]    
    }
    
    
}
```

错误（没有操作权限）

```json
{
    "Code": 10
}
```

##### 3.3.6 /course/schedule 排课求解器，使老师绑定课程的最优解， 老师有且只能绑定一个课程

URL：180.184.74.1/api/v1/course/schedule

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数     | 是否必选 | 类型   |说明     |
| -------- | -------- | ------ | -------- |
| TeacherCourseRelationShip | true     | map\[string][]string | key 为 teacherID , val 为老师期望绑定的课程 courseID 数组   |

**返回字段说明：**

| 参数  | 字段类型 | 字段说明                    |
| ----- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |
| Data | map[string]string    | key 为 teacherID , val 为老师最终绑定的课程 courseID |

**测试数据：**

此接口暂时未写测试数据

```json

```

**返回结果：**

正确

```json
{
    "Code": 0
}
```

错误（没有操作权限）

```json
{
    "Code": 10
}
```



#### 3.4 抢课接口设计

##### 3.4.1 /student/book_course 学生抢课

URL：180.184.74.1/api/v1/student/book_course

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数      | 是否必选 | 类型   | 说明   |
| --------- | -------- | ------ | ------ |
| StudentID | true     | string | 学生id |
| CourseID  | true     | string | 课程id |

**返回字段说明：**

| 参数 | 字段类型 | 字段说明                    |
| ---- | -------- | --------------------------- |
| Code | int      | 请求状态（参考上面ErrNo表） |


**测试数据：**

```json
StudentID:7
CourseID:8
```

**返回结果：**

正确

```json
{
    "Code": 0
}
```



错误（课程已满）

```json
{
    "Code": 7
}
```

##### 3.4.2 /student/course 获取学生课表

URL：180.184.74.1/api/v1/student/book_course

支持请求：JSON

HTTP请求方式：POST

**请求参数**

| 参数      | 是否必选 | 类型   | 说明   |
| --------- | -------- | ------ | ------ |
| StudentID | true     | string | 学生id |

**返回字段说明：**

| 参数  | 字段类型    | 字段说明                    |
| ----- | ----------- | --------------------------- |
| Code | int         | 请求状态（参考上面ErrNo表） |
| Data | TCourse数组   | 请求状态（参考上面ErrNo表） |

**测试数据：**

```json
StudentID:7
```

**返回结果：**

正确

```json
{
    "Code": 0,
    "Data":{
    	"CourseList":[
    		{
                "CourseID": "8",
                "Name": "高等数学",
                "TeacherID": "张三"
            }
            ,
            {
                "CourseID": "9",
                "Name": "线性代数",
                "TeacherID": "张三"
            }
    	]
    }
}
```



错误（未知操作）

```json
{
    "Code": 255
}
```
