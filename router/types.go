package router

type Router string
type webhook Router
type api Router

// Name of the router
var Webhook webhook = "webhook"
var API api = "api"
