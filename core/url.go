package core

type UrlType string

const (
	RegisterUrl UrlType = "polaris/v1/agent/register"
	HeatBeatUrl UrlType = "polaris/v1/agent/beat"
)
