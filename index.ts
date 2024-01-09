window.addEventListener("load", () => {
  document.addEventListener("htmx:afterOnLoad", (event: any) => {
    console.log(event);
    const msg = document.getElementById("msg");
    if (event.detail.xhr.status === 200 && msg) {
      console.log("success");
      msg.textContent = "";
    }
  });
});
