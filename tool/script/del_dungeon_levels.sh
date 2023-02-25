#!/bin/bash
go build -o main ../../app/campaigns
./main -method selected_level_id -level_file ../../doc/campaign/selected_level.txt
awk -F"\t" '{print $2}' ../../doc/campaign/level_with_id.txt | while read row
do
  cmd='db.campaigns.update({"_id":ObjectId("549f07f7e21e041139ef28c7")}, {"$unset":{"levels.'$row'":""}})'
  mongo --eval '$cmd' coco
done
rm main
