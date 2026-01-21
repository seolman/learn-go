package main

import ()

type emailStatus int

const (
	EmailBounced = iota
	EmailInvalid
	EmailDelivered
	EmailOpened
)
