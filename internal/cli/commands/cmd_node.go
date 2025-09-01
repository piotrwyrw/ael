package commands

import (
	"ael/internal/config"
	"context"
	"fmt"
	"net"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"
)

type AelContext struct {
	Config *config.AelConfig
}

func (ael *AelContext) ensureConfiguration() error {
	if ael.Config == nil {
		return cli.Exit("Could not find an AEL configuration in current directory", 1)
	}

	return nil
}

func (ael *AelContext) CmdListNodes(ctx context.Context, cmd *cli.Command) error {
	if err := ael.ensureConfiguration(); err != nil {
		return err
	}

	nodes := ael.Config.Nodes
	if len(nodes) == 0 {
		fmt.Println("No nodes found")
		return nil
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Name", "Address"})

	for _, node := range nodes {
		t.AppendRow(table.Row{node.Name, node.IP.String()})
	}

	t.Render()

	return nil
}

func (ael *AelContext) CmdRemoveNode(ctx context.Context, cmd *cli.Command) error {
	if err := ael.ensureConfiguration(); err != nil {
		return err
	}

	cfg := ael.Config

	nodeName := cmd.String("name")

	if nodeName == "" {
		return cli.Exit("Node name must not be empty", 1)
	}

	node, nodeIndex := cfg.FindNodeByName(nodeName)

	if node == nil {
		return cli.Exit("Node \""+nodeName+"\" does not exist", 1)
	}

	cfg.Nodes = append(cfg.Nodes[:nodeIndex], cfg.Nodes[nodeIndex+1:]...)
	err := cfg.StoreConfiguration()
	if err != nil {
		return err
	}

	fmt.Printf("Removed node \"%s\"\n", nodeName)
	return nil
}

func (ael *AelContext) CmdAddNode(ctx context.Context, cmd *cli.Command) error {
	if err := ael.ensureConfiguration(); err != nil {
		return err
	}

	cfg := ael.Config

	nodeName := cmd.String("name")
	nodeAddr := cmd.String("address")

	if nodeName == "" {
		return cli.Exit("Node name must not be empty", 1)
	}

	if nodeAddr == "" {
		return cli.Exit("Node address must not be empty", 1)
	}

	existingNode, _ := cfg.FindNodeByName(nodeName)
	if existingNode != nil {
		return cli.Exit("Node \""+nodeName+"\" already exists", 1)
	}

	nodeIP := net.ParseIP(nodeAddr)
	if nodeIP == nil {
		return cli.Exit("Invalid node address", 1)
	}

	existingNode, _ = cfg.FindNodeByAddress(nodeIP)
	if existingNode != nil {
		return cli.Exit("Node at "+nodeIP.String()+" already exists: "+nodeName, 1)
	}

	fmt.Printf("Added node \"%s\" on %s\n", nodeName, nodeIP)

	cfg.Nodes = append(cfg.Nodes, config.Node{
		Name: nodeName,
		IP:   nodeIP,
	})

	err := cfg.StoreConfiguration()

	return err
}
