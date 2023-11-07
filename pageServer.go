package isles

import (
	"context"
	_ "embed"
	"net/http"

	"github.com/SamHennessy/hlive"
	"github.com/SamHennessy/hlive/hlivekit"
	"github.com/rs/zerolog"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed assets/tailwind.js
var tailwindjs hlive.HTML

//go:embed assets/daisyui.css
var daisyui hlive.HTML

type PageServer struct {
	pageFN   func(ctx context.Context) *hlive.Page
	logger   zerolog.Logger
	ctxWails context.Context
}

func NewPageServer(pageFN func(ctxWails context.Context) *hlive.Page) *PageServer {
	newPageFN := func(ctxWails context.Context) *hlive.Page {
		page := pageFN(ctxWails)
		pubSub := hlivekit.NewPubSub()

		page.DOM().Head().Add(
			hlivekit.InstallPubSub(pubSub),
			hlive.T("script", hlive.Attrs{"src": "/wails/ipc.js"}),
			hlive.T("script", hlive.Attrs{"src": "/wails/runtime.js"}),
			hlive.T("style", daisyui),
			hlive.T("script", tailwindjs),
		)

		return page
	}

	return &PageServer{
		pageFN: newPageFN,
		// TODO: get logger from page
		logger: hlive.LoggerDev.Level(zerolog.DebugLevel),
	}
}

func (ps *PageServer) Serve(ctxWails context.Context) {
	ps.ctxWails = ctxWails

	hlive.LoggerDev.Debug().Msg("Serve")

	runtime.EventsOn(ctxWails, "connect", func(optionalData ...interface{}) {
		ps.logger.Debug().Msg("Connect")

		fromWails := make(chan hlive.MessageWS)
		toWails := make(chan hlive.MessageWS)

		runtime.EventsOn(ctxWails, "out", func(optionalData ...interface{}) {
			var message []byte

			if len(optionalData) > 0 {
				str, ok := optionalData[0].(string)
				if ok {
					message = []byte(str)
				}
			}

			// TODO: file upload
			fromWails <- hlive.MessageWS{
				Message:  message,
				IsBinary: false,
			}
		})

		go func() {
			for {
				select {
				case <-ctxWails.Done():
					return
				case msg := <-toWails:
					runtime.EventsEmit(ctxWails, "in", string(msg.Message))
				}
			}
		}()

		go func() {
			err := ps.pageFN(ps.ctxWails).ServeWS(ctxWails, "wails", toWails, fromWails)
			if err != nil {
				hlive.LoggerDev.Err(err).Msg("ServeWS")
			}

			hlive.LoggerDev.Info().Msg("ServeWS: done")
		}()
	})
}

func (ps *PageServer) AssetsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/hlive.html" {
			if _, err := ps.pageFN(ps.ctxWails).RunRenderPipeline(r.Context(), w); err != nil {
				hlive.LoggerDev.Err(err).Msg("pageFN.RunRenderPipeline")

				return
			}
		}
	})
}

func (ps *PageServer) OnStartup(ctx context.Context) {
	go ps.Serve(ctx)
}
