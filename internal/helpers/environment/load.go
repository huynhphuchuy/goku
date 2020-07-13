package environment

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	runtimeEnv := os.Getenv("RUNTIME_ENV")
	if runtimeEnv == "" {
		runtimeEnv = "development"
	}

	if runtimeEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	root := filepath.Join(filepath.Dir(b), "../../..")

	// Overload env
	godotenv.Overload(root+"/env/.env."+runtimeEnv, root+"/.env")

	private, _ := ioutil.ReadFile(root + "/internal/helpers/environment/rsa/rsa")
	os.Setenv("PRIVATE_TOKEN_SECRET", string(private))

	public, _ := ioutil.ReadFile(root + "/internal/helpers/environment/rsa/rsa.pub")
	os.Setenv("PUBLIC_TOKEN_SECRET", string(public))
}
