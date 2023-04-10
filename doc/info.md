[toc]
# 索引
## levels collection
```
> db.levels.getIndexes()
[
    {
    "v" : 1,
    "key" : {
    "_id" : 1
    },
    "name" : "_id_",
    "ns" : "coco.levels"
    },
    
    {
    "v" : 1,
    "unique" : true,
    "key" : {
    "original" : 1,
    "version.major" : -1,
    "version.minor" : -1
    },
    "name" : "version index",
    "ns" : "coco.levels",
    "background" : true,
    "safe" : null
    },
    
    {
    "v" : 1,
    "unique" : true,
    "key" : {
    "slug" : 1
    },
    "name" : "slug index",
    "ns" : "coco.levels",
    "sparse" : true
    },
    
    {
    "v" : 1,
    "key" : {
    "index" : 1,
    "_fts" : "text",
    "_ftsx" : 1
    },
    "name" : "search index",
    "ns" : "coco.levels",
    "language_override" : "searchLanguage",
    "weights" : {
    "description" : 1,
    "name" : 1
    },
    "default_language" : "english",
    "sparse" : true,
    "background" : true,
    "safe" : null,
    "textIndexVersion" : 2
    },
    
    {
    "v" : 1,
    "key" : {
    "index" : 1
    },
    "name" : "index index",
    "sparse" : true,
    "background" : true,
    "safe" : null,
    "ns" : "coco.levels"
    }
]
```