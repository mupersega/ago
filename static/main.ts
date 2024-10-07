// add the beginning of your app entry
import "vite/modulepreload-polyfill";

import htmx from "htmx.org";
import _hyperscript from "hyperscript.org";
_hyperscript.browserInit();

import "./style.scss";
import "./three-dee.ts";
import "./data.ts";
import { canvasResize, hideCanvases, revealCanvases } from "./two-dee.ts";

import { MapBuildData } from "./interface.ts";

import { appState } from "./state.ts";

declare global {
  interface Window {
    SetTileMapSize: () => void;
    getTileMap: () => Promise<MapBuildData>;
    appState: any;
    resizeDebounce: number;
  }
}

function ReTrigger(e) {
  if (
    (e.type === "mouseover" || e.type === "click") &&
    e.shiftKey === false &&
    e.ctrlKey === false
  ) {
    return;
  }
  var target = e.currentTarget;
  if (!(target instanceof HTMLElement)) {
    return;
  }
  if (target.nodeName !== "SPAN") {
    return;
  }
  var magnitude = 1;
  var pmag = window.$pMag || -1;
  var shapeEvent = new CustomEvent("shape", {
    detail: {
      magnitude: e.shiftKey ? magnitude : -1 * magnitude,
      prescribedMagnitude: pmag,
    },
  });
  target.dispatchEvent(shapeEvent);
}
htmx.on("htmx:load", function (e) {
  var element = e.detail.elt;

  // Add event listeners to all 'span.tile' elements within the loaded element
  element.querySelectorAll("span.tile").forEach(function (span) {
    span.addEventListener("click", ReTrigger);
    span.addEventListener("mouseover", ReTrigger);
  });

  // If the loaded element itself is a 'span.tile', add event listeners to it
  if (element.matches("span.tile")) {
    element.addEventListener("click", ReTrigger);
    element.addEventListener("mouseover", ReTrigger);
  }
});
// add control and shift keydown listeners
window.addEventListener("keydown", (e) => {
  var pmag = window.$pMag || -1;
  if (e.key === "Shift") {
    if (pmag === -1) {
      document.getElementById("lift").classList.add("active");
      document.getElementById("mini-map").classList.add("active");
    } else {
      const magSelected = document.querySelector(".mag-selected");
      if (magSelected) {
        magSelected.classList.add("active");
      }
    }
  }
  if (e.key === "Control") {
    if (pmag === -1) {
      document.getElementById("mini-map").classList.add("active");
      document.getElementById("lower").classList.add("active");
    } else {
      const magSelected = document.querySelector(".mag-selected");
      if (magSelected) {
        magSelected.classList.add("active");
      }
    }
  }
});

window.addEventListener("keyup", (e) => {
  if (e.key === "Shift") {
    document.getElementById("lift").classList.remove("active");
    document.getElementById("mini-map").classList.remove("active");
    const magSelected = document.querySelector(".mag-selected");
    if (magSelected) {
      magSelected.classList.remove("active");
    }
  }
  if (e.key === "Control") {
    document.getElementById("lower").classList.remove("active");
    document.getElementById("mini-map").classList.remove("active");
    const magSelected = document.querySelector(".mag-selected");
    if (magSelected) {
      magSelected.classList.remove("active");
    }
  }
});

function listen() {
  // store all main listener code here
  let resizeDebounce;

  // resize listener
  window.addEventListener("resize", () => {
    hideCanvases();
    clearTimeout(resizeDebounce);
    resizeDebounce = setTimeout(() => {
      canvasResize();
      revealCanvases();
    }, 400);
  });
}

const setTileMapSize = () => {
  const mapWrapper = document.getElementById("mapWrapper");
  const tileMap = document.getElementById("mini-map");

  if (!mapWrapper || !tileMap) {
    return;
  }

  tileMap.style.display = "none";

  const min = Math.min(mapWrapper.offsetWidth, mapWrapper.offsetHeight);

  tileMap.style.width = `${min}px`;
  tileMap.style.height = `${min}px`;
  tileMap.style.display = "grid";
};

window.addEventListener("resize", () => {
  setTileMapSize();
  console.log("resize");
});

listen();

window.SetTileMapSize = setTileMapSize;
(window.appState as any) = appState;
