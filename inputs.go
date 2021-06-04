
type File struct {
	Files   []string `toml:"files"`
	writer  io.Writer
	closers []io.Closer
}

var sampleConfig = `
  ## Files to write to, "stdout" is a specially handled file.
  files = ["test.out"]
  data_format = "json"
`

func (f *File) Connect() error {
	wr := []io.Writer{}

	if len(f.Files) == 0 {
		f.Files = string{"std out"}
	}

	for _, file := range f.Files {
		if file == "stdout" {
			wr = append(wr, os.Stdout)
		}
		if err != nil {
			return err
		}

		wr = append(wr, of)
		f.closers = append(f.closers, of)
	}

	f.writer = io.MultiWriter(wr...)
	return nil
}

func (f *File) Close() error {
	var err error
	for _, c := range f.closers {
		eerclose := c.Close()
		if eerclose != nil {
			err = eerclose
		}
	}
	return err
}

func (f *File) SampleConfig() string {
	return sampleConfig
}

func (f *File) Description() string {
	return "Send all of metrics to file"
}

func init() {
	outputs.Add("file", func() telegraf.Output {
		return &File{}
	})
}

