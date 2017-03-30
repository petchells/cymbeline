#!/bin/sh

rm -f cscope.out

find . -name '*.go' > cscope.files
cscope -b

# gotags -R *.go > tags
ctags-exuberant --language-force=go -R *.go

#exec gvim --servername ALL -S VIM.vim
gvim -S
