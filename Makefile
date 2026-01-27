# Nome do executável
APP_NAME=app
MAIN=main.go
PID_FILE=.app_pid

# Regra padrão
all: build run

# Compilar o código Go
build:
	@go build -o $(APP_NAME) $(MAIN)

# Executar o binário em background e salvar o PID
run: build
	@./$(APP_NAME) & echo $$! > $(PID_FILE)

# Para Windows (sem PID_FILE, apenas execução direta)
#run:
#   ./$(APP_NAME).exe

# Parar a aplicação anteri
stop:
	@if [ -f $(PID_FILE) ]; then \
		kill `cat $(PID_FILE)` && rm -f $(PID_FILE); \
		echo "Aplicação parada."; \
	fi

# Observa alterações e recompila/roda automaticamente 
# Oculta mensagens de diretório do make
MAKEFLAGS += --no-print-directory

watch:
	@# Sobe o docker totalmente em silêncio
	@docker-compose up -d --quiet-pull > /dev/null 2>&1
	@echo "Aguardando banco de dados..."
	@# Aguarda o banco sem imprimir pontos ou mensagens extras
	@until docker-compose exec -T postgres pg_isready > /dev/null 2>&1; do \
		sleep 1; \
	done
	@$(MAKE) stop > /dev/null 2>&1
	@$(MAKE) run
	@# Trap limpo: apenas executa a limpeza sem ecoar mensagens
	@trap 'docker-compose down > /dev/null 2>&1; $(MAKE) stop > /dev/null 2>&1; exit 0' INT TERM; \
	while true; do \
		inotifywait -q -e modify -r . --exclude '(\.git|$(APP_NAME)|$(PID_FILE))' > /dev/null 2>&1; \
		$(MAKE) stop > /dev/null 2>&1; \
		$(MAKE) run; \
	done

# Limpar arquivos gerados
clean: stop
	rm -f $(APP_NAME)