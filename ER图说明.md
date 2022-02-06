![](https://s3.bmp.ovh/imgs/2022/01/5b77d2f959361d2e.png)

* 实体集：2个，Member(用户成员)实体集和Course(课程)实体集
* 联系集：2个，Bind(绑定课程)是教师用户成员到课程1对1的联系集，Choice（选课)是学生用户成员到课程的多对多的联系集



一共4个关系模式

* Member表，$Member(\underline{UserID},Nickname,Username,Password,UserType) $

  - UserID,Member表的主键，类型int，Member表每条记录UserID唯一
  - Nickname，用户昵称，类型string，不小于4位不超过20位（字节）
  - Username，用户名，类型string，支持大小写，不小于8位不超过20位（字节）
  - Password，密码，类型string，同时包括大小写、数字，不少于8位不超过20位（字节）
  - UserType，用户类型，类型为枚举值
    * 1:管理员
    * 2:学生
    * 3:教师

* Course表，$Course(\underline{CourseID},TeacherID,Name,Capacity)$

  * CourseID,Course表的主键，类型int，Course表每条记录CourseID唯一
  * TeacherID,类型int，作为外键参考Member表的UserID
  * Name，课程名，类型string
  * Capacity，课程容量，类型int，用于抢课

* Bind表，$Bind(\underline{TeacherID},CourseID)$

  * TeacherID,类型int，作为外键参考Member表的UserID
  * CourseID,类型int，作为外键参考Course表的CourseID

<<<<<<< HEAD
  表示一个教师到一个课程的绑定，以TeacherID作为主键，参考Member表的UserID,CourseID参考Course表的CourseID，一个老师只能绑定一门课程，一门课程只能由一个老师绑定
=======
  bind表的记录表示某老师可以上某门课，一个老师能上不同的课，一个课能被不同老师上
>>>>>>> pre

* Choice表，$Choice(\underline{StudentID},\underline{CourseID})$

  * StudentID，类型int，作为外键参考Member表的UserID
  * CourseID,类型int，作为外键参考Course表的CourseID
  * StudentID和CourseID共同作为Choice表的主键

  choice表的记录表示一个学生选了一门课的记录，一个学生能选不同的课，不同学生能选一个课

