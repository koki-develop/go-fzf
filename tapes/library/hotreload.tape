# configuration
Output ./examples/hotreload/demo.gif
Set Shell "bash"
Set FontSize 32
Set Width 1200
Set Height 1000

# setup
Hide
Type "mkdir ./tmp" Enter
Type "cp ./examples/hotreload/main.go ./tmp/main.go" Enter
Type "cd ./tmp" Enter
Ctrl+l
Show

# ---

Type "go run ./main.go" Sleep 750ms  Enter
Sleep 3s

Type "5"
Sleep 3s

Enter

Sleep 3s

# ---

# cleanup
Hide
Type "cd ../" Enter
Type "\rm -rf ./tmp" Enter
