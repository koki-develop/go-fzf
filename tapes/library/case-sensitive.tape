# configuration
Output ./examples/case-sensitive/demo.gif
Set Shell "bash"
Set FontSize 32
Set Width 1200
Set Height 600

# setup
Hide
Type "mkdir ./tmp" Enter
Type "cp ./examples/case-sensitive/main.go ./tmp/main.go" Enter
Type "cd ./tmp" Enter
Ctrl+l
Show

# ---

Type "go run ./main.go" Sleep 750ms  Enter
Sleep 2s

Type@100ms "hello"
Sleep 750ms
Backspace 5
Sleep 750ms

Type@100ms "Hello"
Sleep 750ms
Backspace 5
Sleep 750ms

Type@100ms "HELLO"
Sleep 750ms
Backspace 5
Sleep 750ms

Enter

Sleep 3s

# ---

# cleanup
Hide
Type "cd ../" Enter
Type "\rm -rf ./tmp" Enter
