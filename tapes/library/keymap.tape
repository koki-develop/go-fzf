# configuration
Output ./examples/keymap/demo.gif
Set Shell "bash"
Set FontSize 32
Set Width 1200
Set Height 600

# setup
Hide
Type "mkdir ./tmp" Enter
Type "cp ./examples/keymap/main.go ./tmp/main.go" Enter
Type "cd ./tmp" Enter
Ctrl+l
Show

# ---

Type "go run ./main.go" Sleep 750ms  Enter
Sleep 2s

Ctrl+f
Ctrl+f
Sleep 1s

Ctrl+b
Ctrl+b
Sleep 1s

Enter

Sleep 3s

# ---

# cleanup
Hide
Type "cd ../" Enter
Type "\rm -rf ./tmp" Enter
