package tools

import (
	"encoding/json"
	"fmt"
)

// DataConvBasic 数据类型转换结构体
type DataConvBasic struct{}

// DataConvert 数据类型转换
var DataConvert *DataConvBasic

func init() {
	DataConvert = NewDataConv()
}

// NewDataConv 初始化 DataConvBasic
func NewDataConv() (n *DataConvBasic) {
	n = new(DataConvBasic)
	return
}

// MapToRawJSON 任意map类型转换为原始JSON数据
func (dc *DataConvBasic) MapToRawJSON(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}

// RawMaps2Maps 原始map数组转[]map[string]interface{}
func (dc *DataConvBasic) RawMaps2Maps(data []byte) (m []map[string]interface{}) {
	m = make([]map[string]interface{}, 0)
	_ = json.Unmarshal(data, &m)
	return
}

// RawMap2Map 原始map转map[string]interface{}
func (dc *DataConvBasic) RawMap2Map(data []byte) (m map[string]interface{}) {
	m = make(map[string]interface{})
	_ = json.Unmarshal(data, &m)
	return
}

// RawArray2Array 原始数组类型转[]interface{}
func (dc *DataConvBasic) RawArray2Array(data []byte) (m []interface{}) {
	m = make([]interface{}, 0)
	_ = json.Unmarshal(data, &m)
	return
}

// Print 格式打印输出
func (dc *DataConvBasic) Print(f string, i interface{}) {
	fmt.Printf(f, i)
}

// Float2uint 浮点转整型处理
func (dc *DataConvBasic) Float2uint(f float64) int64 {
	u := int64((f * 100) + 0.5)
	return u
}
