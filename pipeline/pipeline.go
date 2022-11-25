package pipeline

type pipeline[T, R any] struct {
	subscribers []func(R)
	transform   func(T) R
	chout       chan R
}

func (p pipeline[T, R]) Emit(v ...T) {
	go func() {
		for _, _v := range v {
			r := p.transform(_v)
			for i := 0; i < len(p.subscribers); i++ {
				p.subscribers[i](r)
			}
			if p.chout != nil {
				p.chout <- r
			}
		}
	}()
}

func (p *pipeline[T, R]) Subscribe(subscriber func(R)) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p *pipeline[T, R]) GetChan() chan R {
	if p.chout != nil {
		return p.chout
	} else {
		p.chout = make(chan R)
		return p.chout
	}
}

func From[T, R any](transform func(T) R) *pipeline[T, R] {
	return &pipeline[T, R]{
		subscribers: make([]func(R), 0),
		transform:   transform,
	}
}

func Connect[T, R, S any](p1 *pipeline[T, R], p2 *pipeline[R, S]) {
	p1.Subscribe(func(r R) { p2.Emit(r) })
}
