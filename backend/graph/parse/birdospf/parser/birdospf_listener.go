// Code generated from BirdOSPF.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // BirdOSPF
import "github.com/antlr4-go/antlr/v4"

// BirdOSPFListener is a complete listener for a parse tree produced by BirdOSPFParser.
type BirdOSPFListener interface {
	antlr.ParseTreeListener

	// EnterState is called when entering the state production.
	EnterState(c *StateContext)

	// EnterArea is called when entering the area production.
	EnterArea(c *AreaContext)

	// EnterRouter is called when entering the router production.
	EnterRouter(c *RouterContext)

	// EnterNetwork is called when entering the network production.
	EnterNetwork(c *NetworkContext)

	// EnterRouterEntry is called when entering the routerEntry production.
	EnterRouterEntry(c *RouterEntryContext)

	// EnterDistance is called when entering the distance production.
	EnterDistance(c *DistanceContext)

	// ExitState is called when exiting the state production.
	ExitState(c *StateContext)

	// ExitArea is called when exiting the area production.
	ExitArea(c *AreaContext)

	// ExitRouter is called when exiting the router production.
	ExitRouter(c *RouterContext)

	// ExitNetwork is called when exiting the network production.
	ExitNetwork(c *NetworkContext)

	// ExitRouterEntry is called when exiting the routerEntry production.
	ExitRouterEntry(c *RouterEntryContext)

	// ExitDistance is called when exiting the distance production.
	ExitDistance(c *DistanceContext)
}
