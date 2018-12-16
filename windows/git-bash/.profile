GOROOT=/c/Go
GOPATH=/d/Projects/Rebel-L/_goworkspace/

alias viprof="vi /c/Users/lgaub/.profile"
alias scprof="source /c/Users/lgaub/.profile"

alias lsgo="ls -lah $GOPATH/bin"
alias lsgop="ls -lah $GOPATH/pkg"

function la(){
	ls -lah $1
}

function cdproj(){
	cd /d/Projects/Rebel-L/$1
}

function cdgop(){
	cd $GOPATH/src/github.com/rebel-l/$1
}

function gitRebaseFromBranch(){
	BRANCH="master"
	if [ $# -gt 0 ]
	then
		BRANCH=$1
	fi

	git checkout $BRANCH && git pull && git checkout - && git rebase $BRANCH
}

