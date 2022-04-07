package errors

import "errors"

// etcd
var (
	ErrNewEtcdClient     = errors.New("new etcd client failed")
	ErrStartEtcd         = errors.New("start etcd failed")
	ErrStartEtcdCanceled = errors.New("etcd start canceled")
	ErrEtcdURLMap        = errors.New("etcd url map errors")
	ErrEtcdGrantLease    = errors.New("etcd lease failed")
	ErrEtcdTxnInternal   = errors.New("internal etcd transaction errors occurred")
	ErrEtcdTxnConflict   = errors.New("etcd transaction failed, conflicted and rolled back")
	ErrEtcdKVPut         = errors.New("etcd KV put failed")
	ErrEtcdKVDelete      = errors.New("etcd KV delete failed")
	ErrEtcdKVGet         = errors.New("etcd KV get failed")
	ErrEtcdKVGetResponse = errors.New("etcd invalid get value response %v, must only one")
	ErrEtcdGetCluster    = errors.New("etcd get cluster from remote peer failed")
	ErrEtcdMoveLeader    = errors.New("etcd move leader errors")
	ErrEtcdTLSConfig     = errors.New("etcd TLS config errors")
	ErrEtcdWatcherCancel = errors.New("watcher canceled")
	ErrCloseEtcdClient   = errors.New("close etcd client failed")
	ErrEtcdMemberList    = errors.New("etcd member list failed")
	ErrEtcdMemberRemove  = errors.New("etcd remove member failed")
)

var (
	ErrURLParse = errors.New("parse url error")
)
