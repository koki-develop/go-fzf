# configuration
Output ./docs/cli/basic.gif
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
Type "mkdir subdir" Enter
Type "touch hello.go world.ts subdir/foo.ts subdir/bar.js" Enter
Ctrl+l
Show

# ---

Type "gofzf" Sleep 2s  Enter
Sleep 2s

Down 2
Sleep 1s

Up 2
Sleep 1s

Type "foo"
Sleep 750ms

Enter

Sleep 3s

# ---

# cleanup
Hide
Type "cd ../" Enter
Type "\rm -rf ./tmp" Enter
