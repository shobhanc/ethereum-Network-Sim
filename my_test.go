//This test is intended to bring up a bootstrapnode and subsequently bring up
//multiple nodes and discover each other using discovery v4 protocol.
package mytest

import (
        "fmt"
        "net"
        "os"
        "time"
        "testing"
        "strconv"

        //"github.com/ethereum/go-ethereum/internal/testlog"
        "github.com/ethereum/go-ethereum/node"
        "github.com/ethereum/go-ethereum/cmd/utils"
        "github.com/ethereum/go-ethereum/crypto"
        "github.com/ethereum/go-ethereum/p2p/discover"
        "github.com/ethereum/go-ethereum/p2p/enode"
        "github.com/ethereum/go-ethereum/log"
)

var (
	nodeKey, _ = crypto.GenerateKey()
	addr, _ = net.ResolveUDPAddr("udp", ":30301")
	conn, _ = net.ListenUDP("udp", addr)
	realaddr = conn.LocalAddr().(*net.UDPAddr)
        url string
        //Number of nodes to participate in the experiment
        noOfNodes = 2
)

func init(){

        glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
        glogger.Verbosity(log.Lvl(9))
        log.Root().SetHandler(glogger)
	realaddr.IP = net.IP{127, 0, 0, 1}
	url = enode.NewV4(&nodeKey.PublicKey, realaddr.IP, 0, realaddr.Port).URLv4()
}

func TestRun(t *testing.T) {
        t.Parallel()
	t.Run("BootStapNode", startBootStrapNode)
        i :=1
        for i<=noOfNodes {
	  t.Run("NODE"+strconv.Itoa(i), startNode)
          i++
        }
}


func startBootStrapNode(t *testing.T) {
	fmt.Println("startBootStrapNode")
	db, _ := enode.OpenDB("")
	ln := enode.NewLocalNode(db, nodeKey)
	cfg := discover.Config{
		PrivateKey: nodeKey,
                //Log:        testlog.Logger(t, log.LvlTrace),
	}
	_, err := discover.ListenUDP(conn, ln, cfg)
	if err != nil {
		t.Errorf("%v", err)
	}
        time.Sleep(30*time.Second)
        t.Parallel()
        select {}
}

func startNode(t* testing.T){
        t.Parallel()
        cfg := &node.Config{
                 //Logger:        testlog.Logger(t, log.LvlTrace),
               }
        cfg.P2P.BootstrapNodes = make([]*enode.Node, 0, 1)
        nodeurl, err := enode.Parse(enode.ValidSchemes, url)
        if err != nil {
              t.Errorf("Bootstrap URL invalid encode %s err %v", url, err)
        }
        cfg.P2P.BootstrapNodes = append(cfg.P2P.BootstrapNodes, nodeurl)
        cfg.P2P.ListenAddr = ":0"
        cfg.P2P.MaxPeers = node.DefaultConfig.P2P.MaxPeers

	// Create a network node to run protocols with the default values.
	stack, err := node.New(cfg)
	if err != nil {
		t.Errorf("Failed to create network node: %v", err)
	}
/*
        added := make(chan *node, 1)
        var table *discover.Table
        table.nodeAddedHook = func(n *discover.node) { added <- n }
*/

	defer stack.Close()

        utils.StartNode(stack)
        stack.Wait()
        
/*
        // The node should be added to the table shortly after getting the
        // pong packet.
        select {
        case n := <-added:
                fmt.Println("Node added ", testInstance)
                rid := encodePubkey(&test.remotekey.PublicKey).id()
                if n.ID() != rid {
                        t.Errorf("node has wrong ID: got %v, want %v", n.ID(), rid)
                }
                if !n.IP().Equal(test.remoteaddr.IP) {
                        t.Errorf("node has wrong IP: got %v, want: %v", n.IP(), test.remoteaddr.IP)
                }
                if n.UDP() != test.remoteaddr.Port {
                        t.Errorf("node has wrong UDP port: got %v, want: %v", n.UDP(), test.remoteaddr.Port)
                }
                if n.TCP() != int(testRemote.TCP) {
                        t.Errorf("node has wrong TCP port: got %v, want: %v", n.TCP(), testRemote.TCP)
                }
        case <-time.After(2 * time.Second):
                t.Errorf("node was not added within 2 seconds")
        }
*/

}
