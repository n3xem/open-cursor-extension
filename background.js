chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (message.type === "openInCursor") {
    chrome.runtime.sendNativeMessage( // アプリにメッセージを送信
      "com.github.n3xem.open_cursor_extension", // アプリ名
      {
        org: message.org,
        repo: message.repo,
      },
      (response) => {
        if (chrome.runtime.lastError) {
          console.error("Native messaging error:", chrome.runtime.lastError);
          return;
        }
        console.log("Response from native app:", response);
      }
    );
  }
});
