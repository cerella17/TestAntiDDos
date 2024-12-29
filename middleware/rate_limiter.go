package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "sync"
    "time"
    "strings"
    "progetto-ddos/queue"   
    "progetto-ddos/ip_blocker"
)

var (
    requestCounts = make(map[string]int) //traccia il numero di richieste per IP
    mu            sync.Mutex //protegge l'accesso a requestCounts
    whitelist = map[string]bool{
		"127.0.0.1":true,}
    blacklist = map[string]bool{}
)

func getRealIP(c *gin.Context) string {
	xff := c.GetHeader("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0]) // Prendi il primo IP
	}
	return c.ClientIP() // Torna all'IP predefinito
}

func RateLimiter() gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := getRealIP(c)
	if whitelist[ip]{
		c.Next() //accesso senza limiti
		return
	}
	if blacklist[ip]{
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error":"Accesso negato. IP bloccato."})
		return
	}
        mu.Lock()
        requestCounts[ip]++
        count := requestCounts[ip]
        mu.Unlock()

        if count > 10 {//richieste al minuto
		blacklist[ip] = true //aggiungi ip alla  bl
		ip_blocker.BlockIP(ip)
		enqueueErr := queue.EnqueueRequest(queue.ConnectRedis(), c.Request.URL.Path)//metti la  richiesta in coda           
                if enqueueErr != nil{
                   c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Errore durante l'inserimento in coda."})
		   return
		}
	    c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Richieste eccessive. IP bloccato."})
	    return
        }

        go func() {
	 time.Sleep(1 * time.Minute)
	 mu.Lock()
	 requestCounts[ip]-- 
	 mu.Unlock()
        }()

        c.Next()
    }
}
