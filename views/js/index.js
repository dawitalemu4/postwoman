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

    document.addEventListener('focusin', () => {});
};

const dots = () => {

    document.getElementById('request-response').innerHTML = '$  curling...';

    const tokenString = localStorage.getItem('auth');
    const profile = tokenString ? parseJwt(tokenString) : null;
    const historyItems = document.getElementsByClassName('history-item');

    toggleHistoryList();
    document.getElementById('history-modal').style.display = 'none';

    const requestID = historyItems.length === 1 ? historyItems[0].id : historyItems[historyItems.length - 1].id;
    const updatedHistory = profile ? (profile.history ? profile.history.push(Number(requestID)) : [requestID]) : null;

    fetch(`/api/user/history/${requestID}`, { method: "PATCH", body: JSON.stringify({ "username": profile.username, "email": profile.email, "password": profile.password, "history": updatedHistory, "deleted": false })});
};

const emptyForm = () => {
    document.getElementById('new-request').reset();
};

const toggleHistoryList = () => {

    const tokenString = localStorage.getItem('auth');
    const email = tokenString ? parseJwt(tokenString).email : null;
    const favoritesModal = document.getElementById('favorites-modal');
    const historyModal = document.getElementById('history-modal');

    if (favoritesModal.style.display === 'flex') {
        toggleFavoritesList();
    };

    if (historyModal.style.display === 'flex') {
        historyModal.style.display = 'none';
    } else {

        htmx.ajax("GET", `/handle/request/history/${email}`, { target: "#history-modal", swap: "innerHTML" });

        historyModal.style.display = 'flex';

        setTimeout(() => {
            if (document.getElementsByClassName("history-item")[0]) {
                document.getElementsByClassName("history-item")[0].focus();
            };
        }, 100);
    };
};

const toggleFavoritesList = () => {

    const tokenString = localStorage.getItem('auth');
    const email = tokenString ? parseJwt(tokenString).email : null;
    const favorites = tokenString ? parseJwt(tokenString).favorites : null;
    const historyModal = document.getElementById('history-modal');
    const favoritesModal = document.getElementById('favorites-modal');

    if (historyModal.style.display === 'flex') {
        toggleHistoryList();
    };

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

            setTimeout(() => {
                if (document.getElementsByClassName("favorites-item")[0]) {
                    document.getElementsByClassName("favorites-item")[0].focus();
                };
            }, 100);
        };
    };
};

const fillForm = () => {};

const toggleFavoriteItem = async () => {

    const selectedItem = document.activeElement;
    const requestID = Number(selectedItem.id);
    const profile = localStorage.getItem('auth') !== null ? parseJwt(localStorage.getItem('auth')) : console.log('no profile found');

    if (selectedItem.className === 'history-item' || selectedItem.className === 'favorites-item') {

        if (profile.favorites || (profile.favorites && profile.favorites.includes(requestID))) {

            const updatedFavorites = profile.favorites.includes(requestID) ?
                    (profile.favorites.length === 1 ? [] : profile.favorites.splice(profile.favorites.indexOf(requestID), 1))
                :
                    (profile.favorites.length === 0 ? [requestID] : profile.favorites.push(requestID));

            const favoriteRequest = await fetch("/api/user/favorites", { method: "PATCH", body: JSON.stringify({ "username": profile.username, "email": profile.email, "password": profile.password, "favorites": updatedFavorites, "deleted": false })});
            const favoriteResponse = await favoriteRequest.json();
            localStorage.setItem('auth', favoriteResponse);

            if (document.getElementById('favorites-modal').style.display === 'flex') {
                toggleFavoritesList();
                toggleFavoritesList();
            } else {
console.log(document.getElementById(requestID), document.getElementById(requestID).querySelector("removed-favorite"));
                document.getElementById(requestID).querySelector("removed-favorite").style.display = 'flex';

                setTimeout(() => {
                    document.getElementById(requestID).querySelector("removed-favorite").style.display = 'none';
                }, 2000);
            };
        } else {

            const updatedFavorites = [requestID];
            const favoriteRequest = await fetch("/api/user/favorites", { method: "PATCH", body: JSON.stringify({ "username": profile.username, "email": profile.email, "password": profile.password, "favorites": updatedFavorites, "deleted": false })});
            const favoriteResponse = await favoriteRequest.json();
            localStorage.setItem('auth', favoriteResponse);

            if (document.getElementById('favorites-modal').style.display === 'flex') {
                toggleFavoritesList();
                toggleFavoritesList();
            } else {
console.log(document.getElementById(requestID), document.getElementById(requestID).querySelector("added-favorite"));
                document.getElementById(requestID).querySelector("added-favorite").style.display = 'flex';

                setTimeout(() => {
                    document.getElementById(requestID).querySelector("added-favorite").style.display = 'none';
                }, 2000);
            };
        };
    };
};

const hideRequest = async () => {

    const selectedItem = document.activeElement;
    const requestID = selectedItem.id;
    const profile = localStorage.getItem('auth') !== null ? parseJwt(localStorage.getItem('auth')) : console.log('no profile found');
    const updatedHistory = profile.history.push(Number(requestID));

    const hideRequestRequest = await fetch(`/api/user/history/${requestID}?remove=true`, { method: "PATCH", body: JSON.stringify({ "email": profile.email, "password": profile.password, "favorites": updatedHistory, "deleted": false })});
    const hideRequestResponse = await hideRequestRequest.json();
    localStorage.setItem('auth', hideRequestResponse);

    toggleHistoryList();
    toggleFavoritesList();
};

