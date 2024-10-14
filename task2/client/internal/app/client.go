package app

type Client struct {
	host int
}

func NewClient() *Client {
	return &Client{}
}
