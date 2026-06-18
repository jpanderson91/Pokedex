package pokeapi

import (
	"net/http"
)

type Client struct{
	httpClient http.Client
}