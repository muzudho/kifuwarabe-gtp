package ui

import (
	"io/ioutil"

	e "github.com/muzudho/kifuwarabe-gtp/entities"
	u "github.com/muzudho/kifuwarabe-gtp/usecases"
	"github.com/pelletier/go-toml"
)

// LoadEntryConf - ゲーム設定ファイルを読み込みます。
func LoadEntryConf(path string) e.EntryConf {

	// ファイル読込
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(u.G.Log.Fatal("path=%s err=%s", path, err))
	}

	debugPrintToml(fileData)

	// Toml解析
	binary := []byte(string(fileData))
	config := e.EntryConf{}
	toml.Unmarshal(binary, &config)

	return config
}

func debugPrintToml(fileData []byte) {
	// u.G.Log.Trace("<Engine> content=%s", string(fileData))

	// Toml解析
	tomlTree, err := toml.Load(string(fileData))
	if err != nil {
		panic(err)
	}
	u.G.Log.Trace("<Engine> Input:\n")
	u.G.Log.Trace("<Engine> Engine.Komi=%f\n", tomlTree.Get("Engine.Komi").(float64))
	u.G.Log.Trace("<Engine> Engine.BoardSize=%d\n", tomlTree.Get("Engine.BoardSize").(int64))
	u.G.Log.Trace("<Engine> Engine.MaxMoves=%d\n", tomlTree.Get("Engine.MaxMoves").(int64))
	u.G.Log.Trace("<Engine> Engine.BoardData=%s\n", tomlTree.Get("Engine.BoardData").(string))
}
func debugPrintConfig(config e.EntryConf) {
	u.G.Log.Trace("<Engine> Memory:\n")
	u.G.Log.Trace("<Engine> Server.Host=%s\n", config.Server.Host)
	u.G.Log.Trace("<Engine> Server.Port=%d\n", config.Server.Port)
	u.G.Log.Trace("<Engine> User.Name=%s\n", config.User.Name)
	u.G.Log.Trace("<Engine> User.Pass=%s\n", config.User.Pass)
	u.G.Log.Trace("<Engine> Engine.Komi=%f\n", config.Engine.Komi)
	u.G.Log.Trace("<Engine> Engine.BoardSize=%d\n", config.Engine.BoardSize)
	u.G.Log.Trace("<Engine> Engine.MaxMoves=%d\n", config.Engine.MaxMoves)
	u.G.Log.Trace("<Engine> Engine.MaxMoves=%s\n", config.Engine.BoardData)
	u.G.Log.Trace("<Engine> Engine.SentinelBoardMax()=%d\n", config.SentinelBoardMax())
}
