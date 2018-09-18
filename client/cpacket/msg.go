package cpacket

import (
	"strings"
)

type queue_val []string

var msg queue_val

func init() {
	msg = make([]string, 0)
}

func (q *queue_val) Enque(val string) {
	*q = append(*q, val)

}

func (q *queue_val) Flush() string {
	str := strings.Join(*q, "")
	*q = make([]string, 0)
	return str
}
