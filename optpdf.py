import os
import files

files.files.append(("book/book", "book"))
for (fname, title) in files.files:
    print("Processing " + title)
    i = 0
    while True:
        i += 1
        print(f"Iteration {i}: ", end="", flush=True)
        output = os.popen(f"sudo bash ./optpdf.sh {fname}.pdf").read()
        if "Saving" not in output:
            print("Done\n")
            break
        print(output, end="", flush=True)
