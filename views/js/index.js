window.onload = () => {

    const tokenString = localStorage.getItem('auth');
    const newRequests = document.getElementsByClassName("new-request");
    console.log(newRequests);
    const latestRequest = newRequests.length === 0 ? newRequests[0] : newRequests[newRequests.length - 1];

    htmx.ajax("GET", `/handle/navbar/home/${tokenString}`, { target: "#navbar-profile", swap: "innerHTML" });
    htmx.ajax("GET", `/handle/shortcut/${tokenString}`, { target: "#shortcuts", swap: "beforeend" });

    setTimeout(() => {
        htmx.ajax("GET", `/handle/username/${tokenString}`, { target: "#terminal-console", swap: "beforeend" });
    }, 600);

    setTimeout(() => {
        htmx.ajax("GET", `/handle/request/new`, { target: "#terminal-console", swap: "beforeend" });
    }, 1200);

    latestRequest.addEventListener("submit", async (e) => {

        e.preventDefault();

        const requestData = latestRequest.elements;
        console.log(requestData);
        const method = requestData[0];
        console.log(requestData);

        const newRequest = await fetch("/api/request/new", { method: "POST", body: JSON.stringify({"date": `${Date.now()}`, "hidden": false}) });
        const requestResponse = await newRequest.json();

        // htmx.ajax("POST", `/handle/request/response`, { target: "#request-response", swap: "innerHTML" });
        // htmx.ajax("GET", `/handle/request/new`, { target: "#terminal-console", swap: "beforeend" });
    });
};
