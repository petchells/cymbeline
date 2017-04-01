#!/bin/sh

rm -f cscope.out

find . -name '*.go' > cscope.files
cscope -b

GOTAGS=`which gotags 2>/dev/null`
CTAGS=`which ctags-exuberant 2>/dev/null`

if [[ $GOTAGS ]]
then
	$GOTAGS -R *.go > tags
else if [[ $CTAGS ]]
then
	$CTAGS --language-force=go -R *.go
fi
fi

#exec gvim --servername ALL -S VIM.vim
gvim -S
