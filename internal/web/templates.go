package web

import (
	"embed"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

//go:embed templates/*
var templateFiles embed.FS

var templates *template.Template

// Mapeamento de nome do template para caminho do arquivo
var templatePathMap = make(map[string]string) // key: nome do template, value: caminho completo

// Cache para armazenar CSS e JS por pasta de template
var cssCache = make(map[string]template.CSS) // key: caminho da pasta
var jsCache = make(map[string]template.JS)   // key: caminho da pasta

// Executado automaticamente no runtime do Go ao iniciar o programa.
func init() {
    templates = template.New("") // cria um conjunto vazio
    readTemplates("templates")
}

func readTemplates(folder string) {
    entries, err := templateFiles.ReadDir(folder)
    if err != nil {
        panic(err)
    }
    
    for _, e := range entries {
        fullPath := folder + "/" + e.Name()
        
        if e.IsDir() {
            readTemplates(fullPath) // entra recursivamente
        } else if strings.HasSuffix(e.Name(), ".html") {
            // Extrai o diretório do template
            templateDir := filepath.Dir(fullPath)
            
            // Tenta carregar CSS da mesma pasta
            cssContent := loadAssetFromDir(templateDir, ".css")
            if cssContent != "" {
                cssCache[templateDir] = template.CSS(cssContent)
            }
            
            // Tenta carregar JS da mesma pasta
            jsContent := loadAssetFromDir(templateDir, ".js")
            if jsContent != "" {
                jsCache[templateDir] = template.JS(jsContent)
            }
            
            // Carrega o arquivo .html como template
            content, err := templateFiles.ReadFile(fullPath)
            if err != nil {
                panic(err)
            }
            
            // Parse o template para extrair o nome definido
            tmpl, err := templates.New(fullPath).Parse(string(content))
            if err != nil {
                panic(err)
            }
            
            // Encontra o nome do template definido (como "StudentLogin")
            for _, t := range tmpl.Templates() {
                if t.Name() != fullPath && !strings.HasPrefix(t.Name(), "templates/") {
                    // Este é o nome definido no {{define "Nome"}}
                    templateName := t.Name()
                    // Armazena o mapeamento
                    templatePathMap[templateName] = fullPath
                }
            }
        }
    }
}

// Função para carregar assets (CSS/JS) do mesmo diretório do template
func loadAssetFromDir(dir, extension string) string {
    // Lista de nomes possíveis para o arquivo
    possibleNames := []string{
        "style" + extension,
        "main" + extension,
        "app" + extension,
        "script" + extension,
        "index" + extension,
    }
    
    // Tenta cada nome possível
    for _, name := range possibleNames {
        assetPath := dir + "/" + name
        content, err := templateFiles.ReadFile(assetPath)
        if err == nil {
            // Encontrou o arquivo!
            return string(content)
        }
    }
    
    return ""
}

// Estrutura de dados que será passada para os templates
type PageData struct {
    Style  template.CSS
    Script template.JS
    Data   any
}

func RenderTemplate(w http.ResponseWriter, templateName string, data any) {
    // Obtém o caminho do arquivo a partir do nome do template
    filePath, exists := templatePathMap[templateName]
    if !exists {
        // Tenta usar o nome como caminho direto (backward compatibility)
        filePath = templateName
    }
    
    // Obtém o diretório do template
    templateDir := filepath.Dir(filePath)
    
    // Obtém CSS e JS do cache (ou usa padrão se não existir)
    css, hasCSS := cssCache[templateDir]
    if !hasCSS {
        css = template.CSS("body { background-color: #1c1c1c; color: #fff; }")
    }
    
    js, hasJS := jsCache[templateDir]
    if !hasJS {
        js = template.JS("alert('Não foi possível carregar os arquivos javascript !');")
    }
    
    pd := PageData{
        Style:  css,
        Script: js,
        Data:   data,
    }
    
    // Usa o NOME do template (não o caminho)
    err := templates.ExecuteTemplate(w, templateName, pd)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Função para obter o caminho do arquivo de um template pelo nome
func GetTemplatePath(templateName string) (string, bool) {
    path, ok := templatePathMap[templateName]
    return path, ok
}

// Função para obter CSS de um template específico
func GetTemplateCSS(templateName string) template.CSS {
    if path, ok := templatePathMap[templateName]; ok {
        dir := filepath.Dir(path)
        if css, ok := cssCache[dir]; ok {
            return css
        }
    }
    return template.CSS("")
}

// Função para obter JS de um template específico
func GetTemplateJS(templateName string) template.JS {
    if path, ok := templatePathMap[templateName]; ok {
        dir := filepath.Dir(path)
        if js, ok := jsCache[dir]; ok {
            return js
        }
    }
    return template.JS("")
}

// Lista todos os templates disponíveis
func ListTemplates() []string {
    var list []string
    for name := range templatePathMap {
        list = append(list, name)
    }
    return list
}