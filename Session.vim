let SessionLoad = 1
if &cp | set nocp | endif
let s:so_save = &so | let s:siso_save = &siso | set so=0 siso=0
let v:this_session=expand("<sfile>:p")
silent only
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
set shortmess=aoO
badd +1 ~/Projects/golang/src/petchells/cymbeline/README.md
badd +1 ~/Projects/golang/src/petchells/cymbeline/.gitignore
badd +4 ~/Projects/golang/src/petchells/cymbeline/cscope.files
badd +9 ~/Projects/golang/src/petchells/cymbeline/cymbeline.go
badd +54 ~/Projects/golang/src/petchells/cymbeline/board.go
badd +47 ~/Projects/golang/src/petchells/cymbeline/mover.go
badd +1 ~/Projects/golang/src/petchells/cymbeline/doc.go
badd +8 ~/Projects/golang/src/petchells/cymbeline/tags
badd +8 ~/Projects/golang/src/petchells/cymbeline/VIM.sh
argglobal
silent! argdel *
set lines=52 columns=128
edit ~/Projects/golang/src/petchells/cymbeline/board.go
set splitbelow splitright
set nosplitbelow
set nosplitright
wincmd t
set winminheight=1 winheight=1 winminwidth=1 winwidth=1
argglobal
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let s:l = 26 - ((10 * winheight(0) + 25) / 50)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
26
normal! 013|
if exists('s:wipebuf')
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20 shortmess=filnxtToO
set winminheight=1 winminwidth=1
let s:sx = expand("<sfile>:p:r")."x.vim"
if file_readable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &so = s:so_save | let &siso = s:siso_save
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
