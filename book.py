import os

os.makedirs("book", exist_ok=True)
with open("book/book.tex", "w") as o:
    o.write(r"""\documentclass{book}

\input{../util.tex}

\addto\captionsUKenglish{\renewcommand{\chaptername}{Course}}

\title{Notes on the Cambridge University Mathematical Tripos}
\author{}

\begin{document}
\maketitle

\tableofcontents
\newpage

\chapter*{Introduction}
This book contains notes for the maths courses at Cambridge University. Please note that while efforts have been made to ensure completeness and correctness, no guarantees can be made; this is simply a reasonably complete way of collating information about the courses.

This book can be downloaded in PDF form for free at \url{https://thirdsgames.co.uk/gh/maths-compiled/book/book.pdf}, and the source code (for the book itself and for the individual courses) can be accessed at \url{https://github.com/thirdsgames/cambridge-maths-notes}.

You are given the right to download the PDF of the book (or its component parts) for private use. You are permitted to download and modify the source code of the repository (the book and the course notes it contains), but may not distribute these modifications (including object files such as PDFs generated from these modifications) to others. However, you are permitted to make public forks of the repository in order to create pull requests, but this does not grant you permission to distribute object files created from these forked repositories. It must be clear when visiting it that such a repository is a fork of \url{https://github.com/thirdsgames/cambridge-maths-notes}, and must include a link to the original repository. (Forks created on GitHub satisfy this requirement, as the title contains the words `forked from' and then a link to the original repository.)

""")
    files = [
        ("ia/ns", "Numbers and Sets"),
        ("ia/groups", "Groups"),
        ("ia/de", "Differential Equations"),
        ("ia/vm", "Vectors and Matrices"),
        ("ia/analysis", "Analysis I"),
        ("ia/probability", "Probability"),
        ("ia/dr", "Dynamics and Relativity"),
        ("ia/vc", "Vector Calculus")
    ]
    for (fname, title) in files:
        o.write(f"\\chapter{{{title}}}")
        with open(fname + ".tex", "r") as i:
            # Trim out the \begin{document} and stuff from the file.
            for line in i:
                # Check if this line should be deleted.
                to_delete = [r"\begin{document}", r"\end{document}",
                             r"\documentclass", r"\title", r"\author", r"\maketitle", r"\tableofcontents", r"\newpage", r"\input{../util.tex}"]
                if not any([line.startswith(s) for s in to_delete]):
                    o.write(line)
    o.write(r"\end{document}")
    o.write("\n")
