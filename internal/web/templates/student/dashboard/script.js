// Dados das atividades
const activities = Array.from(document.querySelectorAll(".nav-menu__item")).map(
  (item, index) => {
    const link = item.querySelector(".nav-menu__link");
    const scoreElement = item.querySelector(".nav-menu__score");
    const checkElement = item.querySelector(".nav-menu__check");

    let title = "";
    if (link.getAttribute("data-activity")) {
      title = link.getAttribute("data-activity");
    } else {
      // Pega o texto do span que contém o nome (ignorando ícone e pontuação)
      const spans = link.querySelectorAll("span");
      for (let i = 0; i < spans.length; i++) {
        if (
          !spans[i].classList.contains("nav-menu__icon") &&
          !spans[i].classList.contains("nav-menu__score") &&
          !spans[i].classList.contains("nav-menu__check")
        ) {
          title = spans[i].textContent;
          break;
        }
      }
    }

    return {
      id: index + 1,
      element: item,
      link: link,
      title: title,
      score: parseInt(item.getAttribute("data-score")) || 0,
      hasScore: !checkElement || checkElement.style.display === "none",
      completed: false,
    };
  },
);

let currentActivityIndex = 0;
let totalScore = 0;

// Elementos DOM
const currentTitle = document.getElementById("currentActivityTitle");
const totalScoreSpan = document.getElementById("totalScore");
const nextButton = document.getElementById("nextActivity");
const activityArea = document.getElementById("activityArea");
const menuLinks = document.querySelectorAll(".nav-menu__link");
const desktopToggle = document.getElementById("desktopToggle");
const sidebar = document.getElementById("sidebar");

// Função para atualizar pontuação total
function updateTotalScore() {
  totalScore = activities
    .filter((a) => a.completed)
    .reduce((sum, a) => sum + a.score, 0);
  totalScoreSpan.textContent = totalScore;
}

// Função para marcar atividade como concluída
function completeActivity(activityId) {
  const activity = activities.find((a) => a.id === activityId);
  if (activity && !activity.completed) {
    activity.completed = true;

    // Atualizar visual
    const item = activity.element;
    const checkElement = item.querySelector(".nav-menu__check");
    const scoreElement = item.querySelector(".nav-menu__score");

    if (checkElement) {
      checkElement.style.display = "none";
    }
    if (scoreElement) {
      scoreElement.style.display = "inline-block";
    }

    updateTotalScore();
  }
}

// Função para trocar de atividade
function switchActivity(index) {
  if (index < 0 || index >= activities.length) return;

  // Remover active de todos
  menuLinks.forEach((link) => {
    link.classList.remove("nav-menu__link--active");
  });

  // Adicionar active na atividade atual
  const currentActivity = activities[index];
  currentActivity.link.classList.add("nav-menu__link--active");

  // Atualizar título
  currentTitle.textContent = currentActivity.title;

  // Marcar como concluída (simulação)
  completeActivity(currentActivity.id);

  // Habilitar/desabilitar botão próxima
  nextButton.disabled = index >= activities.length - 1;

  // Fechar menu em mobile se estiver aberto
  if (window.innerWidth <= 768) {
    closeMenu();
  }
}

// Event listeners para os links do menu
menuLinks.forEach((link, index) => {
  link.addEventListener("click", (e) => {
    e.preventDefault();
    currentActivityIndex = index;
    switchActivity(index);
  });
});

// Botão próxima atividade
if (nextButton) {
  nextButton.addEventListener("click", () => {
    if (currentActivityIndex < activities.length - 1) {
      currentActivityIndex++;
      switchActivity(currentActivityIndex);
    }
  });
}

// Controle do menu mobile
const menuToggle = document.getElementById("menuToggle");
const overlay = document.getElementById("sidebarOverlay");
const sidebarClose = document.getElementById("sidebarClose");

function openMenu() {
  sidebar.classList.add("open");
  overlay.classList.add("active");
  document.body.style.overflow = "hidden";
}

function closeMenu() {
  sidebar.classList.remove("open");
  overlay.classList.remove("active");
  document.body.style.overflow = "";
}

if (menuToggle) {
  menuToggle.addEventListener("click", openMenu);
}

if (sidebarClose) {
  sidebarClose.addEventListener("click", closeMenu);
}

if (overlay) {
  overlay.addEventListener("click", closeMenu);
}

// Controle do toggle desktop (encolher/expandir)
if (desktopToggle) {
  desktopToggle.addEventListener("click", () => {
    sidebar.classList.toggle("collapsed");

    // Muda a direção da seta
    if (sidebar.classList.contains("collapsed")) {
      desktopToggle.innerHTML = "▶";
      desktopToggle.title = "Expandir menu";
      // Força a posição via JavaScript
      desktopToggle.style.left = "0";
    } else {
      desktopToggle.innerHTML = "◀";
      desktopToggle.title = "Encolher menu";
      // Volta à posição original
      desktopToggle.style.left = "320px";
    }
  });
}

// Fecha o menu ao redimensionar para desktop
window.addEventListener("resize", () => {
  if (window.innerWidth > 768) {
    closeMenu();
  }
});

// Inicialização
updateTotalScore();

// Simular conclusão da primeira atividade após 3 segundos (apenas para demonstração)
setTimeout(() => {
  completeActivity(1);
}, 3000);

// Logout
//const className = {{.Data.ClassName}};
const closeSessionBtn = document.getElementById("closeSession");

if (closeSessionBtn) {
  closeSessionBtn.addEventListener("click", async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`/student-logout/${className}`, {
        method: "POST",
        credentials: "include",
      });

      if (response.ok) {
        window.location.href = `/student-login/${className}`;
      } else {
        console.error("Erro no logout");
      }
    } catch (error) {
      console.error("Erro:", error);
    }
  });
}
