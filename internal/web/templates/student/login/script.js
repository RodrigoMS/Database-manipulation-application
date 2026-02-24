// Função principal para gerenciar o envio do formulário
function setupFormSubmission() {
  // Encontrar o formulário e botão
  const loginForm = document.querySelector(".login-form");
  const submitButton = document.querySelector("#submitButton");

  if (!submitButton || !loginForm) {
    console.error("Elementos do formulário não encontrados");
    return;
  }

  // Adicionar evento de submit ao botão
  submitButton.addEventListener("click", handleFormSubmit);

  // Adicionar eventos de foco/blur aos inputs para classes BEM
  const inputs = document.querySelectorAll(".form-field__input");
  inputs.forEach((input) => {
    // Evento para submit com Enter
    input.addEventListener("keypress", (event) => {
      if (event.key === "Enter") {
        handleFormSubmit();
      }
    });

    // Eventos para adicionar/remover classe de foco
    input.addEventListener("focus", () => {
      input.classList.add("form-field__input--focus");
    });

    input.addEventListener("blur", () => {
      input.classList.remove("form-field__input--focus");
    });
  });
}

// Função para validar os dados do formulário
function validateFormData(formData) {
  const errors = [];

  // Verificar se pelo menos um aluno foi preenchido
  if (!formData.student_1?.trim() && !formData.student_2?.trim()) {
    errors.push("Pelo menos um aluno deve ser informado");
  }

  // Validar comprimento dos nomes
  if (formData.student_1 && formData.student_1.length > 100) {
    errors.push("Nome do aluno 1 é muito longo");
  }

  if (formData.student_2 && formData.student_2.length > 100) {
    errors.push("Nome do aluno 2 é muito longo");
  }

  // Validar caracteres (permitir letras, espaços, acentos e hífens)
  const nameRegex = /^[a-zA-ZÀ-ÿ\s\-']+$/;

  if (formData.student_1 && !nameRegex.test(formData.student_1)) {
    errors.push("Nome do aluno 1 contém caracteres inválidos");
  }

  if (formData.student_2 && !nameRegex.test(formData.student_2)) {
    errors.push("Nome do aluno 2 contém caracteres inválidos");
  }

  return {
    isValid: errors.length === 0,
    errors,
  };
}

// Função para coletar dados do formulário
function collectFormData() {
  const student1Input = document.querySelector("#student_1");
  const student2Input = document.querySelector("#student_2");

  return {
    student_1: student1Input?.value || "",
    student_2: student2Input?.value || "",
  };
}

// Função para exibir mensagens de feedback
function showMessage(message, type = "error") {
  // Remover mensagens anteriores
  const existingMessage = document.querySelector(".form-message");
  if (existingMessage) {
    existingMessage.remove();
  }

  if (!message) return;

  const messageDiv = document.createElement("div");
  messageDiv.className = `form-message form-message--${type}`;
  messageDiv.textContent = message;

  const loginForm = document.querySelector(".login-form");
  loginForm.insertBefore(messageDiv, loginForm.firstChild.nextSibling);

  // Remover mensagem após alguns segundos (apenas para sucesso)
  if (type === "success") {
    setTimeout(() => {
      messageDiv.remove();
    }, 3000);
  }
}

// Função para mostrar loading
function showLoading(show) {
  const submitButton = document.querySelector("#submitButton");

  if (show) {
    submitButton.disabled = true;
    submitButton.textContent = "Enviando...";
    submitButton.classList.add("submit-button--disabled");
    submitButton.classList.remove(
      "submit-button--hover",
      "submit-button--active",
    );
  } else {
    submitButton.disabled = false;
    submitButton.textContent = "Entrar";
    submitButton.classList.remove("submit-button--disabled");
    submitButton.classList.add("submit-button--hover", "submit-button--active");
  }
}

// Função para validar e adicionar classes de validação aos inputs
function applyValidationStyles(inputElement, isValid) {
  if (isValid) {
    inputElement.classList.remove("form-field__input--error");
    inputElement.classList.add("form-field__input--valid");
  } else {
    inputElement.classList.remove("form-field__input--valid");
    inputElement.classList.add("form-field__input--error");
  }
}

// Adicionar estilos CSS dinâmicos para validação
function addValidationStyles() {
  if (!document.querySelector("#validation-styles")) {
    const style = document.createElement("style");
    style.id = "validation-styles";
    style.textContent = `
      .form-field__input--error {
        border-color: #f44336 !important;
        box-shadow: 0 0 0 2px rgba(244, 67, 54, 0.2) !important;
      }
      
      .form-field__input--valid {
        border-color: #4CAF50 !important;
      }
    `;
    document.head.appendChild(style);
  }
}

// Função principal de envio do formulário
async function handleFormSubmit() {
  try {
    const formData = collectFormData();
    const validation = validateFormData(formData);

    const student1Input = document.querySelector("#student_1");
    const student2Input = document.querySelector("#student_2");

    applyValidationStyles(
      student1Input,
      formData.student_1 &&
        formData.student_1.length <= 100 &&
        /^[a-zA-ZÀ-ÿ\s\-']+$/.test(formData.student_1),
    );

    if (formData.student_2) {
      applyValidationStyles(
        student2Input,
        formData.student_2.length <= 100 &&
          /^[a-zA-ZÀ-ÿ\s\-']+$/.test(formData.student_2),
      );
    }

    if (!validation.isValid) {
      showMessage(validation.errors.join(". "), "error");
      return;
    }

    showLoading(true);

    const response = await fetch(`/student-authentication`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        ...formData,
        className: className,
      }),
      signal: AbortSignal.timeout(10000),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Erro do servidor: ${response.status} - ${errorText}`);
    }

    // Resultado
    const result = await response.json();
    showMessage(result.message || "Login realizado com sucesso!", "success");

    setTimeout(() => {
      student1Input.value = "";
      student2Input.value = "";
      student1Input.classList.remove("form-field__input--valid");
      student2Input.classList.remove("form-field__input--valid");
    }, 1000);

    if (result.redirect) {
      window.location.href = result.redirect;
    }

    console.log("Resposta do servidor:", result);

    const event = new CustomEvent("formSubmittedSuccessfully", {
      detail: { formData, serverResponse: result },
    });
    document.dispatchEvent(event);
  } catch (error) {
    console.error("Erro ao enviar formulário:", error);
    let errorMessage = "Erro ao enviar formulário. ";
    if (error.name === "AbortError") {
      errorMessage += "Tempo de conexão excedido.";
    } else if (error.name === "TypeError" && error.message.includes("fetch")) {
      errorMessage += "Não foi possível conectar ao servidor.";
    } else {
      errorMessage += error.message;
    }
    showMessage(errorMessage, "error");
  } finally {
    showLoading(false);
  }
}

// Função para melhorar a UX dos inputs
function enhanceInputUX() {
  // Adicionar estilos de validação
  addValidationStyles();

  const inputs = document.querySelectorAll(".form-field__input");

  inputs.forEach((input) => {
    // Adicionar validação em tempo real
    input.addEventListener("input", () => {
      const value = input.value;
      if (value) {
        const isValid =
          value.length <= 100 && /^[a-zA-ZÀ-ÿ\s\-']+$/.test(value);
        applyValidationStyles(input, isValid);
      } else {
        input.classList.remove(
          "form-field__input--error",
          "form-field__input--valid",
        );
      }
    });

    // Limitar comprimento
    input.maxLength = 100;
  });
}

// Inicialização quando o DOM estiver carregado
document.addEventListener("DOMContentLoaded", () => {
  setupFormSubmission();
  enhanceInputUX();

  // Adicionar fallback para navegadores antigos
  if (!AbortSignal.timeout) {
    AbortSignal.timeout = function (ms) {
      const controller = new AbortController();
      setTimeout(
        () =>
          controller.abort(new DOMException("TimeoutError", "TimeoutError")),
        ms,
      );
      return controller.signal;
    };
  }
});

// Exportar funções para uso em módulos (se necessário)
if (typeof module !== "undefined" && module.exports) {
  module.exports = {
    setupFormSubmission,
    validateFormData,
    collectFormData,
    handleFormSubmit,
  };
}
