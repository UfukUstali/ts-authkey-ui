package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const port = "8924"

type PageData struct {
	Error   string
	Success bool
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="de">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tailscale Auth Key</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 500px;
            margin: 100px auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background: white;
            padding: 30px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 {
            color: #333;
            text-align: center;
            margin-bottom: 30px;
        }
        .form-group {
            margin-bottom: 20px;
        }
        label {
            display: block;
            margin-bottom: 5px;
            color: #555;
            font-weight: bold;
        }
        input[type="text"] {
            width: 100%;
            padding: 12px;
            border: 1px solid #ddd;
            border-radius: 4px;
            font-size: 16px;
            box-sizing: border-box;
        }
        button {
            width: 100%;
            padding: 12px;
            background-color: #007cba;
            color: white;
            border: none;
            border-radius: 4px;
            font-size: 16px;
            cursor: pointer;
        }
        button:hover {
            background-color: #005a87;
        }
        .error {
            color: #d32f2f;
            background-color: #ffebee;
            padding: 10px;
            border-radius: 4px;
            margin-bottom: 20px;
        }
        .success {
            color: #2e7d32;
            background-color: #e8f5e8;
            padding: 10px;
            border-radius: 4px;
            margin-bottom: 20px;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Tailscale Authentifizierung</h1>
        
        {{if .Error}}
            <div class="error">{{.Error}}</div>
        {{end}}
        
        {{if .Success}}
            <div class="success">
                Tailscale wurde erfolgreich gestartet! Diese Anwendung wird sich in wenigen Sekunden beenden.
                <script>
                    setTimeout(function() {
                        window.close();
                    }, 3000);
                </script>
            </div>
        {{else}}
            <form method="POST" action="/submit">
                <div class="form-group">
                    <label for="authkey">Auth Key:</label>
                    <input type="text" id="authkey" name="authkey" required 
                           placeholder="Geben Sie Ihren Tailscale Auth Key ein">
                </div>
                <button type="submit">Tailscale starten</button>
            </form>
        {{end}}
    </div>
</body>
</html>
`

func checkTailscale() error {
	_, err := exec.LookPath("tailscale")
	return err
}

func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	exec.Command(cmd, args...).Start()
}

func main() {
	if err := checkTailscale(); err != nil {
		log.Fatal("Tailscale ist nicht installiert oder nicht im PATH verfügbar")
	}

	tmpl := template.Must(template.New("index").Parse(htmlTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		authKey := strings.TrimSpace(r.FormValue("authkey"))
		if authKey == "" {
			data := PageData{Error: "Auth Key darf nicht leer sein"}
			tmpl.Execute(w, data)
			return
		}

		args := []string{"up", "--auth-key", authKey}
		if runtime.GOOS == "windows" {
			args = append(args, "--unattended")
		}
		cmd := exec.Command("tailscale", args...)
		output, err := cmd.CombinedOutput()

		if err != nil {
			errorMsg := fmt.Sprintf("Fehler beim Starten von Tailscale: %s", string(output))
			data := PageData{Error: errorMsg}
			tmpl.Execute(w, data)
			return
		}

		data := PageData{Success: true}
		tmpl.Execute(w, data)

		go func() {
			time.Sleep(1 * time.Second)
			os.Exit(0)
		}()
	})

	url := fmt.Sprintf("http://localhost:%s", port)
	fmt.Printf("Server läuft auf %s\n", url)

	go openBrowser(url)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
