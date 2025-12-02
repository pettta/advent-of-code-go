package day

type Day interface {
	SolvePart1(input []byte) (string, error)
	SolvePart2(input []byte) (string, error)
}

type YearDayMap map[int]map[int]Day
var Days = make(YearDayMap)


func (d YearDayMap) RegisterDay(year, day int, impl Day) {
	if year < 1 || day < 1 || impl == nil {
		return
	}
	if d[year] == nil {
		d[year] = make(map[int]Day)
	}
	d[year][day] = impl
}

func (d YearDayMap) GetDay(year, day int) Day {
	if year < 1 || day < 1 {
		return nil
	}
	if yearMap := d[year]; yearMap != nil {
		return yearMap[day]
	}
	return nil
}
