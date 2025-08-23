const API_BASE_URL = process.env.API_BASE_URL || "http://localhost:8100/v1";

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
