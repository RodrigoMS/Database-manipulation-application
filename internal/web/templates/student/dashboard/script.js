// Controle do menu mobile
const menuToggle = document.getElementById("menuToggle");
const sidebar = document.getElementById("sidebar");
const overlay = document.getElementById("sidebarOverlay");
const sidebarClose = document.getElementById("sidebarClose");

function openMenu() {
  sidebar.classList.add("open");
  overlay.classList.add("active");
  document.body.style.overflow = "hidden"; // Previne scroll do body
}

function closeMenu() {
  sidebar.classList.remove("open");
  overlay.classList.remove("active");
  document.body.style.overflow = ""; // Restaura scroll
}

menuToggle.addEventListener("click", (e) => {
  e.stopPropagation();
  openMenu();
});

sidebarClose.addEventListener("click", closeMenu);
overlay.addEventListener("click", closeMenu);

// Fecha o menu ao redimensionar para desktop
window.addEventListener("resize", () => {
  if (window.innerWidth > 768) {
    closeMenu();
  }
});

// Fecha o menu ao clicar em um link (opcional)
document.querySelectorAll(".nav-menu__link").forEach((link) => {
  link.addEventListener("click", () => {
    if (window.innerWidth <= 768) {
      closeMenu();
    }
  });
});

// Logout
// const className = {{.Data.ClassName}};
document.getElementById("closeSession").addEventListener("click", async () => {
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
