package util

func ToString(f *string) string {
	if f == nil {
		return ""
	}
	return *f
}
