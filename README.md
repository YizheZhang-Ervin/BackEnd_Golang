# EZExpress

## Steps of Express

1. npm init
2. npm install xxxxxxx --save (this save to common dependencies)
3. npm install xxxxxxx --save-dev (this save todeveloping dependencies)
4. ..MongoDB\bin> mongod --dbpath ..MongoDB\data
5. Create: server.js  
   3.1 create server + listen port  
   3.2 static resources handling + error URL handling when not has given static resources  
   3.3 client request data conversion
   3.4 URL -> data API handling
6. Run: path> node server.js or npm start

## Things need to install

1. npm install express --save
2. npm install mongoose --save

## Basic Knowledge of request & response

- Request
- request.path: requested path name
- request.query: ?parameter information
- request.body: work with middleware body-parser, body save request side's sent information
- request.method: like POST/GET/DELETE...
- request.get: get response header information
-
- Response
- response.end: end response and return contents
- response.json: return contents to client by sending JSON file (convert string/buffer default)
- response.send: return contents by transmitting object/path/buffer/txt (via response body)
- response.status: return status code(via response header)
- response.type: return content type
- response.set: set header information [ res.set(key,value) or res.set({k1:v1,k2:v2,...}) ]

- middleware
- definition: do sth between request and response
- eg: get all request information and save in request.body
- method: app.use((req,res,next)=>{}) // use next to do the following things

## Basic Knowledge of MongoDB

1. Linux start  
   Path: which mongod  
   Create Folder: mkdir -p /data/db  
   Check enough Space: df -lh  
   Start: mongod --dbpath=/data/db --port=27017
   Start Progress: mongod --dbpath=/data/db --port=27017 --fork --syslog (use fork as backend progress wit system log)  
   Start Progress: mongod --dbpath=/data/db --port=27017 --fork --log=/var/log.mongod.log (use fork as backend progress with self-defined log)  
   log: tail -f /var/log/messages  
   End : ctrl+C  
   End Progress: mongod --shutdown
2. Client usage  
   (DB > Collection > document/field/index)
   DB client: mongo 127.0.0.1:27017  
   View tables: show dbs  
   View collections: show collections  
   Switch DB: use xxDB  
   Create collections: use xxDB + db.xxCollections.insert() // after two commands, creation completed  
   Search: db.xxCollections.find('\_id:...')  
   Count results:db.xxCollections.find().count()  
   Insert: db.xxCollections.insert({'key':'value'});  
   Update (change given parameter): db.xxCollections.update({'name':'xx'},{\$set:{'group':'xx'}},{multi:true})  
   Update (change all parameters): db.xxCollections.save({xx:1,yy:2})  
   Delete: db.xxCollections.remove({xx:1},true) // defualt delete multi all,true=only delete the first satisfied document  
   Delete: db.xxCollections.drop(); //delete all index and data
