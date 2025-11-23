document.querySelector("#list_users").addEventListener("click", loadUsers)
document.querySelector("#add_user").addEventListener("click", UserForm)

async function loadUsers() {
    const res = await fetch("http://localhost:8080/api/user/");
    const data = await res.json();

    console.log(data)

    /*  const list = document.querySelector("#list");
     list.innerHTML = ""; */

    /* if (Object.keys(data.data).length !== 0) {
        const divs = document.querySelectorAll(".Users")
        divs.forEach(d => {
            d.style.display = "block"
        })
    } */

    /*  Object.values(data.data).forEach(u => {
         const li = document.createElement("li");
         li.id = u.ID;
         li.innerText = u.Firstname;
 
         li.addEventListener("click", () => {
             selectedID = u.ID;
         });
 
         list.appendChild(li);
     }); */
}

function UserForm() {
    document.querySelector(".container-main").classList.add("d-none")
    document.querySelector(".addUser").style.display = "block"
}

document.querySelector("#UserForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const data = {
        Firstname: e.target.FirstName.value,
        Lastname: e.target.LastName.value,
        Biography: e.target.Biography.value
    };

    await fetch("http://localhost:8080/api/user/", {
        method: "POST",
        body: JSON.stringify(data)
    });

    e.target.reset()
    loadUsers();
});


/* document.querySelector("#EditForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const data = {
        Firstname: e.target.FirstName.value,
        Lastname: e.target.LastName.value,
        Biography: e.target.Biography.value
    };

    await fetch("http://localhost:8080/api/user/" + selectedID, {
        method: "PUT",
        body: JSON.stringify(data)
    });
    e.target.reset()
    loadUsers();
});

document.querySelector("#Delete").addEventListener("click", async () => {
    if (!selectedID) return;

    await fetch("http://localhost:8080/api/user/" + selectedID, {
        method: "DELETE"
    });

    selectedID = null;
    loadUsers();
});


document.querySelector("#SearchForm").addEventListener("submit", async (e) => {
    e.preventDefault()
    var Firstname = e.target.FirstName.value


    const res = await fetch("http://localhost:8080/api/user/search/" + Firstname)

    const result = await res.json()

    if (result.data.Id == 0) {
        return
    }

    const listInfo = document.querySelector("#listInfo")
    listInfo.innerHTML = ""
    li = document.createElement("li")
    li.id = result.data.ID
    li.innerText = `First name:  ${result.data.Firstname}`

    li2 = document.createElement("li")
    li2.id = result.data.ID
    li2.innerText = `Last name:  ${result.data.Lastname}`

    li3 = document.createElement("li")
    li3.id = result.data.ID
    li3.innerText = `Biography:  ${result.data.Biography}`

    listInfo.appendChild(li)
    listInfo.appendChild(li2)
    listInfo.appendChild(li3)
})
 */