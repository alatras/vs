package config

import (
	"os"
	"reflect"
	"testing"
)

func init() {
	Read("../")
}

func Test_Reading_Mongo(t *testing.T) {
	if App.Mongo.DB == "" {
		t.Error("Mongo App.Mongo.DB isn't set/read correctly")
	}
	if App.Mongo.URL == "" {
		t.Error("Mongo App.Mongo.URL isn't set/read correctly")
	}
	if reflect.TypeOf(App.Mongo.RetryMilliseconds).Kind() == reflect.TypeOf("test").Kind() {
		t.Error("Mongo App.Mongo.RetryMilliseconds isn't set/read correctly")
	}
}

func Test_Reading_Server(t *testing.T) {
	if reflect.TypeOf(App.HTTPPort).Kind() == reflect.TypeOf("test").Kind() {
		t.Error("Mongo App.HTTPPort isn't set/read correctly")
	}
}

func Test_Reading_Log(t *testing.T) {
	if App.Log.Format == "" {
		t.Error("App.Log.Format isn't set/read correctly")
	}
}

func Test_Adding_New_Env_Variable(t *testing.T) {
	err := os.Setenv("APP_DYNAMICS_APP_NAME", "Test App Name")
	if err != nil {
		t.Error("Failed to add a new env variable")
	}

	Read("../")

	if App.AppD.AppName != "Test App Name" {
		t.Error("App.AppD.AppName added to env but isn't set/read correctly")
	}
}
