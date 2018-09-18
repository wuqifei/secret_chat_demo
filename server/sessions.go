package server

import (
	"time"

	"github.com/wuqifei/chat/logs"
	"github.com/wuqifei/server_lib/concurrent"
	"github.com/wuqifei/server_lib/libnet2"
	"github.com/wuqifei/server_lib/libtime"
)

var (
	// 连接，会被关闭的池
	connPool *concurrent.ConcurrentIDGroupMap
	// 聊天池
	talkPool *concurrent.ConcurrentIDGroupMap

	wheel *libtime.TimerWheel
)

func init() {
	connPool = concurrent.NewCocurrentIDGroup()
	talkPool = concurrent.NewCocurrentIDGroup()
	wheel = libtime.NewTimerWheel()
	taskcheck := libtime.NewTimerTaskTimeOut("time_check", CheckConnSess)

	wheel.AddTask(time.Duration(1)*time.Minute, -1, taskcheck)
}

// 加入连接
func AddConnSess(sess libnet2.Session2Interface) {
	connPool.Set(sess.GetUniqueID(), sess)
}

// 加入聊天
func AddTalkSess(sess libnet2.Session2Interface) {
	connPool.Del(sess.GetUniqueID())
	talkPool.Set(sess.GetUniqueID(), sess)
}

func CheckConnSess(val interface{}) {
	now := time.Now().UTC()
	logs.Debug("chekconnsess[%v] time[%v]", val, now)
	length := len(connPool.SyncMaps)
	for i := 0; i < length; i++ {
		for _, item := range connPool.SyncMaps[i].Items {
			sess := item.(libnet2.Session2Interface)
			v, _ := sess.Get("conn_time")
			if ((v.(int64)) + 60) < now.Unix() {
				// 删除并且关闭
				connPool.Del(sess.GetUniqueID())
				sess.Close()
			}
		}
	}
}
