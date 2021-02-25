import os

# If you enable this flag, connectives like "hence" and "therefore" will be swapped around randomly.
# Sometimes, archaic connectives will be used, such as "thence" and "whence".
# You can change the proportion of each connective by changing the multiplicity of each connective in the list.
# Please note that there is no real reason to enable this, other than just for fun.
muddle_connectives = False

# The index is a list of keywords, sorted by keyword length.
index = {}
with open("keywords.txt", "r") as i:
    for line in i:
        alternatives = line.split("|")
        word = alternatives[0].strip()
        for alt in alternatives:
            alt = alt.strip()
            if len(alt) not in index:
                index[len(alt)] = {}
            index_entry = index[len(alt)]
            index_entry[alt.lower()] = word


def gen_index(line: str) -> str:
    # Loop through each character in the line. If that character is the start of a keyword, and preceded by whitespace, add in the keyword index entry.
    i = 0
    while i < len(line):
        if i == 0 or line[i-1] in " \t[](){}`":
            # This is preceded by whitespace.
            # Loop through all possible keyword lengths.
            for length in index:
                substring = line[i:i+length]
                index_entry = index[length].get(substring.lower())
                if index_entry is not None:
                    # Append the substring with the index entry string.
                    to_insert = '\\index{' + index_entry + '}'
                    line = line[:i+length] + to_insert + line[i+length:]
                    i += length + len(to_insert)
        i += 1

    if muddle_connectives:
        import random
        import re

        connectives = ["hence", "therefore", "whence", "thence"]
        random.shuffle(connectives)
        for i in range(1, len(connectives)):
            # Just for fun, convert some connectives into random other connectives.
            line = re.sub(
                "\\b" + connectives[i] + "\\b", connectives[i-1], line)
            line = re.sub("\\b" + connectives[i][0].upper() + connectives[i]
                          [1:] + "\\b", connectives[i-1][0].upper() + connectives[i-1][1:], line)
    return line


os.makedirs("book", exist_ok=True)
with open("book/book.tex", "w") as o:
    o.write(r"""\documentclass{book}

\usepackage{imakeidx}
\input{../util.tex}
\makeindex

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
        print("Processing " + title)
        o.write(f"\\chapter{{{title}}}")
        with open(fname + ".tex", "r") as i:
            # Trim out the \begin{document} and stuff from the file.
            for line in i:
                # Check if this line should be deleted.
                to_delete = [r"\begin{document}", r"\end{document}",
                             r"\documentclass", r"\title", r"\author", r"\maketitle", r"\tableofcontents", r"\newpage", r"\input{../util.tex}"]
                if not any([line.startswith(s) for s in to_delete]):
                    o.write(gen_index(line))
    o.write(r"\newpage\printindex")
    o.write("\n")
    o.write(r"\end{document}")
    o.write("\n")
