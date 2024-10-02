package identification

type Generator func() uint64

func DefaultGenerator() func() uint64 {
	var id uint64 = 1
	return func() uint64 {
		id += 2
		return id
	}
}

func AlmostDefaultGenerator() func() uint64 {
	var id uint64 = 0
	return func() uint64 {
		id += 2
		return id
	}
}
