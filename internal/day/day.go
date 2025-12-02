package day

var Days DayMap

type Day interface {
	SolvePart1(input []byte) (string, error)
	SolvePart2(input []byte) (string, error)
}

func init() {
	Days = make(DayMap)
}

type DayMap map[int]Day

func (d DayMap) RegisterDay(dayNum int, day Day) {
	d[dayNum] = day
}

func (d DayMap) GetDay(day int) Day {
	return d[day]
}
