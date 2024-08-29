package tapd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

// -----------------------------------------------------------------------------
// PriorityLabel is a type for priority labels.
// -----------------------------------------------------------------------------

type PriorityLabel string

const (
	High       PriorityLabel = "High"
	Middle     PriorityLabel = "Middle"
	Low        PriorityLabel = "Low"
	NiceToHave PriorityLabel = "Nice To Have"
)

func (p PriorityLabel) String() string {
	return string(p)
}

// -----------------------------------------------------------------------------
// ID is a query encoder for id parameters.
// -----------------------------------------------------------------------------

type ID struct {
	id []int
}

var _ query.Encoder = (*ID)(nil)

func NewID(id ...int) *ID {
	return &ID{id: id}
}

func (i *ID) EncodeValues(key string, v *url.Values) error {
	ids := make([]string, 0, len(i.id))
	for _, id := range i.id {
		ids = append(ids, strconv.Itoa(id))
	}
	if len(ids) > 0 {
		v.Add(key, strings.Join(ids, ","))
	}
	return nil
}

// -----------------------------------------------------------------------------
// Order is a query encoder for order parameters.
// -----------------------------------------------------------------------------

type OrderType string

const (
	OrderTypeAsc  OrderType = "asc"
	OrderTypeDesc OrderType = "desc"
)

// Order is a type for order parameters.
type Order struct {
	Field     string
	OrderType OrderType
}

var (
	_ json.Marshaler   = (*Order)(nil)
	_ json.Unmarshaler = (*Order)(nil)
	_ query.Encoder    = (*Order)(nil)
)

type OrderOption func(*Order)

func WithOrderType(orderType OrderType) OrderOption {
	return func(o *Order) {
		o.OrderType = orderType
	}
}

var (
	OrderAsc  = WithOrderType(OrderTypeAsc)
	OrderDesc = WithOrderType(OrderTypeDesc)
)

func NewOrder(field string, opts ...OrderOption) *Order {
	o := &Order{
		Field:     field,
		OrderType: OrderTypeAsc,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%s %s", o.Field, o.OrderType))
}

func (o *Order) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}
	_, err := fmt.Sscanf(s, "%s %s", &o.Field, &o.OrderType)
	return err
}

func (o *Order) EncodeValues(key string, v *url.Values) error {
	v.Add(key, fmt.Sprintf("%s %s", o.Field, o.OrderType))
	return nil
}

// -----------------------------------------------------------------------------
// Fields is a query encoder for fields parameters.
// -----------------------------------------------------------------------------

// Fields todo: refactor to type Fields []string
type Fields struct {
	fields []string
}

var _ query.Encoder = (*Fields)(nil)

func NewFields(fields ...string) *Fields {
	return &Fields{fields: fields}
}

func (f *Fields) EncodeValues(key string, v *url.Values) error {
	if len(f.fields) > 0 {
		v.Add(key, strings.Join(f.fields, ","))
	}
	return nil
}

// -----------------------------------------------------------------------------
// Enum is a type for enum values.
// Enum{Value1, Value2, Value3} => value1|value2|value3
// -----------------------------------------------------------------------------

type Enum[T any] []T

var _ query.Encoder = (*Enum[string])(nil)

func NewEnum[T any](values ...T) *Enum[T] {
	return (*Enum[T])(&values)
}

func (e Enum[T]) EncodeValues(key string, v *url.Values) error {
	if len(e) > 0 {
		var values []string
		for _, value := range e {
			values = append(values, fmt.Sprint(value))
		}
		v.Add(key, strings.Join(values, "|"))
	}
	return nil
}

// -----------------------------------------------------------------------------
// EntityType is a type for entity types.
// -----------------------------------------------------------------------------

type EntityType string

const (
	EntityTypeTask  EntityType = "task"
	EntityTypeStory EntityType = "story"
	EntityTypeBug   EntityType = "bug"
)

// -----------------------------------------------------------------------------
// OperateType is a type for operate types.
//
// 操作类型，默认为所有，可以填写add,delete,download,upload中的一个
// -----------------------------------------------------------------------------

type OperateType string

const (
	OperateTypeAdd      OperateType = "add"
	OperateTypeDelete   OperateType = "delete"
	OperateTypeDownload OperateType = "download"
	OperateTypeUpload   OperateType = "upload"
)

// -----------------------------------------------------------------------------
// OperateObject is a type for operate objects.
//
// 操作对象，默认为所有，可以填写attachment,board,bug,document,
// iteration,launch,member_activity_log,
// release,story,task,tcase,testplan,wiki中的一个
// -----------------------------------------------------------------------------

type OperateObject string

const (
	OperateObjectAttachment        OperateObject = "attachment"
	OperateObjectBoard             OperateObject = "board"
	OperateObjectBug               OperateObject = "bug"
	OperateObjectDocument          OperateObject = "document"
	OperateObjectIteration         OperateObject = "iteration"
	OperateObjectLaunch            OperateObject = "launch"
	OperateObjectMemberActivityLog OperateObject = "member_activity_log"
	OperateObjectRelease           OperateObject = "release"
	OperateObjectStory             OperateObject = "story"
	OperateObjectTask              OperateObject = "task"
	OperateObjectTestCase          OperateObject = "tcase"
	OperateObjectTestPlan          OperateObject = "testplan"
	OperateObjectWiki              OperateObject = "wiki"
)
