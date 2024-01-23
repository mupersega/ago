/**
 * @param {Event} e
 */

function ReTrigger(e) {
  if (e.type === "mouseover" && e.shiftKey === false) {
    return;
  }
  var target = e.currentTarget;
  if (!(target instanceof HTMLElement)) {
    return;
  }
  if (target.nodeName !== "SPAN") {
    return;
  }
  const shapeEvent = new CustomEvent("shape", {
    detail: {
      magnitude: e.ctrlKey ? magnitude : -1 * magnitude,
    },
  });
  target.dispatchEvent(shapeEvent);
}

htmx.on("htmx:load", function (e) {
  var element = e.detail.elt;
  element.querySelectorAll("span").forEach(function (span) {
    span.addEventListener("click", ReTrigger);
    span.addEventListener("mouseover", ReTrigger);
  });
  if (element.nodeName === "SPAN") {
    element.addEventListener("click", ReTrigger);
    element.addEventListener("mouseover", ReTrigger);
  }
});
