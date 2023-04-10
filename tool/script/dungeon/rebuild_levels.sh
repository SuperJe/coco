#!/bin/bash
# 编译工具
go build -o main ../../../app/campaigns
# 执行重建地牢关卡命令, 输出要被删除的关卡和重新索引的关卡奖励到文件中
./main -method rebuild_all -level_file ../../../doc/campaign/selected_level.txt
# 读取要被删除的关卡, 操作mongo删除
awk -F"\t" '{print $2}' ../../../doc/campaign/deleted_level.txt | while read row
do
  # cmd='db.campaigns.update({"_id":ObjectId("549f07f7e21e041139ef28c7")}, {"$unset":{"levels.'$row'":""}})'
  mongo --eval 'db.campaigns.update({"_id":ObjectId("549f07f7e21e041139ef28c7")}, {"$unset":{"levels.'$row'":""}});' coco
done
# 读取重新索引的奖励关卡文件, 重建奖励
while read line
do
  str=$line
  arr=($(echo $str | tr ":" "\n"))
  slug=${arr[0]}
  acv=${arr[1]}
  # 更新下一个奖励关卡
  mongo --eval 'db.achievements.update({"slug":"'$slug'"}, {$set: {"rewards.levels": ["'$acv'"]}});' coco
  # 关闭订阅
  mongo --eval 'db.levels.update({"original":ObjectId("'$acv'")}, {$set:{"requiresSubscription": false}});' coco
  mongo --eval 'db.campaigns.update({"name":"Dungeon"}, {$set:{"levels.'$acv'.requiresSubscription":false}});' coco
  # 关闭练习模式
  mongo --eval 'db.campaigns.update({"name":"Dungeon"}, {$set:{"levels.'$acv'.practice":false}});' coco
done < ../../../doc/campaign/achievements.txt
# 特殊逻辑, 更新部分不符合统一命名规则的奖励关卡
mongo --eval 'db.achievements.update({"slug":"second-kithmaze-complete"}, {$set: {"rewards.levels": ["54d24c49bf87255405a8f834"]}});' coco
# 重新设置背景图片和背景填充颜色
mongo --eval 'db.campaigns.update({"name":"Dungeon", "backgroundImage.width":1366}, {$set: {"backgroundImage.$.image":"../images/dungeon-bk1.jpg"}});' coco
mongo --eval 'db.campaigns.update({"name":"Dungeon", "backgroundImage.width":1920}, {$set: {"backgroundImage.$.image":"../images/dungeon-bk1.jpg"}});' coco
mongo --eval 'db.campaigns.update({"name":"Dungeon"}, {$set: {"backgroundColor": "rgba(45, 45, 45, 1)"}});' coco
mongo --eval 'db.campaigns.update({"name":"Dungeon"}, {$set: {"backgroundColorTransparent": "rgba(45, 45, 45, 0)"}});' coco
# 重新设置关卡坐标
mongo --eval ';' coco
rm main
