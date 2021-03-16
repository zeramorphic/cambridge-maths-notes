# Maths Notes

LaTeX sources for my notes on the maths courses at Cambridge. To view the compiled PDFs, click [here](https://thirdsgames.co.uk/maths.html). Please note that these are _not_ exhaustive descriptions of the courses themselves - please watch all of the lectures! This is simply a way of collating much of the courses' information in one place.

## Building

- Install a TeX distribution, such as MiKTeX or TeX Live.
- Install Go 1.16.
- Run `go get ./...` to fetch dependencies.
- Run `go run . build` to build the project. You can execute `go run . help` for more information about various other commands.

  - For the `fmt` command, you must install `latexindent` and add it to `PATH`.
  - For the `optimise` command, you must install `ghostscript`.

## Contributing

To contribute to the project (such as making corrections or adding explanations), fork the repository and then submit a pull request. Any contributions must be made under the same license as the project itself.
