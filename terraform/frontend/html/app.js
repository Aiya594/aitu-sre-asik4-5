let token = localStorage.getItem("token") || "";

function log(data) {
    document.getElementById("output").textContent =
        JSON.stringify(data, null, 2);
}

//AUTH

async function register() {
    const username = document.getElementById("reg_user").value;
    const password = document.getElementById("reg_pass").value;

    const res = await fetch("/auth/register", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({ username, password })
    });

    log(await res.json());
}

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

//PRODUCTS

async function createProduct() {
    const name = document.getElementById("p_name").value;
    const price = parseFloat(document.getElementById("p_price").value);
    const stock = parseInt(document.getElementById("p_stock").value);

    const res = await fetch("/products/products", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ name, price, stock })
    });

    log(await res.json());
}

async function getProducts() {
    const res = await fetch("/products/products");
    log(await res.json());
}

//ORDERS

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

    log(await res.json());
}

async function getOrders() {
    const res = await fetch("/orders/orders", {
        headers: {
            "Authorization": "Bearer " + token
        }
    });

    log(await res.json());
}