package connector

import (
	"com.wisecharge/central/internal/package/core"
	"fmt"
)

// createConnector
func (h *handler) CreateConnector() core.HandlerFunc {
	return func(c core.Context) {

		fmt.Println("aaaa")
	}
}

// deleteConnector
func (h *handler) DeleteConnector() core.HandlerFunc {
	return func(c core.Context) {

		fmt.Println("aaaa")
	}
}

// updateConnector
func (h *handler) UpdateConnector() core.HandlerFunc {
	return func(c core.Context) {

		fmt.Println("aaaa")
	}
}

// QueryOneConnector
func (h *handler) QueryOneConnector() core.HandlerFunc {
	return func(c core.Context) {

		fmt.Println("aaaa")
	}
}

// QueryPageConnector
func (h *handler) QueryPageConnector() core.HandlerFunc {
	return func(c core.Context) {

		fmt.Println("aaaa")
	}
}
