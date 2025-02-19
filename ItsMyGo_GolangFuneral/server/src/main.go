package main

import (
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "strings"
    "io/ioutil"
    "net/http"
    "os"
    "os/exec"
    "context"
    "time"
)

type CompileRequest struct {
    Env  map[string]string `json:"env"`
    Code string            `json:"code"`
}

func generateRandomHash() (string, error) {
    bytes := make([]byte, 128)
    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }

    hash := sha256.Sum256(bytes)
    return base64.RawURLEncoding.EncodeToString(hash[:]), nil
}

func ensureDir(dir string) error {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        return os.MkdirAll(dir, os.ModePerm)
    }
    return nil
}

func isDanger(value string, blacklist []string) bool {
	for _, blacklistedWord := range blacklist {
		if strings.Contains(value, blacklistedWord) {
			return true
		}
	}
	return false
}

func mygoooHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var req CompileRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            fmt.Println("Error", err)
            http.Error(w, "Error", http.StatusBadRequest)
            return
        }

        fileHash, err := generateRandomHash()
        if err != nil {
            fmt.Println("Error", err)
            http.Error(w, "Error", http.StatusInternalServerError)
            return
        }

        userFileDir := "./userFile"
        if err := ensureDir(userFileDir); err != nil {
            fmt.Println("Error", err)
            http.Error(w, "Error", http.StatusInternalServerError)
            return
        }

        envFileName := fmt.Sprintf("%s/%s_env.json", userFileDir, fileHash)
        envData, _ := json.Marshal(req.Env)
        if err := ioutil.WriteFile(envFileName, envData, 0644); err != nil {
            fmt.Println("Error", err)
            http.Error(w, "Error", http.StatusInternalServerError)
            return
        }

        codeFileName := fmt.Sprintf("%s/%s.go", userFileDir, fileHash)
        if err := ioutil.WriteFile(codeFileName, []byte(req.Code), 0644); err != nil {
            fmt.Println("Error", err)
            http.Error(w, "Error", http.StatusInternalServerError)
            return
        }

       
        go func() {
      
            ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
            defer cancel()
        
            data, err := ioutil.ReadFile(envFileName)
            if err != nil {
                fmt.Println("Error", err)
                os.Remove(envFileName)
                os.Remove(codeFileName)
                return
            }
        
            var env map[string]string
            if err := json.Unmarshal(data, &env); err != nil {
                fmt.Println("Error", err)
                os.Remove(envFileName)
                os.Remove(codeFileName)
                return
            }
        
            blacklist := []string{"FLAG", "id", "curl", "whoami", "cat", "ls", "cd", "grep", ">", "echo", "bash"}
            for key, value := range env {
                if isDanger(value, blacklist) {
                    fmt.Println("YOU ARE NOT MYGO LOVER!!!YOU TRY TO USE DANGER CHARACTER!!!")
                    os.Remove(envFileName)
                    os.Remove(codeFileName)
                    return
                }
                os.Setenv(key, value)
            }
        
            outputPath := fmt.Sprintf("./userEXE/%s", fileHash)
        
            cmd := exec.CommandContext(ctx, "go", "build", "-o", outputPath, codeFileName)
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            

            if err := cmd.Run(); err != nil {
                if ctx.Err() == context.DeadlineExceeded {
                    fmt.Println("Error", err)
                    http.Error(w, "Error", http.StatusInternalServerError)
                    
                } else {
                    fmt.Println("Error", err)
                    http.Error(w, "Error", http.StatusInternalServerError)

                }
                os.Remove(envFileName)
                os.Remove(codeFileName)
                os.Remove(outputPath)
                return
            }

            fmt.Println("MyGo!!!!!")
            os.Remove(envFileName)
            os.Remove(codeFileName)
            os.Remove(outputPath)
            
        }()
        
    } else {
        http.ServeFile(w, r, "./static/mygolang.html")
    }
}


func mygoHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/mygo.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html")
}

func main() {

    http.HandleFunc("/mygolang", mygoooHandler)
    http.HandleFunc("/itsmygo", mygoHandler) 
    http.HandleFunc("/", indexHandler) 

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    fmt.Println("Server started om port http://localhost:20000")
    if err := http.ListenAndServe("0.0.0.0:20000", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
