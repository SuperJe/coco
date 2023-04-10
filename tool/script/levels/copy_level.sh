#!/bin/bash
if [ $# -lt 1 ]; then
  echo "need at least one param"
  echo "eg: ./copy_level.sh copy-level"
  echo "    ./copy_level.sh make-achievement "
  exit
fi

# 编译 可执行文件
go build -o main ../../../app/levels
str="copy-level"
str1="make-achievement"
param=$1
# 打印配置文件
./main -method print-config -config ./config.toml
if [ $param == $str ]; then
  # 复制关卡并插入到campaign中
  ./main -method copy-level -config ./config.toml
  exit
fi

if [ $param == $str1 ]; then
  # 修改奖励, 形成通路
  if [ $# != 6 ]; then
    echo "eg: ./copy_level.sh make-achievement xx-completed levelID1 yy-completed levelID2 Dungeon"
    exit
  fi
  prevSlug=$2
  level1=$3
  currSlug=$4
  level2=$5
  camp=$6
  mongo --eval 'db.achievements.update({"slug":"'$prevSlug'"}, {$set: {"rewards.levels": ["'$level1'"]}});' coco
  mongo --eval 'db.achievements.update({"slug":"'$currSlug'"}, {$set: {"rewards.levels": ["'$level2'"]}});' coco
  mongo --eval 'db.campaigns.update({"name":"'$camp'"}, {$set:{"levels.'$level1'.original":ObjectId("'$level1'")}});' coco

fi


# 删除可执行文件
rm main