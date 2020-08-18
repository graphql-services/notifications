package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

type NotificationResultType struct {
	EntityResultType
}

type Notification struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	GroupID     *string    `json:"groupID" gorm:"column:groupID"`
	Subject     *string    `json:"subject" gorm:"column:subject;type:text"`
	Message     string     `json:"message" gorm:"column:message;type:text"`
	Seen        bool       `json:"seen" gorm:"column:seen"`
	Channel     *string    `json:"channel" gorm:"column:channel"`
	Recipient   *string    `json:"recipient" gorm:"column:recipient"`
	Principal   *string    `json:"principal" gorm:"column:principal;index:basic_fetch"`
	Reference   *string    `json:"reference" gorm:"column:reference"`
	ReferenceID *string    `json:"referenceID" gorm:"column:referenceID"`
	Date        time.Time  `json:"date" gorm:"column:date;index:basic_fetch"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`
}

func (m *Notification) Is_Entity() {}

type NotificationChanges struct {
	ID          string
	GroupID     *string
	Subject     *string
	Message     string
	Seen        bool
	Channel     *string
	Recipient   *string
	Principal   *string
	Reference   *string
	ReferenceID *string
	Date        time.Time
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
