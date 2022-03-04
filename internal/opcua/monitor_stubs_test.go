package opcua

import (
	"time"

	"github.com/gopcua/opcua"
)

func (m *Monitor) AddSubscription(interval time.Duration, sub Subscription) {
	m.subs[interval] = sub
}

func (m *Monitor) AddMonitoredItems(nodes ...string) {
	for _, n := range nodes {
		l := uint32(len(m.items))
		m.items[l] = n
	}
}

func (m *Monitor) PushNotification(n *opcua.PublishNotificationData) {
	m.notifyCh <- n
}

func (m *Monitor) NotifyChannel() chan *opcua.PublishNotificationData {
	return m.notifyCh
}

func (m *Monitor) Subs() map[time.Duration]Subscription {
	return m.subs
}

func (m *Monitor) Items() map[uint32]string {
	return m.items
}
