package metrics

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	conn   net.Conn
	prefix string
}

// NewClient creates a new Graphite metrics client
func NewClient(host, port, prefix string) (*Client, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		prefix: prefix,
	}, nil
}

// Close closes the connection to Graphite
func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// SendMetric sends a metric to Graphite
func (c *Client) SendMetric(name string, value float64) error {
	if c.conn == nil {
		return fmt.Errorf("no connection to graphite")
	}

	timestamp := time.Now().Unix()
	metric := fmt.Sprintf("%s.%s %.2f %d\n", c.prefix, name, value, timestamp)

	_, err := c.conn.Write([]byte(metric))
	return err
}

// Increment sends a counter increment to Graphite
func (c *Client) Increment(name string) error {
	return c.SendMetric(name, 1)
}

// Gauge sends a gauge value to Graphite
func (c *Client) Gauge(name string, value float64) error {
	return c.SendMetric(name, value)
}

// Timer sends a timing metric to Graphite
func (c *Client) Timer(name string, duration time.Duration) error {
	return c.SendMetric(name, float64(duration.Milliseconds()))
}
