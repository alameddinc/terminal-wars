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
	c.ID = len(*cList)
	c.Host = JoinRandomHost(cList)
	c.Customers = []*Connection{}
	c.Status = true
	c.Hash = GenerateHash()
	*cList = append(*cList, c)
	c.Host.Customers = append(c.Host.Customers, c)
}

func (c *Connection) LeaveNetwork(cList *[]*Connection) {
	// Customer-> Parenta bağla ->
	// yoksa -> rasgele bağla
	for _, v := range c.Customers {
		if c.Host != nil {
			v.Host = c.Host
		} else {
			v.Host = JoinRandomHost(cList)
		}
		v.Host.Customers = append(v.Host.Customers, v)
	}
	c.Customers = nil
	// Host'tan kendi kaydını sil
	var tmpCustomers []*Connection
	for _, v := range c.Host.Customers {
		if v.Hash != c.Hash {
			tmpCustomers = append(tmpCustomers, v)
		}
	}
	c.Host.Customers = tmpCustomers
	// COnnection Listten Kaydını sil
	(*cList)[c.ID] = nil
	*cList = append((*cList)[:c.ID], (*cList)[c.ID+1:]...)
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
func Calculate(ConnectionList []*Connection) {

}
