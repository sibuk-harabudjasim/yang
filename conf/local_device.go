package conf

import (
	"encoding/json"
	"io"
	"os"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/meta/yang"
	"github.com/c2stack/c2g/node"
)

type LocalDevice struct {
	browsers     map[string]*node.Browser
	schemaSource meta.StreamSource
	uiSource     meta.StreamSource
}

func NewDevice(schemaSource meta.StreamSource) *LocalDevice {
	return &LocalDevice{
		schemaSource: schemaSource,
		browsers:     make(map[string]*node.Browser),
	}
}

func NewDeviceWithUi(schemaSource meta.StreamSource, uiSource meta.StreamSource) *LocalDevice {
	return &LocalDevice{
		schemaSource: schemaSource,
		uiSource:     uiSource,
		browsers:     make(map[string]*node.Browser),
	}
}

func (self *LocalDevice) SchemaSource() meta.StreamSource {
	return self.schemaSource
}

func (self *LocalDevice) UiSource() meta.StreamSource {
	return self.uiSource
}

func (self *LocalDevice) Modules() map[string]*meta.Module {
	mods := make(map[string]*meta.Module)
	for _, b := range self.browsers {
		m := b.Meta.(*meta.Module)
		mods[m.GetIdent()] = m
	}
	return mods
}

func (self *LocalDevice) Browser(module string) (*node.Browser, error) {
	return self.browsers[module], nil
}

func (self *LocalDevice) Close() {
}

func (self *LocalDevice) Add(module string, n node.Node) error {
	m, err := yang.LoadModule(self.schemaSource, module)
	if err != nil {
		return err
	}
	self.browsers[module] = node.NewBrowser(m, n)
	return nil
}

func (self *LocalDevice) AddBrowser(b *node.Browser) {
	self.browsers[b.Meta.GetIdent()] = b
}

func (self *LocalDevice) ApplyStartupConfig(config io.Reader) error {
	var cfg map[string]interface{}
	if err := json.NewDecoder(config).Decode(&cfg); err != nil {
		return err
	}
	for module, data := range cfg {
		b, err := self.Browser(module)
		if err != nil {
			return err
		}
		if b == nil {
			return c2.NewErrC("browser not found:"+module, 404)
		}
		moduleCfg := data.(map[string]interface{})
		if err := b.Root().UpsertFrom(node.MapNode(moduleCfg)).LastErr; err != nil {
			return err
		}
	}
	return nil
}

func (self *LocalDevice) ApplyStartupConfigFile(fname string) error {
	cfgRdr, err := os.OpenFile(fname, os.O_RDWR, os.ModeExclusive)
	defer cfgRdr.Close()
	if err != nil {
		panic(err)
	}
	return self.ApplyStartupConfig(cfgRdr)
}