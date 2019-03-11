package commandbus

type sortByPriority []middleware

func (slice sortByPriority) Len() int {
	return len(slice)
}

func (slice sortByPriority) Less(i, j int) bool {
	return slice[i].priority > slice[j].priority
}

func (slice sortByPriority) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
