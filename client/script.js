const isLocal = window.location.hostname === "localhost" || window.location.hostname === "127.0.0.1";
const API_BASE_URL = isLocal ? "http://127.0.0.1:8100/v1" : "http://k8s-myipapp-myipserv-076ab4d9ee-572446307.us-east-1.elb.amazonaws.com";

const ipEl = document.getElementById("ip");
const revEl = document.getElementById("reversed");
const errEl = document.getElementById("error");

async function loadIP() {
  ipEl.textContent = "Loading…";
  revEl.textContent = "—";
  errEl.textContent = "";
  try {
        const res = await fetch(`${API_BASE_URL}/myip`, { cache: "no-store" });
        if (!res.ok) throw new Error("HTTP " + res.status);
        const data = await res.json();
        ipEl.textContent = data.ip || "N/A";
        revEl.textContent = data.reversed || "N/A";
    } catch (err) {
        ipEl.textContent = "Unavailable";
        revEl.textContent = "—";
        errEl.textContent = "Error: " + err.message;
      }
    }

document.getElementById("refresh").addEventListener("click", loadIP);
loadIP(); 
