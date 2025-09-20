package configDirectory

type configDirIntegration interface {
	bind(f *Feature)
	files() []ConfigFile
}

type ConfigFile struct {
	Name string
	Data []byte
}

type ConfigDirectoryIntegration struct {
	ConfigFiles []ConfigFile
}

func (c *ConfigDirectoryIntegration) bind(f *Feature) {
	//TODO implement me
	panic("implement me")
}

func (c *ConfigDirectoryIntegration) files() []ConfigFile {
	//TODO implement me
	panic("implement me")
}
