# configuration
Output ./docs/library/demo.gif
Set Shell "bash"
Set FontSize 32
Set Width 1200
Set Height 600

# setup
Hide
Type "mkdir ./tmp" Enter
Type "cp ./examples/demo/main.go ./tmp/main.go" Enter
Type "cd ./tmp" Enter
Type "touch hello world foo bar" Enter
Ctrl+l
Show

# ---

Type "go run ./main.go" Sleep 750ms  Enter
Sleep 2s
Down 2
Sleep 1s
Up 2
Sleep 1s

Type "world"
Sleep 750ms

Enter

Sleep 3s

# ---
Hide
Type "cd ../" Enter
Type "\rm -rf ./tmp" Enter
