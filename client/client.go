package client

import (
	"encoding/json"
	"errors"
	"karim/http_server/constants"
	"karim/http_server/httpstatus"
	"karim/http_server/logger"
	"karim/http_server/request/types"
	"karim/http_server/response"
	"net"
	"strconv"
)

type Client struct {
	Connection 	net.Conn
	Response   	response.ResponseType
	Request 	types.Request
}


// SendResponse serializes the Response map and sends it to the client
func (c *Client) SendResponse(r response.TmpResponse) error {
	logger.InfoLogger.Printf("Sending the Response to %s", c.Connection.LocalAddr().String())

	if c.Connection == nil {
		return errors.New("no active connection")
	}

	c.Response.StartLine = response.StartLine{
		Version:      constants.PROTOCOL_VERSION,
		StatusCode:   r.Status,
		StatusMessage: httpstatus.StatusCodes[r.Status],
	}

	// Handle marshaling correctly
	var dataMarshalled []byte
	var err error

	if r.Data != nil {
		dataMarshalled, err = json.Marshal(r.Data)
		if err != nil {
			return err
		}
	} else {
		dataMarshalled = []byte("null") // Explicitly set "null" for JSON compatibility
	}

	// Ensure Data is correctly formatted
	c.Response.ResponseBody, err = json.Marshal(response.TmpResponse{
		Data:    json.RawMessage(dataMarshalled), // Prevents double encoding
		Message: r.Message,
		Status:  r.Status,
	})

	if err != nil {
		return err
	}

	c.Response.Headers = map[string]string{
		"Content-Type":   "application/json",
		"Content-Length": strconv.Itoa(len(c.Response.ResponseBody)),
	}

	responseString, err := c.Response.String()
	if err != nil {
		return err
	}

	_, err = c.Connection.Write([]byte(responseString))
	if err != nil {
		return err
	}

	return nil
}

