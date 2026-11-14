
type Event struct {
    Time   int
    Affect int
}

func minMeetingRooms(intervals []Interval) int {
    events := []Event{}
    var maxRooms int
    var counter int

    for _, interval := range intervals {
        start, end := interval.start, interval.end

        startEvent := Event{Time: start, Affect: 1}
        endEvent := Event{Time: end, Affect: -1}

        events = append(events, startEvent)
        events = append(events, endEvent)
    }

    sort.Slice(events, func(i, j int) bool {
        if events[i].Time == events[j].Time {
            return events[i].Affect < events[j].Affect
        }
        return events[i].Time < events[j].Time
    })

    for _, event := range events {
        counter += event.Affect
        maxRooms = max(counter, maxRooms)
    }

    return maxRooms
}
