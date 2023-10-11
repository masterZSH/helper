package helper

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
)

func Copy(src, dst interface{}) {
	copier.CopyWithOption(src, dst, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return s.Format("2006-01-02 15:04:05"), nil
				},
			},
			{
				SrcType: copier.String,
				DstType: time.Time{},
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(string)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
				},
			},
		},
	})
}
