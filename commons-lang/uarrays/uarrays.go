package uarrays

func ContainsByte(slice []byte, item byte) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
