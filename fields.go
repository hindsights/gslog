package gslog

import "fmt"

type Fields map[string]interface{}

func CopyFields(fields Fields) Fields {
	if fields == nil || len(fields) == 0 {
		return nil
	}
	return JoinFields(fields)
}

func JoinFields(fields ...Fields) Fields {
	if len(fields) == 0 {
		return nil
	}
	ret := make(Fields)
	for _, fs := range fields {
		for k, v := range fs {
			ret[k] = v
		}
	}
	return ret
}

func FormatFields(fields ...Fields) []string {
	if len(fields) == 0 {
		return nil
	}
	ret := make([]string, 0, GetFieldCount(fields...))
	for _, fs := range fields {
		for k, v := range fs {
			ret = append(ret, fmt.Sprintf("%s=%v", k, v))
		}
	}
	return ret
}

func GetFieldCount(fields ...Fields) int {
	count := 0
	for _, fs := range fields {
		count += len(fs)
	}
	return count
}
