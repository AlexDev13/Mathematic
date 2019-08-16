package code

import ()

func FindUniq(arr []float32) float32 {
	keys := make(map[float32]int)
	for i := 0; i < len(arr); i++ {
		keys[arr[i]]++
	}
	for key, value := range keys {
		if value == 1 {
			arr[0] = key
		}
	}
	return arr[0]
}
