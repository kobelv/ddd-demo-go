package logit

import (
	"container/list"
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const ctxLogFieldsKey = "_logit-log-field-key"

func AddAllLevel(ctx context.Context, args ...Field) {
	mustLogField(ctx).addFields(args...)
}

// FindMetaField 查找 Field,查找不到会返回 nil
func FindLogField(ctx context.Context, key string) Field {
	fs := findLogFields(ctx)
	if fs == nil {
		return Field(zap.Skip())
	}
	return fs.findField(key)
}

func CopyLogID(ctx context.Context) context.Context {
	c := context.WithValue(context.Background(), ctxLogFieldsKey, newLogContextStructure())
	logid := FindLogField(ctx, LogIDKey)
	AddAllLevel(c, logid)
	return c
}

func initLogFields(ctx context.Context) context.Context {
	if findLogFields(ctx) == nil {
		ctx = context.WithValue(ctx, ctxLogFieldsKey, newLogContextStructure())
	}
	return ctx
}

func findLogFields(ctx context.Context) *logContextStructure {
	if ctx == nil {
		return nil
	}
	if v := ctx.Value(ctxLogFieldsKey); v != nil {
		if fm, ok := v.(*logContextStructure); ok {
			return fm
		}
	}
	return nil
}

func mustLogField(ctx context.Context) *logContextStructure {
	fields := findLogFields(ctx)
	if fields == nil {
		switch ctx.(type) {
		case *gin.Context:
			c := ctx.(*gin.Context)
			c.Set(ctxLogFieldsKey, newLogContextStructure())
			return c.Value(ctxLogFieldsKey).(*logContextStructure)
		case context.Context:
			c := context.WithValue(ctx, ctxLogFieldsKey, newLogContextStructure())
			return c.Value(ctxLogFieldsKey).(*logContextStructure)
		}
	}
	return fields
}

// Range 遍历存储在ctx里的Fields
func rangeFields(ctx context.Context, f func(f Field) error) {
	if fields := findLogFields(ctx); fields != nil {
		fields.rangeFields(f)
	}
}

// context log key 存储对象，用来存储日志字段
type logContextStructure struct {
	entry *list.List               // 按添加顺序存储的Field链表
	keys  map[string]*list.Element // field 的 key 与链表的映射
	mtx   sync.RWMutex
}

func newLogContextStructure() *logContextStructure {
	return &logContextStructure{
		entry: list.New(),
		keys:  make(map[string]*list.Element),
	}
}

func (lcs *logContextStructure) addFields(fs ...Field) {
	lcs.mtx.Lock()
	defer lcs.mtx.Unlock()
	for _, f := range fs {
		if old, ok := lcs.keys[f.Key]; ok {
			lcs.entry.Remove(old)
		}
		lcs.keys[f.Key] = lcs.entry.PushBack(f)
	}
}

func (lcs *logContextStructure) delFields(keys ...string) {
	lcs.mtx.Lock()
	defer lcs.mtx.Unlock()
	for _, key := range keys {
		if f, ok := lcs.keys[key]; ok {
			lcs.entry.Remove(f)
			delete(lcs.keys, key)
		}
	}
}

func (lcs *logContextStructure) rangeFields(rangeFunc func(f Field) error) {
	lcs.mtx.RLock()
	defer lcs.mtx.RUnlock()
	for f := lcs.entry.Front(); f != nil; f = f.Next() {
		f1 := f.Value.(Field)
		if rangeFunc(f1) != nil {
			break
		}
	}
}

func (lcs *logContextStructure) findField(key string) Field {
	lcs.mtx.RLock()
	defer lcs.mtx.RUnlock()
	if f, ok := lcs.keys[key]; ok {
		return f.Value.(Field)
	}
	return Field(zap.Skip())
}
