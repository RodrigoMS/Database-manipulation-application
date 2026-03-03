package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/RodrigoMS/app/cmd/internal/domain"
	"github.com/RodrigoMS/app/cmd/internal/web"
)

// Estrutura para receber os dados do login.
type LoginRequest struct {
    Student1 string `json:"student_1"`
    Student2 string `json:"student_2,omitempty"`
    ClassName string `json:"className"`
}

// Carrega o template html do login dos alunos.
func StudentLogin(w http.ResponseWriter, r *http.Request) {
    pathParts := strings.Split(r.URL.Path, "/")
    class := ""
    if len(pathParts) > 2 {
        class = pathParts[2] // "5A"
    }

    // Passa a string diretamente
    web.RenderTemplate(w, "StudentLogin", class)
}

// Autentica o aluno e gera um token de sessão.
// Estrutura que representa os dados recebidos no corpo da requisição
func StudentAuthentication(w http.ResponseWriter, r *http.Request) {
    var request LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, `{"error":"invalid_json"}`, http.StatusBadRequest)
        return
    }

    class := request.ClassName

    // Aqui você faria a validação real
    if domain.ValidateStudent(request.Student1, class) && 
       (request.Student2 == "" || domain.ValidateStudent(request.Student2, class)) {
        student := domain.Student{
            Name:  request.Student1,
            Class: class,
        }

        // Monta os dados da sessão
        sessionData := map[string]string{
            "student1": request.Student1,
            "student2": request.Student2,
            "class":    class,
        }

        // Converte para JSON e depois para Base64
        sessionValueBytes, _ := json.Marshal(sessionData)
        sessionValue := base64.StdEncoding.EncodeToString(sessionValueBytes)

        // Cria o cookie
        cookie := &http.Cookie{
            Name:     "student_session",
            Value:    sessionValue,
            Path:     "/",
            Expires:  time.Now().Add(1 * time.Hour),
            HttpOnly: true,
            Secure:   true,
            SameSite: http.SameSiteStrictMode,
        }
        http.SetCookie(w, cookie)

        // Resposta JSON
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message":  "Login realizado com sucesso",
            "student":  student,
            "partner":  request.Student2,
            "redirect": "/student-dashboard",
        })
    } else {
        http.Error(w, `{"error":"invalid_credentials"}`, http.StatusUnauthorized)
    }
}

// Função para exibir o dashboard
func StudentDashboard(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("student_session")
    if err != nil || cookie.Value == "" {
        http.Redirect(w, r, "/student-login", http.StatusSeeOther)
        return
    }

    // Decodifica Base64
    decoded, err := base64.StdEncoding.DecodeString(cookie.Value)
    if err != nil {
        http.Redirect(w, r, "/student-login", http.StatusSeeOther)
        return
    }

    // Decodifica JSON
    var sessionData map[string]string
    if err := json.Unmarshal(decoded, &sessionData); err != nil {
        http.Redirect(w, r, "/student-login", http.StatusSeeOther)
        return
    }

    student1 := sessionData["student1"]
    student2 := sessionData["student2"]
    class := sessionData["class"] 

    /*if student2 != "" {
        fmt.Printf("Alunos autenticados: %s e %s da turma %s\n", student1, student2, class)
    } else {
        fmt.Printf("Aluno autenticado: %s da turma %s\n", student1, class)
    }*/

    data := LoginRequest{
        Student1: student1,
        Student2: student2,
        ClassName: class,
    }

    web.RenderTemplate(w, "StudentDashboard", data)
}


func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Pega o cookie atual
    //cookie, err := r.Cookie("student_session")
    //if err == nil {
        // 2. Invalida a sessão no servidor (banco de dados, redis, etc)
        //invalidateSession(cookie.Value)
    //}

    pathParts := strings.Split(r.URL.Path, "/")
    className := ""
    if len(pathParts) > 2 {
        className = pathParts[2]
    }
    
    // 3. Remove o cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "student_session",
        Value:    "",
        Path:     "/",
        MaxAge:   -1,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteStrictMode,
    })
    
    http.Redirect(w, r, "/student-login/"+className, http.StatusSeeOther)
}


//pkg
var secretKey = []byte("chave-super-secreta") // guarde em variável de ambiente

// Gera token assinado
func generateToken(name, class string) string {
    exp := time.Now().Add(1 * time.Hour).Unix()
    payload := fmt.Sprintf("%s|%s|%d", name, class, exp)

    h := hmac.New(sha256.New, secretKey)
    h.Write([]byte(payload))
    signature := base64.URLEncoding.EncodeToString(h.Sum(nil))

    return base64.URLEncoding.EncodeToString([]byte(payload)) + "." + signature
}

// Valida token
func validateToken(token string) (string, string, bool) {
    parts := strings.Split(token, ".")
    if len(parts) != 2 {
        return "", "", false
    }

    payloadBytes, err := base64.URLEncoding.DecodeString(parts[0])
    if err != nil {
        return "", "", false
    }
    payload := string(payloadBytes)

    h := hmac.New(sha256.New, secretKey)
    h.Write([]byte(payload))
    expectedSig := base64.URLEncoding.EncodeToString(h.Sum(nil))

    if !hmac.Equal([]byte(expectedSig), []byte(parts[1])) {
        return "", "", false
    }

    // Extrai dados
    fields := strings.Split(payload, "|")
    if len(fields) != 3 {
        return "", "", false
    }
    name := fields[0]
    class := fields[1]
    exp, _ := strconv.ParseInt(fields[2], 10, 64)

    if time.Now().Unix() > exp {
        return "", "", false
    }

    return name, class, true
}