import os
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


for (fname, name) in files:
    new_contents = ""
    section_bodies = []
    with open(fname + ".tex", 'r') as f:
        # Split each section.
        contents = f.read()
        contents = contents.replace("\\end{document}", "")
        sections = contents.split("\\section")
        preamble = sections[0]
        section_titles = [section.splitlines()[0]
                          for section in sections[1:]]
        section_bodies = ["\n".join(section.splitlines()[1:])
                          for section in sections[1:]]

        new_contents = preamble
        for (title, i) in zip(section_titles, range(len(section_titles))):
            new_contents += "\section" + title + \
                "\n\\input{" + str(i+1).zfill(2) + ".tex}" + "\n"
        new_contents += "\n\\end{document}"
    with open(fname + "/main.tex", 'w') as f:
        f.write(new_contents)
    for (body, i) in zip(section_bodies, range(len(section_bodies))):
        with open(fname + "/" + str(i+1).zfill(2) + ".tex", 'w') as f:
            f.write(body)
