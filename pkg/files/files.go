package files

import "os"

func List(folder_path string) ([]string, error) {
	var file_names []string

	files, err := os.ReadDir(folder_path)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		file_names = append(file_names, file.Name())
	}

	return file_names, nil
}
