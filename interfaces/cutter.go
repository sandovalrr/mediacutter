package interfaces

//CutterRepository template for cutter actions
type CutterRepository interface {
	ProgramCommand() string
	LenCommand(path string) string
	CutterCommand(filePath, outputPath string, start, end int) string
}
