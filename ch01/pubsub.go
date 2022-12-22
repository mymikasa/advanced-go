package ch01

import (
	"sync"
	"time"
)

type (
	subscriber chan any
	topicFunc  func(v any) bool
)

// Publisher 发布者对象
type Publisher struct {
	m sync.RWMutex

	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:  buffer,
		timeout: publishTimeout,

		subscribers: make(map[subscriber]topicFunc),
	}
}

// Subscribe 添加一个新的订阅者， 订阅全部主题
func (p *Publisher) Subscribe() chan any {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic 添加一个新的订阅者， 订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan any {
	ch := make(chan any, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// Evict 退出订阅
func (p *Publisher) Evict(sub chan any) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// Publish 发布一个主题
func (p *Publisher) Publish(v any) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}

}

// Close 关闭发布者对象， 同时关闭所有订阅者通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// 发送主题， 可以容忍一定的超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v any, wg *sync.WaitGroup) {
	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):

	}
}
