package isles

import (
	"context"

	"github.com/SamHennessy/hlive"
)

var Themes = []any{
	Option("", nil, "Pick", AttrDisabled(), AttrSelected()),
	Option("", nil, "light"),
	Option("", nil, "dark"),
	Option("", nil, "cupcake"),
	Option("", nil, "bumblebee"),
	Option("", nil, "emerald"),
	Option("", nil, "corporate"),
	Option("", nil, "synthwave"),
	Option("", nil, "retro"),
	Option("", nil, "cyberpunk"),
	Option("", nil, "valentine"),
	Option("", nil, "halloween"),
	Option("", nil, "garden"),
	Option("", nil, "forest"),
	Option("", nil, "aqua"),
	Option("", nil, "lofi"),
	Option("", nil, "pastel"),
	Option("", nil, "fantasy"),
	Option("", nil, "wireframe"),
	Option("", nil, "black"),
	Option("", nil, "luxury"),
	Option("", nil, "dracula"),
	Option("", nil, "cmyk"),
	Option("", nil, "autumn"),
	Option("", nil, "business"),
	Option("", nil, "acid"),
	Option("", nil, "lemonade"),
	Option("", nil, "night"),
	Option("", nil, "coffee"),
	Option("", nil, "winter"),
}


const TopicPickTheme = "isles-pick-theme"

func ThemePicker(class string, defaultTheme string) (hlive.AttrsLockBox, *hlive.ComponentMountable) {
	theme := hlive.NewLockBox[string](defaultTheme)

	c := Select(class, func(ctx context.Context, e hlive.Event) {
		theme.Set(e.Value)
	}, Themes...)

	return hlive.AttrsLockBox{"data-theme": theme}, c
}
