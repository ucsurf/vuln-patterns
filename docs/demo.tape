# Reproducible demo. Run from the repository root:
#   vhs docs/demo.tape
# Requires semgrep on PATH (pip install semgrep==1.183.0) and vhs.
Output docs/demo.gif
Set Shell bash
Set FontSize 15
Set Width 1200
Set Height 560
Set Theme "Dracula"
Hide
Type "export PS1='$ ' && clear" Enter
Show
Type "cat docs/demo/app.py" Enter
Sleep 1s
Type "semgrep --config rules/ docs/demo/app.py --metrics=off --quiet" Sleep 300ms Enter
Sleep 20s
