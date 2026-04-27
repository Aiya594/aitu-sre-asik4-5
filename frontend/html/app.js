let token = localStorage.getItem("token") || "";

function log(data) {
    document.getElementById("output").textContent =
        JSON.stringify(data, null, 2);
}

// REGISTER
async function register() {
    const username = document.getElementById("reg_user").value;
    const password = document.getElementById("reg_pass").value;

    const res = await fetch("/auth/register", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ username, password })
    });

    const data = await res.json();
    log(data);
}

// LOGIN
async function login() {
    const username = document.getElementById("log_user").value;
    const password = document.getElementById("log_pass").value;

    const res = await fetch("/auth/login", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ username, password })
    });

    const data = await res.json();

    if (data.token) {
        token = data.token;
        localStorage.setItem("token", token);
    }

    log(data);
}

// CREATE ORDER
async function createOrder() {
    const product = document.getElementById("product").value;
    const amount = parseFloat(document.getElementById("amount").value);

    const res = await fetch("/orders/order", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + token
        },
        body: JSON.stringify({ product, amount })
    });

    const data = await res.json();
    log(data);
}

// GET ORDERS
async function getOrders() {
    const res = await fetch("/orders/orders", {
        headers: {
            "Authorization": "Bearer " + token
        }
    });

    const data = await res.json();
    log(data);
}