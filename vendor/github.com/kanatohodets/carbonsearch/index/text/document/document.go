package document

import "fmt"

const N = 4

func Tokenize(term string) ([]uint32, error) {
	if len(term) < N {
		return nil, fmt.Errorf("%s is too short to search on", term)
	}

	// len(term) - 3 for quadgrams
	tokens := make([]uint32, 0, len(term)-3)
	for i := 0; i <= len(term)-N; i++ {
		var chunk [N]byte
		copy(chunk[:], term[i:i+N])
		tokens = append(tokens, ngramize(chunk))
	}

	return tokens, nil
}

func ngramize(s [N]byte) uint32 {
	return uint32(s[0])<<24 | uint32(s[1])<<16 | uint32(s[2])<<8 | uint32(s[3])
}

func Validate(metrics []string) []string {
	validMetrics := make([]string, 0, len(metrics))
	for _, metric := range metrics {
		if len(metric) >= N {
			validMetrics = append(validMetrics, metric)
		}
	}
	return validMetrics
}
