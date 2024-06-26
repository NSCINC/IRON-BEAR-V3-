{\rtf1\ansi\deff0\nouicompat{\fonttbl{\f0\fnil\fcharset0 Calibri;}}
{\*\generator Riched20 10.0.19041}\viewkind4\uc1 
\pard\sa200\sl276\slmult1\f0\fs22\lang22 Aqui est\'e1 um exemplo simples de um controlador em Go para autenticar usu\'e1rios em um banco de dados MySQL usando o Azure:\par
\par
```go\par
package main\par
\par
import (\par
    "database/sql"\par
    "fmt"\par
    "log"\par
    "net/http"\par
\par
    _ "github.com/go-sql-driver/mysql"\par
)\par
\par
// Defina as credenciais do banco de dados MySQL\par
const (\par
    DBUsername = "seu_usuario"\par
    DBPassword = "sua_senha"\par
    DBHost     = "seu_host"\par
    DBPort     = "sua_porta"\par
    DBName     = "seu_banco_de_dados"\par
)\par
\par
// Estrutura para armazenar informa\'e7\'f5es do usu\'e1rio\par
type User struct \{\par
    Username string\par
    Password string\par
\}\par
\par
// Fun\'e7\'e3o para autenticar o usu\'e1rio no banco de dados MySQL\par
func authenticateUser(username, password string) (bool, error) \{\par
    // Conecte-se ao banco de dados MySQL\par
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName))\par
    if err != nil \{\par
        return false, err\par
    \}\par
    defer db.Close()\par
\par
    // Consulte o banco de dados para verificar as credenciais do usu\'e1rio\par
    var storedPassword string\par
    err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)\par
    if err != nil \{\par
        if err == sql.ErrNoRows \{\par
            return false, nil // Usu\'e1rio n\'e3o encontrado\par
        \}\par
        return false, err\par
    \}\par
\par
    // Verifique se a senha fornecida corresponde \'e0 senha armazenada no banco de dados\par
    if storedPassword == password \{\par
        return true, nil // Autentica\'e7\'e3o bem-sucedida\par
    \}\par
\par
    return false, nil // Senha incorreta\par
\}\par
\par
// Handler para autenticar o usu\'e1rio\par
func authenticateHandler(w http.ResponseWriter, r *http.Request) \{\par
    // Extrair informa\'e7\'f5es de autentica\'e7\'e3o do corpo da solicita\'e7\'e3o\par
    username := r.FormValue("username")\par
    password := r.FormValue("password")\par
\par
    // Autenticar o usu\'e1rio\par
    authenticated, err := authenticateUser(username, password)\par
    if err != nil \{\par
        http.Error(w, err.Error(), http.StatusInternalServerError)\par
        return\par
    \}\par
\par
    if authenticated \{\par
        // Usu\'e1rio autenticado com sucesso\par
        fmt.Fprintf(w, "Usu\'e1rio autenticado com sucesso!")\par
    \} else \{\par
        // Falha na autentica\'e7\'e3o\par
        fmt.Fprintf(w, "Falha na autentica\'e7\'e3o. Verifique suas credenciais.")\par
    \}\par
\}\par
\par
func main() \{\par
    // Registrar o manipulador de rota\par
    http.HandleFunc("/authenticate", authenticateHandler)\par
\par
    // Iniciar o servidor na porta 8080\par
    log.Fatal(http.ListenAndServe(":8080", nil))\par
\}\par
```\par
\par
\par
}
 
