package router

import (
	"pkg/controllers"
	"pkg/middleware/jwt"
	logger "pkg/middleware/log"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	r.SetTrustedProxies(nil)
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:  []string{"http://222.186.61.53/","http://localhost/","http://127.0.0.1/","http://150.158.107.119/",},
	// 	AllowMethods:  []string{"POST", "GET"},
	// 	AllowHeaders:  []string{"Content-Type", "Authorization"},
	// 	ExposeHeaders: []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge: 50 * time.Second,
	// }))
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	r.GET("/ws", controllers.ChatController{}.Handler)
	user := r.Group("/user")
	{
		user.POST("/create", controllers.UserController{}.CreateUser)
		user.POST("/login", controllers.UserController{}.UserLogin)
		user.GET("/logout", controllers.UserController{}.UserLogout)
	}
	r.Use(jwt.Auth())
	{
		user.POST("/refresh", controllers.UserController{}.GetAccessToken)
		user.GET("/list", controllers.UserController{}.GetList)
		user.GET("/:id", controllers.UserController{}.GetUserInfo)
	}
	problem := r.Group("/problem")
	{

		problem.GET("/list", controllers.ProblemController{}.GetList)
		problem.POST("/create", controllers.ProblemController{}.CreateProblem)
		problem.POST("/submit", controllers.ProblemController{}.SubmitProblem)
		problem.GET("/:id", controllers.ProblemController{}.GetProblemInfo)
	}
	record := r.Group("/record")
	{
		record.GET("/list", controllers.RecodeController{}.GetList)
		record.GET("/:id", controllers.RecodeController{}.GetRecodeInfo)
	}
	ide := r.Group("/ide")
	{
		ide.POST("/submit", controllers.IDEController{}.JudgeCode)
	}
	blog := r.Group("/blog")
	{
		blog.POST("/create", controllers.BlogController{}.CreateBlog)
		blog.GET("/list", controllers.BlogController{}.GetList)
		blog.POST("/change", controllers.BlogController{}.ChangeBlog)
		blog.GET("/delete", controllers.BlogController{}.DeleteBlog)
		blog.GET("/:id", controllers.BlogController{}.GetBlogInfo)
	}
	comment := r.Group("/comment")
	{
		comment.GET("", controllers.CommentController{}.GetComment)
		comment.POST("/main", controllers.CommentController{}.CreateComment)
		comment.POST("/sub", controllers.CommentController{}.CreateSubComment)
	}
	training := r.Group("/training")
	{
		training.GET("/list", controllers.TrainingController{}.GetList)
	}
	contest := r.Group("/contest")
	{
		contest.POST("/list", controllers.ContestController{}.GetList)
	}
	sv1 := r.Group("/auth/")
	{
		sv1.GET("/time", controllers.GetDataByTime)
	}

	return r
}
