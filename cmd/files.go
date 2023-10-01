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
	{FilePath: "ii/algtop", Name: "Algebraic Topology"},
	{FilePath: "ii/pm", Name: "Probability and Measure"},
	{FilePath: "ii/graph", Name: "Graph Theory"},
	{FilePath: "ii/afl", Name: "Automata and Formal Languages"},
	{FilePath: "ii/galois", Name: "Galois Theory"},
	{FilePath: "ii/cc", Name: "Coding and Cryptography"},
	{FilePath: "ii/qic", Name: "Quantum Information and Computation"},
	{FilePath: "ii/lst", Name: "Logic and Set Theory"},
	{FilePath: "ii/alggeom", Name: "Algebraic Geometry"},
	{FilePath: "ii/nf", Name: "Number Fields"},
	{FilePath: "iii/commalg", Name: "Commutative Algebra"},
	{FilePath: "iii/alggeom", Name: "Algebraic Geometry"},
	{FilePath: "iii/mtncl", Name: "Model Theory and Non-Classical Logic"},
	{FilePath: "iii/cat", Name: "Category Theory"},
}

var BookFiles []TexFile = []TexFile{
	{FilePath: "ia/book", Name: "IA Book"},
	{FilePath: "ib/book", Name: "IB Book"},
	{FilePath: "ii/book", Name: "II Book"},
	{FilePath: "iii/book", Name: "III Book"},
	{FilePath: "ia/vol1", Name: "Volume 1"},
	{FilePath: "ia/vol2", Name: "Volume 2"},
	{FilePath: "ib/vol3", Name: "Volume 3"},
	{FilePath: "ib/vol4", Name: "Volume 4"},
	{FilePath: "ii/vol5", Name: "Volume 5"},
	{FilePath: "ii/vol6", Name: "Volume 6"},
	{FilePath: "iii/vol7", Name: "Volume 7"},
}

var FilesWithBook []TexFile

func init() {
	FilesWithBook = make([]TexFile, len(Files))
	copy(FilesWithBook, Files)
	FilesWithBook = append(FilesWithBook, BookFiles...)
}
