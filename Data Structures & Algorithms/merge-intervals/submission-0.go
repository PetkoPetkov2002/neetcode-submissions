type Event struct{
	Time int
	Increment int
}
func merge(intervals [][]int) [][]int {
	events := []Event{}
	counter := 0
	results := [][]int{}
	merging := true
	for _, interval := range intervals{
		startEvent := Event{Time: interval[0], Increment: 1}
		endEvent := Event{Time: interval[1], Increment: -1}
		events = append(events, startEvent)
		events = append(events, endEvent)
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].Time == events[j].Time {
			return events[i].Increment > events[j].Increment
		}
		return events[i].Time < events[j].Time
	})
	currInterval := events[0].Time

	for _, event := range events{
		counter += event.Increment
		if merging == false && counter == 1{
			currInterval = event.Time
			merging = true
		}
		if counter==0 && merging == true{
			mergedInterval := []int{currInterval,event.Time}
			results = append(results, mergedInterval)
			merging = false
		}

	}
	return results

}
