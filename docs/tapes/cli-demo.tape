# configuration
Output ./docs/cli/demo.gif
Set Shell "bash"
Set FontSize 32
Set Width 1200
Set Height 600

# setup
Hide
Type "go install ./cmd/gofzf" Enter
Type "gofzf --help" Enter
Type "mkdir ./tmp" Enter
Type "cd ./tmp" Enter
Type "touch hello world foo bar" Enter
Ctrl+l
Show

# ---

Type "ls | gofzf --no-limit" Sleep 2s  Enter
Sleep 2s
Down 2
Sleep 1s
Up 2
Sleep 1s

Type "hello"
Sleep 750ms
Tab
Sleep 750ms
Backspace 5
Sleep 750ms

Type "world"
Sleep 750ms
Tab
Sleep 750ms
Backspace 5
Sleep 750ms

Enter

Sleep 3s

# cleanup
Hide
Type "cd ../" Enter
Type "\rm -rf ./tmp" Enter
