package routes

import (
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sanbad36/url-shortner/api/database"
)

func GetTopDomains(c *gin.Context) {
	keys, err := database.Keys("*")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from the database"})
		return
	}

	domainCounts := make(map[string]int)
	for _, key := range keys {
		url, err := database.Get(key)
		if err != nil {
			continue
		}

		domain := extractDomain(url)
		domainCounts[domain]++
	}

	topDomains := make([]DomainCount, 0, len(domainCounts))
	for domain, count := range domainCounts {
		topDomains = append(topDomains, DomainCount{Domain: domain, Count: count})
	}

	sort.Slice(topDomains, func(i, j int) bool {
		return topDomains[i].Count > topDomains[j].Count
	})

	top3 := make(map[string]int)
	for i, domainCount := range topDomains {
		if i >= 3 {
			break
		}
		top3[domainCount.Domain] = domainCount.Count
	}

	c.JSON(http.StatusOK, top3)
}

type DomainCount struct {
	Domain string
	Count  int
}

func extractDomain(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) >= 3 {
		return parts[2]
	}
	return ""
}
