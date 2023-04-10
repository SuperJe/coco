#!/bin/bash
if [ $# != 1 ]; then
  echo "need collection name"
  echo "eg: ./restore.sh campaigns"
  exit
fi

campaigns="campaigns"
if [ $1 == $campaigns ]; then
  echo "copy campaignstmp to campaigns..."
  mongo --eval 'db.campaignstmp.find().forEach(function(doc){ db.campaigns.insert(doc) });' coco
  echo "copy done"
  exit
fi