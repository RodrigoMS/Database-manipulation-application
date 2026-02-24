package web

import (
	"bytes"
	"embed"
	//"encoding/base64"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"

	//"regexp"
	"strings"
)

//go:embed templates/*
var templateFiles embed.FS

var templates *template.Template

var filePath map[string]string

func init() {
    readTemplatesFiles("templates")
}

// Buscar todos os arquivos .html em templates e subpastas
func readTemplatesFiles(folder string) {
    var files []string
    fs.WalkDir(templateFiles, folder, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        // if !d.IsDir() && (path[len(path)-5:] == ".html") {
        if !d.IsDir() && strings.HasSuffix(path, ".html") {
            files = append(files, path)
        }
        return nil
    })

    if len(files) == 0 {
        panic("Nenhum arquivo .html encontrado em templates")
    }

    getPathFile(&files) 

    templates = template.Must(template.ParseFS(templateFiles, files...))

}

// Obtem o caminho do arquivo HTMl para ser utilizado para obter o CSS e JS.
func getPathFile(files *[]string) {
    filePath = make(map[string]string)

    for _, file := range *files {
        t, err := template.ParseFS(templateFiles, file)
        if err != nil {
            panic(err)
        }
        for _, tmpl := range t.Templates() {
            // Aqui usamos apenas o diretório, sem o nome do arquivo.
            dir := filepath.Dir(file)
            filePath[tmpl.Name()] = dir
            //fmt.Println("Template:", tmpl.Name(), "vem de", dir)
        }
    }
}

// Coloca tudo em uma única linha e remove comentários.
func minifyText(text string) string {
    // Remove comentários de linha // ...
    //reLine := regexp.MustCompile(`//.*`)
    //text = reLine.ReplaceAllString(text, "")

    // Remove comentários de bloco /* ... */
    //reBlock := regexp.MustCompile(`/\*.*?\*/`)
    //text = reBlock.ReplaceAllString(text, "")

    // Remove quebras de linha e tabs.
    //text = strings.ReplaceAll(text, "\n", "")
    //text = strings.ReplaceAll(text, "\r", "")
    //text = strings.ReplaceAll(text, "\t", "")

    // Remove espaços duplicados.
    //text = strings.Join(strings.Fields(text), " ")

    // Codifica em Base64 
    /*encoded := base64.StdEncoding.EncodeToString([]byte(text))

    return encoded*/

    return text
}

// Lê e retorna o conteúdo dos arquivos do mesmo diretório do template.
func readAssetFile(templateName string, fileExtension string) (string, error) {
    path, ok := filePath[templateName]
    if !ok {
        return "", fmt.Errorf("template %s não encontrado", templateName)
    }

    // files, err := os.ReadDir(path)
    files, err := fs.ReadDir(templateFiles, path)
    if err != nil {
        return "", fmt.Errorf("erro ao abrir diretório %s: %w", path, err)
    }

    var contentFile strings.Builder
    for _, d := range files {
        if !d.IsDir() && strings.HasSuffix(d.Name(), fileExtension) {
            // content, err := os.ReadFile(filepath.Join(path, d.Name()))
            content, err := templateFiles.ReadFile(filepath.Join(path, d.Name()))
            if err != nil {
                return "", fmt.Errorf("erro ao ler %s: %w", d.Name(), err)
            }
            contentFile.WriteString(string(content))
            contentFile.WriteString("\n")
        }
    }
    return minifyText(contentFile.String()), nil
}

// Busca o template, adiciona o estilo CSS e o script JS e os dados no template.
func RenderTemplate(w http.ResponseWriter, templateName string, data any) {
    css, err := readAssetFile(templateName, ".css")
    if err != nil {
        fmt.Println("Erro ao carregar CSS:", err)
        http.Error(w, "Ocorreu um problema interno no servidor. Tente novamente mais tarde.", http.StatusInternalServerError)

        return
    }

    js, err := readAssetFile(templateName, ".js")
    if err != nil {
        fmt.Println("Erro ao carregar JS:", err)
        http.Error(w, "Ocorreu um problema interno no servidor. Tente novamente mais tarde.", http.StatusInternalServerError)

        return
    }

    content := struct {
        CSS  template.CSS
        JS   template.JS
        Data any
    }{
        CSS:  template.CSS(css),
        JS:   template.JS(js),
        Data: data,
    }

    var buf bytes.Buffer
    if err := templates.ExecuteTemplate(&buf, templateName, content); err != nil {
        buf.Reset()
        // sobrescreve err com o resultado da segunda tentativa
        if err = templates.ExecuteTemplate(&buf, "Error500", nil); err != nil {
            http.Error(w, "ERRO 500\n\nOcorreu um problema interno no servidor. \nTente novamente mais tarde.", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusInternalServerError)
    }
    buf.WriteTo(w)
}