package cron

import (
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
)

type Cron struct {
	loc *time.Location
	job map[uuid.UUID]*job
}

func New(location *time.Location) *Cron {
	return &Cron{
		loc: location,
	}
}

type job struct {
	tick *time.Ticker
	fn   func()

	name string
}

func (c *Cron) Every(duration time.Duration) *job {
	val := &job{
		tick: time.NewTicker(duration),
	}
	if c.job == nil {
		c.job = make(map[uuid.UUID]*job)
	}
	c.job[uuid.New()] = val

	return val
}

func (j *job) Name(name string) *job {
	j.name = name
	return j
}

func (j *job) Do(fn func()) {
	j.fn = fn
}

func (c *Cron) RunAsync() {
	for {
		for _, v := range c.job {
			select {
			case <-v.tick.C:
				go func(v *job) {
					defer func() {
						if r := recover(); r != nil {
							buf := make([]byte, 4096)
							runtime.Stack(buf, false)
							fmt.Println(string(buf))
						}
					}()

					v.fn()
				}(v)
			default:
			}
		}
	}
}

func (c *Cron) RunBlocking() {
	for _, v := range c.job {
		select {
		case <-v.tick.C:
			v.fn()
		default:
		}
	}
}

func (c *Cron) Stop() {
	for _, v := range c.job {
		v.tick.Stop()
	}
}
