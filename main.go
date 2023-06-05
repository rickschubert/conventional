package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var ticketRegex = regexp.MustCompile(`\/[A-Z]+-[0-9]+\/`)

func main() {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	stdout, _ := cmd.Output()
	branchName := string(stdout)
	ticketNumberWithSlashes := ticketRegex.FindString(branchName)
	withoutSlashes := strings.ReplaceAll(ticketNumberWithSlashes, "/", "")

	changeType := os.Args[1]
	messageContent := os.Args[2:]

	var plainMessage strings.Builder
	for _, piece := range messageContent {
		plainMessage.WriteString(fmt.Sprintf("%s ", piece))
	}

	ticketSection := fmt.Sprintf("(%s)", withoutSlashes)

	var commitLine string
	// Check if the ticket section is empty; don't append it in that case
	if ticketSection == "()" {
		commitLine = fmt.Sprintf("%s: %s", changeType, strings.TrimRight(plainMessage.String(), " "))
	} else {
		commitLine = fmt.Sprintf("%s%s: %s", changeType, ticketSection, strings.TrimRight(plainMessage.String(), " "))
	}

	addAllFilesCmd := exec.Command("git", "add", "-A")
	_, err := addAllFilesCmd.Output()
	if err != nil {
		fmt.Println(fmt.Errorf("an error occured adding all files: %s", err).Error())
		os.Exit(122)
	}
	commitCmd := exec.Command("git", "commit", "-m", fmt.Sprintf(`"%s"`, commitLine))
	_, err = commitCmd.Output()
	if err != nil {
		fmt.Println(fmt.Errorf("an error occured commiting with message, see here: %s", err).Error())
		os.Exit(122)
	}

	fmt.Println(commitLine)
}
