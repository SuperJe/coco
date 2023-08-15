[toc]
# mongo
## 更新权限
```
# 从数组中删除
db.achievements.update( {"slug":"shadow-guard-complete"}, {$pull:{ "rewards.levels": "54ca592de4983255055a5478"}})
db.achievements.update( {"slug":"shadow-guard-complete"}, {$pull:{ "rewards.levels": "54ca592de4983255055a5478"}})
# 修改数组某元素值
db.levels.update(
  {"original":ObjectId("5a5f6fe6dfafb6002b2b5b36"), "permissions.access":"read"},
  {'$set':{'permissions.$.target':'63e999f122b0ec015d2df745"'}}
)
```

## 更新订阅状态
```
db.campaigns.update({"name":"Dungeon"}, {$set:{"levels.54527a6257e83800009730c7.requiresSubscription":false}})
db.levels.update({"original":ObjectId("54527a6257e83800009730c7")}, $set:{"requiresSubscription": false})
```

##  更新
## 删除某个奖励关卡
```
db.achievements.update( {"slug":"shadow-guard-complete"}, {$pull:{ "rewards.levels": "54ca592de4983255055a5478"}})
```
db.achievements.update( {"slug":"shadow-guard-complete"}, {$pull:{ "rewards.levels": "54ca592de4983255055a5478"}})

## 更新宝石数量
```
db.users.update({"name":"teacher007"}, {$set: {"earned.gems":800}})
```

## 删除
### 删除地牢的某个关卡
```
db.campaigns.update(
  {"_id":ObjectId("549f07f7e21e041139ef28c7")},
  {"$unset":{"levels.64d4b02ce7cdd400c2b06194":""}}
)
```

### 删除levels集合中的某个文档
```
db.levels.deleteOne({"_id":ObjectId("64d5fa9620759400c82e055d")})
```

## copy某个集合
### copy campaigns
```
db.campaigns.find().forEach(function(doc){ db.campaignstmp.insert(doc) })
```

### copy achievements
```
db.achievements.find().forEach(function(doc){ db.achievementstmp.insert(doc) })
```

### copy levels
```
db.levels.find().forEach(function(doc){ db.levelstmp.insert(doc) })
```

```
db.levels.find().forEach(function(doc){ db.levelsback.insert(doc) })
```

# Docker
```
# 拷贝进容器
docker cp 本地路径 容器ID:容器路径
# 拷贝出容器
docker cp 容器ID:容器路径 本地路径
# 挂载共享卷, 暴露多端口
docker run -itd --name cc -v /Users/jianli.yue/MyCode/docker_share:/home/coco/codecombat/data  -p 0.0.0.0:3020:3000 -p0.0.0.0:27018:27017 operepo/ope-codecombat /bin/bash
# 运行mysql镜像
docker run -itd --name mysql-admin -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

```

# npm
```
npm install --save @vue/composition-api@1.3.0
npm install vue-echarts@4
```