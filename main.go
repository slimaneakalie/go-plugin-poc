package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshake,
		Plugins: map[string]plugin.Plugin{
			"dest_grpc": &DestPluginImpl{},
		},
		Cmd:              exec.Command("sh", "-c", os.Getenv("DEST_PLUGIN_PATH")),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("dest_grpc")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// We should have a KV store now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	pl := raw.(DestinationPlugin)

	reqExample := UpsertUsersRequest{
		Users: []*User{
			{Id: "id1", FullName: "full name 1", FirstName: "first name 1"},
			{Id: "id2", FullName: "full name 2", FirstName: "first name 2"},
			{Id: "id3", FullName: "full name 3", FirstName: "first name 3"},
		},
		Mappings: []byte(`{"mkey1": "value1", "mkey2": 153}`),
		Settings: []byte(`{"skey1": "value1", "skey2": 153}`),
	}

	os.Args = os.Args[1:]
	switch os.Args[0] {
	case "upsert":
		pl.BatchUpsertUsers()

	case "update":

	case "delete":

	default:
		fmt.Printf("Please only use 'upsert', 'update', or 'delete', given: %q", os.Args[0])
		os.Exit(1)
	}

	os.Exit(0)
}

var handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}
