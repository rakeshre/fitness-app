package api

import (
	db "Gym-backend/db/sqlc"
	"Gym-backend/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	store  db.Store
	//tokenMaker token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	//tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	//if err != nil {
	//	return nil, fmt.Errorf("cannot create token maker: %w", err)
	//}

	server := &Server{
		config: config,
		store:  store,
		//tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(cors.Default())

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"https://foo.com"},
	//	AllowMethods:     []string{"PUT", "PATCH"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return true //origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/h", HealthCheck)

	router.GET("/users", server.getUser)
	router.POST("/users", server.createUser)
	router.POST("/usersV2", server.createUserV2)

	router.GET("/employee", server.getEmployee)
	router.POST("/employee", server.createEmployee)

	router.GET("/location", server.getLocation)
	router.POST("/location", server.createLocation)
	router.GET("/alllocations", server.getAllLocations)

	router.GET("/device", server.getDevice)
	router.POST("/device", server.createDevice)

	router.GET("/getOverallAnalytics", server.getOverallAnalytics)

	router.GET("/getDayWiseUserActivity", server.getDayWiseUserActivity)
	router.GET("/getPastWorkoutData", server.getPastWorkoutData)
	router.GET("/userActivity", server.getUserActivity)
	router.POST("/userActivity", server.createUserActivity)
	router.POST("/startActivity", server.createStartActicityRecord)
	router.POST("/endActivity", server.createEndActivityRecord)

	router.GET("/checkinActivity", server.getCheckinActivity)
	router.POST("/checkinActivity", server.createCheckinActivity)
	router.POST("/checkinRecord", server.createCheckinRecord)
	router.POST("/checkoutRecord", server.createCheckOutRecord)

	router.GET("/class", server.getClass)
	router.POST("/class", server.createClass)
	router.GET("/getClasses", server.getClasses)
	router.GET("/getClassesForEmployee", server.getClassesForEmployee)
	router.GET("/getUpcomingClasses", server.getUpcomingClasses)

	router.GET("/classCatalogue", server.getClassCatalogue)
	router.POST("/classCatalogue", server.createClassCatalogue)

	router.POST("/login", server.loginUser)

	// enrolment analytics
	router.GET("/classesOfferedAndAttendes", server.getClassesOfferedAndAttendees)
	router.GET("/allClassesOfferedAndAttendes", server.getAllClassesOfferedAndAttendees)
	router.GET("/allClassesOfferedAndAttendesPerWeek", server.getAllClassesOfferedAndAttendeesPerWeek)
	router.GET("/allTopAttendedClass", server.getAllTopAttendedClass)
	router.GET("/classesOfferedAndAttendesPerWeek", server.getClassesOfferedAndAttendeesPerWeek)
	router.GET("/mostPopularHourForClassesOnWeekdays", server.getMostPopularHourForClassesOnWeekdays)
	router.GET("/mostPopularHourForClassesOnWeekends", server.getMostPopularHourForClassesOnWeekends)

	//clubusage analytics
	router.GET("/averageVisitorsPerHourWeekdays", server.getAverageVisitorsPerHourWeekdays)
	router.GET("/averageVisitorsPerHourWeekends", server.getAverageVisitorsPerHourWeekends)
	router.GET("/busiestTimeByHourAndDayOfWeek", server.getBusiestTimeByHourAndDayOfWeek)
	router.GET("/hoursSpentInGymByDay", server.getHoursSpentInGymByDay)

	//revenue analytics
	router.GET("/revenueGenerateByMemberships", server.getRevenueGenerateByMemberships)
	router.GET("/classRevenueGeneratedByLocation", server.getClassRevenueGenerateByLocation)

	//member retention  analytics
	router.GET("/dailyNewMemberEnrolments", server.getDailyNewMemberEnrolments)
	router.GET("/kMostFrequentMembers", server.getKMostFrequentMembers)
	router.GET("/membershipCountsByType", server.getMembershipCountsByType)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
