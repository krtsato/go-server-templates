package config

// Web アプリケーション外部の Web API に関する設定
type Web struct {
	Port      string `yaml:"port"`
	AuthCheck bool   `yaml:"authCheck"`
}
