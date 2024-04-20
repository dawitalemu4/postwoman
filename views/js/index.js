const parseJwt = (token) => {

    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};

window.onload = () => {

    const tokenString = localStorage.getItem("auth");
    const email = tokenString ? parseJwt(tokenString).email : null;

    htmx.ajax("GET", `/handle/navbar/home/${tokenString}`, { target: "#navbar-profile", swap: "innerHTML" });
    htmx.ajax("GET", `/handle/shortcut/${tokenString}`, { target: "#shortcuts", swap: "beforeend" });

    setTimeout(() => {
        htmx.ajax("GET", `/handle/username/${tokenString}`, { target: "#terminal-console", swap: "beforeend" });
    }, 600);

    setTimeout(() => {
        htmx.ajax("GET", `/handle/request/new/${email}`, { target: "#terminal-console", swap: "beforeend" });
    }, 1200);

    document.addEventListener("focusin", () => {});
};

const loading = () => {
    document.getElementById("request-response").innerHTML = "$  curling...";
};

const formatResponse = () => {

    const responseTextarea = document.getElementById("response-textarea");

    if (responseTextarea.textContent.charAt(0) === "{" || responseTextarea.textContent.charAt(0) === "[") {
        responseTextarea.textContent = JSON.stringify(JSON.parse(responseTextarea.textContent), null, 4) + "\n";
    };
};

const fillForm = () => {

    const selectedItem = document.activeElement;
    const curlForm = document.getElementById("new-request");

    const methodField = curlForm.children.method;
    const headersField = curlForm.children.headers;
    const originField = curlForm.children.origin;
    const bodyField = curlForm.children.body;
    const urlField = curlForm.children.url;

    if (selectedItem.className === "history-item" || selectedItem.className === "favorites-item") {

        methodField.value = selectedItem.children[0].children[1].innerHTML;
        headersField.value = selectedItem.children.headers.value;
        originField.value = selectedItem.children.origin.value;
        bodyField.value = selectedItem.children.body.value;
        urlField.value = selectedItem.children[1].children[0].innerHTML;

        document.getElementById("history-modal").style.display = "none";
        document.getElementById("favorites-modal").style.display = "none";

        methodField.focus();
    };
};

const emptyForm = () => {
    document.getElementById("new-request").reset();
};

const removeItem = (array, id) => {
    array.splice(array.indexOf(id), 1);
    return array;
};

const toggleHistoryList = () => {

    const tokenString = localStorage.getItem("auth");
    const email = tokenString ? parseJwt(tokenString).email : null;
    const favoritesModal = document.getElementById("favorites-modal");
    const historyModal = document.getElementById("history-modal");

    if (favoritesModal.style.display === "flex") {
        toggleFavoritesList();
    };

    if (historyModal.style.display === "flex") {
        historyModal.style.display = "none";
    } else {

        htmx.ajax("GET", `/handle/request/history/${email}`, { target: "#history-modal", swap: "innerHTML" });

        historyModal.style.display = "flex";

        setTimeout(() => {
            if (document.getElementsByClassName("history-item")[0]) {
                document.getElementsByClassName("history-item")[0].focus();
            };
        }, 100);
    };
};

const toggleFavoritesList = () => {

    const tokenString = localStorage.getItem("auth");
    const email = tokenString ? parseJwt(tokenString).email : null;
    const historyModal = document.getElementById("history-modal");
    const favoritesModal = document.getElementById("favorites-modal");

    if (historyModal.style.display === "flex") {
        toggleHistoryList();
    };

    if (favoritesModal.style.display === "flex") {
        favoritesModal.style.display = "none";
    } else {

        if (email === null) {

            document.getElementById("favorites-modal").innerHTML = `
                <br />
                <p style="margin-left:15px;">$  sign in to save favorites</p>
            `;

            favoritesModal.style.display = "flex";

            setTimeout(() => {
                favoritesModal.style.display = "none";
            }, 2000);
        } else {

            htmx.ajax("GET", `/handle/request/favorites/${email}`, { target: "#favorites-modal", swap: "innerHTML" });

            favoritesModal.style.display = "flex";

            setTimeout(() => {
                if (document.getElementsByClassName("favorites-item")[0]) {
                    document.getElementsByClassName("favorites-item")[0].focus();
                };
            }, 100);
        };
    };
};

const toggleFavoriteItem = async () => {

    const selectedItem = document.activeElement;
    const requestID = Number(selectedItem.id);
    const profile = localStorage.getItem("auth") !== null ? parseJwt(localStorage.getItem("auth")) : null;

    if (selectedItem.className === "history-item" || selectedItem.className === "favorites-item") {

        if (!profile) {

            document.getElementById(requestID).children[4].style.display = "flex";

            setTimeout(() => {
                document.getElementById(requestID).children[4].style.display = "none"; 
            }, 1000);

        } else if (profile.favorites || (profile.favorites && profile.favorites.includes(requestID))) {

            const favoritesBeforeUpdate = profile.favorites;
            const updatedFavorites = profile.favorites.indexOf(requestID) !== -1 ? (
                profile.favorites.length > 1 ? removeItem(profile.favorites, requestID) : [] 
            ) : ( 
                profile.favorites.length === 0 ? [requestID] : [...profile.favorites, requestID]
            );

            const favoriteRequest = await fetch("/api/user/favorites", { method: "PATCH", body: JSON.stringify({ "username": profile.username, "email": profile.email, "password": profile.password, "favorites": updatedFavorites, "deleted": false })});
            const favoriteResponse = await favoriteRequest.json();
            localStorage.setItem("auth", favoriteResponse);

            if (document.getElementById("favorites-modal").style.display === "flex") {
                toggleFavoritesList();
                toggleFavoritesList();
            } else {

                const responseMessage = updatedFavorites.length > favoritesBeforeUpdate.length ? document.getElementById(requestID).children[2] : document.getElementById(requestID).children[3];
                responseMessage.style.display = "flex";

                setTimeout(() => {
                    responseMessage.style.display = "none"; 
                }, 1000);
            };
        } else {

            const updatedFavorites = [requestID];
            const favoriteRequest = await fetch("/api/user/favorites", { method: "PATCH", body: JSON.stringify({ "username": profile.username, "email": profile.email, "password": profile.password, "favorites": updatedFavorites, "deleted": false })});
            const favoriteResponse = await favoriteRequest.json();
            localStorage.setItem("auth", favoriteResponse);

            if (document.getElementById("favorites-modal").style.display === "flex") {
                toggleFavoritesList();
                toggleFavoritesList();
            } else {

                document.getElementById(requestID).children[2].style.display = "flex";

                setTimeout(() => {
                    document.getElementById(requestID).children[2].style.display = "none"; 
                }, 1000);
            };
        };
    };
};

const hideRequest = async () => {

    const selectedItem = document.activeElement;
    const tokenString = localStorage.getItem("auth");
    const email = tokenString ? parseJwt(tokenString).email : null;

    if (selectedItem.className === "history-item" || selectedItem.className === "favorites-item") {

        const requestID = selectedItem.id;

        fetch(`/api/request/delete/${email}/${requestID}`, { method: "DELETE" });

        if (selectedItem.className === "history-item") {
            selectedItem.children[5].style.display = "flex";
        } else if (selectedItem.className === "favorites-item") {
            selectedItem.children[4].style.display = "flex";
        };
    };
};
