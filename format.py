import os
import shutil
import files

all_files = files.files
all_files.append(("util", "utilities"))
for (fname, title) in files.files:
    print("Processing " + title)
    shutil.copyfile(fname + ".tex", "latexindent.tex")
    os.system(f"latexindent -s latexindent.tex -o={fname}.tex")
    os.remove("latexindent.tex")
