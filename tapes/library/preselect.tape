# configuration
Output ./examples/preselect/demo.gif
Set Shell "bash"
Set FontSize 32
Set Width 1200
Set Height 600

# setup
Hide
Type "mkdir ./tmp" Enter
Type "cp ./examples/preselect/main.go ./tmp/main.go" Enter
Type "cd ./tmp" Enter
Ctrl+l
Show

# ---

Type "go run ./main.go" Sleep 750ms  Enter
Sleep 2s

Down 2
Sleep 1s

Up 2
Sleep 1s

Enter

Sleep 3s

# ---

# cleanup
Hide
Type "cd ../" Enter
Type "\rm -rf ./tmp" Enter
