# configuration
Output ./examples/preview-window/demo.gif
Set Shell "bash"
Set FontSize 32
Set Width 1400
Set Height 600

# setup
Hide
Type "mkdir ./tmp" Enter
Type "cp ./examples/preview-window/main.go ./tmp/main.go" Enter
Type "cd ./tmp" Enter
Type "echo hello > hello.go" Enter
Type "echo world > world.ts" Enter
Type "echo foo > foo.ts" Enter
Type "echo bar > bar.js" Enter
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

# cleanup
Hide
Type "cd ../" Enter
Type "\rm -rf ./tmp" Enter
