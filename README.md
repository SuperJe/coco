# coco

## 部署事项
+ 在云服务器上启动magic-admin
+ 在云服务器上启动magic-admin-ui，修改VUE_APP_BASE_API的ip地址, 修改所属教师的id。
+ 在云服务器上启动data_sync
+ 拿一台公司的Windows机器作为cc游戏的机器，也属于生产环境的机器，操作要小心。
  - 此PC也是保存学生游戏的mongo数据库的机器。
  - 现在admin平台上的就属于脏数据了, 到时候使用此PC的时候要清空脏数据(用户表和进度表)，再在此PC上执行upload程序。
  - 此PC上启动upload程序经过data_sync服务将游戏数据同步到magic-admin的数据库中。
