package utils

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func UInt64ToString(n uint64) string {
	return strconv.FormatUint(uint64(n), 10)
}

func Decimal(value float64) float64 {
	return math.Round(value*1000000) / 1000000
}
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func StringToUint64(s string) uint64 {
	ret, _ := strconv.ParseUint(s, 10, 64)
	return ret
}

func StringToInt64(s string) int64 {
	ret, _ := strconv.ParseInt(s, 10, 64)
	return ret
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 6, 64)
}

func StringToFloat(s string) float64 {
	ret, _ := strconv.ParseFloat(s, 64)
	return ret
}

func ValuesPlaceHolder(start int, count int) string {
	var ret bytes.Buffer
	ret.WriteByte('(')
	for i := 1; i <= count; i++ {
		ret.WriteByte('$')
		ret.WriteString(Int64ToString(int64(start + i)))
		if i != count {
			ret.WriteByte(',')
		}
	}
	ret.WriteByte(')')
	return ret.String()
}

func ValuePlaceHolderAndGeometry(start int, count int, geometryPos ...int) string {
	var ret bytes.Buffer
	ret.WriteByte('(')
	for i := 1; i <= count; i++ {
		isGeometry := false
		for _, pos := range geometryPos {
			if i == pos {
				isGeometry = true
				break
			}
		}
		if isGeometry {
			ret.WriteString(fmt.Sprintf(`ST_SetSRID(st_geomfromgeojson($%s),4326)`, Int64ToString(int64(start+i))))
		} else {
			ret.WriteByte('$')
			ret.WriteString(Int64ToString(int64(start + i)))
		}
		if i != count {
			ret.WriteByte(',')
		}
	}
	ret.WriteByte(')')
	return ret.String()
}

func CondSql(first bool) string {
	if first {
		return " where"
	}
	return " and"
}

func SqlStringIn(ids ...string) string {
	var buffer bytes.Buffer
	for _, v := range ids {
		buffer.WriteString("'")
		buffer.WriteString(v)
		buffer.WriteString("'")
		buffer.WriteString(",")
	}
	temp := buffer.String()
	size := len(temp)
	if size > 0 {
		return temp[:size-1]
	}
	return `''`
}

func SqlIntegerIn(ids ...int) string {
	var buffer bytes.Buffer
	for _, v := range ids {
		buffer.WriteString("'")
		buffer.WriteString(Int64ToString(int64(v)))
		buffer.WriteString("'")
		buffer.WriteString(",")
	}
	temp := buffer.String()
	size := len(temp)
	if size > 0 {
		return temp[:size-1]
	}
	return `''`
}

func SqlIntegerArray(ids ...int) string {
	var buffer bytes.Buffer
	buffer.WriteString("array[")
	for i, v := range ids {
		//buffer.WriteString("'")
		buffer.WriteString(Int64ToString(int64(v)))
		//buffer.WriteString("'")
		if i < len(ids)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]::integer[]")
	return buffer.String()
}

func SqlStringArray(ids ...string) string {
	var buffer bytes.Buffer
	buffer.WriteString("array[")
	for i, v := range ids {
		buffer.WriteString("'")
		buffer.WriteString(v)
		buffer.WriteString("'")
		if i < len(ids)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]::varchar[]")
	return buffer.String()
}

func SqlWithIn(lng, lat float64, radius int) string {
	return fmt.Sprintf(`ST_DWithin(ST_GeomFromText('POINT(%f %f)', 4326)::geography, geometry::geography, %d)`, lng, lat, radius)
}

func ToInt64(v interface{}) int64 {
	if ret, ok := v.(int64); ok {
		return ret
	}
	return 0
}

func ToFloat64(v interface{}) float64 {
	if ret, ok := v.(float64); ok {
		return ret
	}
	return 0
}

func ToString(v interface{}) string {
	if ret, ok := v.(string); ok {
		return ret
	}

	if ret, ok := v.([]byte); ok {
		return string(ret)
	}
	return ""
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	var buffer bytes.Buffer
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteByte('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteByte(byte(r))
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
// 大小大小
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func UpperFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母小写
func LowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
