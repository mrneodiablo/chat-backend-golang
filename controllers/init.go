package controllers

var coordinator *Broker

func init() {

	coordinator = NewBroker()
	go coordinator.Start()

}

type Broker struct {
	stopChannel    chan struct{}
	publishChannel chan interface{}
	subChannel     chan chan interface{}
	unsubChannel   chan chan interface{}
}

func NewBroker() *Broker {
	return &Broker{
		stopChannel:    make(chan struct{}),
		publishChannel: make(chan interface{}, 1),
		subChannel:     make(chan chan interface{}, 1),
		unsubChannel:   make(chan chan interface{}, 1),
	}
}

func (b *Broker) Start() {
	subs := map[chan interface{}]struct{}{}
	for {
		select {
		case <-b.stopChannel:
			return
		case msgCh := <-b.subChannel:
			subs[msgCh] = struct{}{}
		case msgCh := <-b.unsubChannel:
			delete(subs, msgCh)
		case msg := <-b.publishChannel:
			for msgCh := range subs {
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

func (b *Broker) Stop() {
	close(b.stopChannel)
}

func (b *Broker) Subscribe() chan interface{} {
	msgCh := make(chan interface{}, 5)
	b.subChannel <- msgCh
	return msgCh
}

func (b *Broker) Unsubscribe(msgCh chan interface{}) {
	b.unsubChannel <- msgCh
}

func (b *Broker) Publish(msg interface{}) {
	b.publishChannel <- msg
}
