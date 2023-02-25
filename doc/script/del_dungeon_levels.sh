#!/bin/bash
go build -o main ../../app/campaign
./main -method selected_level_id -level_file ../campaign/selected_level.txt
awk -F"\t" '{print $2}' ../campaign/level_with_id.txt | while read row
do
  echo $row
  #cmd='db.campaigns.update({"_id":ObjectId("549f07f7e21e041139ef28c7")}, {"$unset":{"levels.'$row'":""}})'
  #echo $cmd
done
rm main
