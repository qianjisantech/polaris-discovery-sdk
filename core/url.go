package core

type UrlType string

const (
	RegisterUrl UrlType = "/polaris/v1/discovery/agent/register"
	HeatBeatUrl UrlType = "/polaris/v1/discovery/agent/heatbeat/"
)
