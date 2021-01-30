package mining

import "testing"

// belirlenen bir key bir stringin md5 hash halinde 2 veya daha fazla geçiyor mu ?
func TestCheckMining(t *testing.T) {
	results := []struct {
		Key    string
		Text   string
		Result bool
	}{
		{"afe", "TbKPEj8OebWsojKzbqny", true},
		{"afe", "V7wdvYr27h5bTpzcJJqm", true},
		{"192", "6yqjZqvjstKeump41bEt", true},
		{"192", "6yqjZqvjstKeump41bEe", false},
		{"165", "6yqjZqvjstKeump41bEt", false},
	}

	for _, v := range results {
		if CheckMining(v.Key, v.Text) != v.Result {
			t.Error("Failed Status: " + v.Key + " : " + v.Text)
			t.Error("Result " + ConvertToHash(v.Key))
		}
	}
}
