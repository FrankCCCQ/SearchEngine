package config

// type IReader interface {
// 	readConfig() ([]byte, error)
// }

// type ConfigReader struct {
// 	FileName string
// }

// func (r *ConfigReader) readConfig() ([]byte, error) {
// 	file, err := os.ReadFile(r.FileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return file, err
// }

// func InitConfigForTest(reader IReader) {
// 	file, err := reader.readConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = yaml.Unmarshal(file, &Conf)
// 	if err != nil {
// 		panic(err)
// 	}
// }
