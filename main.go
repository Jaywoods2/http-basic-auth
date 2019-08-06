package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	podNamespace  string
	configMapPath *string
	data          map[string]string
}

var config Config

func main() {

	podNamespace := os.Getenv("POD_NAMESPACE")
	if podNamespace == "" {
		podNamespace = "default"
	}
	configMapPath := flag.String("configmap", fmt.Sprintf("%s/accounts", podNamespace), "指定认证配置")
	log.Println(fmt.Sprintf("read config from %s", *configMapPath))
	config = Config{
		podNamespace:  podNamespace,
		configMapPath: configMapPath,
	}
	var ch1 = make(chan map[string]string)
	go loadConfig(ch1)
	config.data = <-ch1
	log.Println("Configure reload success...")

	//go func() {
	//	for {
	//		d := <-ch1
	//		config.mu.Lock()
	//		config.data = d
	//		config.mu.Unlock()
	//		log.Println("Configure reload success...")
	//		for k, v := range d {
	//			fmt.Println(fmt.Sprintf("%s=%s", k, v))
	//		}
	//	}
	//
	//}()

	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(config.data))

	authorized.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Status": "Ok",
		})
	})

	if err := r.Run(); err != nil {
		panic(err.Error())
	}

}

func loadConfig(ch chan map[string]string) {

	c, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	cs, err := kubernetes.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}

	t := strings.Split(*config.configMapPath, "/")
	//for {
	configMap, err := cs.CoreV1().ConfigMaps(t[0]).Get(t[1], v1.GetOptions{})
	if err != nil {
		log.Println("获取configmap失败")
		log.Fatalln(err)
	}
	ch <- configMap.Data
	//if strings.Compare(config.version, configMap.ResourceVersion) != 0 {
	//	ch <- configMap.Data
	//	config.mu.Lock()
	//	config.version = configMap.ResourceVersion
	//	config.mu.Unlock()
	//}

	//time.Sleep(10 * time.Second)
	//}

}
