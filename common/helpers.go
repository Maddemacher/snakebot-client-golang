package common

func (s Positions) Contains(e int) (bool, int) {
	for index, a := range s {
		if a == e {
			return true, index
		}
	}
	return false, -1
}

func (s Ids) Contains(e Id) (bool, int) {
	for index, a := range s {
		if a == e {
			return true, index
		}
	}
	return false, -1
}
