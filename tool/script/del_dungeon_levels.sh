#!/bin/bash
# 编译工具
go build -o main ../../app/campaigns
# 执行重建地牢关卡命令, 输出要被删除的关卡和重新索引的关卡奖励到文件中
./main -method rebuild_all -level_file ../../doc/campaign/selected_level.txt
# 读取要被删除的关卡, 操作mongo删除
awk -F"\t" '{print $2}' ../../doc/campaign/deleted_level.txt | while read row
do
  cmd='db.campaigns.update({"_id":ObjectId("549f07f7e21e041139ef28c7")}, {"$unset":{"levels.'$row'":""}})'
  mongo --eval 'db.campaigns.update({"_id":ObjectId("549f07f7e21e041139ef28c7")}, {"$unset":{"levels.'$row'":""}});' coco
done
# 读取重新索引的奖励关卡文件, 重建奖励
cat ../../doc/campaign/achievements.txt | while read line
do
  array=(${string//:/ })
  echo $array[0]
done
rm main
