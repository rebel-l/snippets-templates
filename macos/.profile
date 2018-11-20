export GOPATH=$(go env GOPATH)
export CDPATH=.:~:$GOPATH/src:$GOPATH/src/github.com
export PATH=.:$PATH:$GOPATH/bin

export PS1="\[\033[35m\]\d \t\[\033[m\] \[\033[33m\]|\[\033[m\] \[\033[36m\]\W\[\033[m\] 🖖 > "
export CLICOLOR=1
export LSCOLORS=GxFxCxDxBxegedabagaced
export GREP_COLORS='sl=1;31'

alias viprof="vi ~/.profile"
alias scprof=". ~/.profile"

alias dockerClean="docker system prune -a --volumes"

alias showSSHCFG="cat ~/.ssh/config"

function ll(){
	ls -lah $1
}
