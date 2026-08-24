package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

// sectionItemLimit caps how many items each "Latest" README section shows.
const sectionItemLimit = 5

type repo struct {
	Name  string
	URL   string
	Stars int
}

type item struct {
	Number    int
	Title     string
	URL       string
	State     string
	CreatedAt time.Time
	RepoName  string
	RepoURL   string
}

func renderReadme(now time.Time, repos []repo, pullRequests, issues, discussions []item) string {
	var b strings.Builder
	fmt.Fprintf(
		&b,
		"The information below is updated daily. Last updated at **%s UTC**\n",
		now.UTC().Format("2006/01/02 15:04"),
	)

	b.WriteString("\n## My most famous *repositories*\n\n")
	for _, r := range repos {
		fmt.Fprintf(&b, "- [%s](%s) (%d stars)\n", r.Name, r.URL, r.Stars)
	}

	writeSection(&b, "Latest *pull requests* I created", "docs/pull-requests.md", pullRequests)
	writeSection(&b, "Latest *issues* I participated in", "docs/issues.md", issues)
	writeSection(&b, "Latest *discussions* I participated in", "docs/discussions.md", discussions)

	return b.String()
}

func writeSection(b *strings.Builder, heading, seeAllPath string, items []item) {
	fmt.Fprintf(b, "\n## %s [[see all](%s)]\n\n", heading, seeAllPath)
	for _, it := range items[:min(len(items), sectionItemLimit)] {
		b.WriteString(itemLine(it))
		b.WriteString("\n")
	}
}

func renderSeeAll(title string, items []item) string {
	sorted := slices.Clone(items)
	slices.SortFunc(sorted, func(a, b item) int {
		return b.CreatedAt.Compare(a.CreatedAt)
	})

	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n", title)
	year := 0
	for _, it := range sorted {
		if y := it.CreatedAt.UTC().Year(); y != year {
			year = y
			fmt.Fprintf(&b, "\n## %d\n\n", year)
		}
		b.WriteString(itemLine(it))
		b.WriteString("\n")
	}
	return b.String()
}

func itemLine(it item) string {
	line := fmt.Sprintf(
		"- [%s](%s) • [`#%d` %s](%s)",
		it.RepoName,
		it.RepoURL,
		it.Number,
		escapeText(it.Title),
		it.URL,
	)
	if it.State != "" {
		line += fmt.Sprintf(" (%s)", it.State)
	}
	return line
}

var textEscaper = strings.NewReplacer(
	`\`, `\\`,
	"[", `\[`,
	"]", `\]`,
	"`", "\\`",
	"*", `\*`,
	"_", `\_`,
	"<", `\<`,
	">", `\>`,
	"~", `\~`,
)

// escapeText renders titles literally so markdown or HTML in them cannot
// break the link syntax or the page formatting.
func escapeText(s string) string {
	return textEscaper.Replace(s)
}
