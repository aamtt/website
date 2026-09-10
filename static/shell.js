(() => {
  const picker = document.querySelector(".shells");
  const list = document.querySelector(".shells [role='tablist']");
  const recipes = document.querySelectorAll(".recipe[data-shell]");
  if (!picker || !list || !recipes.length) return;

  const buttons = Array.from(list.querySelectorAll(".shell"));
  const glide = list.querySelector(".glide");
  const panel = document.querySelector(".snippet");
  const ghost = list.querySelector(".ghost");

  const place = (mark, button) => {
    if (!mark || !button) return;
    mark.style.setProperty("--x", `${button.offsetLeft}px`);
    mark.style.setProperty("--w", `${button.offsetWidth}px`);
  };

  const settle = () => {
    const picked = buttons.find(
      (button) => button.getAttribute("aria-selected") === "true",
    );
    place(glide, picked);
    const joined = picked === buttons[0];
    list.classList.toggle("at-start", joined);
    if (panel) panel.classList.toggle("is-joined", joined);
  };
  const known = new Set(buttons.map((button) => button.dataset.shell));
  const storageKey = "shell";

  const onWindows = () => {
    const hinted = navigator.userAgentData && navigator.userAgentData.platform;
    if (hinted) return hinted === "Windows";
    return /Win/i.test(navigator.platform || navigator.userAgent);
  };

  const show = (wanted) => {
    recipes.forEach((recipe) => {
      recipe.hidden = recipe.dataset.shell !== wanted;
    });
    buttons.forEach((button) => {
      const picked = button.dataset.shell === wanted;
      button.setAttribute("aria-selected", String(picked));
      button.tabIndex = picked ? 0 : -1;
    });
    settle();
  };

  const remember = (wanted) => {
    try {
      localStorage.setItem(storageKey, wanted);
    } catch {}
  };

  let remembered = null;
  try {
    remembered = localStorage.getItem(storageKey);
  } catch {}

  show(
    known.has(remembered) ? remembered : onWindows() ? "powershell" : "posix",
  );
  list.classList.add("is-settling");
  picker.hidden = false;
  settle();
  requestAnimationFrame(() => {
    requestAnimationFrame(() => list.classList.remove("is-settling"));
  });

  buttons.forEach((button) => {
    button.addEventListener("pointerenter", () => {
      place(ghost, button);
      list.classList.add("is-hovering");
    });
    button.addEventListener("focus", () => {
      place(ghost, button);
      list.classList.add("is-hovering");
    });
  });

  const clearGhost = () => list.classList.remove("is-hovering");
  list.addEventListener("pointerleave", clearGhost);
  list.addEventListener("focusout", clearGhost);

  addEventListener("resize", settle);
  if (document.fonts && document.fonts.ready) {
    document.fonts.ready.then(settle);
  }

  buttons.forEach((button) => {
    button.addEventListener("click", () => {
      show(button.dataset.shell);
      remember(button.dataset.shell);
    });
  });

  list.addEventListener("keydown", (event) => {
    const step =
      event.key === "ArrowRight" || event.key === "ArrowDown"
        ? 1
        : event.key === "ArrowLeft" || event.key === "ArrowUp"
          ? -1
          : event.key === "Home"
            ? -buttons.length
            : event.key === "End"
              ? buttons.length
              : 0;
    if (!step) return;

    event.preventDefault();
    const at = buttons.findIndex(
      (button) => button.getAttribute("aria-selected") === "true",
    );
    const next =
      buttons[Math.min(buttons.length - 1, Math.max(0, at + step))] ||
      buttons[at];
    show(next.dataset.shell);
    remember(next.dataset.shell);
    next.focus();
  });
})();
