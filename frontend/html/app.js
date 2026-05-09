let token = localStorage.getItem("token");

// Base paths (works inside nginx container)
const API = {
    auth: "/auth",
    orders: "/orders",
    products: "/products",
    payments: "/payments",
    notifications: "/notifications",
    profiles: "/profiles"
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

// ---------------- PROFILE ----------------

async function createProfile() {

    try {

        const user_id = parseInt(
            document.getElementById("profile_user_id").value
        );

        const email = document.getElementById("profile_email").value;

        const phone = document.getElementById("profile_phone").value;

        const address = document.getElementById("profile_address").value;

        const res = await fetch(`${API.profiles}/profile`, {
            method: "POST",

            headers: {
                "Content-Type": "application/json"
            },

            body: JSON.stringify({
                user_id,
                email,
                phone,
                address
            })
        });

        log(await res.json());

    } catch (err) {

        handleError(err);
    }
}

async function getProfile() {

    try {

        const userID = document.getElementById(
            "profile_get_user_id"
        ).value;

        const res = await fetch(
            `${API.profiles}/profile/${userID}`
        );

        log(await res.json());

    } catch (err) {

        handleError(err);
    }
}

async function updateProfile() {

    try {

        const userID = document.getElementById(
            "profile_update_user_id"
        ).value;

        const email = document.getElementById(
            "profile_update_email"
        ).value;

        const phone = document.getElementById(
            "profile_update_phone"
        ).value;

        const address = document.getElementById(
            "profile_update_address"
        ).value;

        const res = await fetch(
            `${API.profiles}/profile/${userID}`,
            {
                method: "PUT",

                headers: {
                    "Content-Type": "application/json"
                },

                body: JSON.stringify({
                    email,
                    phone,
                    address
                })
            }
        );

        log(await res.json());

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

// ---------------- PAYMENTS ----------------

async function createPayment() {

    try {

        const user_id = parseInt(
            document.getElementById(
                "payment_user_id"
            ).value
        );

        const order_id = parseInt(
            document.getElementById(
                "payment_order_id"
            ).value
        );

        const amount = parseFloat(
            document.getElementById(
                "payment_amount"
            ).value
        );

        const res = await fetch(
            `${API.payments}/payment`,
            {
                method: "POST",

                headers: {
                    "Content-Type":
                        "application/json"
                },

                body: JSON.stringify({
                    user_id,
                    order_id,
                    amount
                })
            }
        );

        log(await res.json());

    } catch (err) {

        handleError(err);
    }
}


// ---------------- NOTIFICATIONS ----------------

// async function getNotifications() {

//     try {

//         const userID = document.getElementById(
//             "notification_user_id"
//         ).value;

//         const res = await fetch(
//             `${API.notifications}/notifications/${userID}`
//         );

//         log(await res.json());

//     } catch (err) {

//         handleError(err);
//     }
// }