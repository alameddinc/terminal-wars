package repair

import (
	"strconv"
	"testing"
)

// 10m ile 100m arasında olacak bir sayının md5 hali veriliyor, cevap olarak o sayı bekleniyor.
func TestCheckRepair(t *testing.T) {
	results := []struct {
		Number    int
		Md5Result string
		Result    bool
	}{
		{15765233, "2f0899f61805126a241830dc8f0b3d70", true},
		{64176233, "d723f8c7e136c804c32aa19236dba4d3", true},
		{32165233, "4407ddc7463403b98f25ae796bba07f8", true},
		{12, "4407ddc7463403b98f25ae796bba07f8", false},
		{0, "4407ddc7463403b98f25ae796bba07f8", false},
		{123456758, "4407ddc7463403b98f25ae796bba07f8", false},
	}
	for _, v := range results {
		if CheckRepair(v.Number, v.Md5Result) != v.Result {
			t.Error(strconv.Itoa(v.Number) + " result shouldn't be " + v.Md5Result)
		}
	}
}
