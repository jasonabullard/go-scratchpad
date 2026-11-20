const form = document.getElementById("add-form")
const input = document.getElementById("new-todo")
const deleteTodoButtons = document.querySelectorAll(".btnDeleteTodo")
const completeTodoButtons = document.querySelectorAll(".btnCompleteTodo")

form.addEventListener("submit", async (event) => {
  event.preventDefault();

  await fetch("/api/v1/todos", {
    method: "POST",
    body: JSON.stringify({ body: input.value })
  })

  location.reload()
})

deleteTodoButtons.forEach((el) => {
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

completeTodoButtons.forEach((el) => {
  el.addEventListener("click", async (event) => {
    event.preventDefault();
    const id = el.getAttribute("data-id")
    const done = el.getAttribute("data-done")

    const resp = await fetch("/api/v1/todos", {
      method: "PATCH",
      body: JSON.stringify({ id, done: done === "true" ? false : true })
    })

    if (resp.ok) {
      location.reload()
    }
  })
})