Conventional Committer
----------------------

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

### Install using golang

```sh
go install github.com/rickschubert/conventional
# Execute the tool from anywhere in your terminal with conventional
conventional feat This is my first message
```

### Download from Releases

- Download the right file for your operating system from the [Releases page](https://github.com/rickschubert/conventional/releases)
- Invoke the script by pointing to the download:

```sh
/path/to/download/conventional feat This is my first message
```

## One-line usage from shell

If you want to be able to fire the script from anywhere in your terminal without invoking `go run` like so: `feat I just built a new feature`, then make sure to set an alias in your shell configuration:

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

Now you can for example type `feat I have created a new feature` in your terminal and it will automatically invoke the script.

Note: We have to use `tests` instead of `test` because `test` is a shell built-in command.
