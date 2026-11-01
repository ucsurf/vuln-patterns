Output docs/demo.gif
Set Shell bash
Set FontSize 15
Set Width 1200
Set Height 560
Set Theme "Dracula"
Hide
Type "export PATH=/home/box/.local/bin:$PATH PS1='$ ' && clear" Enter
Show
Type "cat app.py" Enter
Sleep 1s
Type "semgrep --config vuln-patterns/rules/ app.py --metrics=off --quiet" Sleep 300ms Enter
Sleep 20s
