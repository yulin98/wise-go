package router

import (
	chargePointHandler "com.wisecharge/central/internal/api/chargepoint"
	connectorHandler "com.wisecharge/central/internal/api/connector"
	priceHandler "com.wisecharge/central/internal/api/price"
	stationHandler "com.wisecharge/central/internal/api/station"
)

/*func setApiRouter(r *resource) {
	//设置路由
	router := r.srv.Handler.(*gin.Engine)
	// service interface
	apiGroup := router.Group("/api")
	{
		// connector
		connectorGroup := apiGroup.Group("/connector")
		{
			handler := connectorHandler.New(r.log, r.redis, r.mysql)
			connectorGroup.POST("/create_connector", handler.CreateConnector)
			connectorGroup.POST("/delete_connector", handler.DeleteConnector)
			connectorGroup.POST("/update_connector", handler.UpdateConnector)
			connectorGroup.POST("/query_one_connector", handler.QueryOneConnector)
			connectorGroup.POST("/query_page_connector", handler.QueryPageConnector)
		}

		// chargePoint
		chargePointGroup := apiGroup.Group("/charge_point")
		{
			handler := chargePointHandler.New(r.log, r.redis, r.mysql)
			chargePointGroup.POST("/create_charge_point", handler.CreateChargePoint)
			chargePointGroup.POST("/delete_charge_point", handler.DeleteChargePoint)
			chargePointGroup.POST("/update_charge_point", handler.UpdateChargePoint)
			chargePointGroup.POST("/query_one_charge_point", handler.QueryOneChargePoint)
			chargePointGroup.POST("/query_page_charge_point", handler.QueryPageChargePoint)
		}

		// station
		stationGroup := apiGroup.Group("/station")
		{
			handler := stationHandler.New(r.log, r.redis, r.mysql)
			stationGroup.POST("/create_station", handler.CreateStation)
			stationGroup.POST("/delete_station", handler.DeleteStation)
			stationGroup.POST("/update_station", handler.UpdateStation)
			stationGroup.POST("/query_one_station", handler.QueryOneStation)
			stationGroup.POST("/query_page_station", handler.QueryPageStation)
		}

		// price
		priceGroup := apiGroup.Group("/price")
		{
			handler := priceHandler.New(r.log, r.redis, r.mysql)
			priceGroup.POST("/create_price", handler.CreatePrice)
			priceGroup.POST("/delete_price", handler.DeletePrice)
			priceGroup.POST("/update_price", handler.UpdatePrice)
			priceGroup.POST("/query_one_price", handler.QueryOnePrice)
			priceGroup.POST("/query_page_price", handler.QueryPagePrice)
		}
	}
}
*/

func setApiRouter(r *resource) {
	//设置路由 service interface
	apiGroup := r.mux.Group("/api")
	{
		// connector
		connectorGroup := apiGroup.Group("/connector")
		{
			handler := connectorHandler.New(r.log, r.redis, r.mysql)
			connectorGroup.POST("/create_connector", handler.CreateConnector())
			connectorGroup.POST("/delete_connector", handler.DeleteConnector())
			connectorGroup.POST("/update_connector", handler.UpdateConnector())
			connectorGroup.POST("/query_one_connector", handler.QueryOneConnector())
			connectorGroup.POST("/query_page_connector", handler.QueryPageConnector())
		}

		// chargePoint
		chargePointGroup := apiGroup.Group("/charge_point")
		{
			handler := chargePointHandler.New(r.log, r.redis, r.mysql)
			chargePointGroup.POST("/create_charge_point", handler.CreateChargePoint())
			chargePointGroup.POST("/delete_charge_point", handler.DeleteChargePoint())
			chargePointGroup.POST("/update_charge_point", handler.UpdateChargePoint())
			chargePointGroup.POST("/query_one_charge_point", handler.QueryOneChargePoint())
			chargePointGroup.POST("/query_page_charge_point", handler.QueryPageChargePoint())
		}

		// station
		stationGroup := apiGroup.Group("/station")
		{
			handler := stationHandler.New(r.log, r.redis, r.mysql)
			stationGroup.POST("/create_station", handler.CreateStation())
			stationGroup.POST("/delete_station", handler.DeleteStation())
			stationGroup.POST("/update_station", handler.UpdateStation())
			stationGroup.POST("/query_one_station", handler.QueryOneStation())
			stationGroup.POST("/query_page_station", handler.QueryPageStation())
		}

		// price
		priceGroup := apiGroup.Group("/price")
		{
			handler := priceHandler.New(r.log, r.redis, r.mysql)
			priceGroup.POST("/create_price", handler.CreatePrice())
			priceGroup.POST("/delete_price", handler.DeletePrice())
			priceGroup.POST("/update_price", handler.UpdatePrice())
			priceGroup.POST("/query_one_price", handler.QueryOnePrice())
			priceGroup.POST("/query_page_price", handler.QueryPagePrice())
		}
	}
}
