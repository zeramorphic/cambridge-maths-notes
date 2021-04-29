# unnumberedtotoc

For standard classes, this LaTeX-package provides the commands
`\addpart`, `\addchap` and `\addsec`. Those typeset unnumbered
sectioning levels *and* include them in the table of contents and
set up the running headers correctly. Starred versions are
available as well, which work like the usual starred sectioning
commands for the standard classes but clear the header as well. 


Classes of the KOMA-script bundle provide those commands by
default. This package is pretty useless with a KOMA class.

## Downloading this package 

1. Clone or download the repository by clicking on the big green
   button in the top right.

2. Unzip if necessary and move the three files to your project
   repository. 

3. If you need the functionality of this package more than once,
   have a look at 
   [Where do I place my own `.sty` files, to make them available to all my `.tex` files?](http://tex.stackexchange.com/q/1137)
   to make it available
   to all your projects.

4. Once the the files are properly set up, you can use the
   package like any other LaTeX package.

## Available options

- silent   
Suppress the warning that is printed in the log file and the
terminal output

- indentunnumbered    
Indent the unnumbered entries in the table of contents

## Usage and example file 

```source=LaTeX
\documentclass{book}
\usepackage[
% silent,% suppress the warning
% indentunnumbered% need unnumbered chapters to be indented in the toc?? 
]{unnumberedtotoc}
\usepackage{blindtext}
\begin{document}
\tableofcontents
\addchap{Wombat}
\addsec{Capybara}
\blindtext[10]
\chapter{Numbered chapter}
\addsec{unnumbered}
\section{numbered}
\end{document}
```


## Issues

If you have a general *how to use* question, please ask on
LaTeX-community.org

If you discovered a bug, please open a github issue. Don't forget
to provide a minimal working example to reproduce the problem. 
