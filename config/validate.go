package config

import "errors"

func (c Server) Validate() error {
	if c.HTTPPort == 0 {
		return errors.New("httpPort is missing")
	}

	if err := c.Mongo.Validate(); err != nil {
		return err
	}

	if err := c.AppD.Validate(); err != nil {
		return err
	}

	if err := c.Log.Validate(); err != nil {
		return err
	}

	return nil
}

func (c Mongo) Validate() error {
	if c.URL == "" {
		return errors.New("mongo.url is missing")
	}

	if c.DB == "" {
		return errors.New("mongo.db is missing")
	}

	return nil
}

func (c AppD) Validate() error {
	if c.AppName == "" {
		return errors.New("appd.appName is missing")
	}

	if c.TierName == "" {
		return errors.New("appd.tierName is missing")
	}

	if c.NodeName == "" {
		return errors.New("appd.nodeName is missing")
	}

	if c.InitTimeout == 0 {
		return errors.New("appd.initTimeout is missing")
	}

	if c.Controller.Host == "" {
		return errors.New("appd.controller.host is missing")
	}

	if c.Controller.Port == 0 {
		return errors.New("appd.controller.port is missing")
	}

	if c.Controller.Account == "" {
		return errors.New("appd.controller.account is missing")
	}

	if c.Controller.AccessKey == "" {
		return errors.New("appd.controller.accessKey is missing")
	}

	return nil
}

func (c Log) Validate() error {
	if c.Format == "" {
		return errors.New("log.format is missing")
	}

	if c.Level == "" {
		return errors.New("log.level is missing")
	}

	return nil
}
