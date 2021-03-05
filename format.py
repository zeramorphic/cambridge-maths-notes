import os
import shutil
import files

all_files = files.files
all_files.append(("utilities", "util.tex"))
for (fname, title) in files.files:
    print("Processing " + title)
    shutil.copyfile(fname, "latexindent.tex")
    os.system(f"latexindent latexindent.tex -o={fname}")
    os.remove("latexindent.tex")
