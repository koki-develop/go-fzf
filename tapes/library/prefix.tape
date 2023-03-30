# configuration
Output ./examples/prefix/demo.gif
Set Shell "bash"
Set FontSize 32
Set Width 1200
Set Height 600

# setup
Hide
Type "mkdir ./tmp" Enter
Type "cp ./examples/prefix/main.go ./tmp/main.go" Enter
Type "cd ./tmp" Enter
Ctrl+l
Show

# ---

Type "go run ./main.go" Sleep 750ms  Enter
Sleep 2s

Tab
Sleep 750ms
Down
Sleep 750ms
Tab
Sleep 750ms
Down
Sleep 750ms
Tab
Sleep 2s

Enter

Sleep 3s

# ---

# cleanup
Hide
Type "cd ../" Enter
Type "\rm -rf ./tmp" Enter
