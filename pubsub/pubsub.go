package pubsub

import (
	"fmt"
	"sync"
	"time"
)

type (
	subscriber      chan interface{}
	topicFilterFunc func(v interface{}) bool
)

type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFilterFunc
}

func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFilterFunc),
	}
}

// 订阅全部主题
func (p *Publisher) Subscribe() subscriber {
	return p.SubscribeTopic(nil)
}

// 订阅符合过滤器规则的主题
func (p *Publisher) SubscribeTopic(filter topicFilterFunc) subscriber {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = filter
	p.m.Unlock()
	return ch
}

// 某订阅者退出
func (p *Publisher) Evict(sub subscriber) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发布主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, filter := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, filter, v, &wg)
	}
	wg.Wait()
}

// 发布主题 可以容忍一定的超时
func (p *Publisher) sendTopic(sub subscriber, filter topicFilterFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if filter != nil && !filter(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
		fmt.Println("timeout")
	}
}

// 关闭所有订阅者
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}
