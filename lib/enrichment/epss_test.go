package enrichment

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devops-kung-fu/bomber/models"
)

func TestEnrich(t *testing.T) {
	vulnerabilities := []models.Vulnerability{
		{
			Cve: "CVE-2021-43138",
		},
		{
			Cve: "CVE-2020-15084",
		},
		{
			Cve: "CVE-2020-28282",
		},
		{
			Cve: "sonatype-2020-1214",
		},
	}
	enriched, err := Enrich(vulnerabilities)

	assert.NoError(t, err)
	assert.Len(t, enriched, 4)

	assert.Empty(t, enriched[3].Epss.Cve)
	assert.Equal(t, enriched[0].Epss.Cve, "CVE-2021-43138")

}
