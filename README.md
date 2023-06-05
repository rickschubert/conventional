Conventional Committer
----------------------

Commit your current progress using conventional commits with one command:

Type `fix I have fixed this` while you are on branch `me/ABC-123/important-fix` to commit all your currently changed files with the message `fix(ABC-123): I have fixed this`.

## Example

* Given you are on branch `me/ABC-123/my-change`
* And the ticket number in the branches is surrounded by slashes
* When you execute the script with

```
go run main.go [changeType] rest of the message with spaces, even commas
```

* Then we will run `git add -A` to add all files

* And we will create a commit with message

```
changeType(ABC-123): rest of the message with spaces, even commas
```

## One-line usage from shell

If you want to be able to fire the script from anywhere in your terminal without invoking `go run` like so: `feat I just built a new feature`, then make sure to set an alias in your shell configuration:

Set the following in your `~/.zshrc`:

```sh
# Custom conventional commiter
alias fix="go run /path/to/conventional-committer/main.go fix"
alias feat="go run /path/to/conventional-committer/main.go feat"
alias chore="go run /path/to/conventional-committer/main.go chore"
alias docs="go run /path/to/conventional-committer/main.go docs"
alias tests="go run /path/to/conventional-committer/main.go test"
```

Now you can for example type `feat I have created a new feature` in your terminal and it will automatically invoke the script.

Note: We have to use `tests` instead of `test` because `test` is a shell built-in command.
