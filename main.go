package main

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func solution(schedules []int, timelogs [][]int, startday int) int {
	trunc := func(a int) int {
		return (a / 100) * 100
	}

	answer := 0
	var wg sync.WaitGroup
	mx := &sync.RWMutex{}

	wg.Add(len(schedules))

	for i, v := range schedules {
		go func(idx int, v int) {
			defer wg.Done()
			h := trunc(v) / 100
			m := v - trunc(v)

			m += 10

			if m >= 60 {
				h += 1
				m -= 60
			}
			logs := timelogs[idx]

			isGet := true

			nowDay := startday

			for _, v2 := range logs {
				if nowDay == 6 || nowDay == 7 {
					nowDay += 1
					continue
				}
				logH := trunc(v2) / 100
				logM := v2 - trunc(v2)
				if h < logH || (h == logH && m < logM) {
					isGet = false
					break
				}
				nowDay += 1
			}

			if isGet {
				mx.Lock()
				answer += 1
				mx.Unlock()
			}

		}(i, v)
	}
	wg.Wait()
	log.Println(answer)
	return answer
}

func main() {
	// mail_tool.SendMail("skatn7979@gmail.com", "{error:'internal server error'}", "abc 배치 작업 오류")
	// dfs(map[int][]int{
	// 	0: {1, 3},
	// 	1: {0, 2, 4},
	// 	2: {1},
	// 	3: {0, 4},
	// 	4: {1, 3},
	// }, 0, map[int]bool{}, true)
	solution([]int{730, 855, 700, 720}, [][]int{
		{710, 700, 650, 735, 700, 931, 912}, {908, 901, 805, 815, 800, 831, 835}, {705, 701, 702, 705, 710, 710, 711}, {707, 731, 859, 913, 934, 931, 905},
	}, 1)
}

func dfs(graph map[int][]int, start int, visited map[int]bool) {
	if visited[start] {
		return
	}

	visited[start] = true
	fmt.Println("Visit:", start)

	for _, neighbor := range graph[start] {
		dfs(graph, neighbor, visited)
	}
}

func bfs(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			continue
		}

		visited[node] = true
		fmt.Println("Visit:", node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}
