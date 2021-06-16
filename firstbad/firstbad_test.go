package firstbad

import (
	"reflect"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	result := firstBadVersion(5)
	if !reflect.DeepEqual(result, 4) {
		t.Errorf("firstBadVersion = %d", result)
	}
}
