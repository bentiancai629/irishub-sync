//create database and user
// use sync-iris
// db.createUser(
//     {
//         user:"iris",
//         pwd:"irispassword",
//         roles:[{role:"root",db:"admin"}]
//     }
// )

// create collections
db.createCollection("block");
db.createCollection("sync_task");
db.createCollection("tx_common");
db.createCollection("proposal");
// db.createCollection("tx_msg");
db.createCollection("sync_conf");
db.createCollection("mgo_txn");
db.createCollection("mgo_txn.stash");


// create index
db.account.createIndex({"address": 1}, {"unique": true});
db.block.createIndex({"height": -1}, {"unique": true});

db.sync_task.createIndex({"start_height": 1, "end_height": 1}, {"unique": true});
db.sync_task.createIndex({"status": 1}, {"background": true});

db.tx_common.createIndex({"height": -1});
db.tx_common.createIndex({"time": -1});
db.tx_common.createIndex({"tx_hash": 1}, {"unique": true});
// db.tx_common.createIndex({"from": 1});
// db.tx_common.createIndex({"to": 1});
db.tx_common.createIndex({"type": 1});
db.tx_common.createIndex({"status": 1});
db.tx_common.createIndex({"proposal_id": 1}, {"background": true});
db.tx_common.createIndex({"type": -1, "time": -1, "height": -1}, {"background": true});

db.proposal.createIndex({"proposal_id": 1}, {"unique": true});
db.proposal.createIndex({"status": 1}, {"background": true});
db.proposal.createIndex({"voting_end_time": 1, "deposit_end_time": 1, "status": 1}, {"background": true});

// db.tx_msg.createIndex({"hash": 1}, {"unique": true});

// init data
db.sync_conf.insert({"block_num_per_worker_handle": 50, "max_worker_sleep_time": 120});

// drop collection
// db.account.drop();
// db.block.drop();
// db.proposal.drop();
// db.sync_task.drop();
// db.tx_common.drop();
// db.tx_msg.drop();
// db.mgo_txn.drop();
// db.mgo_txn.stash.drop();

// remove collection data
// db.account.remove({});
// db.block.remove({});
// db.proposal.remove({});
// db.sync_task.remove({});
// db.tx_common.remove({});
// db.tx_msg.remove({});
// db.mgo_txn.remove({});
// db.mgo_txn.stash.remove({});
