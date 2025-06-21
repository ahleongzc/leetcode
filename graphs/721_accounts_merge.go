package graphs

import (
	"slices"
)

func AccountsMerge(accounts [][]string) [][]string {
	parent := make(map[int]int)
	rank := make(map[int]int)

	for i := range accounts {
		parent[i] = i
		rank[i] = 0
	}

	findRep := func(i int) int {
		p := parent[i]
		for {
			if p == parent[p] {
				break
			}
			parent[p] = parent[parent[p]]
			p = parent[p]
		}
		return p
	}

	union := func(i, j int) {
		repI, repJ := findRep(i), findRep(j)
		if repI == repJ {
			return
		}

		if rank[repI] > rank[repJ] {
			parent[repJ] = repI
		} else {
			parent[repI] = repJ
			rank[repJ]++
		}
	}

	// The key is the email, the value is the index
	processedEmails := make(map[string]int)
	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			email := account[j]
			index, exist := processedEmails[email]
			if !exist {
				processedEmails[email] = i
				continue
			}
			union(i, index)
		}
	}

	accountMap := make(map[int][]int)
	for _, index := range processedEmails {
		// This finds the parent to link all the emails together under the same index
		parent := findRep(index)
		if accountMap[parent] == nil {
			accountMap[parent] = make([]int, 0)
		}
		accountMap[parent] = append(accountMap[parent], index)
	}

	result := make([][]string, 0)
	for index, arrOfSameAccount := range accountMap {
		name := accounts[index][0]
		emailMap := make(map[string]struct{})

		for _, accountIndex := range arrOfSameAccount {
			for i := 1; i < len(accounts[accountIndex]); i++ {
				email := accounts[accountIndex][i]
				emailMap[email] = struct{}{}
			}
		}

		emailArr := make([]string, 0)
		for email := range emailMap {
			emailArr = append(emailArr, email)
		}

		slices.Sort(emailArr)
		emailArr = slices.Insert(emailArr, 0, name)
		result = append(result, emailArr)
	}

	return result
}
