package sctx

// sctx is super context, because wendy is super :)

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func New() Context {
	return &superContext{
		startTime: time.Now(),
		ctx:       context.Background(),
		logs:      []*Log{},
	}
}

type Context interface {
	context.Context
	WithTimeout(duration time.Duration) Context
	GetLogs() string
	AddLog(entity, action string) *Log
	SetFinalError(err error)
	End()
}

type superContext struct {
	ctx context.Context

	duration   float64
	startTime  time.Time
	endTime    time.Time
	finalError string
	logs       []*Log
}

// as we know, super context is a context, so it has to implement all of this method of context
func (sctx *superContext) Deadline() (deadline time.Time, ok bool) { return sctx.ctx.Deadline() }
func (sctx *superContext) Done() <-chan struct{}                   { return sctx.ctx.Done() }
func (sctx *superContext) Err() error                              { return sctx.ctx.Err() }
func (sctx *superContext) Value(key interface{}) interface{}       { return sctx.ctx.Value(key) }

func (sctx *superContext) WithTimeout(duration time.Duration) Context {
	sctx.ctx, _ = context.WithTimeout(sctx.ctx, duration)
	return sctx
}

func (sctx *superContext) GetLogs() string {
	marshalledStruct := struct {
		StartTime  time.Time `json:"start_time"`
		EndTime    time.Time `json:"end_time"`
		Duration   float64   `json:"duration"`
		FinalError string    `json:"final_error"`
		Logs       []*Log    `json:"logs"`
	}{
		StartTime:  sctx.startTime,
		EndTime:    sctx.endTime,
		Duration:   sctx.duration,
		FinalError: sctx.finalError,
		Logs:       sctx.logs,
	}

	bytes, err := json.Marshal(marshalledStruct)
	if err != nil {
		return fmt.Sprintf(`{"parse_error": %s, "final_error": "%s","info":"failed to parse full logs"}`, err, sctx.finalError)
	}

	return string(bytes)
}

func (sctx *superContext) AddLog(entity, action string) *Log {
	log := &Log{
		StartTime: time.Now(),
		Entity:    entity,
		Action:    action,
		InfoMaps:  map[string]interface{}{},
	}

	sctx.logs = append(sctx.logs, log)
	return log
}

type Log struct {
	StartTime time.Time              `json:"start_time"`
	EndTime   time.Time              `json:"end_time"`
	Duration  float64                `json:"duration"`
	Entity    string                 `json:"entity"`
	Action    string                 `json:"action"`
	Error     error                  `json:"error"`
	InfoMaps  map[string]interface{} `json:"additional_info"`
}

func (l *Log) WithInfo(key string, value interface{}) *Log {
	l.InfoMaps[key] = value
	return l
}

func (l *Log) EndWithError(err error) *Log {
	l.EndTime = time.Now()
	l.Duration = time.Since(l.StartTime).Seconds()
	l.Error = errors.New("test")

	return l
}

func (sctx *superContext) SetFinalError(err error) {
	if err != nil {
		sctx.finalError = ""
	} else {
		sctx.finalError = err.Error()
	}
}

func (sctx *superContext) End() {
	sctx.endTime = time.Now()
	sctx.duration = time.Since(sctx.startTime).Seconds()
}
