package main

import (
    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "progetto-ddos/middleware"
    "progetto-ddos/queue"
    "log"
    "time"
)

//def metriche
var (
	blockedIPs = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name:"blocked_ips_total",
			Help:"Numero totale di IP bloccati",
					},
	)
	processedRequest = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name:"processed_request_total",
			Help:"Numero totale di richieste elaborate dalla coda",
					},
	)
)
func init() {
//registra le metriche
	prometheus.MustRegister(blockedIPs)
	prometheus.MustRegister(processedRequest)
}
func ProcessQueue() {
	client := queue.ConnectRedis()
	for{
		//dequeue una richiesta
		request,err := queue.DequeueRequest(client)
		if err == nil && request != ""{
			log.Printf("Processing request: %s\n", request)
			processedRequest.Inc()
		}
		time.Sleep(1 * time.Second)
		}
}
func main() {

	r := gin.Default()
	go ProcessQueue()
	//esponi le metriche a prometheus
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
    // Applica il middleware RateLimiter
    r.Use(middleware.RateLimiter())

    // Endpoint di esempio
    r.GET("/api/data", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Benvenuto al backend Go!"})
    })

   if err :=  r.Run(":8080"); err != nil{
	log.Fatalf("Errore avvio: %v\n",err)
	}
}

