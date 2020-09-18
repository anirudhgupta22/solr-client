package connectionFactory


type CollectionConfig struct {
	name          string
	master        *Connection
	slaves        []*Connection
	roundRobinCtr int64
}

func NewEmptyCollectionConfig() *CollectionConfig {
	return &CollectionConfig{}
}

func NewCollectionConfig(name string, master *Connection, slaves []*Connection) *CollectionConfig {
	return &CollectionConfig{name: name, master: master, slaves: slaves}
}

func (c *CollectionConfig) Master() *Connection {
	return c.master
}

func (c *CollectionConfig) SetMaster(master *Connection) {
	c.master = master
}

func (c *CollectionConfig) GetSlavesCount() int64 {
	return int64(len(c.slaves))
}

func (c *CollectionConfig) SetSlaves(slaves []*Connection) {
	c.slaves = slaves
}

func (c *CollectionConfig) AddSlaves(slave *Connection) {
	if c.slaves == nil {
		c.slaves = make([]*Connection, 0)
	}
	c.slaves = append(c.slaves, slave)
}

func (c *CollectionConfig) Slave() *Connection {
	slaveNo := c.roundRobinCtr % c.GetSlavesCount()
	c.roundRobinCtr++
	return c.slaves[slaveNo]
}
