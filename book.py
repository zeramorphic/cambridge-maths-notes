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
