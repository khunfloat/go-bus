package bus

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

// HandlerFunc คือฟังก์ชันกลางที่ใช้ภายใน bus
type HandlerFunc func(ctx context.Context, data any) (any, error)

// RequestBus คือตัวกลางที่เก็บ mapping ของ service.method → handler
type RequestBus struct {
	mu       sync.RWMutex
	handlers map[string]HandlerFunc
}

// NewRequestBus สร้าง instance ใหม่ของ bus
func NewRequestBus() *RequestBus {
	return &RequestBus{
		handlers: make(map[string]HandlerFunc),
	}
}

// RegisterContract ใช้ลงทะเบียน service ที่ implement interface ตาม contract
// เช่น b.RegisterContract("Exam", (*contract.ExamContract)(nil), port)
func (b *RequestBus) RegisterContract(serviceName string, iface any, impl any) {
	ifaceType := reflect.TypeOf(iface).Elem() // ดึง type ของ interface เช่น contract.ExamContract
	implValue := reflect.ValueOf(impl)

	for i := 0; i < ifaceType.NumMethod(); i++ {
		method := ifaceType.Method(i)
		implMethod := implValue.MethodByName(method.Name)

		if !implMethod.IsValid() {
			panic(fmt.Sprintf("❌ %s missing method: %s", serviceName, method.Name))
		}

		key := fmt.Sprintf("%s.%s", serviceName, method.Name)
		fmt.Println("📡 Registered:", key)

		// wrapper function ที่จะถูกเรียกตอน CallContract
		b.handlers[key] = func(ctx context.Context, data any) (any, error) {
			mType := implMethod.Type()
			args := []reflect.Value{reflect.ValueOf(ctx)}

			// ถ้ามี argument ตัวที่สอง (request struct) ให้ส่งเข้าไปด้วย
			if mType.NumIn() > 1 {
				args = append(args, reflect.ValueOf(data))
			}

			// reflect call → port.FindExamById(ctx, req)
			out := implMethod.Call(args)

			var err error
			if len(out) == 2 && !out[1].IsNil() {
				err = out[1].Interface().(error)
			}
			if len(out) >= 1 {
				return out[0].Interface(), err
			}
			return nil, err
		}
	}
}

// CallContract เรียกใช้งาน handler
// resRaw, _ := b.CallContract(ctx, "Exam.FindExamById", contract.ExamRequest{ID: "EX-001"})
// res := resRaw.(*contract.ExamResponse)
func (b *RequestBus) CallContract(ctx context.Context, key string, req any) (any, error) {
	b.mu.RLock()
	h, ok := b.handlers[key]
	b.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("no handler for key: %s", key)
	}
	return h(ctx, req)
}

// Discover แสดงรายการ handler ทั้งหมดที่ลงทะเบียนอยู่ (ใช้ debug/log)
func (b *RequestBus) Discover(prefix string) []string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	keys := []string{}
	for k := range b.handlers {
		if prefix == "" || len(k) >= len(prefix) && k[:len(prefix)] == prefix {
			keys = append(keys, k)
		}
	}
	return keys
}
