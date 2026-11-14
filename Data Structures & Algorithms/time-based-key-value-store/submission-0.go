type TimeMap struct {
	mapOfEvents map[string][]*Event
}
type Event struct{
	emotion string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{mapOfEvents: make(map[string][]*Event)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.mapOfEvents[key] = append(this.mapOfEvents[key], &Event{emotion:value, timestamp: timestamp})
	
}

func (this *TimeMap) Get(key string, timestamp int) string {
	relevantEvents := this.mapOfEvents[key]

	left := 0
	right := len(relevantEvents)-1
	
	mid := left + (right-left)/2
	currentBest := ""
	for left <= right{
		if relevantEvents[mid].timestamp>timestamp{
			right = mid -1 
			mid = left+(right-left)/2
		}else if relevantEvents[mid].timestamp<=timestamp{
			currentBest = relevantEvents[mid].emotion
			left = mid + 1 
			mid = left + (right-left)/2
		}
	}
	return currentBest

	
}
