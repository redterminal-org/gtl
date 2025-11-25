package ui

import (
	"fmt"
	"strings"

	"codeberg.org/fab/gtl/core"
)

func displayStreamGemini(data *core.TlData, limit int) error {
	e := data.RefreshFeeds()
	if e != nil {
		return fmt.Errorf("Couldn't refresh feeds")
	}

	stream := data.Stream

	var max int
	max = len(stream.Items) - 1
	if limit > 0 && limit < max {
		max = limit
	}

	output := ""

	for i := 0; i < max; i++ {
		output += "## "
		output += stream.Items[i].Published.Format(data.Config.Date_format)
		output += "\n### author: "
		output += stream.Items[i].Author
		output += "\n" + stream.Items[i].Content + "\n\n"
	}

	var authors []string
	for _, f := range data.Feeds {
		authors = append(authors, "=> "+f.Link+" "+f.DisplayName)
	}
	output += "Agregated tinylogs: \n" + strings.Join(authors, "\n")

	protector := "  "
	output = strings.ReplaceAll(output, protector + "```", "```")
	fmt.Println(output)
	return nil
}
