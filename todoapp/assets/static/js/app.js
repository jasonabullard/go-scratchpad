const form = document.getElementById("add-form")
const input = document.getElementById("new-todo")
const btnElements = document.querySelectorAll(".btn-delete-todo")

form.addEventListener("submit", async (event) => {
  event.preventDefault();

  await fetch("/api/v1/todos", {
    method: "POST",
    body: JSON.stringify({ body: input.value })
  })

  location.reload()
})

btnElements.forEach((el) => {
  el.addEventListener("click", async (event) => {
    event.preventDefault();
    const id = el.getAttribute("data-id")

    const resp = await fetch(`/api/v1/todos/${id}`, {
      method: "DELETE"
    })

    if (resp.ok) {
      location.reload()
    }
  })
})