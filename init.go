package chardet

import (
	"sort"
	"strings"
)

func Detect(data []byte) ResultDict {
	detector := NewUniversalDetector(ALL, false)
	detector.Feed(data)
	return detector.Close()
}

func DetectAll(data []byte, ignoreThreshold bool, shouldRenameLegacy bool) []ResultDict {
	detector := NewUniversalDetector(ALL, shouldRenameLegacy)
	detector.Feed(data)
	detector.Close()

	results := make([]ResultDict, 0)
	if detector.InputState() == HIGH_BYTE {
		for _, prober := range detector.allProbers {
			if ignoreThreshold || prober.getConfidence() > MINIMUM_THRESHOLD {
				charName := prober.charName()
				lowerCharName := strings.ToLower(charName)

				if strings.HasPrefix(lowerCharName, "iso-8859") {
					nameTmp, ok := ISO_WIN_MAP[lowerCharName]
					if ok {
						charName = nameTmp
					}
				}

				if shouldRenameLegacy {
					nameTmp, ok := LEGACY_MAP[strings.ToLower(charName)]
					if ok {
						charName = nameTmp
					}
				}

				results = append(results, ResultDict{
					encoding:   charName,
					confidence: prober.getConfidence(),
					language:   prober.language(),
				})
			}
		}
		if len(results) > 0 {
			sort.SliceStable(results, func(i, j int) bool {
				return results[i].confidence > results[j].confidence
			})
			return results
		}
	}
	results = append(results, detector.result)
	return results
}

func DetectBest(data []byte) ResultDict {
	results := DetectAll(data, true, false)
	if len(results) > 0 {
		return results[0]
	}
	return ResultDict{}
}
