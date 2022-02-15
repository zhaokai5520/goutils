package global

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type BaseModel struct {
	ID         int        `json:"id" gorm:"primarykey"` // 主键ID
	CreateTime TstampTime `json:"createTime" gorm:"column:create_time;autoCreateTime"`
	UpdateTime TstampTime `json:"updateTime" gorm:"column:update_time;autoUpdateTime"`
}

type TstampTime int64

var timeFormat = "2006-01-02 15:04:05"

func (t *TstampTime) UnmarshalJSON(data []byte) (err error) {
	time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	return
}

func (t TstampTime) MarshalJSON() ([]byte, error) {
	tint := int64(t)
	if tint == 0 {
		return []byte(`""`), nil
	}
	tf := time.Unix(tint, 0).Format(`"` + timeFormat + `"`)
	return []byte(tf), nil
}

func (t TstampTime) Value() (driver.Value, error) {
	var tm int64
	if t == 0 {
		tm = time.Now().Unix()
	} else {
		tm = int64(t)
	}
	return tm, nil
}

func (t *TstampTime) Scan(v interface{}) (err error) {
	var value int
	switch v.(type) {
	case []uint8:
		value, err = strconv.Atoi(string(v.([]uint8)))
		if err != nil {
			return fmt.Errorf("can not convert %v to timestamp", v)
		}
		*t = TstampTime(value)
	case int64:
		value, ok := v.(int64)
		if ok {
			*t = TstampTime(value)
			return nil
		}
		return fmt.Errorf("can not convert %v to timestamp", v)
	}
	return nil
}

// func (t *TstampTime) Scan(v interface{}) (err error) {
// 	var value int
// 	switch v.(type) {
// 	case []uint8:
// 		value, err = strconv.Atoi(string(v.([]uint8)))
// 	case int64:
// 		value, ok := v.(int64)
// 		if ok {
// 			*t = TstampTime(value)
// 			return nil
// 		}
// 	}
// 	if err != nil {
// 		return fmt.Errorf("can not convert %v to timestamp", v)
// 	}
// 	*t = TstampTime(value)
// 	return nil
// }
