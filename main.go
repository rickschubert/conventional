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
	branchName, err := cmd.Output()
	if err != nil {
		fmt.Println(fmt.Errorf("an error occured trying to retrieve the branch name: %s", err).Error())
		os.Exit(1)
	}

	ticketIdWithSeparators := ticketRegex.FindString(string(branchName))
	ticketId := strings.ReplaceAll(ticketIdWithSeparators, "/", "")

	changeType := os.Args[1]
	messageContent := os.Args[2:]

	var messageBuilder strings.Builder
	for _, piece := range messageContent {
		messageBuilder.WriteString(fmt.Sprintf("%s ", piece))
	}
	message := strings.TrimRight(messageBuilder.String(), " ")

	scope := fmt.Sprintf("(%s)", ticketId)

	var commitLine string
	// Only add the scope section if it's filled out
	if scope == "()" {
		commitLine = fmt.Sprintf("%s: %s", changeType, message)
	} else {
		commitLine = fmt.Sprintf("%s%s: %s", changeType, scope, message)
	}

	addAllFilesCmd := exec.Command("git", "add", "-A")
	_, err = addAllFilesCmd.Output()
	if err != nil {
		fmt.Println(fmt.Errorf("an error occured adding all files: %s", err).Error())
		os.Exit(1)
	}

	commitCmd := exec.Command("git", "commit", "-m", fmt.Sprintf(`"%s"`, commitLine))
	_, err = commitCmd.Output()
	if err != nil {
		fmt.Println(fmt.Errorf("an error occured commiting with message, see here: %s", err).Error())
		os.Exit(1)
	}

	fmt.Println(commitLine)
}
