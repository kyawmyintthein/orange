# Orange (Beta)
### RESTful framework

**Orange** is  a [httprouter](https://github.com/julienschmidt/httprouter) based framework for api only. <strong> Org-cli </strong> is a tool to manage  *Orange*  projects and to generate project skeleton.


**Features**
 - Lightweight framework
 - RESTful only
 - Middleware
 - Configuration
 
**Example**

> server.go

   
    package main
    import "github.com/kyawmyintthein/orange"
    import "net/http"
  
    var App *orange.App
    var ns_v1 *orange.Router
    var config *orange.Config
    
    func main() {
    	App.Start(config.GetString("app.dev.address"))
    }
    
    func init() {
    	App = orange.NewApp("Test")
		config = App.AppConfig()	

    	ns_v1 = App.Namespace("/v1")
    	var objectController = ns_v1.Controller("/objects")
    	
    	objectController.GET("/", func(ctx *orange.Context) {
    		ctx.JSON(http.StatusOK, map[string]interface{}{"Object": "Value"})
    	})
    
    	objectController.GET("/:name", func(ctx *orange.Context) {
    		name := ctx.Param("name")
    		ctx.JSON(http.StatusOK, map[string]interface{}{"name": name})
    	})
    }

> applicaiton.yaml

    app:
      name: "App Name"
      envs: ["test", "dev", "prod"]
      env: "dev"
      version: 1.0.0
      dev: 
        address: localhost:3000
