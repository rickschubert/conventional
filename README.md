![](assets/heroimg.jpg "Fox committing conventional commits")

conventional
============

Commit your current progress using conventional commits with one command:

Type `fix I have fixed this` while you are on branch `me/ABC-123/important-fix` to commit all your currently changed files with the message `fix(ABC-123): I have fixed this`.

## Example

* Given you are on branch `me/ABC-123/my-change`
* And the ticket number in the branches is surrounded by slashes
* When you execute the tool like so:
```
conventional [changeType] rest of the message with spaces, even commas
```
* Then we will run `git add -A` to add all files
* And we will create a commit with message
```
changeType(ABC-123): rest of the message with spaces, even commas
```

## Installation

### Download as release

- Download the right file for your operating system from the [releases page](https://github.com/rickschubert/conventional/releases)
- Invoke the script by pointing to the download:

```sh
/path/to/download/conventional feat This is my first message
```

### Install using Golang

This step assumes that you have Golang installed locally on your machine.

```sh
go install github.com/rickschubert/conventional@latest
# Execute the tool from anywhere in your terminal with conventional
conventional feat This is my first message
```

## One-line usage from shell

If you want to be able to fire the script from anywhere in your terminal without having to type `conventional` first, then make sure to set an alias in your shell configuration:

Set the following in your `~/.zshrc`:

```sh
# If you installed it with go install:
alias fix="/path/to/conventional fix"
alias feat="/path/to/conventional feat"
alias chore="/path/to/conventional chore"
alias docs="/path/to/conventional docs"
alias tests="/path/to/conventional test"

# If you downloaded a release:
alias fix="conventional fix"
alias feat="conventional feat"
alias chore="conventional chore"
alias docs="conventional docs"
alias tests="conventional test"
```

Now you can type `feat I have created a new feature` in your terminal and it will automatically run our tool!

Note: We want to alias `tests` for `test` because `test` is a shell built-in command which we don't want to override.
