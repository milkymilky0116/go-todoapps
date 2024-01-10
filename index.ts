window.addEventListener("load", () => {
  document.addEventListener("htmx:afterOnLoad", (event: any) => {
    console.log(event);
    const context = document.getElementById("context") as HTMLInputElement;
    const msg = document.getElementById("msg");
    if (event.detail.xhr.status === 200 && msg && context) {
      context.value = "";
      console.log("success");
      msg.textContent = "";
    }
  });
});
