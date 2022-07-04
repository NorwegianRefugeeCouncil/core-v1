package devinit

import (
	"os"
	"path"
)

type Config struct {
	SpiceDBToken   string
	ZanzibarPrefix string
	ZanzibarToken  string
}

func (c *Config) MakeConfig() error {

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	RootDir := cwd
	CredsDir := path.Join(RootDir, "creds")
	ZanzibarDir := path.Join(CredsDir, "zanzibar")
	SpiceDBDir := path.Join(CredsDir, "spicedb")

	c.SpiceDBToken, err = getOrCreateRandomSecretStr(32, SpiceDBDir, "token")
	if err != nil {
		return err
	}

	c.ZanzibarToken, err = getOrCreateRandomSecretStr(32, ZanzibarDir, "token")
	if err != nil {
		return err
	}

	c.ZanzibarPrefix, err = getOrCreateRandomSecretStr(32, ZanzibarDir, "prefix")
	if err != nil {
		return err
	}

	return nil
}
