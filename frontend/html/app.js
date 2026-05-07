let token = localStorage.getItem("token");

// Base paths (works inside nginx container)
const API = {
    auth: "/auth",
    orders: "/orders",
    products: "/products"
};

function log(data) {
    document.getElementById("output").textContent =
        JSON.stringify(data, null, 2);
}

function handleError(err) {
    log({ error: err.message || err });
}

// ---------------- AUTH ----------------

async function register() {
    try {
        const username = document.getElementById("reg_user").value;
        const password = document.getElementById("reg_pass").value;

        const res = await fetch(`${API.auth}/register`, {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({ username, password })
        });

        log(await res.json());
    } catch (err) {
        handleError(err);
    }
}

async function login() {
    try {
        const username = document.getElementById("log_user").value;
        const password = document.getElementById("log_pass").value;

        const res = await fetch(`${API.auth}/login`, {
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
    } catch (err) {
        handleError(err);
    }
}

// ---------------- PRODUCTS ----------------

async function createProduct() {
    try {
        const name = document.getElementById("p_name").value;
        const price = parseFloat(document.getElementById("p_price").value);
        const stock = parseInt(document.getElementById("p_stock").value);

        const res = await fetch(`${API.products}/product`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ name, price, stock })
        });

        log(await res.json());
    } catch (err) {
        handleError(err);
    }
}

async function getProducts() {
    try {
        const res = await fetch(`${API.products}/products`);
        log(await res.json());
    } catch (err) {
        handleError(err);
    }
}

// ---------------- ORDERS (FIXED) ----------------

async function createOrder() {
    try {
        const product_id = parseInt(document.getElementById("product_id").value);
        const amount = parseFloat(document.getElementById("amount").value);

        const res = await fetch(`${API.orders}/order`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + token
            },
            body: JSON.stringify({ product_id, amount })
        });

        log(await res.json());
    } catch (err) {
        handleError(err);
    }
}

async function getOrders() {
    try {
        const res = await fetch(`${API.orders}/orders`, {
            headers: {
                "Authorization": "Bearer " + token
            }
        });

        log(await res.json());
    } catch (err) {
        handleError(err);
    }
}