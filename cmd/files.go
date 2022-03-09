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
	{FilePath: "ib/markov", Name: "Markov Chains"},
	{FilePath: "ib/antop", Name: "Analysis and Topology"},
	{FilePath: "ib/methods", Name: "Methods"},
	{FilePath: "ib/quantum", Name: "Quantum Mechanics"},
	{FilePath: "ib/linalg", Name: "Linear Algebra"},
	{FilePath: "ib/grm", Name: "Groups, Rings and Modules"},
	{FilePath: "ib/ca", Name: "Complex Analysis"},
	{FilePath: "ib/geom", Name: "Geometry"},
	{FilePath: "ib/stats", Name: "Statistics"},
}

var BookFiles []TexFile = []TexFile{
	{FilePath: "ia/book", Name: "IA Book"},
	{FilePath: "ib/book", Name: "IB Book"},
}

var FilesWithBook []TexFile

func init() {
	FilesWithBook = make([]TexFile, len(Files))
	copy(FilesWithBook, Files)
	FilesWithBook = append(FilesWithBook, BookFiles...)
}
