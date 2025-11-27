package orders

import (
	"fmt"
	"strings"
)

func FlattenJSON(prefix string, v interface{}, out *[]string) {
	if v == nil {
		*out = append(*out, fmt.Sprintf("[yellow]%s:[white] %s", prefix, "NOT AVAILABLE"))
		return
	}

	switch val := v.(type) {

	// Map (meta, raw JSON)
	case map[string]interface{}:
		if len(val) == 0 {
			*out = append(*out, fmt.Sprintf("[yellow]%s:[white] %s", prefix, "NOT AVAILABLE"))
			return
		}
		for k, sub := range val {
			FlattenJSON(prefix+"."+k, sub, out)
		}

	// Slice (tags)
	case []interface{}:
		if len(val) == 0 {
			*out = append(*out, fmt.Sprintf("[yellow]%s:[white] %s", prefix, "NOT AVAILABLE"))
			return
		}
		for i, sub := range val {
			FlattenJSON(fmt.Sprintf("%s[%d]", prefix, i), sub, out)
		}

	// String
	case string:
		if strings.TrimSpace(val) == "" {
			val = "NOT AVAILABLE"
		}
		*out = append(*out, fmt.Sprintf("[yellow]%s:[white] %s", prefix, val))

	// Numbers / bool
	case float64, float32, int64, int32, int, bool:
		*out = append(*out, fmt.Sprintf("[yellow]%s:[white] %v", prefix, val))

	// Fallback
	default:
		*out = append(*out, fmt.Sprintf("[yellow]%s:[white] %v", prefix, val))
	}
}
