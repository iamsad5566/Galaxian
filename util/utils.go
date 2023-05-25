package util

func MediumPosition(parentObjX float64, parentObjWidth int, childObjWidth int) float64 {
	return parentObjX + float64(parentObjWidth-childObjWidth)/2
}

func IsBulletExist(screenY, bulletY float64) bool {
	return screenY >= bulletY
}
