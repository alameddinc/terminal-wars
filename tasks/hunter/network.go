package hunter

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type Connection struct {
	ID        int
	Username  string
	Hash      string
	Host      *Connection
	Customers []*Connection
	Status    bool
}

func (c *Connection) JoinNetwork(cList *[]*Connection) {
	c.ID = len(*cList) + 1
	c.Host = JoinRandomHost(cList)
	c.Customers = []*Connection{}
	c.Status = true
	c.Hash = GenerateHash()
	*cList = append(*cList, c)
	c.Host.Customers = append(c.Host.Customers, c)
}

func InitRoot(cList *[]*Connection) {
	root := Connection{
		ID:        0,
		Hash:      GenerateHash(),
		Host:      nil,
		Customers: []*Connection{},
		Status:    true,
	}
	*cList = append(*cList, &root)
}

func InitNetwork(cList *[]*Connection) {
	InitRoot(cList)
	for i := 0; i < 100; i++ {
		tmp := Connection{
			ID:        i + 1,
			Hash:      GenerateHash(),
			Host:      JoinRandomHost(cList),
			Customers: []*Connection{},
			Status:    true,
		}
		*cList = append(*cList, &tmp)
		tmp.Host.Customers = append(tmp.Host.Customers, &tmp)
	}
	fmt.Print("Started")
}

func JoinRandomHost(cList *[]*Connection) *Connection {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	rand := random.Int() % len(*cList)
	return (*cList)[rand]
}

func GenerateHash() string {
	return "TR-" + uuid.New().String()
}

// Todo - AO - AÇ - Bir düğümden başka düğüme ulaşmanın yolunu hesaplayan fonksiyon

func (c *Connection) DrawWay(to Connection) ([]*Connection, bool) {
	responseConnections := []*Connection{}
	for {
		if c.Host == nil {
			break
		}
		responseConnections = append(responseConnections, c)
		if c.Hash == to.Hash {
			return responseConnections, true
		}
		c = c.Host
	}
	childArr, _ := c.RunOnTheRoad(to)
	responseConnections = append(responseConnections, c)
	responseConnections = append(responseConnections, childArr...)
	return responseConnections, true
}

func (c *Connection) RunOnTheRoad(to Connection) ([]*Connection, bool) {
	// altına bak
	if c.Hash == to.Hash {
		return nil, false
	}
	responseConnections := []*Connection{}
	for _, v := range c.Customers {
		if to.Hash == v.Hash {
			responseConnections = append(responseConnections, v)
			return responseConnections, true
		}
		childArr, status := v.RunOnTheRoad(to)
		if status {
			responseConnections = append(responseConnections, v)
			for _, v2 := range childArr {
				responseConnections = append(responseConnections, v2)
			}
			return responseConnections, true
		}
	}
	return nil, false
}
