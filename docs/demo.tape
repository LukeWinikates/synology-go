Output docs/demo.gif
Require synoctl
Set Shell "bash"
Set FontSize 18
Set Width 1000
Set Height 800
Type@50ms "synoctl docker container list" Sleep 1500ms  Enter
Sleep 5s
Type@50ms "synoctl docker container logs --name hello-world --simple" Sleep 1500ms  Enter
Sleep 12s
