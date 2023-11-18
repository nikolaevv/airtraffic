package converter

import (
	"github.com/jinzhu/copier"
	"time"
)

const (
	defaultDateFormat = "2006-01-02 15:04:05"
)

var TimeConverter = copier.TypeConverter{
	SrcType: time.Time{},
	DstType: "",
	Fn: func(src interface{}) (dst interface{}, err error) {
		return src.(time.Time).Format(defaultDateFormat), nil
	},
}

var DefaultConverterOptions = copier.Option{
	IgnoreEmpty: true,
	DeepCopy:    true,
	Converters: []copier.TypeConverter{
		TimeConverter,
	},
}
