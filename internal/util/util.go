package util

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// SliceEqual compares two slices and checks whether they have equal content
func SliceEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// SliceSubtract subtracts the content of slice a2 from slice a1
func SliceSubtract(a1, a2 []string) []string {
	a := []string{}

	for _, e1 := range a1 {
		found := false

		for _, e2 := range a2 {
			if e1 == e2 {
				found = true
				break
			}
		}

		if !found {
			a = append(a, e1)
		}
	}

	return a
}

// StringMapSubtract subtracts the content of structmap m2 from structmap m1
func StringMapSubtract(m1, m2 map[string]string) map[string]string {
	m := map[string]string{}

	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok {
			if v2 != v1 {
				m[k1] = v1
			}
		} else {
			m[k1] = v1
		}
	}

	return m
}

// StructMapSubtract subtracts the content of structmap m2 from structmap m1
func StructMapSubtract(m1, m2 map[string]struct{}) map[string]struct{} {
	m := map[string]struct{}{}

	for k1, v1 := range m1 {
		if _, ok := m2[k1]; !ok {
			m[k1] = v1
		}
	}

	return m
}

// ComputeMaxMemoryPerContainerInByte computes the max memory in byte from the given arg
func ComputeMaxMemoryPerContainerInByte(paramLimit string) (int64, error) {
	paramLimit = strings.ToUpper(paramLimit)
	if strings.HasSuffix(paramLimit, "G") || strings.HasSuffix(paramLimit, "M") || strings.HasSuffix(paramLimit, "K") {
		memUnit := paramLimit[len(paramLimit)-1:]
		memValue, err := strconv.ParseInt(paramLimit[:len(paramLimit)-1], 0, 64)
		if err != nil {
			log.Errorf("Error while extracting the memory. Root cause:=%s", err)
			return memValue, err
		}
		log.Infof("The configured max memory limit is =%s, value without unit=%d", paramLimit, memValue)
		if strings.EqualFold(memUnit, "G") {
			memValue = memValue * (1024 * 1024 * 1024)
		} else if strings.EqualFold(memUnit, "M") {
			memValue = memValue * (1024 * 1024)
		} else if strings.EqualFold(memUnit, "K") {
			memValue = memValue * 1024
		}
		return memValue, nil
	}
	value, error := strconv.ParseInt(paramLimit, 0, 64)
	if error != nil {
		log.Errorf("Error while extracting the memory. Root cause:=%s", error)
	}
	return value, nil
}
