
type Edge struct{
	To int
	Weight int
}

type MinHeap []Edge

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i].Weight < h[j].Weight
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(Edge))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func minimumSpanningTree(n int, edges [][]int) int {
	graph := make([][]Edge, n)
	totalWeight := 0
	visited := make(map[int]bool)
	for _, edge := range edges{
		outEdge := Edge{To: edge[1], Weight: edge[2]}
		backEdge := Edge{To: edge[0], Weight: edge[2]}

		graph[edge[0]] = append(graph[edge[0]], outEdge)
		graph[edge[1]] = append(graph[edge[1]], backEdge)
	}
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Edge{To: 0, Weight: 0})

	for len(*h)>0 && len(visited)<n{
		edge := heap.Pop(h).(Edge)
		if visited[edge.To]{
			continue
		}
		weight := edge.Weight
		visited[edge.To] = true
		totalWeight += weight
		for _, neighbour := range graph[edge.To]{
			if !visited[neighbour.To]{
				heap.Push(h, neighbour)
			}
		}
	}

	if len(visited)!=n{
		return -1
	}
	return totalWeight
	
}
