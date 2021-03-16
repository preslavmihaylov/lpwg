package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	minTime := time.Unix(-2208988800, 0) // Jan 1, 1900
	maxTime := minTime.Add(1<<63 - 1)
	minDate, maxDate := maxTime, minTime

	avgEmailsPerPerson := map[string]float64{}
	for i := 0; i < n; i++ {
		line, err = r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		tokens := strings.Split(strings.TrimSpace(line), " ")
		email := tokens[1]
		date, err := time.Parse("2006-01-02", tokens[0])
		if err != nil {
			panic(err)
		}

		avgEmailsPerPerson[email]++
		if date.Before(minDate) {
			minDate = date
		}

		if date.After(maxDate) {
			maxDate = date
		}
	}

	totalDays := maxDate.Add(24*time.Hour).Sub(minDate).Hours() / 24
	for email := range avgEmailsPerPerson {
		avgEmailsPerPerson[email] /= float64(totalDays)
	}

	emails := []string{}
	for email := range avgEmailsPerPerson {
		emails = append(emails, email)
	}

	sort.Slice(emails, func(i, j int) bool {
		return avgEmailsPerPerson[emails[i]] > avgEmailsPerPerson[emails[j]]
	})

	spammers := []string{}
	for _, email := range emails {
		fmt.Printf("%s -> %.2f\n", email, avgEmailsPerPerson[email])
		if avgEmailsPerPerson[email] > 3 {
			spammers = append(spammers, email)
		}
	}

	fmt.Println()
	fmt.Println("The spammers are:")
	for _, spammer := range spammers {
		fmt.Println(spammer)
	}
}
