package routeUtil

import (
	"github.com/gin-gonic/gin"
)

func DealRoutes(r *gin.RouterGroup) {
	r.POST("/trades/create/:zipcode", CreateTrade())
	r.PUT("/trades/:tid", EditTrade())
	r.GET("/trades/:tid", FetchTrade())
	r.GET("/trades", FetchTrades())
	r.GET("/trades/user/:uid", GetTradesByUserId())
	r.GET("/trades/organization/:oid", GetTradesByOrgId())
	r.GET("/trades/tcpreference/:id/:orgid", FetchTradeById())
	r.GET("/trades/checkOwner/:tradeid", AuthTradeOwner())
	r.POST("/trades/acceptOffer/:tid/:uid", AcceptOffer())
	r.POST("/trades/offer/:id", AddDealOffer())
	r.GET("/trades/offer/:tid/amnt", GetOfferAmount())
	r.GET("/trades/offer/:tid/:id", GetOffer())
	r.GET("/trades/offersFetch/:tid/", FetchTradeOffers())
	//r.DELETE("/trades/:tid", DeleteTrade())
	//delete trade doesnt exist
}
