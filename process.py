import re
import os

def collate(root: str):
    """
    Packs together every source file into a single document.
    """
    # Read main.tex.
    contents = ""
    with open(f"{root}/main.tex") as f:
        contents = f.read()

    # Replace \input with the actual file, unless it's "../../util.tex".
    while True:
        m = re.search("\\\\input\\{([a-zA-Z0-9\\.]*)\\}", contents)
        if m is None:
            break
        file_to_input = m.group(1)
        new_input = ""
        with open(f"{root}/{file_to_input}") as f:
            new_input = f.read()
        contents = contents[:m.start(0)] + new_input + contents[m.end(0):]

    i1 = contents.find("\\begin{document}")
    i2 = contents.find("\\end{document}")
    contents = contents[i1 + len("\\begin{document}"): i2]
    contents = contents.replace("\\maketitle", "")
    contents = contents.replace("\\tableofcontentsnewpage{}", "")
    contents = contents.strip()
    with open(f"{root}/collated.tex", 'w') as f:
        f.write(contents)

def uncollate(root: str):
    """
    Extracts the sections from the input file, and puts them into their own files.
    """
    contents = ""
    with open(f"{root}/collated.tex") as f:
        contents = f.read()
    # Remove old files.
    for file in list(os.walk(root))[0][2]:
        if file[0].isnumeric() and file[1].isnumeric():
            os.remove(f"{root}/{file}")
    m = list(re.finditer("\\\\section\\{([^\\}]*)\\}", contents))
    section_names = []
    for i in range(len(m)):
        start = m[i]
        end = m[i+1] if i+1<len(m) else None
        section_name = start.group(1)
        section_contents = contents[start.end(0)+1:end.start(0) if end is not None else len(contents)]
        section_contents = section_contents.strip() + "\n"
        section_fname = f"{i+1:02}" + "_" + re.sub("[^0-9a-zA-Z]+", "_", section_name).lower() + ".tex"
        section_names.append((section_name, section_fname))
        with open(f"{root}/{section_fname}", 'w') as f:
            f.write(section_contents)
    section_inputs = ""
    for (name, fname) in section_names:
        section_inputs += f"\\section{{{name}}}\n"
        section_inputs += f"\\input{{{fname}}}\n"
    main_contents = ""
    with open(f"{root}/main.tex", 'r') as f:
        main_contents = f.read()
    main_contents = re.sub("\\\\input\\{\\d.*\\}", "", main_contents)
    main_contents = re.sub("\\\\section\\{.*\\}", "", main_contents)
    main_contents = re.sub("\\n(\\n)+", "\\n\\n", main_contents)
    main_contents = main_contents.replace("\\end{document}", section_inputs + "\n\\end{document}")
    with open(f"{root}/main.tex", 'w') as f:
        f.write(main_contents)

def cleanup(root: str):
    for i in range(1, 25):
        fname = f"{root}/{i:02}.tex"
        try:
            os.remove(fname)
        except:
            break
    os.remove(f"{root}/collated.tex")

for entry in ["ia/dr", "ia/probability", "ia/vc"]:
    cleanup(entry)

# collate("ia/ns")
# uncollate("ia/ns")
# cleanup("ia/ns")
