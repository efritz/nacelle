package config

// Sourcer pulls requested names from a variable source. This can be the
// environment, a file, a remote server, etc. This can be done on-demand
// per variable, or a cache of variables can be built on startup and then
// pulled from a cached mapping as requested.
type Sourcer func(envTagValue, contextTagValue string) (string, bool)
