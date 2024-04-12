const parseJwt = (token) => {

    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};

window.onload = () => {

    const tokenString = localStorage.getItem('auth');
    const email = tokenString ? parseJwt(tokenString).email : null;

    htmx.ajax("GET", `/handle/navbar/home/${tokenString}`, { target: "#navbar-profile", swap: "innerHTML" });
    htmx.ajax("GET", `/handle/shortcut/${tokenString}`, { target: "#shortcuts", swap: "beforeend" });

    setTimeout(() => {
        htmx.ajax("GET", `/handle/username/${tokenString}`, { target: "#terminal-console", swap: "beforeend" });
    }, 600);

    setTimeout(() => {
        htmx.ajax("GET", `/handle/request/new/${email}`, { target: "#terminal-console", swap: "beforeend" });
    }, 1200);

    document.addEventListener('focusin', (e) => {});
};

const dots = () => {
    document.getElementById('request-response').innerHTML = '$  curling...';
};

const emptyForm = () => {
    document.getElementById('new-request').reset();
};

const toggleHistoryList = () => {

    const tokenString = localStorage.getItem('auth');
    const email = tokenString ? parseJwt(tokenString).email : null;
    const historyModal = document.getElementById('history-modal');

    if (historyModal.style.display === 'flex') {
        historyModal.style.display = 'none';
    } else {

        htmx.ajax("GET", `/handle/request/history/${email}`, { target: "#history-modal", swap: "innerHTML" });
        historyModal.style.display = 'flex';

        const historyItems= document.getElementsByClassName("history-item");
        setTimeout(() => {
            historyItems[0].focus();
        }, 100);
    };
};

const toggleFavoritesList = () => {

    const tokenString = localStorage.getItem('auth');
    const email = tokenString ? parseJwt(tokenString).email : null;
    const favorites = tokenString ? parseJwt(tokenString).favorites : null;
    const favoritesModal = document.getElementById('favorites-modal');

    if (favoritesModal.style.display === 'flex') {
        favoritesModal.style.display = 'none';
    } else {

        if (favorites === null && email === null) {

            document.getElementById('favorites-modal').innerHTML = `
                <br />
                <p style="margin-left:15px;">$  sign in to save favorites</p>
            `;
            favoritesModal.style.display = 'flex';

            setTimeout(() => {
                favoritesModal.style.display = 'none';
            }, 2000);
        } else {

            htmx.ajax("GET", `/handle/request/favorites/${email}/${favorites}`, { target: "#favorites-modal", swap: "innerHTML" });
            favoritesModal.style.display = 'flex';

            const favoritesItems= document.getElementsByClassName("favorites-item");
            setTimeout(() => {
                favoritesItems[0].focus();
            }, 100);
        };
    };
};

const toggleFavoriteItem = () => {

    const selectedItem = document.activeElement;
    const requestID = selectedItem.value;
    const email = parseJwt(localStorage.getItem('auth').email);
    const password = parseJwt(localStorage.getItem('auth').password);
    const favorites = parseJwt(localStorage.getItem('auth').favorites);
    const updatedFavorites = favorites.append(Number(requestID));

    fetch("/api/user/favorites", { method: "PATCH", body: JSON.stringify({ "email": email, "password": password, "favorites": updatedFavorites, "date": "doesntmatter", "deleted": false })});
};

const hideRequest = () => {

    const selectedItem = document.activeElement;
    const requestID = selectedItem.value;
    const email = parseJwt(localStorage.getItem('auth').email);
    const password = parseJwt(localStorage.getItem('auth').password);
    const history = parseJwt(localStorage.getItem('auth').history);
    const updatedHistory = favorites.append(Number(requestID));

    fetch("/api/user/history?remove=true", { method: "PATCH", body: JSON.stringify({ "email": email, "password": password, "favorites": updatedFavorites, "date": "doesntmatter", "deleted": false })});
}
