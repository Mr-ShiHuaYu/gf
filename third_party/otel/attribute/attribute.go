package attribute

type Key string

type Value struct {
	vtype     ValueType
	numeric   int64
	boolVal   bool
	stringVal string
	sliceVal  interface{}
}

type ValueType int

const (
	INVALID ValueType = iota
	BOOL
	INT64
	FLOAT64
	STRING
	ARRAY
)

func (v Value) Type() ValueType {
	return v.vtype
}

func (v Value) AsBool() bool {
	return v.boolVal
}

func (v Value) AsInt64() int64 {
	return v.numeric
}

func (v Value) AsFloat64() float64 {
	return float64(v.numeric)
}

func (v Value) AsString() string {
	return v.stringVal
}

func (v Value) Emit() string {
	return v.stringVal
}

type KeyValue struct {
	Key   Key
	Value Value
}

func String(k, v string) KeyValue {
	return KeyValue{
		Key: Key(k),
		Value: Value{
			vtype:     STRING,
			stringVal: v,
		},
	}
}

func Int64(k string, v int64) KeyValue {
	return KeyValue{
		Key: Key(k),
		Value: Value{
			vtype:   INT64,
			numeric: v,
		},
	}
}

func Int(k string, v int) KeyValue {
	return Int64(k, int64(v))
}

func Bool(k string, v bool) KeyValue {
	return KeyValue{
		Key: Key(k),
		Value: Value{
			vtype:   BOOL,
			boolVal: v,
		},
	}
}

func Float64(k string, v float64) KeyValue {
	return KeyValue{
		Key: Key(k),
		Value: Value{
			vtype:   FLOAT64,
			numeric: int64(v),
		},
	}
}

func Any(k string, v interface{}) KeyValue {
	return String(k, "")
}

func (k Key) String(v string) KeyValue {
	return String(string(k), v)
}

func (k Key) Int64(v int64) KeyValue {
	return Int64(string(k), v)
}

func (k Key) Int(v int) KeyValue {
	return Int(string(k), v)
}

func (k Key) Bool(v bool) KeyValue {
	return Bool(string(k), v)
}

func (k Key) Float64(v float64) KeyValue {
	return Float64(string(k), v)
}
