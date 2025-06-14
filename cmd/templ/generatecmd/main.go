package generatecmd

import (
	"context"
	_ "embed"
	"log/slog"

	_ "net/http/pprof"

	"github.com/a-h/templ/generator"
)

type Arguments struct {
	FileName                        string
	FileWriter                      FileWriterFunc
	Path                            string
	Watch                           bool
	WatchPattern                    string
	OpenBrowser                     bool
	Command                         string
	ProxyBind                       string
	ProxyPort                       int
	Proxy                           string
	NotifyProxy                     bool
	WorkerCount                     int
	GenerateSourceMapVisualisations bool
	IncludeVersion                  bool
	IncludeTimestamp                bool
	Stripspace                      bool
	// PPROFPort is the port to run the pprof server on.
	PPROFPort         int
	KeepOrphanedFiles bool
	Lazy              bool
}

func Run(ctx context.Context, log *slog.Logger, args Arguments) (err error) {
	g, err := NewGenerate(log, args)
	if err != nil {
		return err
	}
	generator.Stripspace = args.Stripspace
	return g.Run(ctx)
}
