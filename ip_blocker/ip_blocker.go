package ip_blocker

import (
    "os/exec"
)

func BlockIP(ip string) {
   // Crea un comando per aggiungere una regola ad iptables
	// - "iptables": il comando principale per gestire le regole del firewall.
	// - "-A INPUT": aggiunge ("-A") una regola alla catena "INPUT".
	//               La catena "INPUT" gestisce tutto il traffico in entrata sul sistema.
	// - "-s <ip>": specifica l'indirizzo IP sorgente da bloccare.
	// - "-j DROP": indica che i pacchetti provenienti dall'IP specificato devono essere ignorati ("DROP").
    cmd := exec.Command("iptables", "-A", "INPUT", "-s", ip, "-j", "DROP")
    cmd.Run()
}
