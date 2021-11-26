package rename

func Do(files []string, name string, prefix string, suffix string, devider string) error {
	return nil
}

func WithPrefix(name string, prefix string) string {
	return prefix + name
}

func WithSuffix(name string, suffix string) string {
	return name + suffix
}
