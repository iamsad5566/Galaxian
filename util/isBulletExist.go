package util

func IsBulletExist(screenY, bulletY float64) bool {
	if screenY >= bulletY {
		return true
	}
	return false
}
