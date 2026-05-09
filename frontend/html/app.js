let token = localStorage.getItem("token") || "";
let currentUserId = null;

const API = {
    auth: "/auth/",
    orders: "/orders/",
    products: "/products/",
    payment: "/payment/",
    profile: "/profile/",
};

function log(data) {
    const output = document.getElementById("output");
    output.textContent = typeof data === "string" ? data : JSON.stringify(data, null, 2);
}

function handleError(err) {
    log({ error: err.message || err });
}

function setAuthState() {
    const status = document.getElementById("auth-status");

    if (token) {
        const claims = parseJwt(token);
        currentUserId = claims?.user_id || null;
        status.textContent = `Logged in as user ${currentUserId || "unknown"}`;
        document.getElementById("logout-button").style.display = "inline-flex";
    } else {
        currentUserId = null;
        status.textContent = "Not logged in";
        document.getElementById("logout-button").style.display = "none";
    }
}

function parseJwt(token) {
    try {
        const base64Payload = token.split(".")[1];
        const decodedPayload = atob(base64Payload.replace(/-/g, "+").replace(/_/g, "/"));
        return JSON.parse(decodeURIComponent(decodedPayload.split("").map(function(c) {
            return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        }).join("")));
    } catch (err) {
        return null;
    }
}

async function apiFetch(path, options = {}) {
    const headers = options.headers || {};
    if (!headers["Content-Type"]) {
        headers["Content-Type"] = "application/json";
    }

    if (token) {
        headers["Authorization"] = `Bearer ${token}`;
    }

    const response = await fetch(path, { ...options, headers });

    const contentType = response.headers.get("content-type") || "";
    const body = contentType.includes("application/json")
        ? await response.json()
        : await response.text();

    if (!response.ok) {
        throw new Error(body?.error || body || `HTTP ${response.status}`);
    }

    return body;
}

async function register() {
    try {
        const username = document.getElementById("log_user").value;
        const password = document.getElementById("log_pass").value;

        const result = await apiFetch(`${API.auth}register`, {
            method: "POST",
            body: JSON.stringify({ username, password }),
        });

        log(result);
    } catch (err) {
        handleError(err);
    }
}

async function login() {
    try {
        const username = document.getElementById("log_user").value;
        const password = document.getElementById("log_pass").value;

        const result = await apiFetch(`${API.auth}login`, {
            method: "POST",
            body: JSON.stringify({ username, password }),
        });

        if (result.token) {
            token = result.token;
            localStorage.setItem("token", token);
            setAuthState();
            log({ status: "Login successful" });
        } else {
            log(result);
        }
    } catch (err) {
        handleError(err);
    }
}

function logout() {
    token = "";
    currentUserId = null;
    localStorage.removeItem("token");
    setAuthState();
    log("Logged out");
}

async function loadProducts() {
    try {
        const products = await apiFetch(`${API.products}products`);
        renderProductList(products);
        renderProductOptions(products);
        log({ products });
    } catch (err) {
        handleError(err);
    }
}

async function createProduct() {
    try {
        const name = document.getElementById("p_name").value;
        const price = parseFloat(document.getElementById("p_price").value);
        const stock = parseInt(document.getElementById("p_stock").value, 10);

        const result = await apiFetch(`${API.products}product`, {
            method: "POST",
            body: JSON.stringify({ name, price, stock }),
        });

        await loadProducts();
        log(result);
    } catch (err) {
        handleError(err);
    }
}

function renderProductList(products) {
    const select = document.getElementById("product-select");
    select.innerHTML = "";

    products.forEach((product) => {
        const option = document.createElement("option");
        option.value = product.id;
        option.textContent = `${product.id}: ${product.name} - $${product.price} (${product.stock} in stock)`;
        select.appendChild(option);
    });
}

function renderProductOptions(products) {
    const select = document.getElementById("order-product-select");
    select.innerHTML = "";

    products.forEach((product) => {
        const option = document.createElement("option");
        option.value = `${product.id}|${product.name}`;
        option.textContent = `${product.name} [ID ${product.id}]`;
        select.appendChild(option);
    });
}

async function createOrder() {
    try {
        if (!token) {
            throw new Error("Login required to create an order");
        }

        const selected = document.getElementById("order-product-select").value;
        if (!selected) {
            throw new Error("Choose a product first");
        }

        const [productId, productName] = selected.split("|");
        const amount = parseFloat(document.getElementById("order_amount").value);

        const result = await apiFetch(`${API.orders}order`, {
            method: "POST",
            body: JSON.stringify({
                product_id: parseInt(productId, 10),
                product_name: productName,
                amount,
            }),
        });

        await loadOrders();
        log(result);
    } catch (err) {
        handleError(err);
    }
}

async function loadOrders() {
    try {
        if (!token) {
            throw new Error("Login required to view orders");
        }

        const orders = await apiFetch(`${API.orders}orders`, { method: "GET" });
        const list = document.getElementById("order-list");
        list.innerHTML = "";

        orders.forEach((order) => {
            const item = document.createElement("li");
            item.textContent = `#${order.id}: ${order.product} — ${order.amount}`;
            list.appendChild(item);
        });

        log({ orders });
    } catch (err) {
        handleError(err);
    }
}

async function processPayment() {
    try {
        const orderId = document.getElementById("payment_order_id").value;
        const amount = parseFloat(document.getElementById("payment_amount").value);

        const result = await apiFetch(`/payment/payment`, {
            method: "POST",
            body: JSON.stringify({
                order_id: orderId,
                user_id: currentUserId || 0,
                amount,
            }),
        });

        await loadPayments();
        log(result);
    } catch (err) {
        handleError(err);
    }
}

async function loadPayments() {
    try {
        const payments = await apiFetch(`/payment/payments`, { method: "GET" });
        const list = document.getElementById("payment-list");
        list.innerHTML = "";

        payments.forEach((payment) => {
            const item = document.createElement("li");
            item.textContent = `#${payment.id}: order=${payment.order_id} user=${payment.user_id} amount=${payment.amount} status=${payment.status}`;
            list.appendChild(item);
        });

        log({ payments });
    } catch (err) {
        handleError(err);
    }
}

async function saveProfile() {
    try {
        if (!currentUserId) {
            throw new Error("Login required to save profile");
        }

        const email = document.getElementById("profile_email").value;
        const phone = document.getElementById("profile_phone").value;
        const address = document.getElementById("profile_address").value;

        const result = await apiFetch(`${API.profile}`, {
            method: "POST",
            body: JSON.stringify({
                user_id: currentUserId,
                email,
                phone,
                address,
            }),
        });

        log(result);
    } catch (err) {
        handleError(err);
    }
}

async function loadProfile() {
    try {
        if (!currentUserId) {
            throw new Error("Login required to load profile");
        }

        const profile = await apiFetch(`${API.profile}${currentUserId}`, { method: "GET" });

        document.getElementById("profile_email").value = profile.email || "";
        document.getElementById("profile_phone").value = profile.phone || "";
        document.getElementById("profile_address").value = profile.address || "";

        log(profile);
    } catch (err) {
        handleError(err);
    }
}

window.addEventListener("load", () => {
    setAuthState();
    loadProducts().catch(() => {});
});
