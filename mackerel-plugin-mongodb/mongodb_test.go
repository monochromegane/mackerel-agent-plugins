package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphDefinition(t *testing.T) {
	var mongodb MongoDBPlugin

	graphdef := mongodb.GraphDefinition()
	if len(graphdef) != 4 {
		t.Errorf("GetTempfilename: %d should be 4", len(graphdef))
	}
}

func TestParse22(t *testing.T) {
	var mongodb MongoDBPlugin
	stub_2_2_7 := `
{"asserts":{"msg":0,"regular":0,"rollovers":0,"user":0,"warning":0},"backgroundFlushing":{"average_ms":0,"flushes":0,"last_finished":"1970-01-01T09:00:00+09:00","last_ms":0,"total_ms":0},"connections":{"available":818,"current":1},"cursors":{"clientCursors_size":0,"timedOut":0,"totalOpen":0},"dur":{"commits":30,"commitsInWriteLock":0,"compression":0,"earlyCommits":0,"journaledMB":0,"timeMs":{"dt":3074,"prepLogBuffer":0,"remapPrivateView":0,"writeToDataFiles":0,"writeToJournal":0},"writeToDataFilesMB":0},"extra_info":{"heap_usage_bytes":25585584,"note":"fields vary by platform","page_faults":136},"globalLock":{"activeClients":{"readers":0,"total":0,"writers":0},"currentQueue":{"readers":0,"total":0,"writers":0},"lockTime":1638,"totalTime":35489000},"host":"58a1c98acba3","indexCounters":{"btree":{"accesses":0,"hits":5,"missRatio":0,"misses":0,"resets":0}},"localTime":"2015-08-17T15:08:02.677+09:00","locks":{".":{"timeAcquiringMicros":{"R":1593,"W":279},"timeLockedMicros":{"R":1906,"W":1638}},"admin":{"timeAcquiringMicros":{},"timeLockedMicros":{}},"local":{"timeAcquiringMicros":{"r":9,"w":0},"timeLockedMicros":{"r":44,"w":0}}},"mem":{"bits":64,"mapped":0,"mappedWithJournal":0,"resident":30,"supported":true,"virtual":128},"network":{"bytesIn":510,"bytesOut":2319,"numRequests":9},"ok":1,"opcounters":{"command":10,"delete":0,"getmore":0,"insert":0,"query":0,"update":0},"pid":1,"process":"mongod","recordStats":{"accessesNotInMemory":0,"local":{"accessesNotInMemory":0,"pageFaultExceptionsThrown":0},"pageFaultExceptionsThrown":0},"uptime":35,"uptimeEstimate":34,"uptimeMillis":35489,"version":"2.2.7","writeBacksQueued":false}
`

	var v interface{}
	err := json.Unmarshal([]byte(stub_2_2_7), &v)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	bsonStats, err := bson.Marshal(v)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	var m bson.M
	err = bson.Unmarshal(bsonStats, &m)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}

	stat, err := mongodb.ParseStatus(m)
	fmt.Println(stat)
	assert.Nil(t, err)
	// Mongodb Stats
	assert.EqualValues(t, reflect.TypeOf(stat["btree_hits"]).String(), "float64")
	assert.EqualValues(t, stat["btree_hits"], 5.0)
}

func TestParse24(t *testing.T) {
	var mongodb MongoDBPlugin
	stub_2_4_14 := `
{"asserts":{"msg":0,"regular":0,"rollovers":0,"user":0,"warning":0},"backgroundFlushing":{"average_ms":0.6153846153846154,"flushes":26,"last_finished":"2015-08-17T14:55:58.622+09:00","last_ms":0,"total_ms":16},"connections":{"available":818,"current":1,"totalCreated":10},"cursors":{"clientCursors_size":0,"timedOut":0,"totalOpen":0},"dur":{"commits":30,"commitsInWriteLock":0,"compression":0,"earlyCommits":0,"journaledMB":0,"timeMs":{"dt":3074,"prepLogBuffer":0,"remapPrivateView":0,"writeToDataFiles":0,"writeToJournal":0},"writeToDataFilesMB":0},"extra_info":{"heap_usage_bytes":62256840,"note":"fields vary by platform","page_faults":181},"globalLock":{"activeClients":{"readers":0,"total":0,"writers":0},"currentQueue":{"readers":0,"total":0,"writers":0},"lockTime":143869,"totalTime":1603601000},"host":"bcd5355930ff","indexCounters":{"accesses":0,"hits":5,"missRatio":0,"misses":0,"resets":0},"localTime":"2015-08-17T14:56:42.209+09:00","locks":{".":{"timeAcquiringMicros":{"R":66884,"W":12244},"timeLockedMicros":{"R":86058,"W":143869}},"admin":{"timeAcquiringMicros":{},"timeLockedMicros":{}},"local":{"timeAcquiringMicros":{"r":513,"w":0},"timeLockedMicros":{"r":11886,"w":0}}},"mem":{"bits":64,"mapped":80,"mappedWithJournal":160,"resident":38,"supported":true,"virtual":341},"metrics":{"document":{"deleted":0,"inserted":1,"returned":0,"updated":0},"getLastError":{"wtime":{"num":0,"totalMillis":0},"wtimeouts":0},"operation":{"fastmod":0,"idhack":0,"scanAndOrder":0},"queryExecutor":{"scanned":0},"record":{"moves":0},"repl":{"apply":{"batches":{"num":0,"totalMillis":0},"ops":0},"buffer":{"count":0,"maxSizeBytes":268435456,"sizeBytes":0},"network":{"bytes":0,"getmores":{"num":0,"totalMillis":0},"ops":0,"readersCreated":0},"oplog":{"insert":{"num":0,"totalMillis":0},"insertBytes":0},"preload":{"docs":{"num":0,"totalMillis":0},"indexes":{"num":0,"totalMillis":0}}},"ttl":{"deletedDocuments":0,"passes":26}},"network":{"bytesIn":1940,"bytesOut":18064,"numRequests":33},"ok":1,"opcounters":{"command":35,"delete":0,"getmore":0,"insert":1,"query":26,"update":0},"opcountersRepl":{"command":0,"delete":0,"getmore":0,"insert":0,"query":0,"update":0},"pid":1,"process":"mongod","recordStats":{"accessesNotInMemory":0,"local":{"accessesNotInMemory":0,"pageFaultExceptionsThrown":0},"pageFaultExceptionsThrown":0},"uptime":1604,"uptimeEstimate":1581,"uptimeMillis":1603600,"version":"2.4.14","writeBacksQueued":false}
`

	var v interface{}
	err := json.Unmarshal([]byte(stub_2_4_14), &v)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	bsonStats, err := bson.Marshal(v)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	var m bson.M
	err = bson.Unmarshal(bsonStats, &m)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}

	stat, err := mongodb.ParseStatus(m)
	fmt.Println(stat)
	assert.Nil(t, err)
	// Mongodb Stats
	assert.EqualValues(t, reflect.TypeOf(stat["btree_hits"]).String(), "float64")
	assert.EqualValues(t, stat["btree_hits"], 5.0)
}

func TestParse26(t *testing.T) {
	var mongodb MongoDBPlugin
	stub_2_6_11 := `
{"asserts":{"msg":0,"regular":0,"rollovers":0,"user":0,"warning":0},"backgroundFlushing":{"average_ms":0,"flushes":0,"last_finished":"1970-01-01T09:00:00+09:00","last_ms":0,"total_ms":0},"connections":{"available":818,"current":1,"totalCreated":1},"cursors":{"clientCursors_size":0,"note":"deprecated, use server status metrics","pinned":0,"timedOut":0,"totalNoTimeout":0,"totalOpen":0},"dur":{"commits":30,"commitsInWriteLock":0,"compression":0,"earlyCommits":0,"journaledMB":0,"timeMs":{"dt":3074,"prepLogBuffer":0,"remapPrivateView":0,"writeToDataFiles":0,"writeToJournal":0},"writeToDataFilesMB":0},"extra_info":{"heap_usage_bytes":62512008,"note":"fields vary by platform","page_faults":228},"globalLock":{"activeClients":{"readers":0,"total":0,"writers":0},"currentQueue":{"readers":0,"total":0,"writers":0},"lockTime":68622,"totalTime":6583000},"host":"08ea07b5a8fd","indexCounters":{"accesses":2,"hits":5,"missRatio":0,"misses":0,"resets":0},"localTime":"2015-08-17T15:08:44.187+09:00","locks":{".":{"timeAcquiringMicros":{"R":254,"W":94},"timeLockedMicros":{"R":520,"W":68622}},"admin":{"timeAcquiringMicros":{"r":28,"w":0},"timeLockedMicros":{"r":338,"w":0}},"local":{"timeAcquiringMicros":{"r":22,"w":0},"timeLockedMicros":{"r":46,"w":0}}},"mem":{"bits":64,"mapped":80,"mappedWithJournal":160,"resident":36,"supported":true,"virtual":342},"metrics":{"cursor":{"open":{"noTimeout":0,"pinned":0,"total":0},"timedOut":0},"document":{"deleted":0,"inserted":1,"returned":0,"updated":0},"getLastError":{"wtime":{"num":0,"totalMillis":0},"wtimeouts":0},"operation":{"fastmod":0,"idhack":0,"scanAndOrder":0},"queryExecutor":{"scanned":0,"scannedObjects":0},"record":{"moves":0},"repl":{"apply":{"batches":{"num":0,"totalMillis":0},"ops":0},"buffer":{"count":0,"maxSizeBytes":268435456,"sizeBytes":0},"network":{"bytes":0,"getmores":{"num":0,"totalMillis":0},"ops":0,"readersCreated":0},"preload":{"docs":{"num":0,"totalMillis":0},"indexes":{"num":0,"totalMillis":0}}},"storage":{"freelist":{"search":{"bucketExhausted":0,"requests":6,"scanned":11}}},"ttl":{"deletedDocuments":0,"passes":0}},"network":{"bytesIn":224,"bytesOut":380,"numRequests":4},"ok":1,"opcounters":{"command":6,"delete":0,"getmore":0,"insert":1,"query":1,"update":0},"opcountersRepl":{"command":0,"delete":0,"getmore":0,"insert":0,"query":0,"update":0},"pid":1,"process":"mongod","recordStats":{"accessesNotInMemory":0,"admin":{"accessesNotInMemory":0,"pageFaultExceptionsThrown":0},"local":{"accessesNotInMemory":0,"pageFaultExceptionsThrown":0},"pageFaultExceptionsThrown":0},"uptime":7,"uptimeEstimate":6,"uptimeMillis":6581,"version":"2.6.11","writeBacksQueued":false}
`

	var v interface{}
	err := json.Unmarshal([]byte(stub_2_6_11), &v)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	bsonStats, err := bson.Marshal(v)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	var m bson.M
	err = bson.Unmarshal(bsonStats, &m)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}

	stat, err := mongodb.ParseStatus(m)
	fmt.Println(stat)
	assert.Nil(t, err)
	// Mongodb Stats
	assert.EqualValues(t, reflect.TypeOf(stat["btree_hits"]).String(), "float64")
	assert.EqualValues(t, stat["btree_hits"], 5.0)
}

func TestParse30(t *testing.T) {
	var mongodb MongoDBPlugin
	stub_3_0_5 := `
{"asserts":{"msg":0,"regular":0,"rollovers":0,"user":0,"warning":0},"backgroundFlushing":{"average_ms":0,"flushes":0,"last_finished":"1970-01-01T09:00:00+09:00","last_ms":0,"total_ms":3},"connections":{"available":818,"current":1,"totalCreated":1},"cursors":{"clientCursors_size":0,"note":"deprecated, use server status metrics","pinned":0,"timedOut":0,"totalNoTimeout":0,"totalOpen":0},"dur":{"commits":31,"commitsInWriteLock":0,"compression":1.1555931725208068,"earlyCommits":0,"journaledMB":0.024576,"timeMs":{"commits":1,"commitsInWriteLock":0,"dt":3037,"prepLogBuffer":0,"remapPrivateView":0,"writeToDataFiles":0,"writeToJournal":3},"writeToDataFilesMB":0.020941},"extra_info":{"heap_usage_bytes":62891464,"note":"fields vary by platform","page_faults":189},"globalLock":{"activeClients":{"readers":0,"total":9,"writers":0},"currentQueue":{"readers":0,"total":0,"writers":0},"totalTime":5609000},"host":"db625bac64b5","localTime":"2015-08-17T15:09:12.821+09:00","locks":{"Collection":{"acquireCount":{"R":7}},"Database":{"acquireCount":{"W":2,"r":7}},"Global":{"acquireCount":{"W":5,"r":21,"w":2}},"MMAPV1Journal":{"acquireCount":{"R":54,"r":5,"w":10}},"Metadata":{"acquireCount":{"W":4}}},"mem":{"bits":64,"mapped":80,"mappedWithJournal":160,"resident":52,"supported":true,"virtual":358},"metrics":{"commands":{"\u003cUNKNOWN\u003e":0,"_getUserCacheGeneration":{"failed":0,"total":0},"_isSelf":{"failed":0,"total":0},"_mergeAuthzCollections":{"failed":0,"total":0},"_migrateClone":{"failed":0,"total":0},"_recvChunkAbort":{"failed":0,"total":0},"_recvChunkCommit":{"failed":0,"total":0},"_recvChunkStart":{"failed":0,"total":0},"_recvChunkStatus":{"failed":0,"total":0},"_transferMods":{"failed":0,"total":0},"aggregate":{"failed":0,"total":0},"appendOplogNote":{"failed":0,"total":0},"applyOps":{"failed":0,"total":0},"authSchemaUpgrade":{"failed":0,"total":0},"authenticate":{"failed":0,"total":0},"availableQueryOptions":{"failed":0,"total":0},"buildInfo":{"failed":0,"total":0},"checkShardingIndex":{"failed":0,"total":0},"cleanupOrphaned":{"failed":0,"total":0},"clone":{"failed":0,"total":0},"cloneCollection":{"failed":0,"total":0},"cloneCollectionAsCapped":{"failed":0,"total":0},"collMod":{"failed":0,"total":0},"collStats":{"failed":0,"total":0},"compact":{"failed":0,"total":0},"connPoolStats":{"failed":0,"total":0},"connPoolSync":{"failed":0,"total":0},"connectionStatus":{"failed":0,"total":0},"convertToCapped":{"failed":0,"total":0},"copydb":{"failed":0,"total":0},"copydbgetnonce":{"failed":0,"total":0},"copydbsaslstart":{"failed":0,"total":0},"count":{"failed":0,"total":0},"create":{"failed":0,"total":0},"createIndexes":{"failed":0,"total":0},"createRole":{"failed":0,"total":0},"createUser":{"failed":0,"total":0},"currentOpCtx":{"failed":0,"total":0},"cursorInfo":{"failed":0,"total":0},"dataSize":{"failed":0,"total":0},"dbHash":{"failed":0,"total":0},"dbStats":{"failed":0,"total":0},"delete":{"failed":0,"total":0},"diagLogging":{"failed":0,"total":0},"distinct":{"failed":0,"total":0},"driverOIDTest":{"failed":0,"total":0},"drop":{"failed":0,"total":0},"dropAllRolesFromDatabase":{"failed":0,"total":0},"dropAllUsersFromDatabase":{"failed":0,"total":0},"dropDatabase":{"failed":0,"total":0},"dropIndexes":{"failed":0,"total":0},"dropRole":{"failed":0,"total":0},"dropUser":{"failed":0,"total":0},"eval":{"failed":0,"total":0},"explain":{"failed":0,"total":0},"features":{"failed":0,"total":0},"filemd5":{"failed":0,"total":0},"find":{"failed":0,"total":0},"findAndModify":{"failed":0,"total":0},"forceerror":{"failed":0,"total":0},"fsync":{"failed":0,"total":0},"geoNear":{"failed":0,"total":0},"geoSearch":{"failed":0,"total":0},"getCmdLineOpts":{"failed":0,"total":0},"getLastError":{"failed":0,"total":0},"getLog":{"failed":0,"total":0},"getParameter":{"failed":0,"total":0},"getPrevError":{"failed":0,"total":0},"getShardMap":{"failed":0,"total":0},"getShardVersion":{"failed":0,"total":0},"getnonce":{"failed":0,"total":1},"grantPrivilegesToRole":{"failed":0,"total":0},"grantRolesToRole":{"failed":0,"total":0},"grantRolesToUser":{"failed":0,"total":0},"group":{"failed":0,"total":0},"handshake":{"failed":0,"total":0},"hostInfo":{"failed":0,"total":0},"insert":{"failed":0,"total":0},"invalidateUserCache":{"failed":0,"total":0},"isMaster":{"failed":0,"total":1},"listCollections":{"failed":0,"total":0},"listCommands":{"failed":0,"total":0},"listDatabases":{"failed":0,"total":0},"listIndexes":{"failed":0,"total":0},"logRotate":{"failed":0,"total":0},"logout":{"failed":0,"total":0},"mapReduce":{"failed":0,"total":0},"mapreduce":{"shardedfinish":{"failed":0,"total":0}},"medianKey":{"failed":0,"total":0},"mergeChunks":{"failed":0,"total":0},"moveChunk":{"failed":0,"total":0},"parallelCollectionScan":{"failed":0,"total":0},"ping":{"failed":0,"total":2},"planCacheClear":{"failed":0,"total":0},"planCacheClearFilters":{"failed":0,"total":0},"planCacheListFilters":{"failed":0,"total":0},"planCacheListPlans":{"failed":0,"total":0},"planCacheListQueryShapes":{"failed":0,"total":0},"planCacheSetFilter":{"failed":0,"total":0},"profile":{"failed":0,"total":0},"reIndex":{"failed":0,"total":0},"renameCollection":{"failed":0,"total":0},"repairCursor":{"failed":0,"total":0},"repairDatabase":{"failed":0,"total":0},"replSetElect":{"failed":0,"total":0},"replSetFreeze":{"failed":0,"total":0},"replSetFresh":{"failed":0,"total":0},"replSetGetConfig":{"failed":0,"total":0},"replSetGetRBID":{"failed":0,"total":0},"replSetGetStatus":{"failed":0,"total":0},"replSetHeartbeat":{"failed":0,"total":0},"replSetInitiate":{"failed":0,"total":0},"replSetMaintenance":{"failed":0,"total":0},"replSetReconfig":{"failed":0,"total":0},"replSetStepDown":{"failed":0,"total":0},"replSetSyncFrom":{"failed":0,"total":0},"replSetUpdatePosition":{"failed":0,"total":0},"resetError":{"failed":0,"total":0},"resync":{"failed":0,"total":0},"revokePrivilegesFromRole":{"failed":0,"total":0},"revokeRolesFromRole":{"failed":0,"total":0},"revokeRolesFromUser":{"failed":0,"total":0},"rolesInfo":{"failed":0,"total":0},"saslContinue":{"failed":0,"total":0},"saslStart":{"failed":0,"total":0},"serverStatus":{"failed":0,"total":1},"setParameter":{"failed":0,"total":0},"setShardVersion":{"failed":0,"total":0},"shardConnPoolStats":{"failed":0,"total":0},"shardingState":{"failed":0,"total":0},"shutdown":{"failed":0,"total":0},"splitChunk":{"failed":0,"total":0},"splitVector":{"failed":0,"total":0},"top":{"failed":0,"total":0},"touch":{"failed":0,"total":0},"unsetSharding":{"failed":0,"total":0},"update":{"failed":0,"total":0},"updateRole":{"failed":0,"total":0},"updateUser":{"failed":0,"total":0},"usersInfo":{"failed":0,"total":0},"validate":{"failed":0,"total":0},"whatsmyuri":{"failed":0,"total":0},"writebacklisten":{"failed":0,"total":0}},"cursor":{"open":{"noTimeout":0,"pinned":0,"total":0},"timedOut":0},"document":{"deleted":0,"inserted":0,"returned":0,"updated":0},"getLastError":{"wtime":{"num":0,"totalMillis":0},"wtimeouts":0},"operation":{"fastmod":0,"idhack":0,"scanAndOrder":0,"writeConflicts":0},"queryExecutor":{"scanned":0,"scannedObjects":0},"record":{"moves":0},"repl":{"apply":{"batches":{"num":0,"totalMillis":0},"ops":0},"buffer":{"count":0,"maxSizeBytes":268435456,"sizeBytes":0},"network":{"bytes":0,"getmores":{"num":0,"totalMillis":0},"ops":0,"readersCreated":0},"preload":{"docs":{"num":0,"totalMillis":0},"indexes":{"num":0,"totalMillis":0}}},"storage":{"freelist":{"search":{"bucketExhausted":0,"requests":8,"scanned":0}}},"ttl":{"deletedDocuments":0,"passes":0}},"network":{"bytesIn":224,"bytesOut":381,"numRequests":4},"ok":1,"opcounters":{"command":5,"delete":0,"getmore":0,"insert":0,"query":1,"update":0},"opcountersRepl":{"command":0,"delete":0,"getmore":0,"insert":0,"query":0,"update":0},"pid":1,"process":"mongod","storageEngine":{"name":"mmapv1"},"uptime":5,"uptimeEstimate":5,"uptimeMillis":5616,"version":"3.0.5","writeBacksQueued":false}
`

	var v interface{}
	err := json.Unmarshal([]byte(stub_3_0_5), &v)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	bsonStats, err := bson.Marshal(v)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	var m bson.M
	err = bson.Unmarshal(bsonStats, &m)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}

	stat, err := mongodb.ParseStatus(m)
	fmt.Println(stat)
	assert.Nil(t, err)
	// Mongodb Stats
	assert.EqualValues(t, reflect.TypeOf(stat["duration_ms"]).String(), "float64")
	assert.EqualValues(t, stat["duration_ms"], 3.0)
}
