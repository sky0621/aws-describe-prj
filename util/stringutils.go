package util

func ToString(f *string) string {
	if f == nil {
		return ""
	}
	return *f
}

func ToInt64(f *int64) int64 {
	if f == nil {
		return 0
	}
	return *f
}
