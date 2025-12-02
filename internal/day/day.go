package day

var YearDays YearDayMap

type Day interface {
	SolvePart1(input []byte) (string, error)
	SolvePart2(input []byte) (string, error)
}

func init() {
	YearDays = make(YearDayMap)
}

type YearDayMap map[int]map[int]Day // year -> day -> Day implementation

func (d YearDayMap) RegisterDay(year int, dayNum int, day Day) {
	if d[year] == nil {
		d[year] = make(map[int]Day)
	}
	d[year][dayNum] = day
}

func (d YearDayMap) GetDay(year int, day int) Day {
	return d[year][day]
}
