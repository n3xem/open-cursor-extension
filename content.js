function getRepoInfo() {
  const pathParts = window.location.pathname.split("/").filter(Boolean);
  if (pathParts.length >= 2) {
    return {
      org: pathParts[0],
      repo: pathParts[1],
    };
  }
  return null;
}

// GitHubページにボタンを追加
function addOpenButton() {
  const repoInfo = getRepoInfo();
  if (!repoInfo) return;

  const nav = document.querySelector('nav[aria-label="Repository"]');
  if (!nav) return;

  const button = document.createElement("button");
  button.textContent = "Open in Cursor";
  button.className = "btn btn-sm";
  button.style.marginLeft = "8px";

  button.addEventListener("click", () => {
    chrome.runtime.sendMessage({
      type: "openInCursor",
      org: repoInfo.org,
      repo: repoInfo.repo,
    });
  });

  nav.appendChild(button);
}

addOpenButton();
