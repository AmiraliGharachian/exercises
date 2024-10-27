package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	emails := make(map[string]map[string]int)
	days := make(map[string]struct{})
	senders := make(map[string]int)

	for i := 0; i < n; i++ {
		var date, sender, subject string
		fmt.Scan(&date, &sender)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		subject = scanner.Text()

		if _, ok := emails[sender]; !ok {
			emails[sender] = make(map[string]int)
		}
		emails[sender][date]++
		days[date] = struct{}{}
		senders[sender]++
	}

	type senderInfo struct {
		sender string
		avg    float64
	}

	var spammers []senderInfo
	totalDays := len(days)

	for sender, count := range senders {
		if float64(count)/float64(totalDays) > 3 {
			spammers = append(spammers, senderInfo{sender, float64(count) / float64(totalDays)})
		}
	}

	sort.Slice(spammers, func(i, j int) bool {
		return spammers[i].avg > spammers[j].avg
	})

	fmt.Println("The spammers:")
	for _, spammer := range spammers {
		fmt.Printf("%s %.2f\n", spammer.sender, spammer.avg)
	}
}
