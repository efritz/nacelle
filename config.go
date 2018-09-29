package nacelle

import (
	"github.com/efritz/nacelle/config"
	"github.com/efritz/nacelle/config/tag"
)

type (
	Config        = config.Config
	ConfigSourcer = config.Sourcer
	TagModifier   = tag.Modifier
)

var (
	NewConfig           = config.NewConfig
	NewEnvSourcer       = config.NewEnvSourcer
	NewMapSourcer       = config.NewMapSourcer
	NewYAMLFileSourcer  = config.NewYAMLFileSourcer
	NewConfigMapSourcer = config.NewConfigMapSourcer
	NewMultiSourcer     = config.NewMultiSourcer
	NewEnvTagPrefixer   = tag.NewEnvTagPrefixer
	NewDefaultTagSetter = tag.NewDefaultTagSetter
)
