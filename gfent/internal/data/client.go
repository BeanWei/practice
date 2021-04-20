package data

import "github.com/BeanWei/gfent/internal/data/smc"

// GetSmcMasterClient .
func GetSmcMasterClient() smc.Client {
	return smc.NewClient(true)
}

// GetSmcSlaveClient .
func GetSmcSlaveClient() smc.Client {
	return smc.NewClient(false)
}
