package cmd

type TexFile struct {
	FilePath string
	Name     string
}

var Files []TexFile = []TexFile{
	{FilePath: "ia/ns", Name: "Numbers and Sets"},
	{FilePath: "ia/groups", Name: "Groups"},
	{FilePath: "ia/de", Name: "Differential Equations"},
	{FilePath: "ia/vm", Name: "Vectors and Matrices"},
	{FilePath: "ia/analysis", Name: "Analysis I"},
	{FilePath: "ia/probability", Name: "Probability"},
	{FilePath: "ia/dr", Name: "Dynamics and Relativity"},
	{FilePath: "ia/vc", Name: "Vector Calculus"},
	{FilePath: "ib/opt", Name: "Optimisation"},
	{FilePath: "ib/vp", Name: "Variational Principles"},
}

var FilesWithBook []TexFile

func init() {
	FilesWithBook = make([]TexFile, len(Files))
	copy(FilesWithBook, Files)
	FilesWithBook = append(FilesWithBook, TexFile{FilePath: "book", Name: "Book"})
}
