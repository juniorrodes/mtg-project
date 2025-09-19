let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/Documents/go_projects/mtg-project
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +14 components/forms/advanced_search.templ
badd +15 pkg/main.go
badd +14 ~/Documents/go_projects/mtg-project/pkg/router/router.go
badd +23 ~/Documents/go_projects/mtg-project/pkg/api/routes.go
badd +50 ~/Documents/go_projects/mtg-project/pkg/controller/card-controller.go
badd +17 ~/Documents/go_projects/mtg-project/components/cards/card-view_templ.go
badd +10 components/cards/card-view.templ
badd +3 ~/Documents/go_projects/mtg-project/components/layout/layout.templ
argglobal
%argdel
edit ~/Documents/go_projects/mtg-project/pkg/api/routes.go
argglobal
balt ~/Documents/go_projects/mtg-project/pkg/controller/card-controller.go
setlocal foldmethod=manual
setlocal foldexpr=0
setlocal foldmarker={{{,}}}
setlocal foldignore=#
setlocal foldlevel=99
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal foldenable
silent! normal! zE
sil! 3,9fold
sil! 12,23fold
let &fdl = &fdl
3
sil! normal! zc
let s:l = 23 - ((22 * winheight(0) + 21) / 42)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 23
normal! 038|
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
nohlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
